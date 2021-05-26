package routes

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"io/ioutil"
	"log"
	"strconv"
	"fmt"
	"math/big"
	"net/http"
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
	result := db.Preload(clause.Associations).First(&inboundOracle, i)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No valid oracle id!"})
		return
	}
	fmt.Println(inboundOracle)

	data, _ := ioutil.ReadAll(ctx.Request.Body)
	fmt.Println(string(data))

	var bodyData map[string]interface{} // map[string]interface{}
	if e := json.Unmarshal(data, &bodyData); e != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": e.Error()})
		return
	}

	// TODO: get this address from inbound oracle
	client, err := ethclient.Dial("http://eth-test-net:8545")
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

	// TODO: get this address from InboundOracleTempalte
	address := common.HexToAddress("0xe4EFfB267484Cd790143484de3Bae7fDfbE31F00")
	// TODO: get inputs from InboundOracleTemplate
	inputs := "{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}"
	// TODO: get name from InboundOracleTemplate
	name := inboundOracle.InboundOracleTemplate.ContractName // "set"
	abi := "[{\"inputs\":[" + inputs + "],\"name\":\"" + name + "\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

	instance, err := NewStore(address, client, abi)
	if err != nil {
		log.Fatal(err)
	}

	var parameters []interface{}
	for _, v := range bodyData {
		// TODO: transform parameter type depending on the definition in the db
		parameter, err := transformParameterType(v)
		if err != nil {
			fmt.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Parameters have wrong types!"})
			return
		}
		parameters = append(parameters, parameter)
	}
	tx, err := instance.StoreTransactor.contract.Transact(auth, "set", parameters...)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("tx sent: %s", tx.Hash().Hex())

	// TODO: write result to db that the transaction was submitted
}

func transformParameterType(parameter interface{}) (interface{}, error) {
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
