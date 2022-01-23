package models

import (
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

const (
	STATUS_STARTED = "STARTED"
	STATUS_STOPPED = "STOPPED"
)

type Subscription struct {
	gorm.Model
	Name             string
	Status           string
	UserID           uint
	User             User
	Events           []Event
	ParameterFilters []ParameterFilter
}

func GetSubscriptionByID(ID uint) *Subscription {
	db := utils.DBConnection()
	var subscription Subscription
	db.Find(&subscription, ID)
	return &subscription
}

func (o *Subscription) CreateEvent(eventJson []byte) *Event {
	return CreateEvent(eventJson, o.ID)
}

func (o *Subscription) CheckInput(event *Event) bool {
	parameterFilters := o.GetParameterFilters()
	for _, parameterFilter := range parameterFilters {
		value := event.GetEventValueByParameterName(parameterFilter.ID)
		if valid := parameterFilter.Check(value); !valid {
			return false
		}
	}
	return true
}

func (o *Subscription) GetParameterFilters() []ParameterFilter {
	db := utils.DBConnection()
	var parameterFilters []ParameterFilter
	db.Find(&parameterFilters, "subscription_id = ?", o.ID)
	return parameterFilters
}

func (o *Subscription) setStatus(status string) {
	db := utils.DBConnection()

	o.Status = status
	db.Save(&o)
}

func (o *Subscription) Stop() {
	o.setStatus(STATUS_STOPPED)
}

func (o *Subscription) Start() {
	o.setStatus(STATUS_STARTED)
}

func (o *Subscription) IsStarted() bool {
	return o.Status == STATUS_STARTED
}

func (o *Subscription) IsStopped() bool {
	return o.Status == STATUS_STOPPED
}

func (o *Subscription) GetUser() *User {
	db := utils.DBConnection()

	var user *User
	db.Find(&user, o.UserID)
	return user
}

func (o *Subscription) Save() {
	db := utils.DBConnection()

	db.Save(o)
}
