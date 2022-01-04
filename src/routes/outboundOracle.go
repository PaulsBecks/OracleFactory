package routes

import (
	"fmt"
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
)

// Subscribe godoc
// @Summary      Create Event
// @Description  Subscribe your Smart Contract to a topic. This endpoint should usually only be called by blockchain listeners.
// @Tags         subscriptions
// @Param		 subscription body forms.SubscriptionBody true "The content of the subscription"
// @Param		 outboundOracleID path int true "the ID of the outboundOracle you want to send the subscription to."
// @Produce      json
// @Success      200  {object} models.Subscription
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /outboundOracle/{outboundOracleID}/subscribe [post]
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

// Unsubscribe godoc
// @Summary      Unsubscribe from Topic
// @Description  Unsubscribe your smart contract from a topic. This endpoint should usually only be called by blockchain listeners.
// @Tags         subscriptions
// @Param		 unsubscription body forms.UnsubscriptionBody true "The information to identify the unsubscription"
// @Param		 outboundOracleID path int true "the ID of the outbound oracle you want to receive the unsubscription."
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /outboundOracle/{outboundOracleID}/unsubscribe [post]
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

// StartOuboundOracle godoc
// @Summary      Start an Outbound Oracle
// @Description  Start the specified outbound oracle. This will be called from the frontend, when a user wants to use an oracle for a blockchain conenction.
// @Tags         outboundOracles
// @Param		 outboundOracleID path int true "the ID of the outbound oracle you want to start."
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /outboundOracle/{outboundOracleID}/start [post]
func StartOutboundOracle(ctx *gin.Context) {
	outboundOracleID := ctx.Param("outboundOracleID")
	log.Info(fmt.Sprintf("Start outbound oracle with"))
	outboundOracle := models.GetOutboundOracleByID(outboundOracleID)
	outboundOracle.StartOracle()
	ctx.JSON(http.StatusOK, gin.H{})
}

// StopOutboundOracle godoc
// @Summary      Stop an outbound oracle
// @Description  Stop the specified outbound oracle. This will be called from the frontend, when a user wants to stop an oracle for a blockchain conenction.
// @Tags         outboundOracles
// @Param		 outboundOracleID path int true "the ID of the outbound oracle you want to stop."
// @Produce      json
// @Success      200 {string} string "ok"
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /outboundOracle/{outboundOracleID}/stop [post]
func StopOutboundOracle(ctx *gin.Context) {
	outboundOracleID := ctx.Param("outboundOracleID")
	outboundOracle := models.GetOutboundOracleByID(outboundOracleID)
	outboundOracle.StopOracle()
	ctx.JSON(http.StatusOK, gin.H{})
}
