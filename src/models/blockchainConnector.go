package models

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

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

const (
	HYPERLEDGER_BLOCKCHAIN = "Hyperledger"
	ETHEREUM_BLOCKCHAIN    = "Ethereum"
)

type BlockchainConnector interface {
	GetConnectionString() string
	GetCopyFilesString() string
	GetBlockchainName() string
	CreateTransaction(contractAddress string, methodName string, eventData map[string]interface{}) error
	GetOutboundOracle() *OutboundOracle
}

func ParseValues(eventData map[string]interface{}) []interface{} {
	var parameters []interface{}
	for _, eventValue := range eventData {
		parameters = append(parameters, eventValue)
	}
	return parameters
}

func StartConnector(connector BlockchainConnector) {
	connector.GetOutboundOracle().StartOracle(connector)
}

type EthereumConnector struct {
	gorm.Model
	OutboundOracle     OutboundOracle
	EthereumPrivateKey string
	EthereumAddress    string
}

func (e EthereumConnector) GetConnectionString() string {
	return `\"` + e.EthereumAddress + `\"`
}

func (e EthereumConnector) GetCopyFilesString() string {
	return ""
}

func (e EthereumConnector) GetBlockchainName() string {
	return ETHEREUM_BLOCKCHAIN
}

func (e EthereumConnector) CreateTransaction(contractAddress string, methodName string, eventData map[string]interface{}) error {
	// make sure every user is only creating one transaction at a time in order they arrive

	client, err := ethclient.Dial(e.EthereumAddress)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(e.EthereumPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := common.HexToAddress(contractAddress)
	inputs := e.GetEventParameterJSON(eventData)
	name := methodName
	abi := "[{\"inputs\":" + inputs + ",\"name\":\"" + name + "\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	fmt.Println(abi)
	instance, err := ethereum.NewStore(address, client, abi)
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

	unlock := keyedMutex.Lock(e.ID)
	defer unlock()

	sendTransaction := func() error {
		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			return err
		}
		cachedNonce, nonceFound := latestNonceByAddress.Load(e.ID)
		if nonceFound && cachedNonce != nil && nonce < cachedNonce.(uint64) {
			nonce = cachedNonce.(uint64)
		}
		if err != nil {
			log.Fatal(err)
		}
		auth.Nonce = big.NewInt(int64(nonce))

		tx, err := instance.StoreTransactor.Contract.Transact(auth, name, ParseValues(eventData)...)
		if err != nil {
			return err
		}
		latestNonceByAddress.Store(e.ID, nonce+1)
		fmt.Printf("INFO: tx sent %s\n", tx.Hash().Hex())
		return nil
	}
	return retry(sendTransaction, 10)
}

func (e *EthereumConnector) GetEventParameterJSON(eventData map[string]interface{}) string {
	json := "["
	for key, _ := range eventData {
		// TODO: FIX THIS
		//e.MapType(value.(type))
		json += "{\"internalType\":\"" + "string" + "\",\"name\":\"" + key + "\",\"type\":\"" + "string" + "\"},"
	}
	json = strings.TrimRight(json, ",")
	json += "]"
	return json
}

func (e EthereumConnector) GetOutboundOracle() *OutboundOracle {
	return &e.OutboundOracle
}

type HyperledgerConnector struct {
	gorm.Model
	OutboundOracle
	HyperledgerConfig           string
	HyperledgerCert             string
	HyperledgerKey              string
	HyperledgerOrganizationName string
	HyperledgerChannel          string
}

func (h HyperledgerConnector) GetConnectionString() string {
	// TODO: Describe how this can be extended to add additional blockchains
	return `{
	\"connection.yaml\",
	\"server.key\",
	\"server.crt\",
	\"` + h.HyperledgerOrganizationName + `\",
	\"` + h.HyperledgerChannel + `\"
	}`
}

func (h HyperledgerConnector) GetCopyFilesString() string {
	copyFilesToContainerCommand := echoStringToFile(h.HyperledgerCert, "server.crt")
	copyFilesToContainerCommand += echoStringToFile(h.HyperledgerConfig, "connection.yaml")
	copyFilesToContainerCommand += echoStringToFile(h.HyperledgerKey, "server.key")
	return copyFilesToContainerCommand
}

func (h HyperledgerConnector) GetBlockchainName() string {
	return HYPERLEDGER_BLOCKCHAIN
}

func (h HyperledgerConnector) CreateTransaction(contractAddress string, methodName string, eventData map[string]interface{}) error {
	//_ = os.Setenv("DISCOVERY_AS_LOCALHOST", "false")
	organizationName := h.HyperledgerOrganizationName
	cert := h.HyperledgerCert
	key := h.HyperledgerKey
	gatewayConfig := h.HyperledgerConfig
	channel := h.HyperledgerChannel

	contractName := methodName
	parameters := []string{}
	for _, eventValue := range eventData {
		parameters = append(parameters, fmt.Sprintf("%v", eventValue))
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

func (e HyperledgerConnector) GetOutboundOracle() *OutboundOracle {
	return &e.OutboundOracle
}

func GetBlockchainConnectorByOutboundOracleID(outboundOracleID interface{}) BlockchainConnector {
	db := utils.DBConnection()
	var ethereumConnector EthereumConnector
	result := db.Find(&ethereumConnector, "outbound_oracle_id = ?", outboundOracleID)
	if result.Error == nil {
		return ethereumConnector
	}
	var hyperledgerConnector HyperledgerConnector
	result = db.Find(&hyperledgerConnector, "outbound_oracle_id = ?", outboundOracleID)
	return hyperledgerConnector
}
