package models

import (
	"fmt"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InboundOracle struct {
	gorm.Model
	Oracle                   Oracle
	OracleID                 uint
	SmartContractPublisher   SmartContractPublisher
	SmartContractPublisherID uint
	WebServiceListener       WebServiceListener
	WebServiceListenerID     uint
}

func (i *InboundOracle) GetOracle() *Oracle {
	return GetOracleByID(i.OracleID)
}

func (i *InboundOracle) GetSmartContractPublisher() *SmartContractPublisher {
	return GetSmartContractPublisherByID(i.SmartContractPublisherID)
}

func (i *InboundOracle) HandleEvent(eventJson []byte) {
	fmt.Printf("Info: Inbound Oracle with ID %d handles event %s\n", i.ID, string(eventJson))
	// create Event
	oracle := i.GetOracle()
	event := oracle.CreateEvent(eventJson)
	publisher := i.GetSmartContractPublisher()
	valuesMap, err := utils.GetMapInterfaceFromJson(eventJson)
	event.ParseEventValues(valuesMap, publisher.ListenerPublisherID)

	// check input against filter rules
	ok := oracle.CheckInput(event)
	if !ok {
		fmt.Printf("Debug: Inbound Oracle with ID %d denied event, as checks fail.\n", i.ID)
		return
	}

	// publish event
	err = publisher.Publish(i, event)
	if err != nil {
		fmt.Printf("Error: error while publishing the event: %v\n", err)
		return
	}
	event.SetSuccess(true)
}

func GetInboundOracleByID(id interface{}) (*InboundOracle, error) {
	var inboundOracle InboundOracle
	db := utils.DBConnection()
	result := db.Preload(clause.Associations).Preload("SmartContractPublisher.ListenerPublisher.EventParameters").First(&inboundOracle, id)
	if result.Error != nil {
		return &inboundOracle, result.Error
	}
	return &inboundOracle, nil
}
