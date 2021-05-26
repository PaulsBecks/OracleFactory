package routes

import (
	"fmt"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
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
}
