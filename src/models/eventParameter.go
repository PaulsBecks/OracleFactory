package models

import (
	"gorm.io/gorm"
)

type EventParameter struct {
	gorm.Model
	Name                     string
	Type                     string
	OutboundOracleTemplateID uint
	OutboundOracleTemplate   OutboundOracleTemplate
	InboundOracleTemplate    InboundOracleTemplate
	InboundOracleTemplateID  uint
}

func (e *EventParameter) String() string {
	return e.Type + " " + e.Name
}
