package models

import (
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProviderConsumer struct {
	gorm.Model
	Name            string
	Description     string
	EventParameters []EventParameter
	UserID          uint
	User            User
	Private         bool
}

func GetProviderConsumerByID(id uint) (providerConsumer ProviderConsumer) {
	db := utils.DBConnection()
	db.Preload(clause.Associations).Find(&providerConsumer, id)
	return providerConsumer
}

func (l *ProviderConsumer) GetEventParameters() []EventParameter {
	db := utils.DBConnection()

	var eventParameters []EventParameter
	db.Find(&eventParameters, "provider_consumer_id = ?", l.ID)
	return eventParameters
}
