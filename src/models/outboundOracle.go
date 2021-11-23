package models

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

type OutboundOracle struct {
	gorm.Model
	UserID          uint
	User            User
	DockerContainer string
	IsActive        bool
	Subscriptions   []Subscription
}

func (o *OutboundOracle) createManifest(blockchain, connection string) string {
	return `SET BLOCKCHAIN \"` + blockchain + `\";

SET OUTPUT FOLDER \"./output\";
SET EMISSION MODE \"streaming\";

SET CONNECTION ` + connection + `;

BLOCKS (CURRENT) (CONTINUOUS) {
	LOG ENTRIES (\"ANY\") (OracleFactory(string kind, string token, string topic, string filter, string callback, address smartContractAddress)) {
		if (kind == \"subscribe\") {
			EMIT HTTP REQUEST (\"http://oracle-factory:8080/outboundOracles/` + fmt.Sprint(o.ID) + `/subscribe\") (token, topic, filter, callback, smartContractAddress);
		}
		if (kind == \"unsubscribe\") {
			EMIT HTTP REQUEST (\"http://oracle-factory:8080/outboundOracles/` + fmt.Sprint(o.ID) + `/unsubscribe\") (token, topic, filter, callback, smartContractAddress);
		}
	}
}`
}

func echoStringToFile(content, path string) string {
	return fmt.Sprintf(" echo \"%s\" > %s; ", content, path)
}

func (o *OutboundOracle) StartOracle() error {
	connector := o.GetBlockchainConnector()
	if o.IsActive {
		return fmt.Errorf("Oracle is running already!")
	}
	manifest := o.createManifest(connector.GetBlockchainName(), connector.GetConnectionString())
	copyFilesToContainerCommand := echoStringToFile(manifest, "manifest.bloql")
	copyFilesToContainerCommand += connector.GetCopyFilesString()
	cmd := exec.Command(
		"docker",
		"run",
		"-d",
		"--network=oracle-factory-network",
		"oracle_blueprint",
		"/bin/bash",
		"-c",
		copyFilesToContainerCommand+"cat manifest.bloql; java -jar Blockchain-Logging-Framework/target/blf-cmd.jar extract manifest.bloql")
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	o.DockerContainer = strings.Trim(string(out), "\n")
	o.Save()
	return nil
}

func (o *OutboundOracle) StopOracle() error {
	if !o.IsActive {
		return fmt.Errorf("Oracle is stopped already!")
	}
	cmd := exec.Command(
		"docker",
		"stop",
		o.DockerContainer)
	out, err := cmd.Output()
	fmt.Println("INFO: " + string(out))
	return err
}

func (o *OutboundOracle) Save() {
	db := utils.DBConnection()
	db.Save(o)
}

func GetOutboundOracleByID(outboundOracleID string) *OutboundOracle {
	db := utils.DBConnection()
	var outboundOracle OutboundOracle
	db.Find(&outboundOracle, outboundOracleID)
	return &outboundOracle
}

func (o *OutboundOracle) CreateSubscription(topic, filter, callback, smartContractAddress string) *Subscription {
	db := utils.DBConnection()
	subscription := &Subscription{
		OutboundOracleID:     o.ID,
		Topic:                topic,
		Filter:               filter,
		Callback:             callback,
		SmartContractAddress: smartContractAddress,
	}
	db.Create(subscription)
	return subscription
}

func (o *OutboundOracle) DeleteSubscription(topic string) {
	db := utils.DBConnection()
	var subscription Subscription
	db.Find(&subscription, "outbound_oracle_id = ? AND topic = ?", o.ID, topic)
	db.Delete(&subscription)
}

func (o *OutboundOracle) GetBlockchainConnector() BlockchainConnector {
	return GetBlockchainConnectorByOutboundOracleID(o.ID)
}
