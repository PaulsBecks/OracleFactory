package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/PaulsBecks/OracleFactory/src/services/ethereum"
	"github.com/PaulsBecks/OracleFactory/src/services/hyperledger"
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

	//user := models.UserFromContext(ctx)
	user := inboundOracle.GetOracle().GetUser()

	data, _ := ioutil.ReadAll(ctx.Request.Body)
	var bodyData map[string]interface{}
	if e := json.Unmarshal(data, &bodyData); e != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": e.Error()})
		return
	}

	inboundEvent := models.CreateEvent(inboundOracle.GetOracle().ID, data)

	eventValues, err := models.ParseEventValues(bodyData, *inboundEvent, inboundOracle.InboundOracleTemplate.OracleTemplateID)
	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Parameters have wrong types!"})
	}
	inboundEvent.EventValues = eventValues

	// Check if event should be filtered
	if valid := inboundOracle.GetOracle().CheckInput(bodyData); !valid {
		ctx.JSON(http.StatusAccepted, gin.H{"msg": "Event data is not passing filters!"})
		return
	}

	if inboundOracle.InboundOracleTemplate.OracleTemplate.BlockchainName == "Ethereum" {
		err = ethereum.CreateTransaction(inboundOracle, user, inboundEvent)
	}
	if inboundOracle.InboundOracleTemplate.OracleTemplate.BlockchainName == "Hyperledger" {
		err = hyperledger.CreateTransaction(inboundOracle, user, inboundEvent)
	}

	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Unable to create transaction."})
		inboundEvent.SetSuccess(false)
		return
	}

	inboundEvent.SetSuccess(true)
	ctx.JSON(http.StatusOK, gin.H{})
}

func GetInboundOracles(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	db, err := utils.DBConnection()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	var inboundOracles []models.InboundOracle
	db.Preload(clause.Associations).Preload("InboundOracleTemplate.OracleTemplate").Joins("Oracle").Find(&inboundOracles, "Oracle.user_id = ?", user.ID)

	ctx.JSON(http.StatusOK, gin.H{"inboundOracles": inboundOracles})
}

func GetInboundOracle(ctx *gin.Context) {
	id := ctx.Param("inboundOracleId")
	db, err := utils.DBConnection()
	if err != nil {
		panic(err)
	}
	var inboundOracle models.InboundOracle
	result := db.Preload("Oracle.Events.EventValues.EventParameter").Preload("InboundOracleTemplate.OracleTemplate.EventParameters").Preload(clause.Associations).First(&inboundOracle, id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No inbound Oracle with this ID available."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"inboundOracle": inboundOracle})
}

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
