package models

import (
	"gorm.io/gorm"
	"strings"
)

type OutboundOracleTemplate struct {
	gorm.Model
	Blockchain        string
	BlockchainAddress string
	Address           string
	EventName         string
	EventParameters   []EventParameter
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

func (o *OutboundOracleTemplate) GetConnectionString() string {
	return o.BlockchainAddress
}
