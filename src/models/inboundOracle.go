package models

import (
	"gorm.io/gorm"
)

type InboundOracle struct {
	gorm.Model
	InboundOracleTemplate InboundOracleTemplate
	InboundOracleTemplateID uint
}
