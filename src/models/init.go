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
	db.AutoMigrate(
		&EthereumConnector{},
		&HyperledgerConnector{},
		&Subscription{},
		&OutboundOracle{},
		&User{},
		&Provider{},
	)
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
		Email:    "test@example.com",
		Password: utils.HashAndSalt([]byte("test")),
	}
	user.CreateEthereumConnector(
		"b28c350293dcf09cc5b5a9e5922e2f73e48983fe8d325855f04f749b1a82e0e6",
		"ws://eth-test-net:8545/",
	)
	user.CreateHyperledgerConnector(
		"Org1MSP",
		"mychannel",
		config,
		cert,
		key,
	)
	db.Create(&user)

	// create hyperledger performance test oracles
	initHyperledgerOracles(db, user)

	// create ethereum performance test oracle
	initEthereumOracles(db, user)

}
func initEthereumOracles(db *gorm.DB, user User) {
	user.CreateProvider(
		"Token Give Away",
		"/token/giveaway",
		"Continuous stream of receivers and amount of tokens.",
		true,
	)
	user.CreateProvider(
		"Token transferal",
		"Continuous stream of receivers and amount of tokens.",
		"/token/transfers",
		true,
	)
}

func initHyperledgerOracles(db *gorm.DB, user User) {
	user.CreateProvider("New Assets Endpoint", "assets/create", "This listener receives newly created assets.", true)
}
