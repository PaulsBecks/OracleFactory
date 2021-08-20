package models

import (
	"log"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"gorm.io/gorm"
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
