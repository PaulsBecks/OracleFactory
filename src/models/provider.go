package models

import (
	"fmt"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"github.com/cloudflare/cfssl/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Provider struct {
	gorm.Model
	Name        string
	Description string
	Topic       string `gorm:"index:unique"`
	Private     bool
	UserID      uint
	User        User
	Events      []ProviderEvent
}

func GetProviderByID(ID interface{}) (Provider, error) {
	db := utils.DBConnection()
	var provider Provider
	tx := db.Preload(clause.Associations).Find(&provider, ID)
	if tx.Error != nil {
		fmt.Printf(tx.Error.Error())
		return provider, fmt.Errorf("Unable to find Provider with ID %d", ID)
	}
	return provider, nil
}

func (w *Provider) HandleEvent(body []byte) {
	log.Info(fmt.Sprintf("Provider %d is handling event %s", w.ID, string(body)))
	w.CreateProviderEvent(body)
	//https://github.com/iancoleman/orderedmap
	eventData, err := utils.GetMapInterfaceFromJson(body)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	for _, oracle := range GetSubsriptionsMatchingTopic(w.Topic) {
		log.Info(fmt.Sprintf("Topic %s: found oracle %d interested with topic %s", w.Topic, oracle.ID, oracle.Topic))
		oracle.Publish(eventData)
	}

	for _, blockchainConnectionOutboundOracle := range GetOnChainOracleConnections() {
		blockchainConnectionOutboundOracle.
			GetBlockchainConnector().
			CreateOnChainTransaction(blockchainConnectionOutboundOracle.PubSubOracleAddress, w.Topic, eventData)
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
