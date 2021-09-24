package models

import (
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

type EventParameter struct {
	gorm.Model
	Name             string
	Type             string
	OracleTemplateID uint
	OracleTemplate   OracleTemplate
}

func (e *EventParameter) String() string {
	return e.Type + " " + e.Name
}

func GetEventParameterByID(id interface{}) EventParameter {
	db := utils.DBConnection()

	eventParameter := &EventParameter{}
	db.First(eventParameter, id)
	return *eventParameter
}
