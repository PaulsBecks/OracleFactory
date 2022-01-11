package routes

import (
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

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	var oracles []models.OutboundOracle
	db.Preload(clause.Associations).Preload("SmartContractListener.SmartContract").Preload("SmartContractListener.ListenerPublisher").Preload("WebServicePublisher.ListenerPublisher").Joins("Oracle").Find(&oracles, "Oracle.user_id = ?", user.ID)
	fmt.Println(oracles)

	ctx.JSON(http.StatusOK, gin.H{"outboundOracles": oracles})
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
	result := db.Preload("Oracle.Events.EventValues.EventParameter").Preload("SmartContractListener.ListenerPublisher.EventParameters").Preload("Oracle.ParameterFilters.Filter").Preload(clause.Associations).First(&outboundOracle, i)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No outbound Oracle with this ID available."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"outboundOracle": outboundOracle})
}

func UpdateOutboundOracle(ctx *gin.Context) {
	id := ctx.Param("outboundOracleId")
	outboundOracle, err := models.GetOutboundOracleById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No outbound Oracle with this ID available."})
		return
	}
	var outboundOraclePostBody forms.OutboundOraclePostBody
	if err = ctx.ShouldBind(&outboundOraclePostBody); err != nil || !outboundOraclePostBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}
	outboundOracle.Save()

	oracle := outboundOracle.Oracle
	oracle.Name = outboundOraclePostBody.Oracle.Name
	oracle.Save()

	ctx.JSON(http.StatusOK, gin.H{"outboundOracle": outboundOracle})
}

func DeleteOutboundOracle(ctx *gin.Context) {
	// TODO: delete oracle for the provided id
	ctx.JSON(http.StatusNotImplemented, gin.H{"body": "Hi there, deletion is not implemented yet!"})
}

func StartOutboundOracle(ctx *gin.Context) {
	id := ctx.Param("outboundOracleId")
	outboundOracle, err := models.GetOutboundOracleById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No outbound Oracle with this ID available."})
		return
	}
	err = outboundOracle.StartOracle()
	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Unable to start oracle, try again later."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Oracle got started successfully."})
}

func StopOutboundOracle(ctx *gin.Context) {
	id := ctx.Param("outboundOracleId")
	outboundOracle, err := models.GetOutboundOracleById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No outbound Oracle with this ID available."})
		return
	}
	fmt.Print(outboundOracle)
	err = outboundOracle.StopOracle()
	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Unable to stop oracle, try again later."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Oracle got stopped successfully."})
}

func PostOutboundOracleEvent(ctx *gin.Context) {
	id := ctx.Param("outboundOracleId")
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid oracle id!"})
		return
	}

	outboundOracle, _ := models.GetOutboundOracleById(i)

	data, _ := ioutil.ReadAll(ctx.Request.Body)
	var bodyData map[string]interface{}
	if err := json.Unmarshal(data, &bodyData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	event := models.CreateEvent(data, outboundOracle.GetOracle().ID)
	event.ParseEventValues(bodyData, outboundOracle.GetSmartContractListener().ListenerPublisherID)
	// TODO: Filter outbound oracle event.

	webServicePublisher := outboundOracle.GetWebServicePublisher()
	webServicePublisher.Publish(*event)
}