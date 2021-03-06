package models

import (
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

type EventParameter struct {
	gorm.Model
	Name               string
	Type               string
	Indexed            bool
	ProviderConsumerID uint
	ProviderConsumer   ProviderConsumer
}

func (e *EventParameter) String() string {
	result := e.Type + " "
	if e.Indexed {
		result += "indexed "
	}
	result += e.Name
	return result
}

func GetEventParameterByID(id interface{}) EventParameter {
	db := utils.DBConnection()

	eventParameter := &EventParameter{}
	db.First(eventParameter, id)
	return *eventParameter
}
