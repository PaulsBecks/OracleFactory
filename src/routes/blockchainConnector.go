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
	ethereumConnector := user.CreateEthereumConnector(ethereumConnectorBody.IsOnChain, ethereumConnectorBody.EthereumPrivateKey, ethereumConnectorBody.EthereumAddress)
	ctx.JSON(http.StatusOK, gin.H{"ethereumConnector": ethereumConnector})
}

func PostHyperledgerBlockchainConnector(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	var hyperledgerConnectorBody forms.HyperledgerConnectorBody
	if err := ctx.ShouldBind(&hyperledgerConnectorBody); err != nil || !hyperledgerConnectorBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}
	hyperledgerConnector := user.CreateHyperledgerConnector(
		hyperledgerConnectorBody.IsOnChain,
		hyperledgerConnectorBody.HyperledgerOrganizationName,
		hyperledgerConnectorBody.HyperledgerChannel,
		hyperledgerConnectorBody.HyperledgerConfig,
		hyperledgerConnectorBody.HyperledgerCert,
		hyperledgerConnectorBody.HyperledgerKey,
	)
	ctx.JSON(http.StatusOK, gin.H{"hyperledgerConnector": hyperledgerConnector})

}

func GetEthereumConnectors(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	ethereumConnectors := user.GetEthereumConnectors()
	ctx.JSON(http.StatusOK, gin.H{"ethereumConnectors": ethereumConnectors})
}

func GetEthereumConnector(ctx *gin.Context) {
	ethereumConnectorID := ctx.Param("ethereumConnectorID")
	ethereumConnector := models.GetEthereumConnectorByID(ethereumConnectorID)
	ctx.JSON(http.StatusOK, gin.H{"ethereumConnector": ethereumConnector})
}

func GetHyperledgerConnectors(ctx *gin.Context) {
	user := models.UserFromContext(ctx)
	hyperledgerConnectors := user.GetHyperledgerConnectors()
	ctx.JSON(http.StatusOK, gin.H{"hyperledgerConnectors": hyperledgerConnectors})
}

func GetHyperledgerConnector(ctx *gin.Context) {
	hyperledgerConnectorID := ctx.Param("hyperledgerConnectorID")
	hyperledgerConnector := models.GetHyperledgerConnectorByID(hyperledgerConnectorID)
	ctx.JSON(http.StatusOK, gin.H{"hyperledgerConnector": hyperledgerConnector})
}

func DeleteEthereumBlockchainConnector(ctx *gin.Context) {
	log.Info("Not implemented yet")
}

func DeleteHyperledgerBlockchainConnector(ctx *gin.Context) {
	log.Info("Not implemented yet")
}
