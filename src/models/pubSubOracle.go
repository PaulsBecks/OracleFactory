package models

import (
	"fmt"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PubSubOracle struct {
	gorm.Model
	Oracle        Oracle
	OracleID      uint
	Consumer      Consumer
	ConsumerID    uint
	Provider      Provider
	ProviderID    uint
	SubOracle     OutboundOracle
	SubOracleID   uint
	UnsubOracle   OutboundOracle
	UnsubOracleID uint
}

func (p *PubSubOracle) Subscribe() {
	p.GetOracle().Start()
}

func (p *PubSubOracle) Unsubscribe() {
	p.GetOracle().Stop()
}

func (p *PubSubOracle) GetOracle() *Oracle {
	return GetOracleByID(p.OracleID)
}

func (p *PubSubOracle) GetConsumer() *Consumer {
	return GetConsumerByID(p.ConsumerID)
}

func (p *PubSubOracle) HandleEvent(eventJson []byte) {
	fmt.Printf("Info: Pub-Sub Oracle with ID %d handles event %s\n", p.ID, string(eventJson))
	// create Event
	oracle := p.GetOracle()
	event := oracle.CreateEvent(eventJson)
	publisher := p.GetConsumer()
	valuesMap, err := utils.GetMapInterfaceFromJson(eventJson)
	event.ParseEventValues(valuesMap, publisher.ListenerPublisherID)

	// check input against filter rules
	ok := oracle.CheckInput(event)
	if !ok {
		fmt.Printf("Debug: Pub-Sub Oracle with ID %d denied event, as checks fail.\n", p.ID)
		return
	}

	// publish event
	err = publisher.Publish(p, event)
	if err != nil {
		fmt.Printf("Error: error while publishing the event: %v\n", err)
		return
	}
	event.SetSuccess(true)
}

func GetPubSubOracleByID(id interface{}) (*PubSubOracle, error) {
	var pubSubOracle PubSubOracle
	db := utils.DBConnection()
	result := db.Preload(clause.Associations).Preload("Consumer.ListenerPublisher.EventParameters").First(&pubSubOracle, id)
	if result.Error != nil {
		return &pubSubOracle, result.Error
	}
	return &pubSubOracle, nil
}
