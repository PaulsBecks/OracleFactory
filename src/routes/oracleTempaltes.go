package routes

import (
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetOracleTemplateEventParameters(ctx *gin.Context) {
	oracleTemplateID := ctx.Param("oracleTemplateID")
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}
	var eventParameters []models.EventParameter
	db.Find(&eventParameters, "oracle_template_id=?", oracleTemplateID)
	ctx.JSON(http.StatusOK, gin.H{"eventParameters": eventParameters})
}
