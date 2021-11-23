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

	ctx.JSON(http.StatusOK, gin.H{"subscription": subscription})
}

func GetSubscriptions(ctx *gin.Context) {
	db := utils.DBConnection()
	user := models.UserFromContext(ctx)
	var subscriptions []models.Subscription
	db.Preload(clause.Associations).Joins("JOIN outbound_oracles ON outbound_oracles.id = subscriptions.outbound_oracle_id").Find(&subscriptions, "outbound_oracles.user_id = ?", user.ID)
	ctx.JSON(http.StatusOK, gin.H{"subscriptions": subscriptions})
}
