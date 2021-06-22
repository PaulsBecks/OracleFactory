package models

import (
	"strings"

	"gorm.io/gorm"
)

const (
	HYPERLEDGER_BLOCKCHAIN = "Hyperledger"
	ETHEREUM_BLOCKCHAIN    = "Ethereum"
)

type OutboundOracleTemplate struct {
	gorm.Model
	Blockchain        string
	BlockchainAddress string
	Address           string
	EventName         string
	EventParameters   []EventParameter
	OutboundOracles   []OutboundOracle
}

func (o *OutboundOracleTemplate) GetEventParametersString() string {
	eventParameterStrings := []string{}
	for _, eventParameter := range o.EventParameters {
		eventParameterStrings = append(eventParameterStrings, eventParameter.String())
	}
	return strings.Join(eventParameterStrings, ", ")
}

func (o *OutboundOracleTemplate) GetEventParameterNamesString() string {
	eventParameterStrings := []string{}
	for _, eventParameter := range o.EventParameters {
		eventParameterStrings = append(eventParameterStrings, eventParameter.Name)
	}
	return strings.Join(eventParameterStrings, ", ")
}
