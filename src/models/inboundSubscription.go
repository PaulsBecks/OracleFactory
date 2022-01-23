package models

import (
	"fmt"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InboundSubscription struct {
	gorm.Model
	Subscription            Subscription
	SubscriptionID          uint
	SmartContractConsumer   SmartContractConsumer
	SmartContractConsumerID uint
	WebServiceProvider      WebServiceProvider
	WebServiceProviderID    uint
}

func (i *InboundSubscription) GetSubscription() *Subscription {
	return GetSubscriptionByID(i.SubscriptionID)
}

func (i *InboundSubscription) GetSmartContractConsumer() *SmartContractConsumer {
	return GetSmartContractConsumerByID(i.SmartContractConsumerID)
}

func (i *InboundSubscription) HandleEvent(eventJson []byte) {
	fmt.Printf("Info: Inbound Subscription with ID %d handles event %s\n", i.ID, string(eventJson))
	// create Event
	subscription := i.GetSubscription()
	event := subscription.CreateEvent(eventJson)
	consumer := i.GetSmartContractConsumer()
	valuesMap, err := utils.GetMapInterfaceFromJson(eventJson)
	event.ParseEventValues(valuesMap, consumer.ProviderConsumerID)

	// check input against filter rules
	ok := subscription.CheckInput(event)
	if !ok {
		fmt.Printf("Debug: Inbound Subscription with ID %d denied event, as checks fail.\n", i.ID)
		return
	}

	// publish event
	err = consumer.Publish(i, event)
	if err != nil {
		fmt.Printf("Error: error while publishing the event: %v\n", err)
		return
	}
	event.SetSuccess(true)
}

func GetInboundSubscriptionByID(id interface{}) (*InboundSubscription, error) {
	var inboundSubscription InboundSubscription
	db := utils.DBConnection()
	result := db.Preload(clause.Associations).Preload("SmartContractConsumer.ProviderConsumer.EventParameters").First(&inboundSubscription, id)
	if result.Error != nil {
		return &inboundSubscription, result.Error
	}
	return &inboundSubscription, nil
}
