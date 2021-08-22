package models

import (
	"log"

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

func (o *Oracle) CheckInput(input map[string]interface{}) bool {
	parameterFilters := o.GetParameterFilters()
	for _, parameterFilter := range parameterFilters {
		name := parameterFilter.GetEventParameter().Name
		value := input[name]
		if valid := parameterFilter.Check(value.(string)); !valid {
			return false
		}
	}
	return true
}

func (o *Oracle) GetParameterFilters() []ParameterFilter {
	db, err := utils.DBConnection()
	if err != nil {
		log.Fatal(err)
	}
	var parameterFilters []ParameterFilter
	db.Find(&parameterFilters, "oracle_id = ?", o.ID)
	return parameterFilters
}

func (o *Oracle) setStatus(status string) {
	db, err := utils.DBConnection()
	if err != nil {
		log.Fatal(err)
	}
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
	db, err := utils.DBConnection()
	if err != nil {
		log.Fatal(err)
	}
	var user *User
	db.Find(&user, o.UserID)
	return user
}

func (o *Oracle) Save() {
	db, err := utils.DBConnection()
	if err != nil {
		panic(err)
	}
	db.Save(o)
}
