package models

import (
	"gorm.io/gorm"
)

type EventValue struct {
	gorm.Model
	Value            string
	EventParameterID uint
	EventParameter   EventParameter
	OutboundEventID  uint
	OutboundEvent    OutboundEvent
	InboundEventID   uint
	InboundEvent     InboundEvent
}
