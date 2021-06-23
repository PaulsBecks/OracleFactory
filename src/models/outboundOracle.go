package models

import (
	"fmt"
	"os/exec"

	"gorm.io/gorm"
)

type OutboundOracle struct {
	gorm.Model
	URI                      string
	Name                     string
	UserID                   uint
	User                     User
	OutboundOracleTemplateID uint
	OutboundOracleTemplate   OutboundOracleTemplate
	OutboundEvents           []OutboundEvent
}

func (o *OutboundOracle) GetConnectionString() string {
	// TODO: Describe how this can be extended to add additional blockchains
	switch o.OutboundOracleTemplate.Blockchain {
	case HYPERLEDGER_BLOCKCHAIN:
		return `{
	\"connection.yaml\",
	\"server.key\",
	\"server.crt\",
	\"` + o.User.HyperledgerOrganizationName + `\",
	\"` + o.User.HyperledgerChannel + `\"
}`
	case ETHEREUM_BLOCKCHAIN:
		return `\"` + o.User.EthereumAddress + `\"`
	}
	// TODO: Handle the issue if there is no blockchain with corresponding
	return ""
}

func (o *OutboundOracle) createManifest() string {
	return `SET BLOCKCHAIN \"` + o.OutboundOracleTemplate.Blockchain + `\";

SET OUTPUT FOLDER \"./output\";
SET EMISSION MODE \"streaming\";

SET CONNECTION ` + o.GetConnectionString() + `;


BLOCKS (CURRENT) (CONTINUOUS) {
	LOG ENTRIES (\"` + o.OutboundOracleTemplate.Address + `\") (` + o.OutboundOracleTemplate.EventName + `(` + o.OutboundOracleTemplate.GetEventParametersString() + `)) {
		EMIT HTTP REQUEST (\"` + o.oracleFactoryOutboundEventLink() + `\") (` + o.OutboundOracleTemplate.GetEventParameterNamesString() + `);
	}
}`
}

func echoStringToFile(content, path string) string {
	return fmt.Sprintf(" echo \"%s\" > %s; ", content, path)
}

func (o *OutboundOracle) CreateOracle() error {
	manifest := o.createManifest()
	copyFilesToContainerCommand := echoStringToFile(manifest, "manifest.bloql")
	if o.OutboundOracleTemplate.Blockchain == "Hyperledger" {
		copyFilesToContainerCommand += echoStringToFile(o.User.HyperledgerCert, "server.crt")
		copyFilesToContainerCommand += echoStringToFile(o.User.HyperledgerConfig, "connection.yaml")
		copyFilesToContainerCommand += echoStringToFile(o.User.HyperledgerKey, "server.key")
	}
	cmd := exec.Command(
		"docker",
		"run",
		"-d",
		"--network=oracle-factory-network",
		"oracle_blueprint",
		"/bin/bash",
		"-c",
		copyFilesToContainerCommand+"cat manifest.bloql; java -jar Blockchain-Logging-Framework/target/blf-cmd.jar extract manifest.bloql")
	return cmd.Run()
}

func (o *OutboundOracle) oracleFactoryOutboundEventLink() string {
	return "http://oracle-factory:8080/outboundOracles/" + fmt.Sprint(o.ID) + "/events"
}
