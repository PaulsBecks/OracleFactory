package models

import (
	"gorm.io/gorm"
)

type InboundOracleTemplate struct {
	gorm.Model
	OracleTemplateID uint
	OracleTemplate   OracleTemplate
	InboundOracles   []InboundOracle
}

func (iot *InboundOracleTemplate) GetEventParameterJSON() string {
	json := "["
	for i, v := range iot.OracleTemplate.EventParameters {
		json += "{\"internalType\":\"" + v.Type + "\",\"name\":\"" + v.Name + "\",\"type\":\"" + v.Type + "\"}"
		if i < len(iot.OracleTemplate.EventParameters)-1 {
			json += ","
		}
	}
	json += "]"
	return json
}
