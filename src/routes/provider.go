package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/lock"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/gin-gonic/gin"
)

type ProviderResponse struct {
	Provider models.Provider
}

type ProvidersResponse struct {
	Providers []models.Provider
}

// Register godoc
// @Summary      Create Event
// @Description  Create a new event to be handled and propagated to subscribers.
// @Tags         events
// @Param		 eventContent body object true "The content of the event"
// @Param		 providerID path int true "the ID of the provider you want to send the event to."
// @Produce      json
// @Success      200  {object}  ProvidersResponse
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /provider/{providerID} [get]
func HandleProviderEvent(ctx *gin.Context) {
	// only allow one event at a time, to prevent
	lock.PipeLock.Lock()

	providerID := ctx.Param("providerID")
	provider, err := models.GetProviderByID(providerID)
	if err != nil {
		lock.PipeLock.Unlock() // Unlock on failure
		ctx.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
		return
	}
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	provider.HandleEvent(data) // Otherwise unlock inhere
}

// Register godoc
// @Summary      Get Providers
// @Description  Get all providers for registered user.
// @Tags         providers
// @Produce      json
// @Success      200  {object}  ProvidersResponse
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /providers [get]
func GetProviders(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	ctx.JSON(http.StatusOK, ProvidersResponse{Providers: user.GetProviders()})
}

// Register godoc
// @Summary      Get Providers
// @Description  Get a provider for registered user.
// @Tags         providers
// @Param		 providerID path int true "the ID of the provider you want to retrieve"
// @Produce      json
// @Success      200  {object}  ProvidersResponse
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      404  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /provider/{providerID} [get]
func GetProvider(ctx *gin.Context) {
	providerID := ctx.Param("providerID")
	// add check whether this is the right user
	//user := models.UserFromContext(ctx)
	provider, err := models.GetProviderByID(providerID)
	if err != nil {
		fmt.Printf(err.Error())
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No provider with ID " + providerID})
		return
	}
	ctx.JSON(http.StatusOK, ProviderResponse{Provider: provider})
}

// Register godoc
// @Summary      Create Provider
// @Description  Create a providers for registered user.
// @Tags         providers
// @Produce      json
// @Param		 auth body forms.ProviderBody true "provider to be created"
// @Success      200  {object}  ProvidersResponse
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /providers [post]
func PostProvider(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	var providerBody forms.ProviderBody
	if err := ctx.ShouldBind(&providerBody); err != nil || !providerBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}
	provider := user.CreateProvider(providerBody.Name, providerBody.Topic, providerBody.Description, providerBody.Private)
	ctx.JSON(http.StatusOK, ProviderResponse{Provider: provider})
}
