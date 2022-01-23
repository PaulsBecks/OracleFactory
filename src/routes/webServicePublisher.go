package routes

import (
	"fmt"
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/gin-gonic/gin"
)

// GetWebServicePublishers godoc
// @Summary      Retrieves all web service publishers for a user
// @Description  Retrieves all web service publishers for a user. This service will be called by the frontend to retrieve all web service publishers of the user signed in.
// @Tags         webServicePublisher
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /webServicePublishers [get]
func GetWebServicePublishers(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	webServicePublishers := user.GetWebServicePublishers()
	ctx.JSON(http.StatusOK, gin.H{"webServicePublishers": webServicePublishers})
}

// GetWebServicePublisher godoc
// @Summary      Retrieves a web service publisher for a user
// @Description  Retrieves the web service publisher specified. This service will be called by the frontend to retrieve a specific publishers of the user signed in.
// @Tags         webServicePublisher
// @Param		 webServicePublisherID path int true "the ID of the web service publisher to send data to."
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /webServicePublishers/{webServicePublisherID} [get]
func GetWebServicePublisher(ctx *gin.Context) {
	webServicePublisherID := ctx.Param("webServicePublisherID")
	// add check whether this is the right user
	//user := models.UserFromContext(ctx)
	webServicePublisher, err := models.GetWebServicePublisherByID(webServicePublisherID)
	if err != nil {
		fmt.Printf(err.Error())
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No web service publisher with ID " + webServicePublisherID})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"webServicePublisher": webServicePublisher})
}

// PostWebServicePublisher godoc
// @Summary      Creates a web service publishers for a user
// @Description  Creates a web service publishers for a user. This service will be called by the frontend to when a user filled out the web service publisher form.
// @Tags         webServicePublisher
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /webServicePublishers [post]
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
