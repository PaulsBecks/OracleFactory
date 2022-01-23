package models

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OutboundSubscription struct {
	gorm.Model
	URI                     string
	SubscriptionID          uint
	Subscription            Subscription
	SmartContractProviderID uint
	SmartContractProvider   SmartContractProvider
	WebServiceConsumerID    uint
	WebServiceConsumer      WebServiceConsumer
	DockerContainer         string
}

func (o *OutboundSubscription) GetWebServiceConsumer() *WebServiceConsumer {
	db := utils.DBConnection()
	var webServiceConsumer *WebServiceConsumer
	db.Find(&webServiceConsumer, o.WebServiceConsumerID)
	return webServiceConsumer
}

func (o *OutboundSubscription) GetSmartContractProvider() *SmartContractProvider {
	db := utils.DBConnection()
	var smartContractProvider *SmartContractProvider
	db.Find(&smartContractProvider, o.SmartContractProviderID)
	return smartContractProvider
}

func (o *OutboundSubscription) GetConnectionString() string {
	// TODO: Describe how this can be extended to add additional blockchains
	user := o.GetSubscription().GetUser()
	switch o.GetSmartContractProvider().GetSmartContract().BlockchainName {
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

func (o *OutboundSubscription) createManifest() string {
	smartContractProvider := o.GetSmartContractProvider()
	smartContract := smartContractProvider.GetSmartContract()
	return `SET BLOCKCHAIN \"` + smartContract.BlockchainName + `\";

SET OUTPUT FOLDER \"./output\";
SET EMISSION MODE \"streaming\";

SET CONNECTION ` + o.GetConnectionString() + `;


BLOCKS (CURRENT) (CONTINUOUS) {
	LOG ENTRIES (\"` + smartContract.ContractAddress + `\") (` + smartContract.EventName + `(` + smartContractProvider.GetEventParametersString() + `)) {
		EMIT HTTP REQUEST (\"` + o.oracleFactoryOutboundEventLink() + `\") (` + smartContractProvider.GetEventParameterNamesString() + `);
	}
}`
}

func echoStringToFile(content, path string) string {
	return fmt.Sprintf(" echo \"%s\" > %s; ", content, path)
}

func (o *OutboundSubscription) StartSubscription() error {
	subscription := o.GetSubscription()
	if subscription.IsStarted() {
		return fmt.Errorf("Subscription is running already!")
	}
	manifest := o.createManifest()
	copyFilesToContainerCommand := echoStringToFile(manifest, "manifest.bloql")
	user := subscription.GetUser()
	if o.GetSmartContractProvider().GetSmartContract().BlockchainName == "Hyperledger" {
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
	subscription.Start()
	return nil
}

func (o *OutboundSubscription) StopSubscription() error {
	subscription := o.GetSubscription()
	if subscription.IsStopped() {
		return fmt.Errorf("Subscription is stopped already!")
	}
	cmd := exec.Command(
		"docker",
		"stop",
		o.DockerContainer)
	fmt.Println(cmd.Args)
	out, err := cmd.Output()
	fmt.Println("INFO: " + string(out))
	if err == nil {
		subscription.Stop()
	}
	return err
}

func (o *OutboundSubscription) oracleFactoryOutboundEventLink() string {
	return "http://oracle-factory:8080/outboundSubscriptions/" + fmt.Sprint(o.ID) + "/events"
}

func (o *OutboundSubscription) GetSubscription() *Subscription {
	db := utils.DBConnection()

	var subscription *Subscription
	db.Find(&subscription, o.SubscriptionID)
	return subscription
}

func (o *OutboundSubscription) Save() {
	db := utils.DBConnection()

	db.Save(o)
}

func GetOutboundSubscriptionById(id interface{}) (*OutboundSubscription, error) {
	db := utils.DBConnection()

	var outboundSubscription *OutboundSubscription
	result := db.Preload(clause.Associations).First(&outboundSubscription, id)
	if result.Error != nil {
		return outboundSubscription, result.Error
	}
	return outboundSubscription, nil
}
