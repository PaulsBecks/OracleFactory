package models

import (
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

type EventValue struct {
	gorm.Model
	Value            string
	EventParameterID uint
	EventParameter   EventParameter
	EventID          uint
	Event            Event
}

func (e *EventValue) GetEventParameter() EventParameter {
	var eventParameter EventParameter
	db := utils.DBConnection()
	db.Find(&eventParameter, e.EventParameterID)
	return eventParameter
}
