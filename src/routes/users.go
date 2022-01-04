package routes

import (
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/PaulsBecks/OracleFactory/src/responses"
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TokenResponse struct {
	Token string
}

type ErrorResponse struct {
	Body string
}

func GetCurrentUserDetail(ctx *gin.Context) {
	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)
	ctx.JSON(http.StatusOK, gin.H{"user": user})
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
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Body: "Ups there was a mistake!"})
		return
	}

	var authBody forms.AuthBody
	if err = ctx.ShouldBind(&authBody); err != nil || !authBody.Valid() {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{Body: "No valid body send!"})
		return
	}

	var user models.User

	result := db.Find(&user, "email=?", authBody.Email)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{Body: "No user found with email or password."})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authBody.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{Body: "No user found with email or password."})
		return
	}

	token := user.GetJWT()
	ctx.JSON(http.StatusOK, responses.TokenResponse{Token: token})
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
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Body: "Ups there was a mistake!"})
		return
	}

	var authBody forms.AuthBody
	if err = ctx.ShouldBind(&authBody); err != nil || !authBody.Valid() {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{Body: "No valid body send!"})
		return
	}

	var user models.User

	result := db.Find(&user, "email=?", authBody.Email)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{Body: "Email already in DB"})
		return
	}

	user = models.User{
		Email:    authBody.Email,
		Password: utils.HashAndSalt([]byte(authBody.Password)),
	}

	db.Create(&user)
	token := user.GetJWT()
	ctx.JSON(http.StatusOK, responses.TokenResponse{Token: token})
}
