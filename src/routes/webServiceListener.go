package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/gin-gonic/gin"
)

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

func GetWebServiceListeners(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	webServiceListeners := user.GetWebServiceListeners()
	ctx.JSON(http.StatusOK, gin.H{"webServiceListeners": webServiceListeners})
}

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
