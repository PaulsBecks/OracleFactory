package models

import (
	"fmt"
	"time"

	"github.com/PaulsBecks/OracleFactory/src/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Claims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	Email                       string
	Password                    string
	EthereumPrivateKey          string
	EthereumPublicKey           string
	EthereumAddress             string
	HyperledgerConfig           string
	HyperledgerCert             string
	HyperledgerKey              string
	HyperledgerOrganizationName string
	HyperledgerChannel          string
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

func (u *User) CreateWebServiceConsumer(name string, description string, url string, private bool) WebServiceConsumer {
	db := utils.DBConnection()
	providerConsumer := ProviderConsumer{
		Name:        name,
		Description: description,
		Private:     private,
		UserID:      u.ID,
		User:        *u,
	}
	db.Create(&providerConsumer)
	webServiceConsumer := WebServiceConsumer{
		ProviderConsumer: providerConsumer,
		Url:              url,
	}
	db.Create(&webServiceConsumer)
	return webServiceConsumer
}

func (u *User) CreateWebServiceProvider(name string, description string, private bool) WebServiceProvider {
	db := utils.DBConnection()
	providerConsumer := ProviderConsumer{
		Name:        name,
		Description: description,
		Private:     private,
		UserID:      u.ID,
		User:        *u,
	}
	db.Create(&providerConsumer)
	webServiceProvider := WebServiceProvider{
		ProviderConsumer: providerConsumer,
	}
	db.Create(&webServiceProvider)
	return webServiceProvider
}

func (u *User) CreateInboundSubscription(name string, smartContractConsumerID, webServiceProviderID uint) InboundSubscription {
	db := utils.DBConnection()
	ethereumSubscription := Subscription{
		Name:   name,
		UserID: u.ID,
	}
	db.Create(&ethereumSubscription)
	inboundSubscription := InboundSubscription{
		Subscription:            ethereumSubscription,
		SmartContractConsumerID: smartContractConsumerID,
		WebServiceProviderID:    webServiceProviderID,
	}
	db.Create(&inboundSubscription)
	return inboundSubscription
}

func (u *User) CreateOutboundSubscription(name string, smartContractProviderID, webServiceConsumerID uint) *OutboundSubscription {
	db := utils.DBConnection()
	subscription := Subscription{
		Name: name,
		User: *u,
	}
	db.Create(&subscription)
	outboundSubscription := &OutboundSubscription{
		SmartContractProviderID: smartContractProviderID,
		WebServiceConsumerID:    webServiceConsumerID,
		Subscription:            subscription,
	}
	db.Create(&outboundSubscription)
	return outboundSubscription
}

func (u *User) CreateSmartContractConsumer(blockchainName string, eventName string, contractAddress string, contractAddressSynonym string, consumerName string, description string, private bool, eventParameters []NameTypePair) SmartContractConsumer {
	db := utils.DBConnection()
	smartContract := SmartContract{
		BlockchainName:         blockchainName,
		EventName:              eventName,
		ContractAddress:        contractAddress,
		ContractAddressSynonym: contractAddressSynonym,
	}
	db.Create(&smartContract)
	providerConsumer := ProviderConsumer{
		Name:        consumerName,
		Description: description,
		Private:     private,
		UserID:      u.ID,
		User:        *u,
	}
	db.Create(&providerConsumer)
	smartContractConsumer := SmartContractConsumer{
		SmartContract:    smartContract,
		ProviderConsumer: providerConsumer,
	}
	db.Create(&smartContractConsumer)
	for _, nameType := range eventParameters {
		fmt.Println(nameType)
		eventParameter := EventParameter{
			Name:               nameType.Name,
			Type:               nameType.Type,
			ProviderConsumerID: providerConsumer.ID,
		}
		db.Create(&eventParameter)
	}
	return smartContractConsumer
}
func (u *User) CreateSmartContractProvider(blockchainName string, eventName string, contractAddress string, contractAddressSynonym string, providerName string, description string, private bool, eventParameters []NameTypePair) SmartContractProvider {
	db := utils.DBConnection()
	smartContract := SmartContract{
		BlockchainName:         blockchainName,
		EventName:              eventName,
		ContractAddress:        contractAddress,
		ContractAddressSynonym: contractAddressSynonym,
	}
	db.Create(&smartContract)
	providerConsumer := ProviderConsumer{
		Name:        providerName,
		Description: description,
		Private:     private,
		UserID:      u.ID,
		User:        *u,
	}
	db.Create(&providerConsumer)
	smartContractProvider := SmartContractProvider{
		SmartContract:    smartContract,
		ProviderConsumer: providerConsumer,
	}
	db.Create(&smartContractProvider)
	for _, nameType := range eventParameters {
		eventParameter := EventParameter{
			Name:               nameType.Name,
			Type:               nameType.Type,
			ProviderConsumerID: providerConsumer.ID,
		}
		db.Create(&eventParameter)
	}
	return smartContractProvider
}

func (u *User) GetWebServiceProviders() []WebServiceProvider {
	db := utils.DBConnection()
	var webServiceProvider []WebServiceProvider
	db.Joins("ProviderConsumer").Find(&webServiceProvider, "ProviderConsumer.private = 0 OR ProviderConsumer.user_id = ?", u.ID)
	return webServiceProvider
}

func (u *User) GetWebServiceConsumers() []WebServiceConsumer {
	db := utils.DBConnection()
	var webServiceConsumer []WebServiceConsumer
	db.Joins("ProviderConsumer").Find(&webServiceConsumer, "ProviderConsumer.private = 0 OR ProviderConsumer.user_id = ?", u.ID)
	return webServiceConsumer
}
