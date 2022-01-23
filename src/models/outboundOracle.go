package models

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OutboundOracle struct {
	gorm.Model
	URI                     string
	OracleID                uint
	Oracle                  Oracle
	SmartContractListenerID uint
	SmartContractListener   SmartContractListener
	WebServicePublisherID   uint
	WebServicePublisher     WebServicePublisher
	DockerContainer         string
}

func (o *OutboundOracle) GetWebServicePublisher() *WebServicePublisher {
	db := utils.DBConnection()
	var webServicePublisher *WebServicePublisher
	db.Find(&webServicePublisher, o.WebServicePublisherID)
	return webServicePublisher
}

func (o *OutboundOracle) GetSmartContractListener() *SmartContractListener {
	db := utils.DBConnection()
	var smartContractListener *SmartContractListener
	db.Find(&smartContractListener, o.SmartContractListenerID)
	return smartContractListener
}

func (o *OutboundOracle) GetConnectionString() string {
	// TODO: Describe how this can be extended to add additional blockchains
	user := o.GetOracle().GetUser()
	switch o.GetSmartContractListener().GetSmartContract().BlockchainName {
	case HYPERLEDGER_BLOCKCHAIN:
		return `{
	\"connection.yaml\",
	\"server.key\",
	\"server.crt\",
	\"` + user.HyperledgerOrganizationName + `\",
	\"` + user.HyperledgerChannel + `\"
}`
	case ETHEREUM_BLOCKCHAIN:
		return `\"` + user.EthereumAddress + `\"`
	}
	// TODO: Handle the issue if there is no blockchain with corresponding
	return ""
}

func (o *OutboundOracle) createManifest() string {
	smartContractListener := o.GetSmartContractListener()
	smartContract := smartContractListener.GetSmartContract()
	return `SET BLOCKCHAIN \"` + smartContract.BlockchainName + `\";

SET OUTPUT FOLDER \"./output\";
SET EMISSION MODE \"streaming\";

SET CONNECTION ` + o.GetConnectionString() + `;


BLOCKS (CURRENT) (CONTINUOUS) {
	LOG ENTRIES (\"` + smartContract.ContractAddress + `\") (` + smartContract.EventName + `(` + smartContractListener.GetEventParametersString() + `)) {
		EMIT HTTP REQUEST (\"` + o.oracleFactoryOutboundEventLink() + `\") (` + smartContractListener.GetEventParameterNamesString() + `);
	}
}`
}

func echoStringToFile(content, path string) string {
	return fmt.Sprintf(" echo \"%s\" > %s; ", content, path)
}

func (o *OutboundOracle) StartOracle() error {
	oracle := o.GetOracle()
	if oracle.IsStarted() {
		return fmt.Errorf("Oracle is running already!")
	}
	manifest := o.createManifest()
	copyFilesToContainerCommand := echoStringToFile(manifest, "manifest.bloql")
	user := oracle.GetUser()
	if o.GetSmartContractListener().GetSmartContract().BlockchainName == "Hyperledger" {
		copyFilesToContainerCommand += echoStringToFile(user.HyperledgerCert, "server.crt")
		copyFilesToContainerCommand += echoStringToFile(user.HyperledgerConfig, "connection.yaml")
		copyFilesToContainerCommand += echoStringToFile(user.HyperledgerKey, "server.key")
	}
	cmd := exec.Command(
		"docker",
		"run",
		"-d",
		"--network=oracle-factory-network",
		"paulsbecks/blf-outbound-oracle",
		"/bin/bash",
		"-c",
		copyFilesToContainerCommand+"cat manifest.bloql; java -jar Blockchain-Logging-Framework/target/blf-cmd.jar extract manifest.bloql")
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	o.DockerContainer = strings.Trim(string(out), "\n")
	o.Save()
	oracle.Start()
	return nil
}

func (o *OutboundOracle) StopOracle() error {
	oracle := o.GetOracle()
	if oracle.IsStopped() {
		return fmt.Errorf("Oracle is stopped already!")
	}
	cmd := exec.Command(
		"docker",
		"stop",
		o.DockerContainer)
	fmt.Println(cmd.Args)
	out, err := cmd.Output()
	fmt.Println("INFO: " + string(out))
	if err == nil {
		oracle.Stop()
	}
	return err
}

func (o *OutboundOracle) oracleFactoryOutboundEventLink() string {
	return "http://oracle-factory:8080/outboundOracles/" + fmt.Sprint(o.ID) + "/events"
}

func (o *OutboundOracle) GetOracle() *Oracle {
	db := utils.DBConnection()

	var oracle *Oracle
	db.Find(&oracle, o.OracleID)
	return oracle
}

func (o *OutboundOracle) Save() {
	db := utils.DBConnection()

	db.Save(o)
}

func GetOutboundOracleById(id interface{}) (*OutboundOracle, error) {
	db := utils.DBConnection()

	var outboundOracle *OutboundOracle
	result := db.Preload(clause.Associations).First(&outboundOracle, id)
	if result.Error != nil {
		return outboundOracle, result.Error
	}
	return outboundOracle, nil
}
