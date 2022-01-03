package models

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

var testUrl = "http://host.docker.internal:7890"

func InitDB() {
	fmt.Println("Init Database")
	db := utils.DBConnection()

	// Check if table exists - if not create it
	db.AutoMigrate(
		&EthereumConnector{},
		&HyperledgerConnector{},
		&Subscription{},
		&OutboundOracle{},
		&User{},
		&Provider{},
		&ProviderEvent{},
	)
	env := os.Getenv("ENV")
	if env == "PERFORMANCE_TEST" {
		fmt.Println("Init Performance Tests")
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
		Email:    "test@example.com",
		Password: utils.HashAndSalt([]byte("test")),
	}
	db.Create(&user)

	offChainEthereumConnector := user.CreateEthereumConnector(
		false,
		"b28c350293dcf09cc5b5a9e5922e2f73e48983fe8d325855f04f749b1a82e0e6",
		"ws://eth-test-net:8545/",
	)
	offChainEthereumConnector.OutboundOracle.StartOracle()
	fmt.Println(offChainEthereumConnector);
	offChainHyperledgerConnector := user.CreateHyperledgerConnector(
		false,
		"Org1MSP",
		"mychannel",
		config,
		cert,
		key,
	)
	offChainHyperledgerConnector.OutboundOracle.StartOracle()
	fmt.Println(offChainHyperledgerConnector);
	user.CreateProvider(
		"Test Endpoint",
		"test-topic",
		"Endpoint to test the oracles",
		true,
	)
}
