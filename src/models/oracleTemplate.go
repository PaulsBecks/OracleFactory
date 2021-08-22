package models

import (
	"log"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

type OracleTemplate struct {
	gorm.Model
	BlockchainName  string
	EventName       string
	ContractAddress string
	EventParameters []EventParameter
	UserID          uint
	User            User
	Private         bool
}

func (o *OracleTemplate) GetEventParameters() []EventParameter {
	db, err := utils.DBConnection()
	if err != nil {
		log.Fatal(err)
	}
	var eventParameters []EventParameter
	db.Find(&eventParameters, "oracle_template_id = ?", o.ID)
	return eventParameters
}
