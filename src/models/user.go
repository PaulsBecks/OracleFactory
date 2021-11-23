package models

import (
	"time"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Claims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	Email    string
	Password string
}

func (user *User) GetJWT() string {
	expirationTime := time.Now().Add(5 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(utils.JWT_SECRET))
	return tokenString
}

func UserFromContext(ctx *gin.Context) User {
	userInterface, _ := ctx.Get("user")
	return userInterface.(User)
}

func (u *User) CreateProvider(name string, topic string, description string, private bool) Provider {
	db := utils.DBConnection()
	provider := Provider{
		Name:        name,
		Topic:       topic,
		Description: description,
		Private:     private,
	}
	db.Create(&provider)
	return provider
}

func (u *User) CreateEthereumConnector(ethereumPrivateKey, ethereumAddress string) *EthereumConnector {
	db := utils.DBConnection()
	ethereumConnector := &EthereumConnector{
		OutboundOracle:     *u.CreateOutboundOracle(),
		EthereumPrivateKey: ethereumPrivateKey,
		EthereumAddress:    ethereumAddress,
	}
	db.Create(ethereumConnector)
	return ethereumConnector
}

func (u *User) GetEthereumConnectors() []EthereumConnector {
	db := utils.DBConnection()
	var ethereumConnectors []EthereumConnector
	db.Preload(clause.Associations).Joins("JOIN outbound_oracles ON outbound_oracles.id = ethereum_connectors.outbound_oracle_id").Where("outbound_oracles.user_id = ?", u.ID).Find(&ethereumConnectors)
	return ethereumConnectors
}

func (u *User) CreateHyperledgerConnector(orgName, channel, config, cert, key string) *HyperledgerConnector {
	db := utils.DBConnection()
	hyperledgerConnector := &HyperledgerConnector{
		OutboundOracle:              *u.CreateOutboundOracle(),
		HyperledgerOrganizationName: orgName,
		HyperledgerChannel:          channel,
		HyperledgerConfig:           config,
		HyperledgerCert:             cert,
		HyperledgerKey:              key,
	}
	db.Create(hyperledgerConnector)
	return hyperledgerConnector

}

func (u *User) GetHyperledgerConnectors() []HyperledgerConnector {
	db := utils.DBConnection()
	var hyperledgerConnectors []HyperledgerConnector
	//.Joins("JOIN outbound_oracles ON outbound_oracles.id = hyperledger_connectors.outbound_oracle_id").Where("outbound_oracles.user_id = ?", u.ID)
	db.Preload(clause.Associations).Joins("JOIN outbound_oracles ON outbound_oracles.id = hyperledger_connectors.outbound_oracle_id").Where("outbound_oracles.user_id = ?", u.ID).Find(&hyperledgerConnectors)
	return hyperledgerConnectors
}

func (u *User) CreateOutboundOracle() *OutboundOracle {
	outboundOracle := &OutboundOracle{
		UserID: u.ID,
	}
	return outboundOracle
}

func (u *User) GetProviders() []Provider {
	db := utils.DBConnection()
	var provider []Provider
	db.Find(&provider, "private = 0 OR user_id = ?", u.ID)
	return provider
}
