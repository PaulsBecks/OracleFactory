package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/gin-gonic/gin"
)

// HandleWebServiceProviderEvent godoc
// @Summary      Handles the event send from a web service provider
// @Description  Handles the event send from a web service provider. This endpoint will be called from an external service, that provides data to the artifact.
// @Tags         webServiceProvider
// @Param		 webServiceProviderID path int true "the ID of the web service provider to send data to."
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /webServiceProviders/{webServiceProviderID}/events [post]
func HandleWebServiceProviderEvent(ctx *gin.Context) {
	webServiceProviderID := ctx.Param("webServiceProviderID")
	webServiceProvider, err := models.GetWebServiceProviderByID(webServiceProviderID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
		return
	}
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	webServiceProvider.HandleEvent(data)
}

// GetWebServiceProviders godoc
// @Summary      Retrieves all web service providers for a user
// @Description  Retrieves all web service providers for a user. This service will be called by the frontend to retrieve all web service providers of the user signed in.
// @Tags         webServiceProvider
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /webServiceProviders [get]
func GetWebServiceProviders(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	webServiceProviders := user.GetWebServiceProviders()
	ctx.JSON(http.StatusOK, gin.H{"webServiceProviders": webServiceProviders})
}

// GetWebServiceProvider godoc
// @Summary      Retrieves a web service provider for a user
// @Description  Retrieves the web service provider specified. This service will be called by the frontend to retrieve a specific providers of the user signed in.
// @Tags         webServiceProvider
// @Param		 webServiceProviderID path int true "the ID of the web service provider to send data to."
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /webServiceProviders/{webServiceProviderID} [get]
func GetWebServiceProvider(ctx *gin.Context) {
	webServiceProviderID := ctx.Param("webServiceProviderID")
	// add check whether this is the right user
	//user := models.UserFromContext(ctx)
	webServiceProvider, err := models.GetWebServiceProviderByID(webServiceProviderID)
	if err != nil {
		fmt.Printf(err.Error())
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No web service provider with ID " + webServiceProviderID})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"webServiceProvider": webServiceProvider})
}

// PostWebServiceProvider godoc
// @Summary      Creates a web service providers for a user
// @Description  Creates a web service providers for a user. This service will be called by the frontend to when a user filled out the web service provider form.
// @Tags         webServiceProvider
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /webServiceProviders [post]
func PostWebServiceProvider(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	var eventParameterBody forms.WebServiceProviderBody
	if err := ctx.ShouldBind(&eventParameterBody); err != nil || !eventParameterBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}
	webServiceProvider := user.CreateWebServiceProvider(eventParameterBody.Name, eventParameterBody.Description, eventParameterBody.Private)
	ctx.JSON(http.StatusOK, gin.H{"webServiceProvider": webServiceProvider})
}
