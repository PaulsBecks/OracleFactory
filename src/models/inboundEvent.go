package models

import (
	"gorm.io/gorm"
)

type InboundEvent struct {
	gorm.Model
	Success         bool
	InboundOracle   InboundOracle
	InboundOracleID uint
	EventValues     []EventValue
}
