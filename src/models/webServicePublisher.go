package models

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

type WebServicePublisher struct {
	gorm.Model
	Url                 string
	ListenerPublisherID uint
	ListenerPublisher   ListenerPublisher
	OutboundOracles     []OutboundOracle
}

/*
func (w *WebServicePublisher) GetOutboundOracle() (outboundOracle *OutboundOracle) {
	db := utils.DBConnection()
	db.Find(outboundOracle, w.OutboundOracleID)
	return outboundOracle
}*/

func (w *WebServicePublisher) Publish(event Event) {
	_ = event.GetEventValues()
	fmt.Println("INFO: post data to: " + w.Url)
	http.Post(w.Url, "application/json", bytes.NewBuffer(event.Body))

}

func GetWebServicePublisherByID(ID interface{}) (WebServicePublisher, error) {
	fmt.Printf("Debug: %s", ID.(string))
	db := utils.DBConnection()
	var webServicePublisher WebServicePublisher
	tx := db.Preload("ListenerPublisher").Preload("OutboundOracles.Oracle").Preload("OutboundOracles.SmartContractListener.SmartContract").Preload("OutboundOracles.SmartContractListener.ListenerPublisher").Preload("OutboundOracles.WebServicePublisher.ListenerPublisher").Find(&webServicePublisher, ID)
	if tx.Error != nil {
		fmt.Printf(tx.Error.Error())
		return webServicePublisher, fmt.Errorf("Unable to find WebServicePublisher with ID %d", ID)
	}
	return webServicePublisher, nil
}
