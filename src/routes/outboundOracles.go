package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetOutboundOracles(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var outboundOracles []models.OutboundOracle
	db.Preload(clause.Associations).Find(&outboundOracles)
	fmt.Println(outboundOracles)

	ctx.JSON(http.StatusOK, gin.H{"outboundOracles": outboundOracles})
}

func GetOutboundOracle(ctx *gin.Context) {
	id := ctx.Param("outboundOracleId")
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid oracle id!"})
		return
	}

	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var outboundOracle models.OutboundOracle
	result := db.Preload("OutboundEvents.EventValues.EventParameter").Preload(clause.Associations).First(&outboundOracle, i)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No outbound Oracle with this ID available."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"outboundOracle": outboundOracle})
}

func UpdateOutboundOracle(ctx *gin.Context) {
	id := ctx.Param("outboundOracleId")
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid oracle id!"})
		return
	}

	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var outboundOracle models.OutboundOracle
	result := db.First(&outboundOracle, i)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No outbound Oracle with this ID available."})
		return
	}
	var outboundOraclePostBody forms.OutboundOraclePostBody
	if err = ctx.ShouldBind(&outboundOraclePostBody); err != nil || !outboundOraclePostBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	outboundOracle.Name = outboundOraclePostBody.Name
	outboundOracle.URI = outboundOraclePostBody.URI

	db.Save(&outboundOracle)
	ctx.JSON(http.StatusOK, gin.H{"outboundOracle": outboundOracle})
}

func DeleteOutboundOracle(ctx *gin.Context) {
	// TODO: delete oracle for the provided id
	ctx.JSON(http.StatusOK, gin.H{"body": "Hi there, deletion is not implemented yet!"})
}

func PostOutboundOracleEvent(ctx *gin.Context) {
	id := ctx.Param("outboundOracleId")
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid oracle id!"})
		return
	}

	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ups there was a mistake!"})
		return
	}

	var outboundOracle models.OutboundOracle
	db.First(&outboundOracle, i)

	fmt.Println(outboundOracle)

	// TODO: Filter outbound oracle event.
	outboundEvent := &models.OutboundEvent{
		OutboundOracle:   outboundOracle,
		OutboundOracleID: outboundOracle.ID,
	}
	db.Create(&outboundEvent)

	data, _ := ioutil.ReadAll(ctx.Request.Body)
	var bodyData map[string]interface{}
	if err := json.Unmarshal(data, &bodyData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	fmt.Println(bodyData)

	for key, value := range bodyData {
		var eventParameter models.EventParameter
		// Add .Where("OutboundOracleTemplateID = ?", outboundOracle.OutboundOracleTemplateID)
		db.Where("Name = ?", key).First(&eventParameter)
		sValue := fmt.Sprintf("%v", value)
		fmt.Println(key, sValue)
		eventValue := models.EventValue{
			EventParameterID: eventParameter.ID,
			Value:            sValue,
			OutboundEventID:  outboundEvent.ID,
		}
		fmt.Println(eventValue)
		db.Create(&eventValue)
	}

	output, err := json.Marshal(bodyData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	fmt.Println("INFO: post data to: " + outboundOracle.URI)
	http.Post(outboundOracle.URI, "application/json", bytes.NewBuffer(output))
}
