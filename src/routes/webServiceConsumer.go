package routes

import (
	"fmt"
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/gin-gonic/gin"
)

// GetWebServiceConsumers godoc
// @Summary      Retrieves all web service consumers for a user
// @Description  Retrieves all web service consumers for a user. This service will be called by the frontend to retrieve all web service consumers of the user signed in.
// @Tags         webServiceConsumer
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /webServiceConsumers [get]
func GetWebServiceConsumers(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	webServiceConsumers := user.GetWebServiceConsumers()
	ctx.JSON(http.StatusOK, gin.H{"webServiceConsumers": webServiceConsumers})
}

// GetWebServiceConsumer godoc
// @Summary      Retrieves a web service consumer for a user
// @Description  Retrieves the web service consumer specified. This service will be called by the frontend to retrieve a specific consumers of the user signed in.
// @Tags         webServiceConsumer
// @Param		 webServiceConsumerID path int true "the ID of the web service consumer to send data to."
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /webServiceConsumers/{webServiceConsumerID} [get]
func GetWebServiceConsumer(ctx *gin.Context) {
	webServiceConsumerID := ctx.Param("webServiceConsumerID")
	// add check whether this is the right user
	//user := models.UserFromContext(ctx)
	webServiceConsumer, err := models.GetWebServiceConsumerByID(webServiceConsumerID)
	if err != nil {
		fmt.Printf(err.Error())
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No web service consumer with ID " + webServiceConsumerID})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"webServiceConsumer": webServiceConsumer})
}

// PostWebServiceConsumer godoc
// @Summary      Creates a web service consumers for a user
// @Description  Creates a web service consumers for a user. This service will be called by the frontend to when a user filled out the web service consumer form.
// @Tags         webServiceConsumer
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /webServiceConsumers [post]
func PostWebServiceConsumer(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	var webServiceConsumerBody forms.WebServiceConsumerBody
	if err := ctx.ShouldBind(&webServiceConsumerBody); err != nil || !webServiceConsumerBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}
	webServiceConsumer := user.CreateWebServiceConsumer(
		webServiceConsumerBody.Name,
		webServiceConsumerBody.Description,
		webServiceConsumerBody.URL,
		webServiceConsumerBody.Private,
	)
	ctx.JSON(http.StatusOK, gin.H{"webServiceConsumer": webServiceConsumer})
}
