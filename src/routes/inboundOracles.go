package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func PostInboundOracleEvent(ctx *gin.Context) {
	inboundOracleID := ctx.Param("inboundOracleID")
	inboundOracle, err := models.GetInboundOracleByID(inboundOracleID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No valid oracle id!"})
		return
	}
	fmt.Printf("Event submitted for inbound oracle %s\n", inboundOracleID)
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	inboundOracle.HandleEvent(data)
	ctx.JSON(http.StatusOK, gin.H{})
}

// GetInboundOracles godoc
// @Summary      Retrieves all inbound oracle of a user
// @Description  Retrieve all inbound oracles of a user. This will be called from the frontend, when a user wants to view a list of oracle.
// @Tags         inboundOracles
// @Produce      json
// @Router       /inboundOracles [get]
func GetInboundOracles(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	db := utils.DBConnection()

	var inboundOracles []models.InboundOracle
	db.Preload(clause.Associations).Preload("SmartContractPublisher.SmartContract").Preload("SmartContractPublisher.ListenerPublisher").Preload("WebServiceListener.ListenerPublisher").Joins("Oracle").Find(&inboundOracles, "Oracle.user_id = ?", user.ID)

	ctx.JSON(http.StatusOK, gin.H{"inboundOracles": inboundOracles})
}

// GetInboundOracle godoc
// @Summary      Retrieve an inbound oracle
// @Description  Retrieve the specified inbound oracle. This will be called from the frontend, when a user wants to view an oracle.
// @Tags         inboundOracles
// @Param		 inboundOracleID path int true "the ID of the inbound oracle you want to retrieve."
// @Produce      json
// @Router       /inboundOracles/{inboundOracleID} [get]
func GetInboundOracle(ctx *gin.Context) {
	id := ctx.Param("inboundOracleId")
	db := utils.DBConnection()

	var inboundOracle models.InboundOracle
	result := db.Preload("Oracle.Events.EventValues.EventParameter").Preload("SmartContractPublisher.ListenerPublisher.EventParameters").Preload(clause.Associations).First(&inboundOracle, id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No inbound Oracle with this ID available."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"inboundOracle": inboundOracle})
}

// UpdateInboundOracle godoc
// @Summary      Update an inbound oracle
// @Description  Update the specified inbound oracle. This will be called from the frontend, when a user wants to update an oracle.
// @Tags         inboundOracles
// @Param		 inboundOracleID path int true "the ID of the inbound oracle you want to update."
// @Produce      json
// @Router       /inboundOracles/{inboundOracleID} [put]
func UpdateInboundOracle(ctx *gin.Context) {
	id := ctx.Param("inboundOracleId")
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var inboundOracle models.InboundOracle
	result := db.Preload(clause.Associations).First(&inboundOracle, id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No inbound Oracle with this ID available."})
		return
	}
	var inboundOraclePostBody forms.InboundOracleBody
	if err = ctx.ShouldBind(&inboundOraclePostBody); err != nil || !inboundOraclePostBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	oracle := inboundOracle.Oracle
	oracle.Name = inboundOraclePostBody.Oracle.Name

	db.Save(&oracle)
	ctx.JSON(http.StatusOK, gin.H{"inboundOracle": inboundOracle})
}

// StartInboundOracle godoc
// @Summary      Start an Outbound Oracle
// @Description  Start the specified inbound oracle. This will be called from the frontend, when a user wants to use an oracle for a blockchain conenction.
// @Tags         inboundOracles
// @Param		 inboundOracleID path int true "the ID of the inbound oracle you want to start."
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /inboundOracles/{inboundOracleID}/start [post]
func StartInboundOracle(ctx *gin.Context) {
	id := ctx.Param("inboundOracleID")
	inboundOracle, err := models.GetInboundOracleByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No inbound Oracle with this ID available."})
		return
	}
	inboundOracle.GetOracle().Start()
	ctx.JSON(http.StatusOK, gin.H{"msg": "Oracle got started successfully."})
}

// StopInboundOracle godoc
// @Summary      Stop an inbound oracle
// @Description  Stop the specified inbound oracle. This will be called from the frontend, when a user wants to stop an oracle for a blockchain conenction.
// @Tags         inboundOracles
// @Param		 inboundOracleID path int true "the ID of the inbound oracle you want to stop."
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /inboundOracles/{inboundOracleID}/stop [post]
func StopInboundOracle(ctx *gin.Context) {
	id := ctx.Param("inboundOracleID")
	inboundOracle, err := models.GetInboundOracleByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No inbound Oracle with this ID available."})
		return
	}
	inboundOracle.GetOracle().Stop()
	ctx.JSON(http.StatusOK, gin.H{"msg": "Oracle got stopped successfully."})
}

// PostInboundOracle godoc
// @Summary      Creates an inbound oracle for a user
// @Description  Creates an inbound oracle for a user. This service will be called by the frontend to when a user filled out the inbound oracle form.
// @Tags         inboundOracles
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /inboundOracles [post]
func PostInboundOracle(ctx *gin.Context) {
	var inboundOracleBody forms.InboundOracleBody
	if err := ctx.ShouldBind(&inboundOracleBody); err != nil || !inboundOracleBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	user := models.UserFromContext(ctx)
	inboundOracle := user.CreateInboundOracle(
		inboundOracleBody.Oracle.Name,
		inboundOracleBody.SmartContractPublisherID,
		inboundOracleBody.WebServiceListenerID,
	)
	ctx.JSON(http.StatusOK, gin.H{"inboundOracle": inboundOracle})
}
