package models

import (
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

type SmartContract struct {
	gorm.Model
	BlockchainName         string
	EventName              string
	ContractAddress        string
	ContractAddressSynonym string
}

func GetSmartContractByID(ID uint) *SmartContract {
	db := utils.DBConnection()
	var smartContract *SmartContract
	db.Find(&smartContract, ID)
	return smartContract
}
