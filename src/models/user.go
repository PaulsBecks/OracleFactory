package models

import (
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

func (u *User) CreateWebServicePublisher(name string, description string, url string, private bool) WebServicePublisher {
	db := utils.DBConnection()
	listenerPublisher := ListenerPublisher{
		Name:        name,
		Description: description,
		Private:     private,
		UserID:      u.ID,
		User:        *u,
	}
	db.Create(&listenerPublisher)
	webServicePublisher := WebServicePublisher{
		ListenerPublisher: listenerPublisher,
		Url:               url,
	}
	db.Create(&webServicePublisher)
	return webServicePublisher
}

func (u *User) CreateWebServiceListener(name string, description string, private bool) WebServiceListener {
	db := utils.DBConnection()
	listenerPublisher := ListenerPublisher{
		Name:        name,
		Description: description,
		Private:     private,
		UserID:      u.ID,
		User:        *u,
	}
	db.Create(&listenerPublisher)
	webServiceListener := WebServiceListener{
		ListenerPublisher: listenerPublisher,
	}
	db.Create(&webServiceListener)
	return webServiceListener
}

func (u *User) CreateInboundOracle(name string, smartContractPublisherID, webServiceListenerID uint) InboundOracle {
	db := utils.DBConnection()
	ethereumOracle := Oracle{
		Name:   name,
		UserID: u.ID,
	}
	db.Create(&ethereumOracle)
	inboundOracle := InboundOracle{
		Oracle:                   ethereumOracle,
		SmartContractPublisherID: smartContractPublisherID,
		WebServiceListenerID:     webServiceListenerID,
	}
	db.Create(&inboundOracle)
	return inboundOracle
}

func (u *User) CreateOutboundOracle(name string, smartContractListenerID, webServicePublisherID uint) *OutboundOracle {
	db := utils.DBConnection()
	oracle := Oracle{
		Name: name,
		User: *u,
	}
	db.Create(&oracle)
	outboundOracle := &OutboundOracle{
		SmartContractListenerID: smartContractListenerID,
		WebServicePublisherID:   webServicePublisherID,
		Oracle:                  oracle,
	}
	db.Create(&outboundOracle)
	return outboundOracle
}

func (u *User) CreateSmartContractPublisher(blockchainName string, eventName string, contractAddress string, contractAddressSynonym string, publisherName string, description string, private bool, eventParameters map[string]string) SmartContractPublisher {
	db := utils.DBConnection()
	smartContract := SmartContract{
		BlockchainName:         blockchainName,
		EventName:              eventName,
		ContractAddress:        contractAddress,
		ContractAddressSynonym: contractAddressSynonym,
	}
	db.Create(&smartContract)
	listenerPublisher := ListenerPublisher{
		Name:        publisherName,
		Description: description,
		Private:     private,
		UserID:      u.ID,
		User:        *u,
	}
	db.Create(&listenerPublisher)
	smartContractPublisher := SmartContractPublisher{
		SmartContract:     smartContract,
		ListenerPublisher: listenerPublisher,
	}
	db.Create(&smartContractPublisher)
	for parameterName, parameterType := range eventParameters {
		eventParameter := EventParameter{
			Name:                parameterName,
			Type:                parameterType,
			ListenerPublisherID: listenerPublisher.ID,
		}
		db.Create(&eventParameter)
	}
	return smartContractPublisher
}
func (u *User) CreateSmartContractListener(blockchainName string, eventName string, contractAddress string, contractAddressSynonym string, listenerName string, description string, private bool, eventParameters map[string]string) SmartContractListener {
	db := utils.DBConnection()
	smartContract := SmartContract{
		BlockchainName:         blockchainName,
		EventName:              eventName,
		ContractAddress:        contractAddress,
		ContractAddressSynonym: contractAddressSynonym,
	}
	db.Create(&smartContract)
	listenerPublisher := ListenerPublisher{
		Name:        listenerName,
		Description: description,
		Private:     private,
		UserID:      u.ID,
		User:        *u,
	}
	db.Create(&listenerPublisher)
	smartContractListener := SmartContractListener{
		SmartContract:     smartContract,
		ListenerPublisher: listenerPublisher,
	}
	db.Create(&smartContractListener)
	for parameterName, parameterType := range eventParameters {
		eventParameter := EventParameter{
			Name:                parameterName,
			Type:                parameterType,
			ListenerPublisherID: listenerPublisher.ID,
		}
		db.Create(&eventParameter)
	}
	return smartContractListener
}

func (u *User) GetWebServiceListeners() []WebServiceListener {
	db := utils.DBConnection()
	var webServiceListener []WebServiceListener
	db.Joins("ListenerPublisher").Find(&webServiceListener, "ListenerPublisher.private = 0 OR ListenerPublisher.user_id = ?", u.ID)
	return webServiceListener
}

func (u *User) GetWebServicePublishers() []WebServicePublisher {
	db := utils.DBConnection()
	var webServicePublisher []WebServicePublisher
	db.Joins("ListenerPublisher").Find(&webServicePublisher, "ListenerPublisher.private = 0 OR ListenerPublisher.user_id = ?", u.ID)
	return webServicePublisher
}
