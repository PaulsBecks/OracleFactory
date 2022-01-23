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

// GetCurrentUserDetail godoc
// @Summary      Retrieve signed in user
// @Description  Retrieve the signed in user. This will be called by the frontend to get all information about the user signed in.
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.TokenResponse
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /user [get]
func GetCurrentUserDetail(ctx *gin.Context) {
	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

// UpdateCurrentUser godoc
// @Summary      Update User
// @Description  Update a user. This will be called from the frontend to update the settings
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.TokenResponse
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /user [put]
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

	user.EthereumAddress = userBody.EthereumAddress
	user.EthereumPublicKey = userBody.EthereumPublicKey
	user.EthereumPrivateKey = userBody.EthereumPrivateKey
	user.HyperledgerCert = userBody.HyperledgerCert
	user.HyperledgerChannel = userBody.HyperledgerChannel
	user.HyperledgerConfig = userBody.HyperledgerConfig
	user.HyperledgerOrganizationName = userBody.HyperledgerOrganizationName
	user.HyperledgerKey = userBody.HyperledgerKey

	fmt.Println(user)

	db.Save(&user)
}

// Login godoc
// @Summary      Login User
// @Description  Login a user - get access token.
// @Tags         users
// @Param		 auth body forms.AuthBody true "auth to register"
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.TokenResponse
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /users/login [post]
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

// Register godoc
// @Summary      Register User
// @Description  Register a new user - get access token in return.
// @Tags         users
// @Accept       json
// @Param		 auth body forms.AuthBody true "auth to register"
// @Produce      json
// @Success      200  {object}  responses.TokenResponse
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /users/signup [post]
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
