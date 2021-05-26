package routes

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strconv"
	"gorm.io/gorm/clause"

)

func GetInboundOracleTemplates(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	var inboundOracleTemplates []models.InboundOracleTemplate
	db.Preload(clause.Associations).Find(&inboundOracleTemplates)

	fmt.Println(inboundOracleTemplates)

	ctx.JSON(http.StatusOK, gin.H{"inboundOracleTemplates": inboundOracleTemplates})
}

func PostInboundOracle(ctx *gin.Context) {
	inboundOracleTemplateIDString := ctx.Param("inboundOracleTemplateID")
	inboundOracleTemplateID, err := strconv.Atoi(inboundOracleTemplateIDString)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid oracle id!"})
		return
	}
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	var inboundOracleTemplate models.InboundOracleTemplate
	db.Preload(clause.Associations).First(&inboundOracleTemplate, inboundOracleTemplateID)
	inboundOracle := models.InboundOracle{
		InboundOracleTemplate: inboundOracleTemplate,
		InboundOracleTemplateID: inboundOracleTemplate.ID,
	}
	db.Create(&inboundOracle)
	ctx.JSON(http.StatusOK, gin.H{"inboundOracle": inboundOracle})
}