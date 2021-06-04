package models

import (
	"gorm.io/gorm"
)

type InboundOracle struct {
	gorm.Model
	Name                    string
	InboundOracleTemplate   InboundOracleTemplate
	InboundOracleTemplateID uint
	InboundEvents           []InboundEvent
}
