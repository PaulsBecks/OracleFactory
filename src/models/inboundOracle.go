package models

import (
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InboundOracle struct {
	gorm.Model
	Oracle                  Oracle
	OracleID                uint
	InboundOracleTemplate   InboundOracleTemplate
	InboundOracleTemplateID uint
}

func (i *InboundOracle) GetOracle() *Oracle {
	db := utils.DBConnection()

	var oracle Oracle
	db.Find(&oracle, i.OracleID)
	return &oracle
}

func GetInboundOracleByID(id interface{}) (*InboundOracle, error) {
	var inboundOracle InboundOracle
	db := utils.DBConnection()
	result := db.Preload(clause.Associations).Preload("InboundOracleTemplate.OracleTemplate.EventParameters").First(&inboundOracle, id)
	if result.Error != nil {
		return &inboundOracle, result.Error
	}
	return &inboundOracle, nil
}
