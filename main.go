package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/PaulsBecks/OracleFactory/src/routes"
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func auth(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	authHeader := ctx.GetHeader("Authorization")

	if !strings.HasPrefix(authHeader, "Bearer") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"body": "No credentials provided."})
		return
	}
	tokenString := authHeader[len("Bearer "):]
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, isvalid := t.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", t.Header["alg"])
		}
		return []byte(utils.JWT_SECRET), nil
	})
	if err == nil && token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["id"].(float64))
		var user models.User
		db.First(&user, userID)
		ctx.Set("user", user)
	} else {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}

func main() {
	app := gin.Default()
	middleware(app)
	models.InitDB()

	app.POST("/users/login", routes.Login)
	app.POST("/users/signup", routes.Register)

	app.POST("/outboundOracles/:outboundOracleID/subscribe", routes.PostSubscription)
	app.POST("/outboundOracles/:outboundOracleID/unsubscribe", routes.PostUnsbscription)
	app.POST("/providers/:providerID/events", routes.HandleProviderEvent)

	authorized := app.Group("/", auth)
	{
		authorized.GET("/subscriptions", routes.GetSubscriptions)
		authorized.POST("/subscriptions", routes.PostSubscription)
		authorized.GET("/subscriptions/:subscriptionID", routes.GetSubscription)

		authorized.GET("/providers", routes.GetProviders)
		authorized.GET("/providers/:providerID", routes.GetProvider)
		authorized.POST("/providers", routes.PostProvider)

		authorized.GET("/ethereumConnectors", routes.GetEthereumConnectors)
		authorized.POST("/ethereumConnectors", routes.PostEthereumBlockchainConnector)

		authorized.GET("/hyperledgerConnectors", routes.GetHyperledgerConnectors)
		authorized.POST("/hyperledgerConnectors", routes.PostHyperledgerBlockchainConnector)

		authorized.POST("/outboundOracles/:outboundOracleID/start", routes.StartOutboundOracle)
		authorized.POST("/outboundOracles/:outboundOracleID/stop", routes.StopOutboundOracle)

		authorized.GET("/user", routes.GetCurrentUserDetail)
		authorized.PUT("/user", routes.UpdateCurrentUser)
	}
	app.Run()
}
