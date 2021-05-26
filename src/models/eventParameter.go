package models

import (
	"gorm.io/gorm"
)

type EventParameter struct {
	gorm.Model
	Name                     string `json:name`
	Type                     string `json:type`
	OutboundOracleTemplateID uint
	InboundOracleTemplateID uint
}

func (e *EventParameter) String() string {
	return e.Type + " " + e.Name
}
