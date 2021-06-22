package routes

import (
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetCurrentUserDetail(ctx *gin.Context) {
	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateCurrentUser(ctx *gin.Context) {
	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	var userBody forms.UserBody
	if err = ctx.ShouldBind(&userBody); err != nil || !userBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	user.EthereumAddress = userBody.EthereumAddress
	user.EthereumPublicKey = userBody.EthereumPublicKey
	user.EthereumPrivateKey = userBody.EthereumPrivateKey
	user.HyperledgerCert = userBody.HyperledgerCert
	user.HyperledgerChannel = userBody.HyperledgerChannel
	user.HyperledgerConfig = userBody.HyperledgerConfig
	user.HyperledgerOrganizationName = userBody.HyperledgerOrganizationName
	user.HyperledgerKey = userBody.HyperledgerKey

	db.Save(&user)
}
