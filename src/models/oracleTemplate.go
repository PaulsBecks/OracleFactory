package models

import (
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

type OracleTemplate struct {
	gorm.Model
	BlockchainName         string
	EventName              string
	ContractAddress        string
	ContractAddressSynonym string
	EventParameters        []EventParameter
	UserID                 uint
	User                   User
	Private                bool
}

func (o *OracleTemplate) GetEventParameters() []EventParameter {
	db := utils.DBConnection()

	var eventParameters []EventParameter
	db.Find(&eventParameters, "oracle_template_id = ?", o.ID)
	return eventParameters
}
