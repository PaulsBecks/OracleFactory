package models

import (
	"fmt"

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

func ParseEventValues(bodyData map[string]interface{}, inboundEvent Event, oracleTemplateID uint) ([]EventValue, error) {
	var eventParameters []EventParameter
	db, err := utils.DBConnection()
	if err != nil {
		return nil, err
	}
	db.Find(&eventParameters, "oracle_template_id = ?", oracleTemplateID)
	var eventValues []EventValue
	for _, eventParameter := range eventParameters {
		v := bodyData[eventParameter.Name]
		eventValue := EventValue{EventID: inboundEvent.ID, Event: inboundEvent, Value: fmt.Sprintf("%v", v), EventParameterID: eventParameter.ID, EventParameter: eventParameter}
		db.Create(&eventValue)
		eventValues = append(eventValues, eventValue)
	}
	fmt.Print(eventValues)
	return eventValues, nil
}
