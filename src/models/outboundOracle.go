package models

import (
	"fmt"

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
		return "\"" + o.User.EthereumAddress + "\""
	}
	// TODO: Handle the issue if there is no blockchain with corresponding
	return ""
}

func (o *OutboundOracle) CreateManifest() Manifest {
	return Manifest(`SET BLOCKCHAIN \"` + o.OutboundOracleTemplate.Blockchain + `\";

SET OUTPUT FOLDER \"./output\";
SET EMISSION MODE \"streaming\";

SET CONNECTION ` + o.GetConnectionString() + `;


BLOCKS (CURRENT) (CONTINUOUS) {
	LOG ENTRIES (` + o.OutboundOracleTemplate.Address + `) (` + o.OutboundOracleTemplate.EventName + `(` + o.OutboundOracleTemplate.GetEventParametersString() + `)) {
		EMIT HTTP REQUEST (\"` + o.oracleFactoryOutboundEventLink() + `\") (` + o.OutboundOracleTemplate.GetEventParameterNamesString() + `);
	}
}`)
}

func (o *OutboundOracle) oracleFactoryOutboundEventLink() string {
	return "http://oracle-factory:8080/outboundOracles/" + fmt.Sprint(o.ID) + "/events"
}
