package models

import (
	"gorm.io/gorm"
)

type EventValue struct {
	gorm.Model
	EventParameterID uint
	EventParameter   EventParameter
	OutboundEventID  uint
	OutboundEvent    OutboundEvent
}
