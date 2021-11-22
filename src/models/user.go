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

func (u *User) CreateProvider(name string, description string, private bool) Provider {
	db := utils.DBConnection()
	listenerPublisher := ListenerPublisher{
		Name:        name,
		Description: description,
		Private:     private,
		UserID:      u.ID,
		User:        *u,
	}
	db.Create(&listenerPublisher)
	provider := Provider{
		ListenerPublisher: listenerPublisher,
	}
	db.Create(&provider)
	return provider
}

func (u *User) CreatePubSubOracle(name string, consumerID, providerID, subOracleID, unsubOracleID uint) PubSubOracle {
	db := utils.DBConnection()
	ethereumOracle := Oracle{
		Name:   name,
		UserID: u.ID,
	}
	db.Create(&ethereumOracle)
	pubSubOracle := PubSubOracle{
		Oracle:        ethereumOracle,
		ConsumerID:    consumerID,
		ProviderID:    providerID,
		SubOracleID:   subOracleID,
		UnsubOracleID: unsubOracleID,
	}
	db.Create(&pubSubOracle)
	return pubSubOracle
}

func (u *User) CreateOutboundOracle(name string, blockchainEventID uint, isSubscribing bool) *OutboundOracle {
	db := utils.DBConnection()
	oracle := Oracle{
		Name: name,
		User: *u,
	}
	db.Create(&oracle)
	outboundOracle := &OutboundOracle{
		BlockchainEventID: blockchainEventID,
		Oracle:            oracle,
		IsSubscribing:     isSubscribing,
	}
	db.Create(&outboundOracle)
	return outboundOracle
}

func (u *User) CreateConsumer(blockchainName string, eventName string, contractAddress string, contractAddressSynonym string, publisherName string, description string, private bool, eventParameters []NameTypePair) Consumer {
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
	consumer := Consumer{
		SmartContract:     smartContract,
		ListenerPublisher: listenerPublisher,
	}
	db.Create(&consumer)
	for _, nameType := range eventParameters {
		fmt.Println(nameType)
		eventParameter := EventParameter{
			Name:                nameType.Name,
			Type:                nameType.Type,
			ListenerPublisherID: listenerPublisher.ID,
		}
		db.Create(&eventParameter)
	}
	return consumer
}
func (u *User) CreateBlockchainEvent(blockchainName string, eventName string, contractAddress string, contractAddressSynonym string, listenerName string, description string, private bool, eventParameters []NameTypePair) BlockchainEvent {
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
	blockchainEvent := BlockchainEvent{
		SmartContract:     smartContract,
		ListenerPublisher: listenerPublisher,
	}
	db.Create(&blockchainEvent)
	for _, nameType := range eventParameters {
		eventParameter := EventParameter{
			Name:                nameType.Name,
			Type:                nameType.Type,
			ListenerPublisherID: listenerPublisher.ID,
		}
		db.Create(&eventParameter)
	}
	return blockchainEvent
}

func (u *User) GetProviders() []Provider {
	db := utils.DBConnection()
	var provider []Provider
	db.Joins("ListenerPublisher").Find(&provider, "ListenerPublisher.private = 0 OR ListenerPublisher.user_id = ?", u.ID)
	return provider
}
