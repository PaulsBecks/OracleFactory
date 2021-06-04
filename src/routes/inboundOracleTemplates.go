package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetInboundOracleTemplate(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	inboundOracleTemplateID := ctx.Param("inboundOracleTemplateID")
	i, err := strconv.Atoi(inboundOracleTemplateID)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid oracle id!"})
		return
	}

	var inboundOracleTemplate models.InboundOracleTemplate
	db.Preload(clause.Associations).First(&inboundOracleTemplate, i)

	ctx.JSON(http.StatusOK, gin.H{"inboundOracleTemplate": inboundOracleTemplate})
}

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

	var inboundOracleBody forms.InboundOracleBody
	if err = ctx.ShouldBind(&inboundOracleBody); err != nil || !inboundOracleBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	var inboundOracleTemplate models.InboundOracleTemplate
	db.Preload(clause.Associations).First(&inboundOracleTemplate, inboundOracleTemplateID)
	inboundOracle := models.InboundOracle{
		InboundOracleTemplate:   inboundOracleTemplate,
		InboundOracleTemplateID: inboundOracleTemplate.ID,
		Name:                    inboundOracleBody.Name,
	}
	db.Create(&inboundOracle)
	ctx.JSON(http.StatusOK, gin.H{"inboundOracle": inboundOracle})
}
