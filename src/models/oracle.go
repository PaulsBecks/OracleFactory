package models

import (
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

const (
	STATUS_STARTED = "STARTED"
	STATUS_STOPPED = "STOPPED"
)

type Oracle struct {
	gorm.Model
	Name             string
	Status           string
	UserID           uint
	User             User
	Events           []Event
	ParameterFilters []ParameterFilter
}

func GetOracleByID(ID uint) *Oracle {
	db := utils.DBConnection()
	var oracle Oracle
	db.Find(&oracle, ID)
	return &oracle
}

func (o *Oracle) CreateEvent(eventJson []byte) *Event {
	return CreateEvent(eventJson, o.ID)
}

func (o *Oracle) CheckInput(event *Event) bool {
	parameterFilters := o.GetParameterFilters()
	for _, parameterFilter := range parameterFilters {
		value := event.GetEventValueByParameterName(parameterFilter.ID)
		if valid := parameterFilter.Check(value); !valid {
			return false
		}
	}
	return true
}

func (o *Oracle) GetParameterFilters() []ParameterFilter {
	db := utils.DBConnection()
	var parameterFilters []ParameterFilter
	db.Find(&parameterFilters, "oracle_id = ?", o.ID)
	return parameterFilters
}

func (o *Oracle) setStatus(status string) {
	db := utils.DBConnection()

	o.Status = status
	db.Save(&o)
}

func (o *Oracle) Stop() {
	o.setStatus(STATUS_STOPPED)
}

func (o *Oracle) Start() {
	o.setStatus(STATUS_STARTED)
}

func (o *Oracle) IsStarted() bool {
	return o.Status == STATUS_STARTED
}

func (o *Oracle) IsStopped() bool {
	return o.Status == STATUS_STOPPED
}

func (o *Oracle) GetUser() *User {
	db := utils.DBConnection()

	var user *User
	db.Find(&user, o.UserID)
	return user
}

func (o *Oracle) Save() {
	db := utils.DBConnection()

	db.Save(o)
}