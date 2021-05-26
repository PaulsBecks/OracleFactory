package models

import (
	"gorm.io/gorm"
)

type OutboundEvent struct {
	gorm.Model
	OutboundOracle   OutboundOracle
	OutboundOracleID uint
	EventValues      []EventValue
}
