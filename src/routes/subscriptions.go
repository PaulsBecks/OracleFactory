package routes

import (
	"net/http"
	"strconv"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetSubscriptionParameterFilters(ctx *gin.Context) {
	subscriptionID := ctx.Param("subscriptionID")
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}
	var parameterFilters []models.ParameterFilter
	db.Preload(clause.Associations).Find(&parameterFilters, "subscription_id=?", subscriptionID)
	ctx.JSON(http.StatusOK, gin.H{"parameterFilters": parameterFilters})
}

func PostSubscriptionParameterFilters(ctx *gin.Context) {
	subscriptionID := ctx.Param("subscriptionID")
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	var parameterFilterBody forms.ParameterFilterBody
	if err = ctx.ShouldBind(&parameterFilterBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	parseSubscriptionId, err := strconv.ParseUint(subscriptionID, 10, 32)
	parameterFilter := models.ParameterFilter{
		SubscriptionID:   uint(parseSubscriptionId),
		Scheme:           parameterFilterBody.Scheme,
		EventParameterID: parameterFilterBody.EventParameterID,
		FilterID:         parameterFilterBody.FilterID,
	}
	db.Create(&parameterFilter)
	ctx.JSON(http.StatusCreated, gin.H{})
}

func DeleteSubscriptionParameterFilter(ctx *gin.Context) {
	parameterFilterID := ctx.Param("parameterFilterID")
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	db.Delete(&models.ParameterFilter{}, parameterFilterID)
	ctx.JSON(http.StatusAccepted, gin.H{})
}
