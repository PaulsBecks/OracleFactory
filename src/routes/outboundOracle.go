package routes

import (
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/gin-gonic/gin"
)

func PostSubscription(ctx *gin.Context) {
	var subscriptionBody forms.SubscriptionBody
	if err := ctx.ShouldBind(&subscriptionBody); err != nil || !subscriptionBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	//user := models.UserFromContext(ctx)
	blockchainConnectorID := ctx.Param("blockchainConnectorID")

	outboundOracle := models.GetOutboundOracleByID(blockchainConnectorID)
	subscription := subscriptionBody.CreateSubscription(outboundOracle)
	ctx.JSON(http.StatusOK, gin.H{"subscription": subscription})
}
