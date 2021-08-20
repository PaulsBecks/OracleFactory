package models

import (
	"fmt"
	"os/exec"

	"gorm.io/gorm"
)

type OutboundOracle struct {
	gorm.Model
	URI                      string
	OracleID                 uint
	Oracle                   Oracle
	OutboundOracleTemplateID uint
	OutboundOracleTemplate   OutboundOracleTemplate
}

func (o *OutboundOracle) GetConnectionString() string {
	// TODO: Describe how this can be extended to add additional blockchains
	switch o.OutboundOracleTemplate.OracleTemplate.BlockchainName {
	case HYPERLEDGER_BLOCKCHAIN:
		return `{
	\"connection.yaml\",
	\"server.key\",
	\"server.crt\",
	\"` + o.Oracle.User.HyperledgerOrganizationName + `\",
	\"` + o.Oracle.User.HyperledgerChannel + `\"
}`
	case ETHEREUM_BLOCKCHAIN:
		return `\"` + o.Oracle.User.EthereumAddress + `\"`
	}
	// TODO: Handle the issue if there is no blockchain with corresponding
	return ""
}

func (o *OutboundOracle) createManifest() string {
	return `SET BLOCKCHAIN \"` + o.OutboundOracleTemplate.OracleTemplate.BlockchainName + `\";

SET OUTPUT FOLDER \"./output\";
SET EMISSION MODE \"streaming\";

SET CONNECTION ` + o.GetConnectionString() + `;


BLOCKS (CURRENT) (CONTINUOUS) {
	LOG ENTRIES (\"` + o.OutboundOracleTemplate.OracleTemplate.ContractAddress + `\") (` + o.OutboundOracleTemplate.OracleTemplate.EventName + `(` + o.OutboundOracleTemplate.GetEventParametersString() + `)) {
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
	if o.OutboundOracleTemplate.OracleTemplate.BlockchainName == "Hyperledger" {
		copyFilesToContainerCommand += echoStringToFile(o.Oracle.User.HyperledgerCert, "server.crt")
		copyFilesToContainerCommand += echoStringToFile(o.Oracle.User.HyperledgerConfig, "connection.yaml")
		copyFilesToContainerCommand += echoStringToFile(o.Oracle.User.HyperledgerKey, "server.key")
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
