package routes

import (
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetListenerPublisherEventParameters(ctx *gin.Context) {
	listenerPublisherID := ctx.Param("listenerPublisherID")
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}
	var eventParameters []models.EventParameter
	db.Find(&eventParameters, "listener_publisher_id=?", listenerPublisherID)
	ctx.JSON(http.StatusOK, gin.H{"eventParameters": eventParameters})
}
