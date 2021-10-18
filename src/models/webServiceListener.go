package models

import (
	"fmt"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

type WebServiceListener struct {
	gorm.Model
	ListenerPublisherID uint
	ListenerPublisher   ListenerPublisher
	InboundOracles      []InboundOracle
}

func GetWebServiceListenerByID(ID interface{}) (WebServiceListener, error) {
	db := utils.DBConnection()
	var webServiceListener WebServiceListener
	tx := db.Preload("ListenerPublisher").Preload("InboundOracles.Oracle").Preload("InboundOracles.SmartContractPublisher.SmartContract").Preload("InboundOracles.SmartContractPublisher.ListenerPublisher").Preload("InboundOracles.WebServiceListener.ListenerPublisher").Preload("InboundOracles.SmartContractPublisher.SmartContract").Find(&webServiceListener, ID)
	if tx.Error != nil {
		fmt.Printf(tx.Error.Error())
		return webServiceListener, fmt.Errorf("Unable to find WebServiceListener with ID %d", ID)
	}
	return webServiceListener, nil
}

func (w *WebServiceListener) HandleEvent(body []byte) {
	for _, oracle := range w.InboundOracles {
		oracle.HandleEvent(body)
	}
}
