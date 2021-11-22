package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetBlockchainEvent(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	blockchainEventID := ctx.Param("blockchainEventID")
	i, err := strconv.Atoi(blockchainEventID)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid oracle id!"})
		return
	}

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	var blockchainEvent models.BlockchainEvent
	db.Joins("ListenerPublisher").Preload("ListenerPublisher.EventParameters").Preload("SmartContract").Find(&blockchainEvent, i)

	var outboundOracles []models.OutboundOracle
	db.Joins("Oracle").Preload("BlockchainEvent.SmartContract").Preload("BlockchainEvent.ListenerPublisher").Joins("JOIN pub_sub_oracles ON pub_sub_oracles.sub_oracle_id = outbound_oracles.id OR pub_sub_oracles.unsub_oracle_id = outbound_oracles.id").Find(&outboundOracles, "blockchain_event_id = ? AND Oracle.user_id = ?", blockchainEvent.ID, user.ID)
	blockchainEvent.OutboundOracles = outboundOracles

	ctx.JSON(http.StatusOK, gin.H{"blockchainEvent": blockchainEvent})
}

func GetBlockchainEvents(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	var blockchainEvents []models.BlockchainEvent
	db.Joins("ListenerPublisher").Joins("SmartContract").Preload(clause.Associations).Find(&blockchainEvents, "ListenerPublisher.private = 0 OR ListenerPublisher.user_id = ?", user.ID)

	fmt.Println(blockchainEvents)

	ctx.JSON(http.StatusOK, gin.H{"blockchainEvents": blockchainEvents})
}

func PostBlockchainEvent(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	var blockchainEventBody forms.BlockchainEventBody
	if err = ctx.ShouldBind(&blockchainEventBody); err != nil || !blockchainEventBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	smartContract := models.SmartContract{
		BlockchainName:         blockchainEventBody.BlockchainName,
		ContractAddress:        blockchainEventBody.ContractAddress,
		ContractAddressSynonym: blockchainEventBody.ContractAddressSynonym,
		EventName:              blockchainEventBody.EventName,
	}
	db.Create(&smartContract)

	listenerPublisher := models.ListenerPublisher{
		UserID:      user.ID,
		Private:     blockchainEventBody.Private,
		Name:        blockchainEventBody.Name,
		Description: blockchainEventBody.Description,
	}
	db.Create(&listenerPublisher)

	blockchainEvent := models.BlockchainEvent{SmartContract: smartContract, ListenerPublisher: listenerPublisher}
	db.Create(&blockchainEvent)
	ctx.JSON(http.StatusOK, gin.H{"blockchainEvent": blockchainEvent})
}

func PostOutboundEventParameters(ctx *gin.Context) {
	blockchainEventIDString := ctx.Param("blockchainEventID")
	blockchainEventID, err := strconv.Atoi(blockchainEventIDString)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid oracle id!"})
		return
	}

	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	var blockchainEvent models.BlockchainEvent
	blockchainEventResult := db.Preload(clause.Associations).First(&blockchainEvent, blockchainEventID)
	if blockchainEventResult.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"body": "There is no outbound oracle template with this ID"})
		return
	}
	var eventParameterBody forms.EventParameterBody
	if err = ctx.ShouldBind(&eventParameterBody); err != nil || !eventParameterBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	eventParameter := models.EventParameter{
		Name:                eventParameterBody.Name,
		Type:                eventParameterBody.Type,
		ListenerPublisherID: blockchainEvent.ListenerPublisher.ID,
	}
	db.Create(&eventParameter)
	ctx.JSON(http.StatusOK, gin.H{"eventParameter": eventParameter})
}
