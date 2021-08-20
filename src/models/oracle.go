package models

import (
	"log"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

type Oracle struct {
	gorm.Model
	Name             string
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
	db.Find(parameterFilters, "OracleID = ?", o.ID)
	return parameterFilters
}
