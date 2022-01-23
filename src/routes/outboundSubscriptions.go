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

// GetOutboundSubscriptions godoc
// @Summary      Retrieves all outbound subscription of a user
// @Description  Retrieve all outbound subscriptions of a user. This will be called from the frontend, when a user wants to view a list of subscription.
// @Tags         outboundSubscriptions
// @Produce      json
// @Router       /outboundSubscriptions [get]
func GetOutboundSubscriptions(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	var subscriptions []models.OutboundSubscription
	db.Preload(clause.Associations).Preload("SmartContractProvider.SmartContract").Preload("SmartContractProvider.ProviderConsumer").Preload("WebServiceConsumer.ProviderConsumer").Joins("Subscription").Find(&subscriptions, "Subscription.user_id = ?", user.ID)
	fmt.Println(subscriptions)

	ctx.JSON(http.StatusOK, gin.H{"outboundSubscriptions": subscriptions})
}

// GetOutboundSubscription godoc
// @Summary      Retrieve an outbound subscription
// @Description  Retrieve the specified outbound subscription. This will be called from the frontend, when a user wants to view an subscription.
// @Tags         outboundSubscriptions
// @Param		 outboundSubscriptionID path int true "the ID of the outbound subscription you want to retrieve."
// @Produce      json
// @Router       /outboundSubscriptions/{outboundSubscriptionID} [get]
func GetOutboundSubscription(ctx *gin.Context) {
	id := ctx.Param("outboundSubscriptionId")
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid subscription id!"})
		return
	}

	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var outboundSubscription models.OutboundSubscription
	result := db.Preload("Subscription.Events.EventValues.EventParameter").Preload("SmartContractProvider.ProviderConsumer.EventParameters").Preload("Subscription.ParameterFilters.Filter").Preload(clause.Associations).First(&outboundSubscription, i)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No outbound Subscription with this ID available."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"outboundSubscription": outboundSubscription})
}

// UpdateOutboundSubscription godoc
// @Summary      Update an outbound subscription
// @Description  Update the specified outbound subscription. This will be called from the frontend, when a user wants to update an subscription.
// @Tags         outboundSubscriptions
// @Param		 outboundSubscriptionID path int true "the ID of the outbound subscription you want to update."
// @Produce      json
// @Router       /outboundSubscriptions/{outboundSubscriptionID} [put]
func UpdateOutboundSubscription(ctx *gin.Context) {
	id := ctx.Param("outboundSubscriptionId")
	outboundSubscription, err := models.GetOutboundSubscriptionById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No outbound Subscription with this ID available."})
		return
	}
	var outboundSubscriptionPostBody forms.OutboundSubscriptionPostBody
	if err = ctx.ShouldBind(&outboundSubscriptionPostBody); err != nil || !outboundSubscriptionPostBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}
	outboundSubscription.Save()

	subscription := outboundSubscription.Subscription
	subscription.Name = outboundSubscriptionPostBody.Subscription.Name
	subscription.Save()

	ctx.JSON(http.StatusOK, gin.H{"outboundSubscription": outboundSubscription})
}

func DeleteOutboundSubscription(ctx *gin.Context) {
	// TODO: delete subscription for the provided id
	ctx.JSON(http.StatusNotImplemented, gin.H{"body": "Hi there, deletion is not implemented yet!"})
}

// StartOutboundSubscription godoc
// @Summary      Start an Outbound Subscription
// @Description  Start the specified outbound subscription. This will be called from the frontend, when a user wants to use an subscription for a blockchain conenction.
// @Tags         outboundSubscriptions
// @Param		 outboundSubscriptionID path int true "the ID of the outbound subscription you want to start."
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /outboundSubscriptions/{outboundSubscriptionID}/start [post]
func StartOutboundSubscription(ctx *gin.Context) {
	id := ctx.Param("outboundSubscriptionId")
	outboundSubscription, err := models.GetOutboundSubscriptionById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No outbound Subscription with this ID available."})
		return
	}
	err = outboundSubscription.StartSubscription()
	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Unable to start subscription, try again later."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Subscription got started successfully."})
}

// StopOutboundSubscription godoc
// @Summary      Stop an outbound subscription
// @Description  Stop the specified outbound subscription. This will be called from the frontend, when a user wants to stop an subscription for a blockchain conenction.
// @Tags         outboundSubscriptions
// @Param		 outboundSubscriptionID path int true "the ID of the outbound subscription you want to stop."
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /outboundSubscriptions/{outboundSubscriptionID}/stop [post]
func StopOutboundSubscription(ctx *gin.Context) {
	id := ctx.Param("outboundSubscriptionId")
	outboundSubscription, err := models.GetOutboundSubscriptionById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No outbound Subscription with this ID available."})
		return
	}
	fmt.Print(outboundSubscription)
	err = outboundSubscription.StopSubscription()
	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Unable to stop subscription, try again later."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Subscription got stopped successfully."})
}

func PostOutboundSubscriptionEvent(ctx *gin.Context) {
	id := ctx.Param("outboundSubscriptionId")
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid subscription id!"})
		return
	}

	outboundSubscription, _ := models.GetOutboundSubscriptionById(i)

	data, _ := ioutil.ReadAll(ctx.Request.Body)
	var bodyData map[string]interface{}
	if err := json.Unmarshal(data, &bodyData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	event := models.CreateEvent(data, outboundSubscription.GetSubscription().ID)
	event.ParseEventValues(bodyData, outboundSubscription.GetSmartContractProvider().ProviderConsumerID)
	if outboundSubscription.GetSubscription().CheckInput(event) {
		webServiceConsumer := outboundSubscription.GetWebServiceConsumer()
		webServiceConsumer.Publish(*event)
	}
}

// PostOutboundSubscription godoc
// @Summary      Creates an outbound subscription for a user
// @Description  Creates an outbound subscription for a user. This service will be called by the frontend to when a user filled out the outbound subscription form.
// @Tags         outboundSubscriptions
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /outboundSubscriptions [post]
func PostOutboundSubscription(ctx *gin.Context) {
	var outboundSubscriptionPostBody forms.OutboundSubscriptionPostBody
	if err := ctx.ShouldBind(&outboundSubscriptionPostBody); err != nil || !outboundSubscriptionPostBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	user := models.UserFromContext(ctx)
	outboundSubscription := user.CreateOutboundSubscription(
		outboundSubscriptionPostBody.Subscription.Name,
		outboundSubscriptionPostBody.SmartContractProviderID,
		outboundSubscriptionPostBody.WebServiceConsumerID,
	)
	ctx.JSON(http.StatusOK, gin.H{"outboundSubscription": outboundSubscription})
}
