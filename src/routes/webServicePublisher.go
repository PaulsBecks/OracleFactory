package routes

import (
	"fmt"
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/gin-gonic/gin"
)

func GetWebServicePublishers(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	webServicePublishers := user.GetWebServicePublishers()
	ctx.JSON(http.StatusOK, gin.H{"webServicePublishers": webServicePublishers})
}

func GetWebServicePublisher(ctx *gin.Context) {
	webServicePublisherID := ctx.Param("webServicePublisherID")
	// add check whether this is the right user
	//user := models.UserFromContext(ctx)
	webServicePublisher, err := models.GetWebServicePublisherByID(webServicePublisherID)
	if err != nil {
		fmt.Printf(err.Error())
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No web service listener with ID " + webServicePublisherID})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"webServicePublisher": webServicePublisher})
}

func PostWebServicePublisher(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	var webServicePublisherBody forms.WebServicePublisherBody
	if err := ctx.ShouldBind(&webServicePublisherBody); err != nil || !webServicePublisherBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}
	webServicePublisher := user.CreateWebServicePublisher(
		webServicePublisherBody.Name,
		webServicePublisherBody.Description,
		webServicePublisherBody.URL,
		webServicePublisherBody.Private,
	)
	ctx.JSON(http.StatusOK, gin.H{"webServicePublisher": webServicePublisher})
}
