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

func GetOutboundOracleTemplate(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	outboundOracleTemplateID := ctx.Param("outboundOracleTemplateID")
	i, err := strconv.Atoi(outboundOracleTemplateID)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid oracle id!"})
		return
	}

	var outboundOracleTemplate models.OutboundOracleTemplate
	db.Preload(clause.Associations).First(&outboundOracleTemplate, i)

	ctx.JSON(http.StatusOK, gin.H{"outboundOracleTemplate": outboundOracleTemplate})
}

func GetOutboundOracleTemplates(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	var outboundOracleTemplates []models.OutboundOracleTemplate
	db.Preload(clause.Associations).Find(&outboundOracleTemplates)

	fmt.Println(outboundOracleTemplates)

	ctx.JSON(http.StatusOK, gin.H{"outboundOracleTemplates": outboundOracleTemplates})
}

func PostOutboundOracle(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var outboundOraclePostBody forms.OutboundOraclePostBody
	if err = ctx.ShouldBind(&outboundOraclePostBody); err != nil || !outboundOraclePostBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	outboundOracleTemplateID := ctx.Param("outboundOracleTemplateID")
	var outboundOracleTemplate models.OutboundOracleTemplate
	i, err := strconv.Atoi(outboundOracleTemplateID)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid oracle id!"})
		return
	}
	result := db.Preload(clause.Associations).First(&outboundOracleTemplate, i)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No valid oracle id!"})
		return
	}

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	outboundOracle := &models.OutboundOracle{
		OutboundOracleTemplate:   outboundOracleTemplate,
		OutboundOracleTemplateID: outboundOracleTemplate.ID,
		URI:                      outboundOraclePostBody.URI,
		Name:                     outboundOraclePostBody.Name,
		User:                     user,
	}

	db.Create(&outboundOracle)

	manifest := outboundOracle.CreateManifest()
	err = manifest.Run()

	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Something bad happened :("})
		return
	} else {
		fmt.Println("INFO: new listener started")
	}

	ctx.JSON(http.StatusOK, gin.H{"outboundOracle": outboundOracle})
}

func PostOutboundOracleTemplate(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	var outboundOracleTemplateBody forms.OutboundOracleTemplateBody
	if err = ctx.ShouldBind(&outboundOracleTemplateBody); err != nil || !outboundOracleTemplateBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	outboundOracleTemplate := models.OutboundOracleTemplate{
		BlockchainAddress: outboundOracleTemplateBody.BlockchainAddress,
		Blockchain:        outboundOracleTemplateBody.BlockchainName,
		Address:           outboundOracleTemplateBody.ContractAddress,
		EventName:         outboundOracleTemplateBody.EventName,
	}
	db.Create(&outboundOracleTemplate)
	ctx.JSON(http.StatusOK, gin.H{"outboundOracleTemplate": outboundOracleTemplate})
}

func PostOutboundEventParameters(ctx *gin.Context) {
	outboundOracleTemplateIDString := ctx.Param("outboundOracleTemplateID")
	outboundOracleTemplateID, err := strconv.Atoi(outboundOracleTemplateIDString)
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

	var outboundOracleTemplate models.OutboundOracleTemplate
	outboundOracleTemplateResult := db.Preload(clause.Associations).First(&outboundOracleTemplate, outboundOracleTemplateID)
	if outboundOracleTemplateResult.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"body": "There is no outbound oracle template with this ID"})
		return
	}
	var eventParameterBody forms.EventParameterBody
	if err = ctx.ShouldBind(&eventParameterBody); err != nil || !eventParameterBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	eventParameter := models.EventParameter{
		Name:                     eventParameterBody.Name,
		Type:                     eventParameterBody.Type,
		OutboundOracleTemplateID: outboundOracleTemplate.ID,
	}
	db.Create(&eventParameter)
	ctx.JSON(http.StatusOK, gin.H{"eventParameter": eventParameter})
}
