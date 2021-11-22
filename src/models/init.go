package models

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

var testUrl = "http://host.docker.internal:7890"

func InitDB() {
	db := utils.DBConnection()

	// Check if table exists - if not create it
	db.AutoMigrate(&EventParameter{},
		&PubSubOracle{},
		&BlockchainEvent{},
		&EventValue{},
		&EventParameter{},
		&Event{},
		&Consumer{},
		&OutboundOracle{},
		&User{},
		&Filter{},
		&ParameterFilter{},
		&Provider{},
	)
	InitFilter(db)
	env := os.Getenv("ENV")
	if env == "PERFORMANCE_TEST" {
		initPerformanceTestSetup(db)
	}
}

func fromFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Error: Unable to read data from path %s, %v", path, err)
	}
	return string(data)
}

func initPerformanceTestSetup(db *gorm.DB) {
	// create user
	config := strings.Replace(fromFile("connection-org1.yaml"), "localhost", "peer0.org1.example.com", -1)
	cert := fromFile("hyperledger_cert")
	key := fromFile("hyperledger_key")
	user := User{
		Email:                       "test@example.com",
		Password:                    utils.HashAndSalt([]byte("test")),
		EthereumPrivateKey:          "b28c350293dcf09cc5b5a9e5922e2f73e48983fe8d325855f04f749b1a82e0e6",
		EthereumAddress:             "ws://eth-test-net:8545/",
		HyperledgerOrganizationName: "Org1MSP",
		HyperledgerChannel:          "mychannel",
		HyperledgerConfig:           config,
		HyperledgerCert:             cert,
		HyperledgerKey:              key,
	}
	db.Create(&user)

	// create hyperledger performance test oracles
	initHyperledgerOracles(db, user)

	// create ethereum performance test oracle
	initEthereumOracles(db, user)

}
func initEthereumOracles(db *gorm.DB, user User) {
	mintEthereumConsumer := user.CreateConsumer(
		"Ethereum",
		"mint",
		"0xe4EFfB267484Cd790143484de3Bae7fDfbE31F00",
		"Token",
		"Ethereum Token Minting",
		"Mint a specific amount of Ethereum tokens for a receiver.",
		true,
		[]NameTypePair{
			{Name: "receiver", Type: "address"},
			{Name: "amount", Type: "uint256"},
		},
	)
	provider := user.CreateProvider(
		"Token Give Away",
		"Continuous stream of receivers and amount of tokens.",
		true,
	)
	user.CreatePubSubOracle("Mint tokens on request", mintEthereumConsumer.ID, provider.ID, 0, 0)

	transferEthereumConsumer := user.CreateConsumer(
		"Ethereum",
		"transfer",
		"0xe4EFfB267484Cd790143484de3Bae7fDfbE31F00",
		"Token",
		"Ethereum Token Transfer",
		"Transfer a specific amount of Ethereum tokens to a receiver.",
		true,
		[]NameTypePair{
			{Name: "receiver", Type: "address"},
			{Name: "amount", Type: "uint256"},
		},
	)

	transferProvider := user.CreateProvider(
		"Token transferal",
		"Continuous stream of receivers and amount of tokens.",
		true,
	)
	user.CreatePubSubOracle("Transfer tokens on request", transferEthereumConsumer.ID, transferProvider.ID, 0, 0)

	ethereumBlockchainEvent := user.CreateBlockchainEvent(
		"Ethereum",
		"Transfer",
		"0xe4EFfB267484Cd790143484de3Bae7fDfbE31F00",
		"Token",
		"Ethereum Token Transfer",
		"Listen to ethereum token transfers.",
		true,
		[]NameTypePair{
			{Name: "sender", Type: "address"},
			{Name: "receiver", Type: "address"},
			{Name: "amount", Type: "uint256"},
		},
	)
	outboundOracle := user.CreateOutboundOracle("Outbound Oracle Test", ethereumBlockchainEvent.ID, true)
	outboundOracle.StartOracle()
}

func initHyperledgerOracles(db *gorm.DB, user User) {
	hyperledgerConsumer := user.CreateConsumer(
		"Hyperledger",
		"CreateAsset",
		"events",
		"Events",
		"Create Asset On Hyperledger",
		"This publisher creates an asset in the events smart contract.",
		true,
		[]NameTypePair{
			{Name: "ID", Type: "string"},
			{Name: "Color", Type: "string"},
			{Name: "Size", Type: "string"},
			{Name: "Owner", Type: "string"},
			{Name: "AppraisedValue", Type: "int"},
		},
	)
	provider := user.CreateProvider("New Assets Endpoint", "This listener receives newly created assets.", true)
	user.CreatePubSubOracle("Hyperledger Inbound Test", hyperledgerConsumer.ID, provider.ID, 0, 0)

	hyperledgerBlockchainEvent := user.CreateBlockchainEvent(
		"Hyperledger",
		"CreateAsset",
		"events",
		"Events",
		"Receive newly created assets on Hyperledger",
		"This listener waits for newly created assets.",
		true,
		[]NameTypePair{
			{Name: "ID", Type: "string"},
			{Name: "Color", Type: "string"},
			{Name: "Size", Type: "string"},
			{Name: "Owner", Type: "string"},
			{Name: "AppraisedValue", Type: "int"},
		},
	)
	outboundOracle := user.CreateOutboundOracle("Hyperledger Outbound Test", hyperledgerBlockchainEvent.ID, true)
	outboundOracle.StartOracle()
}
