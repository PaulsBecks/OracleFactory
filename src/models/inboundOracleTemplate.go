package models

import (
	"gorm.io/gorm"
)

type InboundOracleTemplate struct {
	gorm.Model
	BlockchainName  string
	ContractName    string
	ContractAddress string
	EventParameters []EventParameter
	InboundOracles  []InboundOracle
}

func (iot *InboundOracleTemplate) GetEventParameterJSON() string {
	json := "["
	for i, v := range iot.EventParameters {
		json += "{\"internalType\":\"" + v.Type + "\",\"name\":\"" + v.Name + "\",\"type\":\"" + v.Type + "\"}"
		if i < len(iot.EventParameters)-1 {
			json += ","
		}
	}
	json += "]"
	return json
}
