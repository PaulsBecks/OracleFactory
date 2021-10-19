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
		&SmartContractListener{},
		&EventValue{},
		&OutboundOracle{},
		&EventParameter{},
		&Event{},
		&SmartContractPublisher{},
		&InboundOracle{},
		&User{},
		&Filter{},
		&ParameterFilter{},
		&WebServiceListener{},
		&WebServicePublisher{},
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
	mintEthereumSmartContractPublisher := user.CreateSmartContractPublisher(
		"Ethereum",
		"mint",
		"0xe4EFfB267484Cd790143484de3Bae7fDfbE31F00",
		"Token",
		"Ethereum Token Minting",
		"Mint a specific amount of Ethereum tokens for a receiver.",
		true,
		map[string]string{
			"receiver": "address",
			"amount":   "uint256",
		},
	)
	webServiceListener := user.CreateWebServiceListener(
		"Token Give Away",
		"Continuous stream of receivers and amount of tokens.",
		true,
	)
	user.CreateInboundOracle("Mint tokens on request", mintEthereumSmartContractPublisher.ID, webServiceListener.ID)

	transferEthereumSmartContractPublisher := user.CreateSmartContractPublisher(
		"Ethereum",
		"transfer",
		"0xe4EFfB267484Cd790143484de3Bae7fDfbE31F00",
		"Token",
		"Ethereum Token Transfer",
		"Transfer a specific amount of Ethereum tokens to a receiver.",
		true,
		map[string]string{
			"receiver": "address",
			"amount":   "uint256",
		},
	)

	transferWebServiceListener := user.CreateWebServiceListener(
		"Token transferal",
		"Continuous stream of receivers and amount of tokens.",
		true,
	)
	user.CreateInboundOracle("Transfer tokens on request", transferEthereumSmartContractPublisher.ID, transferWebServiceListener.ID)

	ethereumSmartContractListener := user.CreateSmartContractListener(
		"Ethereum",
		"Transfer",
		"0xe4EFfB267484Cd790143484de3Bae7fDfbE31F00",
		"Token",
		"Ethereum Token Transfer",
		"Listen to ethereum token transfers.",
		true,
		map[string]string{
			"sender":   "address",
			"receiver": "address",
			"amount":   "uint256",
		},
	)
	webServicePublisher := user.CreateWebServicePublisher("Notify test setup", "Forward data to the test setup endpoint", testUrl, true)
	outboundOracle := user.CreateOutboundOracle("Outbound Oracle Test", ethereumSmartContractListener.ID, webServicePublisher.ID)
	outboundOracle.StartOracle()
}

func initHyperledgerOracles(db *gorm.DB, user User) {
	hyperledgerSmartContractPublisher := user.CreateSmartContractPublisher(
		"Hyperledger",
		"CreateAsset",
		"events",
		"Events",
		"Create Asset On Hyperledger",
		"This publisher creates an asset in the events smart contract.",
		true,
		map[string]string{
			"assetID":        "string",
			"color":          "string",
			"size":           "string",
			"owner":          "string",
			"appraisedValue": "int",
		},
	)
	webServiceListener := user.CreateWebServiceListener("New Assets Endpoint", "This listener receives newly created assets.", true)
	user.CreateInboundOracle("Hyperledger Inbound Test", hyperledgerSmartContractPublisher.ID, webServiceListener.ID)

	hyperledgerSmartContractListener := user.CreateSmartContractListener(
		"Hyperledger",
		"CreateAsset",
		"events",
		"Events",
		"Receive newly created assets on Hyperledger",
		"This listener waits for newly created assets.",
		true,
		map[string]string{
			"assetID":        "string",
			"color":          "string",
			"size":           "string",
			"owner":          "string",
			"appraisedValue": "int",
		},
	)
	webServicePublisher := user.CreateWebServicePublisher("New Assets Publisher", "This publisher forwards newly created assets.", testUrl, true)
	outboundOracle := user.CreateOutboundOracle("Hyperledger Oubound Test", hyperledgerSmartContractListener.ID, webServicePublisher.ID)
	outboundOracle.StartOracle()
}
