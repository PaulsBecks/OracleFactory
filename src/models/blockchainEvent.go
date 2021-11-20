package models

import (
	"fmt"
	"strings"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

const (
	HYPERLEDGER_BLOCKCHAIN = "Hyperledger"
	ETHEREUM_BLOCKCHAIN    = "Ethereum"
)

type BlockchainEvent struct {
	gorm.Model
	SmartContractID     uint
	SmartContract       SmartContract
	ListenerPublisherID uint
	ListenerPublisher   ListenerPublisher
	OutboundOracles     []OutboundOracle
}

func (o *BlockchainEvent) GetSmartContract() *SmartContract {
	db := utils.DBConnection()
	var smartContract SmartContract
	db.Find(&smartContract, o.SmartContractID)
	return &smartContract
}

func (o *BlockchainEvent) GetListenerPublisher() *ListenerPublisher {
	db := utils.DBConnection()
	var listenerPublisher ListenerPublisher
	db.Find(&listenerPublisher, o.ListenerPublisherID)
	return &listenerPublisher
}

func (o *BlockchainEvent) GetEventParametersString() string {
	eventParameterStrings := []string{}
	for _, eventParameter := range o.GetListenerPublisher().GetEventParameters() {
		eventParameterStrings = append(eventParameterStrings, eventParameter.String())
	}
	fmt.Print(eventParameterStrings)
	return strings.Join(eventParameterStrings, ", ")
}

func (o *BlockchainEvent) GetEventParameterNamesString() string {
	eventParameterStrings := []string{}
	for _, eventParameter := range o.GetListenerPublisher().GetEventParameters() {
		eventParameterStrings = append(eventParameterStrings, eventParameter.Name)
	}
	fmt.Print(eventParameterStrings)
	return strings.Join(eventParameterStrings, ", ")
}
