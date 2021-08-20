package models

import (
	"gorm.io/gorm"
)

type ParameterFilter struct {
	gorm.Model
	Scheme           string
	EventParameter   EventParameter
	EventParameterID uint
	Filter           Filter
	FilterID         uint
	Oracle           Oracle
	OracleID         uint
}

func (p *ParameterFilter) Check(input string) bool {
	return p.Filter.Check(p.Scheme, input)
}

func (p *ParameterFilter) GetEventParameter() EventParameter {
	return GetEventParameterByID(p.EventParameterID)
}
