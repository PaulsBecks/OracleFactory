package models

import (
	"fmt"
	"log"
	"strings"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

const (
	HYPERLEDGER_BLOCKCHAIN = "Hyperledger"
	ETHEREUM_BLOCKCHAIN    = "Ethereum"
)

type OutboundOracleTemplate struct {
	gorm.Model
	OracleTemplateID uint
	OracleTemplate   OracleTemplate
	OutboundOracles  []OutboundOracle
}

func (o *OutboundOracleTemplate) GetOracleTemplate() *OracleTemplate {
	db, err := utils.DBConnection()
	if err != nil {
		log.Fatal(err)
	}
	var oracleTemplate OracleTemplate
	db.Find(&oracleTemplate, o.OracleTemplateID)
	return &oracleTemplate
}

func (o *OutboundOracleTemplate) GetEventParametersString() string {
	eventParameterStrings := []string{}
	for _, eventParameter := range o.GetOracleTemplate().GetEventParameters() {
		eventParameterStrings = append(eventParameterStrings, eventParameter.String())
	}
	fmt.Print(eventParameterStrings)
	return strings.Join(eventParameterStrings, ", ")
}

func (o *OutboundOracleTemplate) GetEventParameterNamesString() string {
	eventParameterStrings := []string{}
	for _, eventParameter := range o.GetOracleTemplate().GetEventParameters() {
		eventParameterStrings = append(eventParameterStrings, eventParameter.Name)
	}
	fmt.Print(eventParameterStrings)
	return strings.Join(eventParameterStrings, ", ")
}
