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

func GetSmartContractListener(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	smartContractListenerID := ctx.Param("smartContractListenerID")
	i, err := strconv.Atoi(smartContractListenerID)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid oracle id!"})
		return
	}

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	var smartContractListener models.SmartContractListener
	db.Joins("ListenerPublisher").Preload("ListenerPublisher.EventParameters").Preload("SmartContract").Preload("OutboundOracles.Oracle").Find(&smartContractListener, i)

	var outboundOracles []models.OutboundOracle
	db.Joins("Oracle").Preload("SmartContractListener.SmartContract").Preload("SmartContractListener.ListenerPublisher").Preload("WebServicePublisher.ListenerPublisher").Find(&outboundOracles, "smart_contract_listener_id = ? AND Oracle.user_id = ?", smartContractListener.ID, user.ID)
	smartContractListener.OutboundOracles = outboundOracles

	ctx.JSON(http.StatusOK, gin.H{"smartContractListener": smartContractListener})
}

func GetSmartContractListeners(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	var smartContractListeners []models.SmartContractListener
	db.Joins("ListenerPublisher").Joins("SmartContract").Preload(clause.Associations).Find(&smartContractListeners, "ListenerPublisher.private = 0 OR ListenerPublisher.user_id = ?", user.ID)

	fmt.Println(smartContractListeners)

	ctx.JSON(http.StatusOK, gin.H{"smartContractListeners": smartContractListeners})
}

func PostOutboundOracle(ctx *gin.Context) {
	var outboundOraclePostBody forms.OutboundOraclePostBody
	if err := ctx.ShouldBind(&outboundOraclePostBody); err != nil || !outboundOraclePostBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	user := models.UserFromContext(ctx)
	outboundOracle := user.CreateOutboundOracle(
		outboundOraclePostBody.Oracle.Name,
		outboundOraclePostBody.SmartContractListenerID,
		outboundOraclePostBody.WebServicePublisherID,
	)
	ctx.JSON(http.StatusOK, gin.H{"outboundOracle": outboundOracle})
}

func PostSmartContractListener(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	var smartContractListenerBody forms.SmartContractListenerBody
	if err = ctx.ShouldBind(&smartContractListenerBody); err != nil || !smartContractListenerBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	smartContract := models.SmartContract{
		BlockchainName:         smartContractListenerBody.BlockchainName,
		ContractAddress:        smartContractListenerBody.ContractAddress,
		ContractAddressSynonym: smartContractListenerBody.ContractAddressSynonym,
		EventName:              smartContractListenerBody.EventName,
	}
	db.Create(&smartContract)

	listenerPublisher := models.ListenerPublisher{
		UserID:      user.ID,
		Private:     smartContractListenerBody.Private,
		Name:        smartContractListenerBody.Name,
		Description: smartContractListenerBody.Description,
	}
	db.Create(&listenerPublisher)

	smartContractListener := models.SmartContractListener{SmartContract: smartContract, ListenerPublisher: listenerPublisher}
	db.Create(&smartContractListener)
	ctx.JSON(http.StatusOK, gin.H{"smartContractListener": smartContractListener})
}

func PostOutboundEventParameters(ctx *gin.Context) {
	smartContractListenerIDString := ctx.Param("smartContractListenerID")
	smartContractListenerID, err := strconv.Atoi(smartContractListenerIDString)
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

	var smartContractListener models.SmartContractListener
	smartContractListenerResult := db.Preload(clause.Associations).First(&smartContractListener, smartContractListenerID)
	if smartContractListenerResult.Error != nil {
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
		Indexed:             eventParameterBody.Indexed,
		ListenerPublisherID: smartContractListener.ListenerPublisher.ID,
	}
	db.Create(&eventParameter)
	ctx.JSON(http.StatusOK, gin.H{"eventParameter": eventParameter})
}
