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

type SmartContractProvider struct {
	gorm.Model
	SmartContractID       uint
	SmartContract         SmartContract
	ProviderConsumerID    uint
	ProviderConsumer      ProviderConsumer
	OutboundSubscriptions []OutboundSubscription
}

func GetSmartContractProviderByID(ID uint) *SmartContractProvider {
	db := utils.DBConnection()
	var smartContractProvider *SmartContractProvider
	db.Preload(clause.Associations).Find(&smartContractProvider, ID)
	return smartContractProvider
}

func (o *SmartContractProvider) GetSmartContract() *SmartContract {
	db := utils.DBConnection()
	var smartContract SmartContract
	db.Find(&smartContract, o.SmartContractID)
	return &smartContract
}

func (o *SmartContractProvider) GetProviderConsumer() *ProviderConsumer {
	db := utils.DBConnection()
	var providerConsumer ProviderConsumer
	db.Find(&providerConsumer, o.ProviderConsumerID)
	return &providerConsumer
}

func (o *SmartContractProvider) GetEventParametersString() string {
	eventParameterStrings := []string{}
	for _, eventParameter := range o.GetProviderConsumer().GetEventParameters() {
		eventParameterString := eventParameter.String()
		eventParameterStrings = append(eventParameterStrings, eventParameterString)
	}
	fmt.Print(eventParameterStrings)
	return strings.Join(eventParameterStrings, ", ")
}

func (o *SmartContractProvider) GetEventParameterNamesString() string {
	eventParameterStrings := []string{}
	for _, eventParameter := range o.GetProviderConsumer().GetEventParameters() {
		eventParameterStrings = append(eventParameterStrings, eventParameter.Name)
	}
	fmt.Print(eventParameterStrings)
	return strings.Join(eventParameterStrings, ", ")
}
