package models

import (
	"encoding/json"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Success     bool
	Oracle      Oracle
	OracleID    uint
	EventValues []EventValue
	Body        []byte
}

func (e *Event) ParseBody() ([]interface{}, error) {
	var bodyData map[string]interface{}

	if e := json.Unmarshal(e.Body, &bodyData); e != nil {
		return nil, e
	}

	var parameters []interface{}
	for _, eventValue := range e.EventValues {
		v := bodyData[eventValue.EventParameter.Name]
		parameter, err := utils.TransformParameterType(v, eventValue.EventParameter.Type)
		if err != nil {
			return nil, err
		}
		parameters = append(parameters, parameter)
	}
	return parameters, nil
}
