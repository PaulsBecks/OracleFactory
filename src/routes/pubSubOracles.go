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

func PostPubSubOracleEvent(ctx *gin.Context) {
	pubSubOracleID := ctx.Param("pubSubOracleID")
	pubSubOracle, err := models.GetPubSubOracleByID(pubSubOracleID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No valid oracle id!"})
		return
	}
	fmt.Printf("Event submitted for pubSub oracle %s\n", pubSubOracleID)
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	pubSubOracle.HandleEvent(data)
	ctx.JSON(http.StatusOK, gin.H{})
}

func GetPubSubOracles(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	db := utils.DBConnection()

	var pubSubOracles []models.PubSubOracle
	db.Preload(clause.Associations).Preload("Consumer.SmartContract").Preload("Consumer.ListenerPublisher").Preload("Provider.ListenerPublisher").Preload("SubOracle.BlockchainEvent").Preload("UnsubOracle.BlockchainEvent").Joins("Oracle").Find(&pubSubOracles, "Oracle.user_id = ?", user.ID)

	ctx.JSON(http.StatusOK, gin.H{"pubSubOracles": pubSubOracles})
}

func GetPubSubOracle(ctx *gin.Context) {
	id := ctx.Param("pubSubOracleId")
	db := utils.DBConnection()

	var pubSubOracle models.PubSubOracle
	result := db.Preload("Consumer.SmartContract").Preload("Provider.ListenerPublisher").Preload("Oracle.Events.EventValues.EventParameter").Preload("Consumer.ListenerPublisher.EventParameters").Preload("SubOracle.BlockchainEvent.ListenerPublisher").Preload("UnsubOracle.BlockchainEvent.ListenerPublisher").Preload(clause.Associations).First(&pubSubOracle, id)
	subOracle, _ := models.GetOutboundOracleById(pubSubOracle.SubOracleID)
	unsubOracle, _ := models.GetOutboundOracleById(pubSubOracle.UnsubOracleID)
	pubSubOracle.SubOracle = *subOracle
	pubSubOracle.UnsubOracle = *unsubOracle
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No pubSub Oracle with this ID available."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"pubSubOracle": pubSubOracle})
}

func UpdatePubSubOracle(ctx *gin.Context) {
	id := ctx.Param("pubSubOracleId")
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var pubSubOracle models.PubSubOracle
	result := db.Preload(clause.Associations).First(&pubSubOracle, id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No pubSub Oracle with this ID available."})
		return
	}
	var pubSubOraclePostBody forms.PubSubOracleBody
	if err = ctx.ShouldBind(&pubSubOraclePostBody); err != nil || !pubSubOraclePostBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	oracle := pubSubOracle.Oracle
	oracle.Name = pubSubOraclePostBody.Oracle.Name

	db.Save(&oracle)
	ctx.JSON(http.StatusOK, gin.H{"pubSubOracle": pubSubOracle})
}

func StartPubSubOracle(ctx *gin.Context) {
	id := ctx.Param("pubSubOracleID")
	pubSubOracle, err := models.GetPubSubOracleByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No pubSub Oracle with this ID available."})
		return
	}
	pubSubOracle.GetOracle().Start()
	ctx.JSON(http.StatusOK, gin.H{"msg": "Oracle got started successfully."})
}

func StopPubSubOracle(ctx *gin.Context) {
	id := ctx.Param("pubSubOracleID")
	pubSubOracle, err := models.GetPubSubOracleByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No pubSub Oracle with this ID available."})
		return
	}
	pubSubOracle.GetOracle().Stop()
	ctx.JSON(http.StatusOK, gin.H{"msg": "Oracle got stopped successfully."})
}

func PostPubSubOracle(ctx *gin.Context) {
	var pubSubOracleBody forms.PubSubOracleBody
	if err := ctx.ShouldBind(&pubSubOracleBody); err != nil || !pubSubOracleBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	user := models.UserFromContext(ctx)
	subOracle := user.CreateOutboundOracle("Subscribe", pubSubOracleBody.SubBlockchainEventID)
	unsubOracle := user.CreateOutboundOracle("Unsubscribe", pubSubOracleBody.UnsubBlockchainEventID)
	pubSubOracle := user.CreatePubSubOracle(
		pubSubOracleBody.Oracle.Name,
		pubSubOracleBody.ConsumerID,
		pubSubOracleBody.ProviderID,
		subOracle.ID,
		unsubOracle.ID,
	)
	ctx.JSON(http.StatusOK, gin.H{"pubSubOracle": pubSubOracle})
}
