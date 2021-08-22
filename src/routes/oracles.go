package routes

import (
	"net/http"
	"strconv"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetOracleParameterFilters(ctx *gin.Context) {
	oracleID := ctx.Param("oracleID")
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}
	var parameterFilters []models.ParameterFilter
	db.Preload(clause.Associations).Find(&parameterFilters, "oracle_id=?", oracleID)
	ctx.JSON(http.StatusOK, gin.H{"parameterFilters": parameterFilters})
}

func PostOracleParameterFilters(ctx *gin.Context) {
	oracleID := ctx.Param("oracleID")
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	var parameterFilterBody forms.ParameterFilterBody
	if err = ctx.ShouldBind(&parameterFilterBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	parseOracleId, err := strconv.ParseUint(oracleID, 10, 32)
	parameterFilter := models.ParameterFilter{
		OracleID:         uint(parseOracleId),
		Scheme:           parameterFilterBody.Scheme,
		EventParameterID: parameterFilterBody.EventParameterID,
		FilterID:         parameterFilterBody.FilterID,
	}
	db.Create(&parameterFilter)
	ctx.JSON(http.StatusCreated, gin.H{})
}

func DeleteOracleParameterFilter(ctx *gin.Context) {
	parameterFilterID := ctx.Param("parameterFilterID")
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	db.Delete(&models.ParameterFilter{}, parameterFilterID)
	ctx.JSON(http.StatusAccepted, gin.H{})
}
