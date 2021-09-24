package models

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() {
	//db, err := sql.Open("sqlite", c.filePath)
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Check if table exists - if not create it
	db.AutoMigrate(&EventParameter{},
		&OutboundOracleTemplate{},
		&EventValue{},
		&OutboundOracle{},
		&EventParameter{},
		&Event{},
		&InboundOracleTemplate{},
		&InboundOracle{},
		&User{},
		&Filter{},
		&ParameterFilter{},
	)
	InitFilter(db)
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

	log.Println(user)

	db.Create(&user)
	log.Println(user)

	oracleTemplate := OracleTemplate{
		BlockchainName:         "Hyperledger",
		EventName:              "CreateAsset",
		ContractAddress:        "events",
		ContractAddressSynonym: "Events",
	}
	db.Create(&oracleTemplate)
	inboundOracleTemplate := InboundOracleTemplate{OracleTemplate: oracleTemplate}
	db.Create(&inboundOracleTemplate)
	oracle := Oracle{
		Name:   "Hyperledger Test",
		UserID: user.ID,
	}
	db.Create(&oracle)
	inboundOracle := InboundOracle{
		Oracle:                  oracle,
		InboundOracleTemplateID: inboundOracleTemplate.ID,
	}
	db.Create(&inboundOracle)
	eventParameter := EventParameter{
		Name:             "assetID",
		Type:             "string",
		OracleTemplateID: oracle.ID,
	}
	db.Create(&eventParameter)

	eventParameter = EventParameter{
		Name:             "color",
		Type:             "string",
		OracleTemplateID: oracle.ID,
	}
	db.Create(&eventParameter)

	eventParameter = EventParameter{
		Name:             "size",
		Type:             "string",
		OracleTemplateID: oracle.ID,
	}
	db.Create(&eventParameter)

	eventParameter = EventParameter{
		Name:             "owner",
		Type:             "string",
		OracleTemplateID: oracle.ID,
	}
	db.Create(&eventParameter)

	eventParameter = EventParameter{
		Name:             "appraisedValue",
		Type:             "int",
		OracleTemplateID: oracle.ID,
	}
	db.Create(&eventParameter)
	/*outboundOracleTemplate := OutboundOracleTemplate{
		OracleTemplate: oracleTemplate,
	}
	db.Create(&outboundOracleTemplate)
	eventParameterOut := EventParameter{
		Name:             "owner",
		Type:             "string",
		OracleTemplateID: oracle.ID,
	}
	db.Create(&eventParameterOut)*/
}

func fromFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Error: Unable to read data from path %s, %v", path, err)
	}
	return string(data)
}
