package models

import (
	"fmt"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

type Provider struct {
	gorm.Model
	Name        string
	Description string
	Topic       string
	Private     bool
	UserID      uint
	User        User
}

func GetProviderByID(ID interface{}) (Provider, error) {
	db := utils.DBConnection()
	var provider Provider
	tx := db.Preload("ListenerPublisher").Preload("PubSubOracles.Oracle").Preload("PubSubOracles.Consumer.SmartContract").Preload("PubSubOracles.Consumer.ListenerPublisher").Preload("PubSubOracles.Provider.ListenerPublisher").Preload("PubSubOracles.Consumer.SmartContract").Find(&provider, ID)
	if tx.Error != nil {
		fmt.Printf(tx.Error.Error())
		return provider, fmt.Errorf("Unable to find Provider with ID %d", ID)
	}
	return provider, nil
}

func (w *Provider) HandleEvent(body []byte) {
	w.CreateProviderEvent(body)
	for _, oracle := range GetSubsriptionsMatchingTopic(w.Topic) {
		oracle.Publish(body)
	}
}

type ProviderEvent struct {
	gorm.Model
	ProviderID uint
	Provider   Provider
	Body       []byte
}

func (w *Provider) CreateProviderEvent(body []byte) *ProviderEvent {
	providerEvent := &ProviderEvent{ProviderID: w.ID, Body: body}
	db := utils.DBConnection()
	db.Create(providerEvent)
	return providerEvent
}
