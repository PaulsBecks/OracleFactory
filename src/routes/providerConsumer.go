package routes

import (
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetProviderConsumerEventParameters(ctx *gin.Context) {
	providerConsumerID := ctx.Param("providerConsumerID")
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}
	var eventParameters []models.EventParameter
	db.Find(&eventParameters, "provider_consumer_id=?", providerConsumerID)
	ctx.JSON(http.StatusOK, gin.H{"eventParameters": eventParameters})
}
