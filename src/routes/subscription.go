package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SubscriptionResponse struct {
	Subscription models.Subscription
}

type SubscriptionsResponse struct {
	Subscriptions []models.Subscription
}

// Register godoc
// @Summary      Get Subscription
// @Description  Get a single subscription by id for registered user.
// @Tags         subscriptions
// @Produce      json
// @Param		 subscriptionID path int true "the ID of the subscription you want to retrieve."
// @Success      200  {object}  SubscriptionResponse
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      404  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /subscription/{subscriptionID} [get]
func GetSubscription(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	subscriptionID := ctx.Param("subscriptionID")
	i, err := strconv.Atoi(subscriptionID)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid oracle id!"})
		return
	}

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)
	fmt.Println(user)

	var subscription models.Subscription
	db.Preload(clause.Associations).Find(&subscription, i)

	ctx.JSON(http.StatusOK, SubscriptionResponse{Subscription: subscription})
}

// Register godoc
// @Summary      Get Subscriptions
// @Description  Get all subscriptions for registered user.
// @Tags         subscriptions
// @Produce      json
// @Success      200  {object}  SubscriptionsResponse
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /subscriptions [get]
func GetSubscriptions(ctx *gin.Context) {
	db := utils.DBConnection()
	user := models.UserFromContext(ctx)
	var subscriptions []models.Subscription
	db.Preload(clause.Associations).Joins("JOIN outbound_oracles ON outbound_oracles.id = subscriptions.outbound_oracle_id").Find(&subscriptions, "outbound_oracles.user_id = ?", user.ID)
	ctx.JSON(http.StatusOK, SubscriptionsResponse{Subscriptions: subscriptions})
}
