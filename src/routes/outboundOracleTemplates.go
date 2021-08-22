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

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	var outboundOracleTemplate models.OutboundOracleTemplate
	db.Joins("OracleTemplate").Preload("OracleTemplate.EventParameters").Preload("OutboundOracles.Oracle").Find(&outboundOracleTemplate, i)

	var outboundOracles []models.OutboundOracle
	db.Joins("Oracle").Preload("OutboundOracleTemplate.OracleTemplate").Find(&outboundOracles, "outbound_oracle_template_id = ? AND Oracle.user_id = ?", outboundOracleTemplate.ID, user.ID)
	outboundOracleTemplate.OutboundOracles = outboundOracles

	ctx.JSON(http.StatusOK, gin.H{"outboundOracleTemplate": outboundOracleTemplate})
}

func GetOutboundOracleTemplates(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	var outboundOracleTemplates []models.OutboundOracleTemplate
	db.Joins("OracleTemplate").Preload(clause.Associations).Find(&outboundOracleTemplates, "OracleTemplate.private = 0 OR OracleTemplate.user_id = ?", user.ID)

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

	oracle := models.Oracle{
		Name: outboundOraclePostBody.Oracle.Name,
		User: user,
	}

	outboundOracle := &models.OutboundOracle{
		OutboundOracleTemplate:   outboundOracleTemplate,
		OutboundOracleTemplateID: outboundOracleTemplate.ID,
		URI:                      outboundOraclePostBody.URI,
		Oracle:                   oracle,
	}

	db.Create(&outboundOracle)

	outboundOracle.StartOracle()

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

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	oracleTemplate := models.OracleTemplate{
		BlockchainName:  outboundOracleTemplateBody.BlockchainName,
		ContractAddress: outboundOracleTemplateBody.ContractAddress,
		EventName:       outboundOracleTemplateBody.EventName,
		UserID:          user.ID,
		Private:         outboundOracleTemplateBody.Private,
	}
	db.Create(&oracleTemplate)

	outboundOracleTemplate := models.OutboundOracleTemplate{OracleTemplate: oracleTemplate}
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
		Name:             eventParameterBody.Name,
		Type:             eventParameterBody.Type,
		OracleTemplateID: outboundOracleTemplate.OracleTemplate.ID,
	}
	db.Create(&eventParameter)
	ctx.JSON(http.StatusOK, gin.H{"eventParameter": eventParameter})
}
