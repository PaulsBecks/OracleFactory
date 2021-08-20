package models

import (
	"time"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Claims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	Email                       string
	Password                    string
	EthereumPrivateKey          string
	EthereumPublicKey           string
	EthereumAddress             string
	HyperledgerConfig           string
	HyperledgerCert             string
	HyperledgerKey              string
	HyperledgerOrganizationName string
	HyperledgerChannel          string
}

func (user *User) GetJWT() string {
	expirationTime := time.Now().Add(5 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(utils.JWT_SECRET))
	return tokenString
}

func UserFromContext(ctx *gin.Context) User {
	userInterface, _ := ctx.Get("user")
	return userInterface.(User)
}
