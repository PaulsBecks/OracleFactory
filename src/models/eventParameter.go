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
	InboundOracleTemplateID  uint
	InboundOracleTemplate    InboundOracleTemplate
}

func (e *EventParameter) String() string {
	return e.Type + " " + e.Name
}
