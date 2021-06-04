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

	outboundOracleTemplateID := ctx.Param("outboundOracleTemplateId")
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

	outboundOracle := &models.OutboundOracle{
		OutboundOracleTemplate:   outboundOracleTemplate,
		OutboundOracleTemplateID: outboundOracleTemplate.ID,
		URI:                      outboundOraclePostBody.URI,
		Name:                     outboundOraclePostBody.Name,
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
