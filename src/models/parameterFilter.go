package models

import (
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
)

type ParameterFilter struct {
	gorm.Model
	Scheme           string
	EventParameter   EventParameter
	EventParameterID uint
	Filter           Filter
	FilterID         uint
	Subscription     Subscription
	SubscriptionID   uint
}

func (p *ParameterFilter) GetFilter() *Filter {
	db := utils.DBConnection()

	var filter Filter
	db.First(&filter, p.FilterID)
	return &filter
}

func (p *ParameterFilter) Check(input string) bool {
	return p.GetFilter().Check(p.Scheme, input)
}

func (p *ParameterFilter) GetEventParameter() EventParameter {
	return GetEventParameterByID(p.EventParameterID)
}
