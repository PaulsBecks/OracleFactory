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

func PostConsumer(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Oh, there was a mistake!"})
		return
	}

	var consumerBody forms.ConsumerBody
	if err = ctx.ShouldBind(&consumerBody); err != nil || !consumerBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	smartContract := models.SmartContract{
		BlockchainName:         consumerBody.BlockchainName,
		ContractAddress:        consumerBody.ContractAddress,
		ContractAddressSynonym: consumerBody.ContractAddressSynonym,
		EventName:              consumerBody.EventName,
	}

	listenerPublisher := models.ListenerPublisher{
		UserID:      user.ID,
		Private:     consumerBody.Private,
		Name:        consumerBody.Name,
		Description: consumerBody.Description,
	}
	db.Create(&listenerPublisher)

	consumer := models.Consumer{
		SmartContract:     smartContract,
		ListenerPublisher: listenerPublisher,
	}
	db.Create(&consumer)

	ctx.JSON(http.StatusOK, gin.H{"consumer": consumer})
}

func GetConsumer(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Oh, there was a mistake!"})
		return
	}

	consumerID := ctx.Param("consumerID")
	i, err := strconv.Atoi(consumerID)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid oracle id!"})
		return
	}

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	var consumer models.Consumer
	db.Joins("SmartContract").Joins("ListenerPublisher").Preload("ListenerPublisher.EventParameters").Find(&consumer, i)

	var pubSubOracles []models.PubSubOracle
	db.Joins("Oracle").Preload("Consumer.SmartContract").Preload("Consumer.ListenerPublisher").Preload("Provider.ListenerPublisher").Find(&pubSubOracles, "consumer_id = ? AND Oracle.user_id = ?", consumer.ID, user.ID)

	consumer.PubSubOracles = pubSubOracles
	ctx.JSON(http.StatusOK, gin.H{"consumer": consumer, "pubSubOracles": pubSubOracles})
}

func GetConsumers(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	var consumers []models.Consumer
	db.Joins("ListenerPublisher").Joins("SmartContract").Find(&consumers, "ListenerPublisher.private = 0 OR ListenerPublisher.user_id = ?", user.ID)

	ctx.JSON(http.StatusOK, gin.H{"consumers": consumers})
}

func PostInboundEventParameters(ctx *gin.Context) {
	consumerIDString := ctx.Param("consumerID")
	consumerID, err := strconv.Atoi(consumerIDString)
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

	var consumer models.Consumer
	consumerResult := db.Preload(clause.Associations).First(&consumer, consumerID)
	if consumerResult.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"body": "There is no pubSub oracle template with this ID"})
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
		ListenerPublisherID: consumer.ListenerPublisher.ID,
	}
	db.Create(&eventParameter)
	ctx.JSON(http.StatusOK, gin.H{"eventParameter": eventParameter})
}
