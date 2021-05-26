package models

import (
	"gorm.io/gorm"
)

type InboundOracleTemplate struct {
	gorm.Model
	BlockchainName string
	BlockchainAddress string
	ContractName string
	ContractAddress         string
	EventParameters []EventParameter
}
