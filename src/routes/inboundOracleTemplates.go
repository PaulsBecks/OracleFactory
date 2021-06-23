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

func PostInboundOracleTemplate(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	var inboundOracleTemplateBody forms.InboundOracleTemplateBody
	if err = ctx.ShouldBind(&inboundOracleTemplateBody); err != nil || !inboundOracleTemplateBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	inboundOracleTemplate := models.InboundOracleTemplate{
		BlockchainName:  inboundOracleTemplateBody.BlockchainName,
		ContractAddress: inboundOracleTemplateBody.ContractAddress,
		ContractName:    inboundOracleTemplateBody.ContractName,
	}
	db.Create(&inboundOracleTemplate)
	ctx.JSON(http.StatusOK, gin.H{"inboundOracleTemplate": inboundOracleTemplate})
}

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

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	var inboundOracleTemplate models.InboundOracleTemplate
	db.Preload(clause.Associations).First(&inboundOracleTemplate, inboundOracleTemplateID)
	inboundOracle := models.InboundOracle{
		InboundOracleTemplate:   inboundOracleTemplate,
		InboundOracleTemplateID: inboundOracleTemplate.ID,
		Name:                    inboundOracleBody.Name,
		User:                    user,
	}
	db.Create(&inboundOracle)
	ctx.JSON(http.StatusOK, gin.H{"inboundOracle": inboundOracle})
}

func PostInboundEventParameters(ctx *gin.Context) {
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
	inboundOracleTemplateResult := db.Preload(clause.Associations).First(&inboundOracleTemplate, inboundOracleTemplateID)
	if inboundOracleTemplateResult.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"body": "There is no inbound oracle template with this ID"})
		return
	}
	var eventParameterBody forms.EventParameterBody
	if err = ctx.ShouldBind(&eventParameterBody); err != nil || !eventParameterBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	eventParameter := models.EventParameter{
		Name:                    eventParameterBody.Name,
		Type:                    eventParameterBody.Type,
		InboundOracleTemplateID: inboundOracleTemplate.ID,
	}
	db.Create(&eventParameter)
	ctx.JSON(http.StatusOK, gin.H{"eventParameter": eventParameter})
}
