package models

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

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

type SmartContractPublisher struct {
	gorm.Model
	SmartContractID     uint
	SmartContract       SmartContract
	ListenerPublisherID uint
	ListenerPublisher   ListenerPublisher
	InboundOracles      []InboundOracle
}

func GetSmartContractPublisherByID(ID uint) *SmartContractPublisher {
	db := utils.DBConnection()
	var smartContractPublisher *SmartContractPublisher
	db.Find(&smartContractPublisher, ID)
	return smartContractPublisher
}

func (s *SmartContractPublisher) GetSmartContract() *SmartContract {
	return GetSmartContractByID(s.SmartContractID)
}

func (iot *SmartContractPublisher) GetEventParameterJSON() string {
	json := "["
	for i, v := range iot.GetListenerPublisher().EventParameters {
		json += "{\"internalType\":\"" + v.Type + "\",\"name\":\"" + v.Name + "\",\"type\":\"" + v.Type + "\"}"
		if i < len(iot.GetListenerPublisher().EventParameters)-1 {
			json += ","
		}
	}
	json += "]"
	return json
}

func (s *SmartContractPublisher) GetListenerPublisher() ListenerPublisher {
	return GetListenerPublisherByID(s.ListenerPublisherID)
}

func (s *SmartContractPublisher) Publish(oracle *InboundOracle, event *Event) error {
	user := oracle.GetOracle().GetUser()
	blockchain := s.GetSmartContract().BlockchainName
	switch blockchain {
	case "Ethereum":
		return s.CreateEthereumTransaction(user, event)
	case "Hyperledger":
		return s.CreateHyperledgerTransaction(user, event)
	}
	return fmt.Errorf("No Blockchain smart contract publisher available for blockchain %s", blockchain)
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

type CreateTransaction func(user *User, event *Event) error

func (s *SmartContractPublisher) CreateEthereumTransaction(user *User, event *Event) error {
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

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	smartContract := s.GetSmartContract()
	address := common.HexToAddress(smartContract.ContractAddress)
	inputs := s.GetEventParameterJSON()
	name := smartContract.EventName
	abi := "[{\"inputs\":" + inputs + ",\"name\":\"" + name + "\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	fmt.Println(abi)
	instance, err := ethereum.NewStore(address, client, abi)
	if err != nil {
		return err
	}

	parameters, err := ParseValues(event)
	if err != nil {
		return err
	}

	tx, err := instance.StoreTransactor.Contract.Transact(auth, name, parameters...)
	if err != nil {
		return err
	}

	fmt.Printf("INFO: tx sent %s\n", tx.Hash().Hex())
	return nil
}

func (s *SmartContractPublisher) CreateHyperledgerTransaction(user *User, event *Event) error {
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
	for _, eventValue := range event.EventValues {
		parameters = append(parameters, eventValue.Value)
	}
	fmt.Println(parameters)

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
