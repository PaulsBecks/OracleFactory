package models

import (
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ListenerPublisher struct {
	gorm.Model
	Name            string
	Description     string
	EventParameters []EventParameter
	UserID          uint
	User            User
	Private         bool
}

func GetListenerPublisherByID(id uint) (listenerPublisher ListenerPublisher) {
	db := utils.DBConnection()
	db.Preload(clause.Associations).Find(&listenerPublisher, id)
	return listenerPublisher
}

func (l *ListenerPublisher) GetEventParameters() []EventParameter {
	db := utils.DBConnection()

	var eventParameters []EventParameter
	db.Find(&eventParameters, "listener_publisher_id = ?", l.ID)
	return eventParameters
}
