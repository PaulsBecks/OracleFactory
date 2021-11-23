package routes

import (
	"fmt"
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	fmt.Println(user)

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
	/*
		user.EthereumAddress = userBody.EthereumAddress
		user.EthereumPublicKey = userBody.EthereumPublicKey
		user.EthereumPrivateKey = userBody.EthereumPrivateKey
		user.HyperledgerCert = userBody.HyperledgerCert
		user.HyperledgerChannel = userBody.HyperledgerChannel
		user.HyperledgerConfig = userBody.HyperledgerConfig
		user.HyperledgerOrganizationName = userBody.HyperledgerOrganizationName
		user.HyperledgerKey = userBody.HyperledgerKey

		fmt.Println(user)
	*/
	db.Save(&user)
}

func Login(ctx *gin.Context) {

	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	var authBody forms.AuthBody
	if err = ctx.ShouldBind(&authBody); err != nil || !authBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	var user models.User

	result := db.Find(&user, "email=?", authBody.Email)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No user found with email or password."})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authBody.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No user found with email or password."})
		return
	}

	token := user.GetJWT()
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func Register(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	var authBody forms.AuthBody
	if err = ctx.ShouldBind(&authBody); err != nil || !authBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	var user models.User

	result := db.Find(&user, "email=?", authBody.Email)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "Email already in DB"})
		return
	}

	user = models.User{
		Email:    authBody.Email,
		Password: utils.HashAndSalt([]byte(authBody.Password)),
	}

	db.Create(&user)
	token := user.GetJWT()
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
