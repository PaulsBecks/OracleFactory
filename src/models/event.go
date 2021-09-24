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

func CreateEvent(oracleID uint, body []byte) *Event {
	db := utils.DBConnection()

	event := &Event{
		OracleID: oracleID,
		Success:  false,
		Body:     body,
	}
	db.Create(event)
	return event
}

func (e *Event) GetEventValues() []EventValue {
	db := utils.DBConnection()
	var events []EventValue
	db.Find(&events, "event_id = ?", e.ID)
	return events
}

func (e *Event) SetSuccess(success bool) {
	e.Success = success
	db := utils.DBConnection()

	db.Save(e)
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
