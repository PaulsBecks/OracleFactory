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

func PostSmartContractPublisher(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Oh, there was a mistake!"})
		return
	}

	var smartContractPublisherBody forms.SmartContractPublisherBody
	if err = ctx.ShouldBind(&smartContractPublisherBody); err != nil || !smartContractPublisherBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	smartContract := models.SmartContract{
		BlockchainName:         smartContractPublisherBody.BlockchainName,
		ContractAddress:        smartContractPublisherBody.ContractAddress,
		ContractAddressSynonym: smartContractPublisherBody.ContractAddressSynonym,
		EventName:              smartContractPublisherBody.ContractName,
	}

	listenerPublisher := models.ListenerPublisher{
		UserID:      user.ID,
		Private:     smartContractPublisherBody.Private,
		Name:        smartContractPublisherBody.Name,
		Description: smartContractPublisherBody.Description,
	}
	db.Create(&listenerPublisher)

	smartContractPublisher := models.SmartContractPublisher{
		SmartContract:     smartContract,
		ListenerPublisher: listenerPublisher,
	}
	db.Create(&smartContractPublisher)

	ctx.JSON(http.StatusOK, gin.H{"smartContractPublisher": smartContractPublisher})
}

func GetSmartContractPublisher(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Oh, there was a mistake!"})
		return
	}

	smartContractPublisherID := ctx.Param("smartContractPublisherID")
	i, err := strconv.Atoi(smartContractPublisherID)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid oracle id!"})
		return
	}

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	var smartContractPublisher models.SmartContractPublisher
	db.Joins("SmartContract").Joins("ListenerPublisher").Preload("ListenerPublisher.EventParameters").Find(&smartContractPublisher, i)

	var inboundOracles []models.InboundOracle
	db.Joins("Oracle").Preload("SmartContractPublisher.SmartContract").Preload("SmartContractPublisher.ListenerPublisher").Preload("WebServiceListener.ListenerPublisher").Find(&inboundOracles, "smart_contract_publisher_id = ? AND Oracle.user_id = ?", smartContractPublisher.ID, user.ID)

	smartContractPublisher.InboundOracles = inboundOracles
	fmt.Println(smartContractPublisher)
	ctx.JSON(http.StatusOK, gin.H{"smartContractPublisher": smartContractPublisher, "inboundOracles": inboundOracles})
}

func GetSmartContractPublishers(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	var smartContractPublishers []models.SmartContractPublisher
	db.Joins("ListenerPublisher").Joins("SmartContract").Find(&smartContractPublishers, "ListenerPublisher.private = 0 OR ListenerPublisher.user_id = ?", user.ID)

	ctx.JSON(http.StatusOK, gin.H{"smartContractPublishers": smartContractPublishers})
}

func PostInboundEventParameters(ctx *gin.Context) {
	smartContractPublisherIDString := ctx.Param("smartContractPublisherID")
	smartContractPublisherID, err := strconv.Atoi(smartContractPublisherIDString)
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

	var smartContractPublisher models.SmartContractPublisher
	smartContractPublisherResult := db.Preload(clause.Associations).First(&smartContractPublisher, smartContractPublisherID)
	if smartContractPublisherResult.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"body": "There is no inbound oracle template with this ID"})
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
		ListenerPublisherID: smartContractPublisher.ListenerPublisher.ID,
	}
	db.Create(&eventParameter)
	ctx.JSON(http.StatusOK, gin.H{"eventParameter": eventParameter})
}
