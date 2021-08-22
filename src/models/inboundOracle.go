package models

import (
	"log"

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
	db, err := utils.DBConnection()
	if err != nil {
		log.Fatal(err)
	}
	var oracle Oracle
	db.Find(&oracle, i.OracleID)
	return &oracle
}

func GetInboundOracleByID(id interface{}) (*InboundOracle, error) {
	var inboundOracle InboundOracle
	db, err := utils.DBConnection()
	if err != nil {
		log.Fatal("No DB connection")
	}
	result := db.Preload(clause.Associations).Preload("InboundOracleTemplate.OracleTemplate.EventParameters").First(&inboundOracle, id)
	if result.Error != nil {
		return &inboundOracle, result.Error
	}
	return &inboundOracle, nil
}
