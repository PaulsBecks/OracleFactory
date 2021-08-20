package models

import (
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
