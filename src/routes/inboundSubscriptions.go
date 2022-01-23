package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func PostInboundSubscriptionEvent(ctx *gin.Context) {
	inboundSubscriptionID := ctx.Param("inboundSubscriptionID")
	inboundSubscription, err := models.GetInboundSubscriptionByID(inboundSubscriptionID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No valid subscription id!"})
		return
	}
	fmt.Printf("Event submitted for inbound subscription %s\n", inboundSubscriptionID)
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	inboundSubscription.HandleEvent(data)
	ctx.JSON(http.StatusOK, gin.H{})
}

// GetInboundSubscriptions godoc
// @Summary      Retrieves all inbound subscription of a user
// @Description  Retrieve all inbound subscriptions of a user. This will be called from the frontend, when a user wants to view a list of subscription.
// @Tags         inboundSubscriptions
// @Produce      json
// @Router       /inboundSubscriptions [get]
func GetInboundSubscriptions(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	db := utils.DBConnection()

	var inboundSubscriptions []models.InboundSubscription
	db.Preload(clause.Associations).Preload("SmartContractConsumer.SmartContract").Preload("SmartContractConsumer.ProviderConsumer").Preload("WebServiceProvider.ProviderConsumer").Joins("Subscription").Find(&inboundSubscriptions, "Subscription.user_id = ?", user.ID)

	ctx.JSON(http.StatusOK, gin.H{"inboundSubscriptions": inboundSubscriptions})
}

// GetInboundSubscription godoc
// @Summary      Retrieve an inbound subscription
// @Description  Retrieve the specified inbound subscription. This will be called from the frontend, when a user wants to view an subscription.
// @Tags         inboundSubscriptions
// @Param		 inboundSubscriptionID path int true "the ID of the inbound subscription you want to retrieve."
// @Produce      json
// @Router       /inboundSubscriptions/{inboundSubscriptionID} [get]
func GetInboundSubscription(ctx *gin.Context) {
	id := ctx.Param("inboundSubscriptionId")
	db := utils.DBConnection()

	var inboundSubscription models.InboundSubscription
	result := db.Preload("Subscription.Events.EventValues.EventParameter").Preload("SmartContractConsumer.ProviderConsumer.EventParameters").Preload(clause.Associations).First(&inboundSubscription, id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No inbound Subscription with this ID available."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"inboundSubscription": inboundSubscription})
}

// UpdateInboundSubscription godoc
// @Summary      Update an inbound subscription
// @Description  Update the specified inbound subscription. This will be called from the frontend, when a user wants to update an subscription.
// @Tags         inboundSubscriptions
// @Param		 inboundSubscriptionID path int true "the ID of the inbound subscription you want to update."
// @Produce      json
// @Router       /inboundSubscriptions/{inboundSubscriptionID} [put]
func UpdateInboundSubscription(ctx *gin.Context) {
	id := ctx.Param("inboundSubscriptionId")
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var inboundSubscription models.InboundSubscription
	result := db.Preload(clause.Associations).First(&inboundSubscription, id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No inbound Subscription with this ID available."})
		return
	}
	var inboundSubscriptionPostBody forms.InboundSubscriptionBody
	if err = ctx.ShouldBind(&inboundSubscriptionPostBody); err != nil || !inboundSubscriptionPostBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	subscription := inboundSubscription.Subscription
	subscription.Name = inboundSubscriptionPostBody.Subscription.Name

	db.Save(&subscription)
	ctx.JSON(http.StatusOK, gin.H{"inboundSubscription": inboundSubscription})
}

// StartInboundSubscription godoc
// @Summary      Start an Outbound Subscription
// @Description  Start the specified inbound subscription. This will be called from the frontend, when a user wants to use an subscription for a blockchain conenction.
// @Tags         inboundSubscriptions
// @Param		 inboundSubscriptionID path int true "the ID of the inbound subscription you want to start."
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /inboundSubscriptions/{inboundSubscriptionID}/start [post]
func StartInboundSubscription(ctx *gin.Context) {
	id := ctx.Param("inboundSubscriptionID")
	inboundSubscription, err := models.GetInboundSubscriptionByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No inbound Subscription with this ID available."})
		return
	}
	inboundSubscription.GetSubscription().Start()
	ctx.JSON(http.StatusOK, gin.H{"msg": "Subscription got started successfully."})
}

// StopInboundSubscription godoc
// @Summary      Stop an inbound subscription
// @Description  Stop the specified inbound subscription. This will be called from the frontend, when a user wants to stop an subscription for a blockchain conenction.
// @Tags         inboundSubscriptions
// @Param		 inboundSubscriptionID path int true "the ID of the inbound subscription you want to stop."
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /inboundSubscriptions/{inboundSubscriptionID}/stop [post]
func StopInboundSubscription(ctx *gin.Context) {
	id := ctx.Param("inboundSubscriptionID")
	inboundSubscription, err := models.GetInboundSubscriptionByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No inbound Subscription with this ID available."})
		return
	}
	inboundSubscription.GetSubscription().Stop()
	ctx.JSON(http.StatusOK, gin.H{"msg": "Subscription got stopped successfully."})
}

// PostInboundSubscription godoc
// @Summary      Creates an inbound subscription for a user
// @Description  Creates an inbound subscription for a user. This service will be called by the frontend to when a user filled out the inbound subscription form.
// @Tags         inboundSubscriptions
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /inboundSubscriptions [post]
func PostInboundSubscription(ctx *gin.Context) {
	var inboundSubscriptionBody forms.InboundSubscriptionBody
	if err := ctx.ShouldBind(&inboundSubscriptionBody); err != nil || !inboundSubscriptionBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	user := models.UserFromContext(ctx)
	inboundSubscription := user.CreateInboundSubscription(
		inboundSubscriptionBody.Subscription.Name,
		inboundSubscriptionBody.SmartContractConsumerID,
		inboundSubscriptionBody.WebServiceProviderID,
	)
	ctx.JSON(http.StatusOK, gin.H{"inboundSubscription": inboundSubscription})
}
