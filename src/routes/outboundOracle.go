package routes

import (
	"fmt"
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
)

func PostSubscription(ctx *gin.Context) {
	var subscriptionBody forms.SubscriptionBody
	if err := ctx.ShouldBind(&subscriptionBody); err != nil || !subscriptionBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	//user := models.UserFromContext(ctx)
	outboundOracleID := ctx.Param("outboundOracleID")

	outboundOracle := models.GetOutboundOracleByID(outboundOracleID)
	subscription := subscriptionBody.CreateSubscription(outboundOracle)
	ctx.JSON(http.StatusOK, gin.H{"subscription": subscription})
}

func PostUnsbscription(ctx *gin.Context) {
	var unsubscriptionBody forms.UnsubscriptionBody
	if err := ctx.ShouldBind(&unsubscriptionBody); err != nil || !unsubscriptionBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	//user := models.UserFromContext(ctx)
	outboundOracleID := ctx.Param("outboundOracleID")

	outboundOracle := models.GetOutboundOracleByID(outboundOracleID)
	unsubscriptionBody.DeleteSubscription(outboundOracle)
	ctx.JSON(http.StatusOK, gin.H{})
}

func StartOutboundOracle(ctx *gin.Context) {
	outboundOracleID := ctx.Param("outboundOracleID")
	log.Info(fmt.Sprintf("Start outbound oracle with"))
	outboundOracle := models.GetOutboundOracleByID(outboundOracleID)
	outboundOracle.StartOracle()
}

func StopOutboundOracle(ctx *gin.Context) {
	outboundOracleID := ctx.Param("outboundOracleID")
	outboundOracle := models.GetOutboundOracleByID(outboundOracleID)
	outboundOracle.StopOracle()
}
