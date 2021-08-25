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

	authorized := app.Group("/", auth)
	{
		authorized.GET("/outboundOracles", routes.GetOutboundOracles)
		authorized.GET("/outboundOracles/:outboundOracleId", routes.GetOutboundOracle)
		authorized.PUT("/outboundOracles/:outboundOracleId", routes.UpdateOutboundOracle)
		authorized.DELETE("/outboundOracles/:outboundOracleId", routes.DeleteOutboundOracle)
		authorized.POST("/outboundOracles/:outboundOracleId/events", routes.PostOutboundOracleEvent)
		authorized.POST("/outboundOracles/:outboundOracleId/start", routes.StartOutboundOracle)
		authorized.POST("/outboundOracles/:outboundOracleId/stop", routes.StopOutboundOracle)

		authorized.GET("/outboundOracleTemplates", routes.GetOutboundOracleTemplates)
		authorized.POST("/outboundOracleTemplates", routes.PostOutboundOracleTemplate)
		authorized.GET("/outboundOracleTemplates/:outboundOracleTemplateID", routes.GetOutboundOracleTemplate)
		authorized.POST("/outboundOracleTemplates/:outboundOracleTemplateID/outboundOracles", routes.PostOutboundOracle)
		authorized.POST("/outboundOracleTemplates/:outboundOracleTemplateID/eventParameters", routes.PostOutboundEventParameters)

		authorized.GET("/inboundOracles/:inboundOracleId", routes.GetInboundOracle)
		authorized.PUT("/inboundOracles/:inboundOracleId", routes.UpdateInboundOracle)
		authorized.POST("/inboundOracles/:inboundOracleID/events", routes.PostInboundOracleEvent)
		authorized.GET("/inboundOracleTemplates/:inboundOracleTemplateID", routes.GetInboundOracleTemplate)
		authorized.GET("/inboundOracles", routes.GetInboundOracles)
		authorized.POST("/inboundOracles/:inboundOracleID/start", routes.StartInboundOracle)
		authorized.POST("/inboundOracles/:inboundOracleID/stop", routes.StopInboundOracle)

		authorized.POST("/inboundOracleTemplates/:inboundOracleTemplateID/inboundOracles", routes.PostInboundOracle)
		authorized.POST("/inboundOracleTemplates/:inboundOracleTemplateID/eventParameters", routes.PostInboundEventParameters)
		authorized.GET("/inboundOracleTemplates", routes.GetInboundOracleTemplates)
		authorized.POST("/inboundOracleTemplates", routes.PostInboundOracleTemplate)

		authorized.GET("/filters", routes.GetFilters)

		authorized.GET("/oracles/:oracleID/parameterFilters", routes.GetOracleParameterFilters)
		authorized.POST("/oracles/:oracleID/parameterFilters", routes.PostOracleParameterFilters)
		authorized.DELETE("/oracles/:oracleID/parameterFilters/:parameterFilterID", routes.DeleteOracleParameterFilter)

		authorized.GET("/oracleTemplates/:oracleTemplateID/eventParameters", routes.GetOracleTemplateEventParameters)

		authorized.GET("/user", routes.GetCurrentUserDetail)
		authorized.PUT("/user", routes.UpdateCurrentUser)
	}
	app.Run()
}
