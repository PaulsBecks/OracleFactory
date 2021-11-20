package main

import (
	"fmt"
	"net/http"

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

	app.POST("/outboundOracles/:outboundOracleId/events", routes.PostOutboundOracleEvent)
	app.POST("/pubSubOracles/:pubSubOracleID/events", routes.PostPubSubOracleEvent)
	app.POST("/providers/:providerID/events", routes.HandleProviderEvent)

	authorized := app.Group("/", auth)
	{
		authorized.GET("/outboundOracles", routes.GetOutboundOracles)
		authorized.GET("/outboundOracles/:outboundOracleId", routes.GetOutboundOracle)
		authorized.DELETE("/outboundOracles/:outboundOracleId", routes.DeleteOutboundOracle)
		authorized.POST("/outboundOracles/:outboundOracleId/start", routes.StartOutboundOracle)
		authorized.POST("/outboundOracles/:outboundOracleId/stop", routes.StopOutboundOracle)

		authorized.GET("/blockchainEvents", routes.GetBlockchainEvents)
		authorized.POST("/blockchainEvents", routes.PostBlockchainEvent)
		authorized.GET("/blockchainEvents/:blockchainEventID", routes.GetBlockchainEvent)
		authorized.POST("/blockchainEvents/:blockchainEventID/eventParameters", routes.PostOutboundEventParameters)

		authorized.GET("/pubSubOracles/:pubSubOracleId", routes.GetPubSubOracle)
		authorized.PUT("/pubSubOracles/:pubSubOracleId", routes.UpdatePubSubOracle)
		authorized.GET("/consumers/:consumerID", routes.GetConsumer)
		authorized.GET("/pubSubOracles", routes.GetPubSubOracles)
		authorized.POST("/pubSubOracles/:pubSubOracleID/start", routes.StartPubSubOracle)
		authorized.POST("/pubSubOracles/:pubSubOracleID/stop", routes.StopPubSubOracle)
		authorized.POST("/pubSubOracles", routes.PostPubSubOracle)

		authorized.POST("/consumers/:consumerID/pubSubOracles", routes.PostPubSubOracle)
		authorized.POST("/consumers/:consumerID/eventParameters", routes.PostInboundEventParameters)
		authorized.GET("/consumers", routes.GetConsumers)
		authorized.POST("/consumers", routes.PostConsumer)

		authorized.GET("/filters", routes.GetFilters)

		authorized.GET("/oracles/:oracleID/parameterFilters", routes.GetOracleParameterFilters)
		authorized.POST("/oracles/:oracleID/parameterFilters", routes.PostOracleParameterFilters)
		authorized.DELETE("/oracles/:oracleID/parameterFilters/:parameterFilterID", routes.DeleteOracleParameterFilter)

		authorized.GET("/providers", routes.GetProviders)
		authorized.GET("/providers/:providerID", routes.GetProvider)
		authorized.POST("/providers", routes.PostProvider)

		authorized.GET("/listenerPublishers/:listenerPublisherID/eventParameters", routes.GetListenerPublisherEventParameters)

		authorized.GET("/user", routes.GetCurrentUserDetail)
		authorized.PUT("/user", routes.UpdateCurrentUser)
	}
	app.Run()
}
