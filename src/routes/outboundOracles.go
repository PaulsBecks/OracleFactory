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

// GetOutboundOracles godoc
// @Summary      Retrieves all outbound oracle of a user
// @Description  Retrieve all outbound oracles of a user. This will be called from the frontend, when a user wants to view a list of oracle.
// @Tags         outboundOracles
// @Produce      json
// @Router       /outboundOracles [get]
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

// GetOutboundOracle godoc
// @Summary      Retrieve an outbound oracle
// @Description  Retrieve the specified outbound oracle. This will be called from the frontend, when a user wants to view an oracle.
// @Tags         outboundOracles
// @Param		 outboundOracleID path int true "the ID of the outbound oracle you want to retrieve."
// @Produce      json
// @Router       /outboundOracles/{outboundOracleID} [get]
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

// UpdateOutboundOracle godoc
// @Summary      Update an outbound oracle
// @Description  Update the specified outbound oracle. This will be called from the frontend, when a user wants to update an oracle.
// @Tags         outboundOracles
// @Param		 outboundOracleID path int true "the ID of the outbound oracle you want to update."
// @Produce      json
// @Router       /outboundOracles/{outboundOracleID} [put]
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

// StartOutboundOracle godoc
// @Summary      Start an Outbound Oracle
// @Description  Start the specified outbound oracle. This will be called from the frontend, when a user wants to use an oracle for a blockchain conenction.
// @Tags         outboundOracles
// @Param		 outboundOracleID path int true "the ID of the outbound oracle you want to start."
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /outboundOracles/{outboundOracleID}/start [post]
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

// StopOutboundOracle godoc
// @Summary      Stop an outbound oracle
// @Description  Stop the specified outbound oracle. This will be called from the frontend, when a user wants to stop an oracle for a blockchain conenction.
// @Tags         outboundOracles
// @Param		 outboundOracleID path int true "the ID of the outbound oracle you want to stop."
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /outboundOracles/{outboundOracleID}/stop [post]
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

// PostOutboundOracle godoc
// @Summary      Creates an outbound oracle for a user
// @Description  Creates an outbound oracle for a user. This service will be called by the frontend to when a user filled out the outbound oracle form.
// @Tags         outboundOracles
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /outboundOracles [post]
func PostOutboundOracle(ctx *gin.Context) {
	var outboundOraclePostBody forms.OutboundOraclePostBody
	if err := ctx.ShouldBind(&outboundOraclePostBody); err != nil || !outboundOraclePostBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	user := models.UserFromContext(ctx)
	outboundOracle := user.CreateOutboundOracle(
		outboundOraclePostBody.Oracle.Name,
		outboundOraclePostBody.SmartContractListenerID,
		outboundOraclePostBody.WebServicePublisherID,
	)
	ctx.JSON(http.StatusOK, gin.H{"outboundOracle": outboundOracle})
}
