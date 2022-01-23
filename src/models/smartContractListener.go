package models

import (
	"fmt"
	"strings"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	HYPERLEDGER_BLOCKCHAIN = "Hyperledger"
	ETHEREUM_BLOCKCHAIN    = "Ethereum"
)

type SmartContractListener struct {
	gorm.Model
	SmartContractID     uint
	SmartContract       SmartContract
	ListenerPublisherID uint
	ListenerPublisher   ListenerPublisher
	OutboundOracles     []OutboundOracle
}

func GetSmartContractListenerByID(ID uint) *SmartContractListener {
	db := utils.DBConnection()
	var smartContractListener *SmartContractListener
	db.Preload(clause.Associations).Find(&smartContractListener, ID)
	return smartContractListener
}

func (o *SmartContractListener) GetSmartContract() *SmartContract {
	db := utils.DBConnection()
	var smartContract SmartContract
	db.Find(&smartContract, o.SmartContractID)
	return &smartContract
}

func (o *SmartContractListener) GetListenerPublisher() *ListenerPublisher {
	db := utils.DBConnection()
	var listenerPublisher ListenerPublisher
	db.Find(&listenerPublisher, o.ListenerPublisherID)
	return &listenerPublisher
}

func (o *SmartContractListener) GetEventParametersString() string {
	eventParameterStrings := []string{}
	for _, eventParameter := range o.GetListenerPublisher().GetEventParameters() {
		eventParameterString := eventParameter.String()
		eventParameterStrings = append(eventParameterStrings, eventParameterString)
	}
	fmt.Print(eventParameterStrings)
	return strings.Join(eventParameterStrings, ", ")
}

func (o *SmartContractListener) GetEventParameterNamesString() string {
	eventParameterStrings := []string{}
	for _, eventParameter := range o.GetListenerPublisher().GetEventParameters() {
		eventParameterStrings = append(eventParameterStrings, eventParameter.Name)
	}
	fmt.Print(eventParameterStrings)
	return strings.Join(eventParameterStrings, ", ")
}
