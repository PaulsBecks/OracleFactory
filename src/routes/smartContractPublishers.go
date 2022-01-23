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

// PostSmartContractPublisher godoc
// @Summary      Creates a smart contract publishers for a user
// @Description  Creates a smart contract publishers for a user. This service will be called by the frontend to when a user filled out the smart contract publisher form.
// @Tags         smartContractPublisher
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /smartContractPublishers [post]
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

// GetSmartContractPublisher godoc
// @Summary      Retrieves a smart contract publisher for a user
// @Description  Retrieves the smart contract publisher specified. This service will be called by the frontend to retrieve a specific publishers of the user signed in.
// @Tags         smartContractPublisher
// @Param		 smartContractPublisherID path int true "the ID of the smart contract publisher to send data to."
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /smartContractPublishers/{smartContractPublisherID} [get]
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

// GetSmartContractPublishers godoc
// @Summary      Retrieves all smart contract publishers for a user
// @Description  Retrieves all smart contract publishers for a user. This service will be called by the frontend to retrieve all smart contract publishers of the user signed in.
// @Tags         smartContractPublisher
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /smartContractPublishers [get]
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
