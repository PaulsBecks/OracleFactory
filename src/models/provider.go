package models

import (
	"fmt"
	"sync"

	"github.com/PaulsBecks/OracleFactory/src/lock"
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
	tx := db.Preload(clause.Associations).First(&provider, ID)
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
	var wg sync.WaitGroup
	func(body []byte, wg *sync.WaitGroup) {
		defer lock.PipeLock.Lock()
		eventData, err := utils.GetMapInterfaceFromJson(body)
		if err != nil {
			log.Fatalf(err.Error())
			return
		}
		// only one event at a time to keep the order this is important yes?
		for _, oracle := range GetSubsriptionsMatchingTopic(w.Topic) {
			log.Info(fmt.Sprintf("Topic %s: found oracle %d interested with topic %s", w.Topic, oracle.ID, oracle.Topic))
			wg.Add(1)
			go func(oracle Subscription) {
				defer wg.Done()
				oracle.Publish(eventData)
			}(oracle)
		}

		for _, blockchainConnectionOutboundOracle := range GetOnChainOracleConnections() {
			log.Info(fmt.Sprintf("Oracle %d interested in topic %s", blockchainConnectionOutboundOracle.ID, w.Topic))
			wg.Add(1)
			go func(blockchainConnectionOutboundOracle OutboundOracle) {
				defer wg.Done()
				blockchainConnectionOutboundOracle.
					GetBlockchainConnector().
					CreateOnChainTransaction(blockchainConnectionOutboundOracle.PubSubOracleAddress, w.Topic, eventData)

			}(blockchainConnectionOutboundOracle)
		}
		lock.PipeLock.Unlock()
	}(body, &wg)
	wg.Wait()
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
