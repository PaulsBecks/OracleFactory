package models

import (
	"log"

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
	Oracle           Oracle
	OracleID         uint
}

func (p *ParameterFilter) GetFilter() *Filter {
	db, err := utils.DBConnection()
	if err != nil {
		log.Fatal(err)
	}
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
