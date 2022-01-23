package models

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

type WebServiceConsumer struct {
	gorm.Model
	Url                   string
	ProviderConsumerID    uint
	ProviderConsumer      ProviderConsumer
	OutboundSubscriptions []OutboundSubscription
}

/*
func (w *WebServiceConsumer) GetOutboundSubscription() (outboundSubscription *OutboundSubscription) {
	db := utils.DBConnection()
	db.Find(outboundSubscription, w.OutboundSubscriptionID)
	return outboundSubscription
}*/

func (w *WebServiceConsumer) Publish(event Event) {
	_ = event.GetEventValues()
	fmt.Println("INFO: post data to: " + w.Url)
	http.Post(w.Url, "application/json", bytes.NewBuffer(event.Body))

}

func GetWebServiceConsumerByID(ID interface{}) (WebServiceConsumer, error) {
	fmt.Printf("Debug: %s", ID.(string))
	db := utils.DBConnection()
	var webServiceConsumer WebServiceConsumer
	tx := db.Preload("ProviderConsumer").Preload("OutboundSubscriptions.Subscription").Preload("OutboundSubscriptions.SmartContractProvider.SmartContract").Preload("OutboundSubscriptions.SmartContractProvider.ProviderConsumer").Preload("OutboundSubscriptions.WebServiceConsumer.ProviderConsumer").Find(&webServiceConsumer, ID)
	if tx.Error != nil {
		fmt.Printf(tx.Error.Error())
		return webServiceConsumer, fmt.Errorf("Unable to find WebServiceConsumer with ID %d", ID)
	}
	return webServiceConsumer, nil
}
