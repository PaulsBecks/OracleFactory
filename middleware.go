package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func middleware(app *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	app.Use(cors.New(config))
}
