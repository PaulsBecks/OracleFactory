package models

import (
	"fmt"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

type Provider struct {
	gorm.Model
	ListenerPublisherID uint
	ListenerPublisher   ListenerPublisher
	PubSubOracles       []PubSubOracle
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
	for _, oracle := range w.PubSubOracles {
		oracle.HandleEvent(body)
	}
}
