package models

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/PaulsBecks/OracleFactory/src/services/ethereum"
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"gorm.io/gorm"
)

type KeyedMutex struct {
	mutexes sync.Map // Zero value is empty and ready for use
}

func (m *KeyedMutex) Lock(key uint) func() {
	value, _ := m.mutexes.LoadOrStore(key, &sync.Mutex{})
	mtx := value.(*sync.Mutex)
	mtx.Lock()

	return func() { mtx.Unlock() }
}

var keyedMutex = &KeyedMutex{}
var latestNonceByAddress = &sync.Map{}

type SmartContractConsumer struct {
	gorm.Model
	SmartContractID      uint
	SmartContract        SmartContract
	ProviderConsumerID   uint
	ProviderConsumer     ProviderConsumer
	InboundSubscriptions []InboundSubscription
}

func GetSmartContractConsumerByID(ID uint) *SmartContractConsumer {
	db := utils.DBConnection()
	var smartContractConsumer *SmartContractConsumer
	db.Find(&smartContractConsumer, ID)
	return smartContractConsumer
}

func (s *SmartContractConsumer) GetSmartContract() *SmartContract {
	return GetSmartContractByID(s.SmartContractID)
}

func (iot *SmartContractConsumer) GetEventParameterJSON() string {
	json := "["
	for i, v := range iot.GetProviderConsumer().EventParameters {
		json += "{\"internalType\":\"" + v.Type + "\",\"name\":\"" + v.Name + "\",\"type\":\"" + v.Type + "\"}"
		if i < len(iot.GetProviderConsumer().EventParameters)-1 {
			json += ","
		}
	}
	json += "]"
	return json
}

func (s *SmartContractConsumer) GetProviderConsumer() ProviderConsumer {
	return GetProviderConsumerByID(s.ProviderConsumerID)
}

func (s *SmartContractConsumer) Publish(subscription *InboundSubscription, event *Event) error {
	user := subscription.GetSubscription().GetUser()
	blockchain := s.GetSmartContract().BlockchainName
	switch blockchain {
	case "Ethereum":
		return s.CreateEthereumTransaction(user, event)
	case "Hyperledger":
		return s.CreateHyperledgerTransaction(user, event)
	}
	return fmt.Errorf("No Blockchain smart contract consumer available for blockchain %s", blockchain)
}

func ParseValues(event *Event) ([]interface{}, error) {
	var bodyData map[string]interface{}

	if e := json.Unmarshal(event.Body, &bodyData); e != nil {
		return nil, e
	}

	var parameters []interface{}
	for _, eventValue := range event.GetEventValues() {
		eventParameter := eventValue.GetEventParameter()
		v := bodyData[eventParameter.Name]
		parameter, err := utils.TransformParameterType(v, eventParameter.Type)
		if err != nil {
			return nil, err
		}
		parameters = append(parameters, parameter)
	}
	return parameters, nil
}

// Blockchain Smart Contract Creation
func retry(callback func() error, retries int) error {
	err := callback()
	if retries <= 0 || err == nil {
		return err
	}
	time.Sleep(200 * time.Millisecond)
	return retry(callback, retries-1)
}

type CreateTransaction func(user *User, event *Event) error

func (s *SmartContractConsumer) CreateEthereumTransaction(user *User, event *Event) error {
	// make sure every user is only creating one transaction at a time in order they arrive

	client, err := ethclient.Dial(user.EthereumAddress)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(user.EthereumPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	smartContract := s.GetSmartContract()
	address := common.HexToAddress(smartContract.ContractAddress)
	inputs := s.GetEventParameterJSON()
	name := smartContract.EventName
	abi := "[{\"inputs\":" + inputs + ",\"name\":\"" + name + "\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	fmt.Println(abi, time.Now())
	instance, err := ethereum.NewStore(address, client, abi)
	if err != nil {
		return err
	}

	parameters, err := ParseValues(event)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	unlock := keyedMutex.Lock(user.ID)
	defer unlock()

	sendTransaction := func() error {
		fmt.Printf("INFO: start prepare transaction %s\n", time.Now())

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			return err
		}
		cachedNonce, nonceFound := latestNonceByAddress.Load(user.ID)
		fmt.Printf("INFO: nonce found %d %d %s\n", nonce, cachedNonce, time.Now())

		if nonceFound && cachedNonce != nil && nonce < cachedNonce.(uint64) {
			nonce = cachedNonce.(uint64)
		}
		fmt.Printf("INFO: nonce found %d %d %s\n", nonce, cachedNonce, time.Now())
		auth.Nonce = big.NewInt(int64(nonce))
		fmt.Printf("INFO: send transaction %s\n", time.Now())
		tx, err := instance.StoreTransactor.Contract.Transact(auth, name, parameters...)
		if err != nil {
			fmt.Printf("INFO: error while sending transaction %s %s\n", err.Error(), time.Now())
			return err
		}
		latestNonceByAddress.Store(user.ID, nonce+1)
		fmt.Printf("INFO: tx sent %s %s\n", tx.Hash().Hex(), time.Now())
		return nil
	}
	return retry(sendTransaction, 10)
}

func (s *SmartContractConsumer) CreateHyperledgerTransaction(user *User, event *Event) error {
	//_ = os.Setenv("DISCOVERY_AS_LOCALHOST", "false")
	organizationName := user.HyperledgerOrganizationName
	cert := user.HyperledgerCert
	key := user.HyperledgerKey
	gatewayConfig := user.HyperledgerConfig
	channel := user.HyperledgerChannel

	smartContract := s.GetSmartContract()
	contractAddress := smartContract.ContractAddress
	contractName := smartContract.EventName
	parameters := []string{}
	for _, eventValue := range event.GetEventValues() {
		parameters = append(parameters, eventValue.Value)
	}

	wallet := gateway.NewInMemoryWallet()
	wallet.Put("appUser", gateway.NewX509Identity(organizationName, string(cert), string(key)))
	config := config.FromRaw([]byte(gatewayConfig), "yaml")
	gw, err := gateway.Connect(
		gateway.WithConfig(config),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		return err
	}
	defer gw.Close()

	network, err := gw.GetNetwork(channel)
	if err != nil {
		return fmt.Errorf("Failed to get network: %v", err)
	}

	contract := network.GetContract(contractAddress)

	_, err = contract.SubmitTransaction(contractName, parameters...)
	if err != nil {
		return fmt.Errorf("Failed to Submit transaction: %v", err)
	}

	return nil
}
