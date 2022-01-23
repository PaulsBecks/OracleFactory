package main

import (
	"fmt"
	"net/http"

	"github.com/PaulsBecks/OracleFactory/docs"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/PaulsBecks/OracleFactory/src/routes"
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin" // swagger embed files
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
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

// @title           Oracle Factory API
// @version         2.0
// @description     This is the Oracle Factory server.

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func main() {
	app := gin.Default()
	middleware(app)
	models.InitDB()

	docs.SwaggerInfo.BasePath = "/"
	app.POST("/users/login", routes.Login)
	app.POST("/users/signup", routes.Register)

	app.POST("/outboundOracles/:outboundOracleId/events", routes.PostOutboundOracleEvent)
	app.POST("/inboundOracles/:inboundOracleID/events", routes.PostInboundOracleEvent)
	app.POST("/webServiceListeners/:webServiceListenerID/events", routes.HandleWebServiceListenerEvent)
	app.POST("/smartContractListeners/:smartContractListenerID/events", routes.HandleSmartContractEvent)

	authorized := app.Group("/", auth)
	{
		authorized.GET("/outboundOracles", routes.GetOutboundOracles)
		authorized.GET("/outboundOracles/:outboundOracleId", routes.GetOutboundOracle)
		authorized.PUT("/outboundOracles/:outboundOracleId", routes.UpdateOutboundOracle)
		authorized.DELETE("/outboundOracles/:outboundOracleId", routes.DeleteOutboundOracle)
		authorized.POST("/outboundOracles/:outboundOracleId/start", routes.StartOutboundOracle)
		authorized.POST("/outboundOracles/:outboundOracleId/stop", routes.StopOutboundOracle)
		authorized.POST("/outboundOracles", routes.PostOutboundOracle)

		authorized.GET("/smartContractListeners", routes.GetSmartContractListeners)
		authorized.POST("/smartContractListeners", routes.PostSmartContractListener)
		authorized.GET("/smartContractListeners/:smartContractListenerID", routes.GetSmartContractListener)
		authorized.POST("/smartContractListeners/:smartContractListenerID/eventParameters", routes.PostOutboundEventParameters)

		authorized.GET("/inboundOracles/:inboundOracleId", routes.GetInboundOracle)
		authorized.PUT("/inboundOracles/:inboundOracleId", routes.UpdateInboundOracle)
		authorized.GET("/smartContractPublishers/:smartContractPublisherID", routes.GetSmartContractPublisher)
		authorized.GET("/inboundOracles", routes.GetInboundOracles)
		authorized.POST("/inboundOracles/:inboundOracleID/start", routes.StartInboundOracle)
		authorized.POST("/inboundOracles/:inboundOracleID/stop", routes.StopInboundOracle)
		authorized.POST("/inboundOracles", routes.PostInboundOracle)

		authorized.POST("/smartContractPublishers/:smartContractPublisherID/inboundOracles", routes.PostInboundOracle)
		authorized.POST("/smartContractPublishers/:smartContractPublisherID/eventParameters", routes.PostInboundEventParameters)
		authorized.GET("/smartContractPublishers", routes.GetSmartContractPublishers)
		authorized.POST("/smartContractPublishers", routes.PostSmartContractPublisher)

		authorized.GET("/filters", routes.GetFilters)

		authorized.GET("/oracles/:oracleID/parameterFilters", routes.GetOracleParameterFilters)
		authorized.POST("/oracles/:oracleID/parameterFilters", routes.PostOracleParameterFilters)
		authorized.DELETE("/oracles/:oracleID/parameterFilters/:parameterFilterID", routes.DeleteOracleParameterFilter)

		authorized.GET("/webServiceListeners", routes.GetWebServiceListeners)
		authorized.GET("/webServiceListeners/:webServiceListenerID", routes.GetWebServiceListener)
		authorized.POST("/webServiceListeners", routes.PostWebServiceListener)

		authorized.GET("/webServicePublishers", routes.GetWebServicePublishers)
		authorized.GET("/webServicePublishers/:webServicePublisherID", routes.GetWebServicePublisher)
		authorized.POST("/webServicePublishers", routes.PostWebServicePublisher)

		authorized.GET("/listenerPublishers/:listenerPublisherID/eventParameters", routes.GetListenerPublisherEventParameters)

		authorized.GET("/user", routes.GetCurrentUserDetail)
		authorized.PUT("/user", routes.UpdateCurrentUser)
	}
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	app.Run()
}
