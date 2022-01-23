package models

import (
	"fmt"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

type WebServiceProvider struct {
	gorm.Model
	ProviderConsumerID   uint
	ProviderConsumer     ProviderConsumer
	InboundSubscriptions []InboundSubscription
}

func GetWebServiceProviderByID(ID interface{}) (WebServiceProvider, error) {
	db := utils.DBConnection()
	var webServiceProvider WebServiceProvider
	tx := db.Preload("ProviderConsumer").Preload("InboundSubscriptions.Subscription").Preload("InboundSubscriptions.SmartContractConsumer.SmartContract").Preload("InboundSubscriptions.SmartContractConsumer.ProviderConsumer").Preload("InboundSubscriptions.WebServiceProvider.ProviderConsumer").Preload("InboundSubscriptions.SmartContractConsumer.SmartContract").Find(&webServiceProvider, ID)
	if tx.Error != nil {
		fmt.Printf(tx.Error.Error())
		return webServiceProvider, fmt.Errorf("Unable to find WebServiceProvider with ID %d", ID)
	}
	return webServiceProvider, nil
}

func (w *WebServiceProvider) HandleEvent(body []byte) {
	for _, subscription := range w.InboundSubscriptions {
		subscription.HandleEvent(body)
	}
}
