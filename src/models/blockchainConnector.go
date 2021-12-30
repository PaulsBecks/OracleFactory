package models

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os/exec"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/PaulsBecks/OracleFactory/src/services/ethereum"
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"github.com/cloudflare/cfssl/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Blockchain Smart Contract Creation
func retry(callback func() error, retries int) error {
	err := callback()
	if retries <= 0 || err == nil {
		return err
	}
	time.Sleep(200 * time.Millisecond)
	return retry(callback, retries-1)
}

const (
	HYPERLEDGER_BLOCKCHAIN = "Hyperledger"
	ETHEREUM_BLOCKCHAIN    = "Ethereum"
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

type BlockchainConnector interface {
	GetConnectionString() string
	GetCopyFilesString() string
	GetBlockchainName() string
	CreateTransaction(contractAddress string, methodName string, eventData map[string]interface{}) error
	GetOutboundOracle() *OutboundOracle
	StartOnChainOracle() (string, error)
	CreateOnChainTransaction(address string, topic string, eventData map[string]interface{}) error
}

func ParseValues(eventData map[string]interface{}) []interface{} {
	var parameters []interface{}
	for _, eventValue := range eventData {
		switch v := eventValue.(type) {
		case float32:
			parameters = append(parameters, big.NewInt(int64(float64(v))))
			break
		case float64:
			parameters = append(parameters, big.NewInt(int64(float64(v))))
			break
		case string:
			parameters = append(parameters, string(v))
			break
		default:
			parameters = append(parameters, eventValue)
		}
	}
	return parameters
}

type EthereumConnector struct {
	gorm.Model
	OutboundOracleID   uint
	OutboundOracle     OutboundOracle
	EthereumPrivateKey string
	EthereumAddress    string
}

func GetEthereumConnectorByID(ID interface{}) EthereumConnector {
	db := utils.DBConnection()
	var ethereumConnector EthereumConnector
	db.Preload(clause.Associations).Find(&ethereumConnector, ID)
	return ethereumConnector
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
		fmt.Printf("NewStore: %v", err.Error())
		return err
	}
	fmt.Println(abi)

	parameters := ParseValues(eventData)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Printf("GasPrice Error: %v", err.Error())
	}
	fmt.Println(abi)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	fmt.Println(abi)

	unlock := keyedMutex.Lock(e.ID)
	defer unlock()

	sendTransaction := func() error {
		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		fmt.Println(abi)
		if err != nil {
			fmt.Printf("Nonce: %v", err.Error())
			return err
		}
		cachedNonce, nonceFound := latestNonceByAddress.Load(e.ID)
		if nonceFound && cachedNonce != nil && nonce < cachedNonce.(uint64) {
			nonce = cachedNonce.(uint64)
		}
		auth.Nonce = big.NewInt(int64(nonce))

		tx, err := instance.StoreTransactor.Contract.Transact(auth, name, parameters...)
		if err != nil {
			fmt.Printf("TX Error: %v", err.Error())
			return err
		}
		latestNonceByAddress.Store(e.ID, nonce+1)
		fmt.Printf("INFO: tx sent %s\n", tx.Hash().Hex())
		return nil
	}
	return retry(sendTransaction, 10)
}

func (e EthereumConnector) CreateOnChainTransaction(address string, topic string, eventData map[string]interface{}) error {
	log.Info(address, topic, eventData)
	onChainOracleCallData := map[string]interface{}{
		"topic": topic,
	}
	methodName := ""
	var value interface{}
	for _, v := range eventData {
		value = v
		onChainOracleCallData["value"] = value
		break
	}
	switch value.(type) {
	case int:
		methodName = "publishInteger"
	case float64:
		methodName = "publishInteger"
	case string:
		methodName = "publishString"
	case bool:
		methodName = "publishBool"
	default:
		fmt.Println("Uff")
	}
	return e.CreateTransaction(address, methodName, onChainOracleCallData)
}

func (e EthereumConnector) StartOnChainOracle() (string, error) {
	cmd := exec.Command("/bin/sh",
		"-c", "cd ethereumOnChainOracle && sh deploy_prod.sh "+e.EthereumPrivateKey+" "+e.EthereumAddress,
	)
	cmd.Wait()
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	r, _ := regexp.Compile("0x([0-9a-fA-F]+)")
	address := r.FindString(string(out))
	log.Info(address)
	if address == "" {
		return "", fmt.Errorf("No address found")
	}
	return strings.Trim(address, "\n"), nil
}

func (e *EthereumConnector) GetEventParameterJSON(eventData map[string]interface{}) string {
	json := "["
	for key, value := range eventData {
		// TODO: FIX THIS
		ethereumType := MapType(value)
		json += "{\"internalType\":\"" + ethereumType + "\",\"name\":\"" + key + "\",\"type\":\"" + ethereumType + "\"},"
	}
	json = strings.TrimRight(json, ",")
	json += "]"
	return json
}

func MapType(value interface{}) string {
	switch value.(type) {
	case int:
		return "uint256"
	case float64:
		return "uint256"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		fmt.Println("Uff")
	}
	return "string"
}

func (e EthereumConnector) GetOutboundOracle() *OutboundOracle {
	return &e.OutboundOracle
}

type HyperledgerConnector struct {
	gorm.Model
	OutboundOracleID            uint
	OutboundOracle              OutboundOracle
	HyperledgerConfig           string
	HyperledgerCert             string
	HyperledgerKey              string
	HyperledgerOrganizationName string
	HyperledgerChannel          string
}

func GetHyperledgerConnectorByID(ID interface{}) HyperledgerConnector {
	db := utils.DBConnection()
	var hyperledgerConnector HyperledgerConnector
	db.Preload(clause.Associations).Find(&hyperledgerConnector, ID)
	return hyperledgerConnector
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

func (h HyperledgerConnector) StartOnChainOracle() (string, error) {
	/*cmd := exec.Command("cd", "ethereumOnChainOracle", "&&", "sh",
		"run",
		"-d",
		"--network=pub-sub-oracle-network",
		"ethereum_onchain_oracle",
		"/bin/bash",
		"-c",
	)
	out, err := cmd.Output()
	if err != nil {
		return ""
	}*/
	log.Info("Not implemented yet.")
	return "", nil
}

func (h HyperledgerConnector) CreateOnChainTransaction(address string, topic string, eventData map[string]interface{}) error {
	return nil
}
