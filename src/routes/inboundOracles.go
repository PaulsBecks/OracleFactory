package routes

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func PostInboundOracleEvent(ctx *gin.Context) {
	inboundOracleID := ctx.Param("inboundOracleID")
	var inboundOracle models.InboundOracle
	i, err := strconv.Atoi(inboundOracleID)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid oracle id!"})
		return
	}
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}
	result := db.Preload(clause.Associations).Preload("InboundOracleTemplate.EventParameters").First(&inboundOracle, i)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No valid oracle id!"})
		return
	}

	data, _ := ioutil.ReadAll(ctx.Request.Body)
	var bodyData map[string]interface{}
	if e := json.Unmarshal(data, &bodyData); e != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": e.Error()})
		return
	}

	client, err := ethclient.Dial(inboundOracle.InboundOracleTemplate.BlockchainAddress)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: get this private key from user account
	privateKey, err := crypto.HexToECDSA("b28c350293dcf09cc5b5a9e5922e2f73e48983fe8d325855f04f749b1a82e0e6")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(inboundOracle.InboundOracleTemplate.ContractAddress)
	inputs := inboundOracle.InboundOracleTemplate.GetEventParameterJSON()
	name := inboundOracle.InboundOracleTemplate.ContractName
	fmt.Println(inputs, name)
	abi := "[{\"inputs\":" + inputs + ",\"name\":\"" + name + "\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

	instance, err := NewStore(address, client, abi)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Unable to generate function stub from template."})
		return
	}

	inboundEvent := models.InboundEvent{
		InboundOracleID: inboundOracle.ID,
		Success:         false,
	}
	db.Create(&inboundEvent)

	var parameters []interface{}
	for k, v := range bodyData {
		var eventParameter models.EventParameter
		// Add .Where("InboundOracleTemplateID = ?", inboundOracle.InboundOracleTemplateID)
		db.Where("Name = ?", k).First(&eventParameter)
		eventValue := models.EventValue{InboundEventID: inboundEvent.ID, Value: fmt.Sprintf("%v", v), EventParameterID: eventParameter.ID}
		db.Create(&eventValue)
		// TODO: transform parameter type depending on the definition in the db
		parameter, err := transformParameterType(v, eventParameter.Type)
		if err != nil {
			fmt.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Parameters have wrong types!"})
			return
		}
		parameters = append(parameters, parameter)
	}
	fmt.Println(parameters...)
	tx, err := instance.StoreTransactor.contract.Transact(auth, "set", parameters...)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured while trying to create transaction!"})
		return
	}

	fmt.Println("INFO: tx sent %s", tx.Hash().Hex())

	inboundEvent.Success = true
	db.Save(&inboundEvent)
	ctx.JSON(http.StatusOK, gin.H{})

}

func transformParameterType(parameter interface{}, parameterType string) (interface{}, error) {
	f, ok := parameter.(float64)
	if ok {
		return big.NewInt(int64(f)), nil
	}
	return nil, fmt.Errorf("Unable to transform paramter of type: %T", parameter)
}

func GetInboundOracles(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"body": "Ups there was a mistake!"})
		return
	}

	var inboundOracles []models.InboundOracle
	db.Preload(clause.Associations).Find(&inboundOracles)

	fmt.Println(inboundOracles)

	ctx.JSON(http.StatusOK, gin.H{"inboundOracles": inboundOracles})
}

func GetInboundOracle(ctx *gin.Context) {
	id := ctx.Param("inboundOracleId")
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid oracle id!"})
		return
	}

	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var inboundOracle models.InboundOracle
	result := db.Preload("InboundEvents.EventValues.EventParameter").Preload(clause.Associations).First(&inboundOracle, i)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No inbound Oracle with this ID available."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"inboundOracle": inboundOracle})
}

func UpdateInboundOracle(ctx *gin.Context) {
	id := ctx.Param("inboundOracleId")
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid oracle id!"})
		return
	}

	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var inboundOracle models.InboundOracle
	result := db.First(&inboundOracle, i)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No inbound Oracle with this ID available."})
		return
	}
	var inboundOraclePostBody forms.InboundOracleBody
	if err = ctx.ShouldBind(&inboundOraclePostBody); err != nil || !inboundOraclePostBody.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"body": "No valid body send!"})
		return
	}

	inboundOracle.Name = inboundOraclePostBody.Name

	db.Save(&inboundOracle)
	ctx.JSON(http.StatusOK, gin.H{"inboundOracle": inboundOracle})
}
