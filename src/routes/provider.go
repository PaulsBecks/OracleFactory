package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/gin-gonic/gin"
)

var mutex sync.RWMutex

func HandleProviderEvent(ctx *gin.Context) {
	// only allow one event at a time, to prevent
	mutex.Lock()
	defer mutex.Unlock()

	providerID := ctx.Param("providerID")
	provider, err := models.GetProviderByID(providerID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
		return
	}
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	provider.HandleEvent(data)
}

func GetProviders(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	ctx.JSON(http.StatusOK, gin.H{"providers": user.GetProviders()})
}

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
	ctx.JSON(http.StatusOK, gin.H{"provider": provider})
}

func PostProvider(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	var providerBody forms.ProviderBody
	if err := ctx.ShouldBind(&providerBody); err != nil || !providerBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}
	provider := user.CreateProvider(providerBody.Name, providerBody.Topic, providerBody.Description, providerBody.Private)
	ctx.JSON(http.StatusOK, gin.H{"provider": provider})
}
