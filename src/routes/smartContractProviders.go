package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// HandleSmartContractProviderEvent godoc
// @Summary      Handles the event send from a smart contract provider
// @Description  Handles the event send from a smart contract provider. This endpoint will be called from the BLF, that provides data to the artifact.
// @Tags         smartContractProvider
// @Param		 smartContractProviderID path int true "the ID of the smart contract provider to send data to."
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /smartContractProviders/{smartContractProviderID}/events [post]
func HandleSmartContractEvent(ctx *gin.Context) {
	id := ctx.Param("smartContractProviderID")
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid provider id!"})
		return
	}
	fmt.Println(id, i)
	smartContractProvider := models.GetSmartContractProviderByID(uint(i))

	data, _ := ioutil.ReadAll(ctx.Request.Body)
	var bodyData map[string]interface{}
	if err := json.Unmarshal(data, &bodyData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	fmt.Println(smartContractProvider)
	for _, outboundSubscription := range smartContractProvider.OutboundSubscriptions {
		fmt.Println(outboundSubscription)
		event := models.CreateEvent(data, outboundSubscription.GetSubscription().ID)
		event.ParseEventValues(bodyData, outboundSubscription.GetSmartContractProvider().ProviderConsumerID)

		webServiceConsumer := outboundSubscription.GetWebServiceConsumer()
		webServiceConsumer.Publish(*event)
	}
}

// GetSmartContractProvider godoc
// @Summary      Retrieves a smart contract provider
// @Description  Retrieves a smart contract provider. This endpoint will be called from the frontend, to display information about a smart contract provider.
// @Tags         smartContractProvider
// @Param		 smartContractProviderID path int true "the ID of the smart contract provider to send data to."
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /smartContractProviders/{smartContractProviderID} [get]
func GetSmartContractProvider(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	smartContractProviderID := ctx.Param("smartContractProviderID")
	i, err := strconv.Atoi(smartContractProviderID)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid subscription id!"})
		return
	}

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	var smartContractProvider models.SmartContractProvider
	db.Joins("ProviderConsumer").Preload("ProviderConsumer.EventParameters").Preload("SmartContract").Preload("OutboundSubscriptions.Subscription").Find(&smartContractProvider, i)

	var outboundSubscriptions []models.OutboundSubscription
	db.Joins("Subscription").Preload("SmartContractProvider.SmartContract").Preload("SmartContractProvider.ProviderConsumer").Preload("WebServiceConsumer.ProviderConsumer").Find(&outboundSubscriptions, "smart_contract_provider_id = ? AND Subscription.user_id = ?", smartContractProvider.ID, user.ID)
	smartContractProvider.OutboundSubscriptions = outboundSubscriptions

	ctx.JSON(http.StatusOK, gin.H{"smartContractProvider": smartContractProvider})
}

// GetSmartContractProviders godoc
// @Summary      Retrieves all smart contract provider of the user signed in.
// @Description  Retrieves all smart contract provider of the user signed in. This endpoint will be called from the frontend, to display information about all smart contract providers of the user signed in.
// @Tags         smartContractProvider
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /smartContractProviders [get]
func GetSmartContractProviders(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	var smartContractProviders []models.SmartContractProvider
	db.Joins("ProviderConsumer").Joins("SmartContract").Preload(clause.Associations).Find(&smartContractProviders, "ProviderConsumer.private = 0 OR ProviderConsumer.user_id = ?", user.ID)

	fmt.Println(smartContractProviders)

	ctx.JSON(http.StatusOK, gin.H{"smartContractProviders": smartContractProviders})
}

// PostSmartContractProvider godoc
// @Summary      Creates a smart contract providers for a user
// @Description  Creates a smart contract providers for a user. This service will be called by the frontend to when a user filled out the smart contract provider form.
// @Tags         smartContractProvider
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /smartContractProviders [post]
func PostSmartContractProvider(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	var smartContractProviderBody forms.SmartContractProviderBody
	if err = ctx.ShouldBind(&smartContractProviderBody); err != nil || !smartContractProviderBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	smartContract := models.SmartContract{
		BlockchainName:         smartContractProviderBody.BlockchainName,
		ContractAddress:        smartContractProviderBody.ContractAddress,
		ContractAddressSynonym: smartContractProviderBody.ContractAddressSynonym,
		EventName:              smartContractProviderBody.EventName,
	}
	db.Create(&smartContract)

	providerConsumer := models.ProviderConsumer{
		UserID:      user.ID,
		Private:     smartContractProviderBody.Private,
		Name:        smartContractProviderBody.Name,
		Description: smartContractProviderBody.Description,
	}
	db.Create(&providerConsumer)

	smartContractProvider := models.SmartContractProvider{SmartContract: smartContract, ProviderConsumer: providerConsumer}
	db.Create(&smartContractProvider)
	ctx.JSON(http.StatusOK, gin.H{"smartContractProvider": smartContractProvider})
}

func PostOutboundEventParameters(ctx *gin.Context) {
	smartContractProviderIDString := ctx.Param("smartContractProviderID")
	smartContractProviderID, err := strconv.Atoi(smartContractProviderIDString)
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

	var smartContractProvider models.SmartContractProvider
	smartContractProviderResult := db.Preload(clause.Associations).First(&smartContractProvider, smartContractProviderID)
	if smartContractProviderResult.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"body": "There is no outbound subscription template with this ID"})
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
		Indexed:            eventParameterBody.Indexed,
		ProviderConsumerID: smartContractProvider.ProviderConsumer.ID,
	}
	db.Create(&eventParameter)
	ctx.JSON(http.StatusOK, gin.H{"eventParameter": eventParameter})
}
