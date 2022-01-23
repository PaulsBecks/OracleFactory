package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/gin-gonic/gin"
)

// HandleWebServiceListenerEvent godoc
// @Summary      Handles the event send from a web service provider
// @Description  Handles the event send from a web service provider. This endpoint will be called from an external service, that provides data to the artifact.
// @Tags         webServiceListener
// @Param		 webServiceListenerID path int true "the ID of the web service listener to send data to."
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /webServiceListeners/{webServiceListenerID}/events [post]
func HandleWebServiceListenerEvent(ctx *gin.Context) {
	webServiceListenerID := ctx.Param("webServiceListenerID")
	webServiceListener, err := models.GetWebServiceListenerByID(webServiceListenerID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
		return
	}
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	webServiceListener.HandleEvent(data)
}

// GetWebServiceListeners godoc
// @Summary      Retrieves all web service listeners for a user
// @Description  Retrieves all web service listeners for a user. This service will be called by the frontend to retrieve all web service listeners of the user signed in.
// @Tags         webServiceListener
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /webServiceListeners [get]
func GetWebServiceListeners(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	webServiceListeners := user.GetWebServiceListeners()
	ctx.JSON(http.StatusOK, gin.H{"webServiceListeners": webServiceListeners})
}

// GetWebServiceListener godoc
// @Summary      Retrieves a web service listener for a user
// @Description  Retrieves the web service listener specified. This service will be called by the frontend to retrieve a specific listeners of the user signed in.
// @Tags         webServiceListener
// @Param		 webServiceListenerID path int true "the ID of the web service listener to send data to."
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /webServiceListeners/{webServiceListenerID} [get]
func GetWebServiceListener(ctx *gin.Context) {
	webServiceListenerID := ctx.Param("webServiceListenerID")
	// add check whether this is the right user
	//user := models.UserFromContext(ctx)
	webServiceListener, err := models.GetWebServiceListenerByID(webServiceListenerID)
	if err != nil {
		fmt.Printf(err.Error())
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No web service listener with ID " + webServiceListenerID})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"webServiceListener": webServiceListener})
}

// PostWebServiceListener godoc
// @Summary      Creates a web service listeners for a user
// @Description  Creates a web service listeners for a user. This service will be called by the frontend to when a user filled out the web service listener form.
// @Tags         webServiceListener
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /webServiceListeners [post]
func PostWebServiceListener(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	var eventParameterBody forms.WebServiceListenerBody
	if err := ctx.ShouldBind(&eventParameterBody); err != nil || !eventParameterBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}
	webServiceListener := user.CreateWebServiceListener(eventParameterBody.Name, eventParameterBody.Description, eventParameterBody.Private)
	ctx.JSON(http.StatusOK, gin.H{"webServiceListener": webServiceListener})
}
