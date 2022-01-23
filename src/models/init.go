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
		&SmartContractProvider{},
		&EventValue{},
		&OutboundSubscription{},
		&EventParameter{},
		&Event{},
		&SmartContractConsumer{},
		&InboundSubscription{},
		&User{},
		&Filter{},
		&ParameterFilter{},
		&WebServiceProvider{},
		&WebServiceConsumer{},
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

	// create hyperledger performance test subscriptions
	initHyperledgerSubscriptions(db, user)

	// create ethereum performance test subscription
	initEthereumSubscriptions(db, user)

}
func initEthereumSubscriptions(db *gorm.DB, user User) {
	mintEthereumSmartContractConsumer := user.CreateSmartContractConsumer(
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
	webServiceProvider := user.CreateWebServiceProvider(
		"Token Give Away",
		"Continuous stream of receivers and amount of tokens.",
		true,
	)
	user.CreateInboundSubscription("Mint tokens on request", mintEthereumSmartContractConsumer.ID, webServiceProvider.ID)

	transferEthereumSmartContractConsumer := user.CreateSmartContractConsumer(
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

	transferWebServiceProvider := user.CreateWebServiceProvider(
		"Token transferal",
		"Continuous stream of receivers and amount of tokens.",
		true,
	)
	user.CreateInboundSubscription("Transfer tokens on request", transferEthereumSmartContractConsumer.ID, transferWebServiceProvider.ID)

	ethereumSmartContractProvider := user.CreateSmartContractProvider(
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
	webServiceConsumer := user.CreateWebServiceConsumer("Notify test setup", "Forward data to the test setup endpoint", testUrl, true)
	outboundSubscription := user.CreateOutboundSubscription("Outbound Subscription Test", ethereumSmartContractProvider.ID, webServiceConsumer.ID)
	outboundSubscription.StartSubscription()
}

func initHyperledgerSubscriptions(db *gorm.DB, user User) {
	hyperledgerSmartContractConsumer := user.CreateSmartContractConsumer(
		"Hyperledger",
		"CreateAsset",
		"events",
		"Events",
		"Create Asset On Hyperledger",
		"This consumer creates an asset in the events smart contract.",
		true,
		[]NameTypePair{
			{Name: "ID", Type: "string"},
			{Name: "Color", Type: "string"},
			{Name: "Size", Type: "string"},
			{Name: "Owner", Type: "string"},
			{Name: "AppraisedValue", Type: "int"},
		},
	)
	webServiceProvider := user.CreateWebServiceProvider("New Assets Endpoint", "This provider receives newly created assets.", true)
	user.CreateInboundSubscription("Hyperledger Inbound Test", hyperledgerSmartContractConsumer.ID, webServiceProvider.ID)

	hyperledgerSmartContractProvider := user.CreateSmartContractProvider(
		"Hyperledger",
		"CreateAsset",
		"events",
		"Events",
		"Receive newly created assets on Hyperledger",
		"This provider waits for newly created assets.",
		true,
		[]NameTypePair{
			{Name: "ID", Type: "string"},
			{Name: "Color", Type: "string"},
			{Name: "Size", Type: "string"},
			{Name: "Owner", Type: "string"},
			{Name: "AppraisedValue", Type: "int"},
		},
	)
	webServiceConsumer := user.CreateWebServiceConsumer("New Assets Consumer", "This consumer forwards newly created assets.", testUrl, true)
	outboundSubscription := user.CreateOutboundSubscription("Hyperledger Oubound Test", hyperledgerSmartContractProvider.ID, webServiceConsumer.ID)
	outboundSubscription.StartSubscription()
}
