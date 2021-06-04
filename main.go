package main

import (
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/PaulsBecks/OracleFactory/src/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() {
	//db, err := sql.Open("sqlite", c.filePath)
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Check if table exists - if not create it
	db.AutoMigrate(&models.OutboundOracleTemplate{})
	db.AutoMigrate(&models.EventValue{})
	db.AutoMigrate(&models.OutboundOracle{})
	db.AutoMigrate(&models.EventParameter{})
	db.AutoMigrate(&models.OutboundEvent{})
	db.AutoMigrate(&models.EventParameter{})
	db.AutoMigrate(&models.InboundOracleTemplate{})
	db.AutoMigrate(&models.InboundOracle{})
	db.AutoMigrate(&models.OutboundEvent{})
	oot := &models.OutboundOracleTemplate{
		Address:           "0xe4EFfB267484Cd790143484de3Bae7fDfbE31F00",
		EventName:         "Stored",
		Blockchain:        "Ethereum",
		BlockchainAddress: "ws://eth-test-net:8545/",
	}
	db.Create(oot)
	db.Create(&models.EventParameter{
		Type:                     "uint256",
		Name:                     "storedData",
		OutboundOracleTemplateID: oot.ID,
	})

	iot := &models.InboundOracleTemplate{
		ContractAddress:   "0xe4EFfB267484Cd790143484de3Bae7fDfbE31F00",
		ContractName:      "set",
		BlockchainName:    "Ethereum",
		BlockchainAddress: "ws://eth-test-net:8545/",
	}
	db.Create(iot)
	db.Create(&models.EventParameter{
		Type:                    "uint256",
		Name:                    "x",
		InboundOracleTemplateID: iot.ID,
	})
}

func main() {
	app := gin.Default()
	middleware(app)
	InitDB()

	app.GET("/outboundOracles", routes.GetOutboundOracles)
	app.GET("/outboundOracles/:outboundOracleId", routes.GetOutboundOracle)
	app.PUT("/outboundOracles/:outboundOracleId", routes.UpdateOutboundOracle)
	app.DELETE("/outboundOracles/:outboundOracleId", routes.DeleteOutboundOracle)
	app.POST("/outboundOracles/:outboundOracleId/events", routes.PostOutboundOracleEvent)

	app.GET("/outboundOracleTemplates", routes.GetOutboundOracleTemplates)
	app.GET("/outboundOracleTemplates/:outboundOracleTemplateID", routes.GetOutboundOracleTemplate)
	app.POST("/outboundOracleTemplates/:outboundOracleTemplateId/outboundOracles", routes.PostOutboundOracle)

	app.GET("/inboundOracles/:inboundOracleId", routes.GetInboundOracle)
	app.PUT("/inboundOracles/:inboundOracleId", routes.UpdateInboundOracle)
	app.POST("/inboundOracles/:inboundOracleID/events", routes.PostInboundOracleEvent)
	app.GET("/inboundOracleTemplates/:inboundOracleTemplateID", routes.GetInboundOracleTemplate)
	app.GET("/inboundOracles", routes.GetInboundOracles)

	app.POST("/inboundOracleTemplates/:inboundOracleTemplateID/inboundOracles", routes.PostInboundOracle)
	app.GET("/inboundOracleTemplates", routes.GetInboundOracleTemplates)

	app.Run()
}
