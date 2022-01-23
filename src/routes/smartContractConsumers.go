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

// PostSmartContractConsumer godoc
// @Summary      Creates a smart contract consumers for a user
// @Description  Creates a smart contract consumers for a user. This service will be called by the frontend to when a user filled out the smart contract consumer form.
// @Tags         smartContractConsumer
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /smartContractConsumers [post]
func PostSmartContractConsumer(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Oh, there was a mistake!"})
		return
	}

	var smartContractConsumerBody forms.SmartContractConsumerBody
	if err = ctx.ShouldBind(&smartContractConsumerBody); err != nil || !smartContractConsumerBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	smartContract := models.SmartContract{
		BlockchainName:         smartContractConsumerBody.BlockchainName,
		ContractAddress:        smartContractConsumerBody.ContractAddress,
		ContractAddressSynonym: smartContractConsumerBody.ContractAddressSynonym,
		EventName:              smartContractConsumerBody.ContractName,
	}

	providerConsumer := models.ProviderConsumer{
		UserID:      user.ID,
		Private:     smartContractConsumerBody.Private,
		Name:        smartContractConsumerBody.Name,
		Description: smartContractConsumerBody.Description,
	}
	db.Create(&providerConsumer)

	smartContractConsumer := models.SmartContractConsumer{
		SmartContract:    smartContract,
		ProviderConsumer: providerConsumer,
	}
	db.Create(&smartContractConsumer)

	ctx.JSON(http.StatusOK, gin.H{"smartContractConsumer": smartContractConsumer})
}

// GetSmartContractConsumer godoc
// @Summary      Retrieves a smart contract consumer for a user
// @Description  Retrieves the smart contract consumer specified. This service will be called by the frontend to retrieve a specific consumers of the user signed in.
// @Tags         smartContractConsumer
// @Param		 smartContractConsumerID path int true "the ID of the smart contract consumer to send data to."
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /smartContractConsumers/{smartContractConsumerID} [get]
func GetSmartContractConsumer(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Oh, there was a mistake!"})
		return
	}

	smartContractConsumerID := ctx.Param("smartContractConsumerID")
	i, err := strconv.Atoi(smartContractConsumerID)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid subscription id!"})
		return
	}

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	var smartContractConsumer models.SmartContractConsumer
	db.Joins("SmartContract").Joins("ProviderConsumer").Preload("ProviderConsumer.EventParameters").Find(&smartContractConsumer, i)

	var inboundSubscriptions []models.InboundSubscription
	db.Joins("Subscription").Preload("SmartContractConsumer.SmartContract").Preload("SmartContractConsumer.ProviderConsumer").Preload("WebServiceProvider.ProviderConsumer").Find(&inboundSubscriptions, "smart_contract_consumer_id = ? AND Subscription.user_id = ?", smartContractConsumer.ID, user.ID)

	smartContractConsumer.InboundSubscriptions = inboundSubscriptions
	fmt.Println(smartContractConsumer)
	ctx.JSON(http.StatusOK, gin.H{"smartContractConsumer": smartContractConsumer, "inboundSubscriptions": inboundSubscriptions})
}

// GetSmartContractConsumers godoc
// @Summary      Retrieves all smart contract consumers for a user
// @Description  Retrieves all smart contract consumers for a user. This service will be called by the frontend to retrieve all smart contract consumers of the user signed in.
// @Tags         smartContractConsumer
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /smartContractConsumers [get]
func GetSmartContractConsumers(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	var smartContractConsumers []models.SmartContractConsumer
	db.Joins("ProviderConsumer").Joins("SmartContract").Find(&smartContractConsumers, "ProviderConsumer.private = 0 OR ProviderConsumer.user_id = ?", user.ID)

	ctx.JSON(http.StatusOK, gin.H{"smartContractConsumers": smartContractConsumers})
}

func PostInboundEventParameters(ctx *gin.Context) {
	smartContractConsumerIDString := ctx.Param("smartContractConsumerID")
	smartContractConsumerID, err := strconv.Atoi(smartContractConsumerIDString)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid subscription id!"})
		return
	}

	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	var smartContractConsumer models.SmartContractConsumer
	smartContractConsumerResult := db.Preload(clause.Associations).First(&smartContractConsumer, smartContractConsumerID)
	if smartContractConsumerResult.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"body": "There is no inbound subscription template with this ID"})
		return
	}
	var eventParameterBody forms.EventParameterBody
	if err = ctx.ShouldBind(&eventParameterBody); err != nil || !eventParameterBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	eventParameter := models.EventParameter{
		Name:               eventParameterBody.Name,
		Type:               eventParameterBody.Type,
		ProviderConsumerID: smartContractConsumer.ProviderConsumer.ID,
	}
	db.Create(&eventParameter)
	ctx.JSON(http.StatusOK, gin.H{"eventParameter": eventParameter})
}
