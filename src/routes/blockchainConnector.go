package routes

import (
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
)

func PostEthereumBlockchainConnector(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	var ethereumConnectorBody forms.EthereumConnectorBody
	if err := ctx.ShouldBind(&ethereumConnectorBody); err != nil || !ethereumConnectorBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}
	user.CreateEthereumConnector(ethereumConnectorBody.EthereumPrivateKey, ethereumConnectorBody.EthereumAddress)
}

func PostHyperledgerBlockchainConnector(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	var hyperledgerConnectorBody forms.HyperledgerConnectorBody
	if err := ctx.ShouldBind(&hyperledgerConnectorBody); err != nil || !hyperledgerConnectorBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}
	user.CreateHyperledgerConnector(
		hyperledgerConnectorBody.HyperledgerOrganizationName,
		hyperledgerConnectorBody.HyperledgerChannel,
		hyperledgerConnectorBody.HyperledgerConfig,
		hyperledgerConnectorBody.HyperledgerCert,
		hyperledgerConnectorBody.HyperledgerKey,
	)
}

func DeleteEthereumBlockchainConnector(ctx *gin.Context) {
	log.Info("Not implemented yet")
}

func DeleteHyperledgerBlockchainConnector(ctx *gin.Context) {
	log.Info("Not implemented yet")
}
