package models

import (
	"gorm.io/gorm"
)

type InboundOracle struct {
	gorm.Model
	Name                    string
	UserID                  uint
	User                    User
	InboundOracleTemplate   InboundOracleTemplate
	InboundOracleTemplateID uint
	InboundEvents           []InboundEvent
}
