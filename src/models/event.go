package models

import (
	"encoding/json"
	"fmt"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Success        bool
	SubscriptionID uint
	Subscription   Subscription
	EventValues    []EventValue
	Body           []byte
}

func CreateEvent(body []byte, subscriptionID uint) *Event {
	db := utils.DBConnection()

	event := &Event{
		SubscriptionID: subscriptionID,
		Success:        false,
		Body:           body,
	}
	db.Create(event)
	return event
}

func (e *Event) GetEventValues() []EventValue {
	db := utils.DBConnection()
	var events []EventValue
	db.Order("id").Find(&events, "event_id = ?", e.ID)
	return events
}

func (e *Event) SetSuccess(success bool) {
	e.Success = success
	db := utils.DBConnection()

	db.Save(e)
}

func (e *Event) ParseBody() ([]interface{}, error) {
	var bodyData map[string]interface{}

	if err := json.Unmarshal(e.Body, &bodyData); err != nil {
		return nil, err
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

func (e *Event) GetEventValueByParameterName(eventParameterID uint) string {
	db := utils.DBConnection()
	var eventValue EventValue
	db.Find(&eventValue, "event_parameter_id = ?", eventParameterID)
	return eventValue.Value
}

func (e *Event) ParseEventValues(bodyData map[string]interface{}, providerConsumerID uint) ([]EventValue, error) {
	var eventParameters []EventParameter
	db := utils.DBConnection()
	db.Find(&eventParameters, "provider_consumer_id = ?", providerConsumerID)
	var eventValues []EventValue
	for _, eventParameter := range eventParameters {
		v := bodyData[eventParameter.Name]
		eventValue := EventValue{EventID: e.ID, Event: *e, Value: fmt.Sprintf("%v", v), EventParameterID: eventParameter.ID, EventParameter: eventParameter}
		db.Create(&eventValue)
		fmt.Println(eventValue)
		eventValues = append(eventValues, eventValue)
	}
	return eventValues, nil
}
