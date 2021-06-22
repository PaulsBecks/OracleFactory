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
	"os"
	"strconv"

	"github.com/PaulsBecks/OracleFactory/src/forms"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
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

	userInterface, _ := ctx.Get("user")
	user, _ := userInterface.(models.User)

	inboundEvent := models.InboundEvent{
		InboundOracleID: inboundOracle.ID,
		Success:         false,
	}
	db.Create(&inboundEvent)

	data, _ := ioutil.ReadAll(ctx.Request.Body)
	var bodyData map[string]interface{}
	if e := json.Unmarshal(data, &bodyData); e != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": e.Error()})
		return
	}

	if inboundOracle.InboundOracleTemplate.BlockchainName == "Ethereum" {
		createEthereumTransaction(ctx, db, &inboundOracle, &inboundEvent, &user, bodyData)
	}
	if inboundOracle.InboundOracleTemplate.BlockchainName == "Hyperledger" {
		createHyperledgerTransaction(ctx, db, &inboundOracle, &inboundEvent, &user, bodyData)
	}

	inboundEvent.Success = true
	db.Save(&inboundEvent)
	ctx.JSON(http.StatusOK, gin.H{})

}

func createHyperledgerTransaction(ctx *gin.Context, db *gorm.DB, inboundOracle *models.InboundOracle, inboundEvent *models.InboundEvent, user *models.User, bodyData map[string]interface{}) {
	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	organizationName := user.HyperledgerOrganizationName
	cert := user.HyperledgerCert
	key := user.HyperledgerKey
	gatewayConfig := user.HyperledgerConfig
	channel := user.HyperledgerChannel
	fmt.Println(cert, key, gatewayConfig)

	contractAddress := inboundOracle.InboundOracleTemplate.ContractAddress
	contractName := inboundOracle.InboundOracleTemplate.ContractName
	parameters := []string{}
	for _, eventParameter := range inboundOracle.InboundOracleTemplate.EventParameters {
		parameterName := eventParameter.Name
		value, ok := bodyData[parameterName]
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("No value for event parameter %s.", parameterName)})
			return
		}
		stringValue, ok := value.(string)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Unable to convert parameter %s to string.", parameterName)})
			return
		}
		parameters = append(parameters, eventParameter.Name, stringValue)
	}
	fmt.Println(parameters)

	wallet := gateway.NewInMemoryWallet()
	wallet.Put("appUser", gateway.NewX509Identity(organizationName, string(cert), string(key)))

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromRaw([]byte(gatewayConfig), "yaml")),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("Failed to connect to gateway: %v", err)})
		return
	}
	defer gw.Close()

	network, err := gw.GetNetwork(channel)
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}

	contract := network.GetContract(contractAddress)

	_, err = contract.SubmitTransaction(contractName, parameters...)
	if err != nil {
		log.Fatalf("Failed to Submit transaction: %v", err)
	}
}

func createEthereumTransaction(ctx *gin.Context, db *gorm.DB, inboundOracle *models.InboundOracle, inboundEvent *models.InboundEvent, user *models.User, bodyData map[string]interface{}) {
	client, err := ethclient.Dial(inboundOracle.User.EthereumAddress)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(user.EthereumPrivateKey)
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

	var eventParameters []models.EventParameter
	db.Find(&eventParameters, "inbound_oracle_template_id=?", inboundOracle.InboundOracleTemplateID)
	fmt.Println(eventParameters)
	var parameters []interface{}
	for _, eventParameter := range eventParameters {
		v := bodyData[eventParameter.Name]
		eventValue := models.EventValue{InboundEventID: inboundEvent.ID, Value: fmt.Sprintf("%v", v), EventParameterID: eventParameter.ID}
		db.Create(&eventValue)
		parameter, err := transformParameterType(v, eventParameter.Type)
		if err != nil {
			fmt.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Parameters have wrong types!"})
			return
		}
		parameters = append(parameters, parameter)
	}
	fmt.Println(parameters...)
	tx, err := instance.StoreTransactor.contract.Transact(auth, name, parameters...)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured while trying to create transaction!"})
		return
	}

	fmt.Println("INFO: tx sent %s", tx.Hash().Hex())
}

func transformParameterType(parameter interface{}, parameterType string) (interface{}, error) {
	switch parameterType {
	case "uint256":
		f, ok := parameter.(float64)
		if ok {
			return big.NewInt(int64(f)), nil
		}
		return nil, fmt.Errorf("Unable to parse parameter. %v", parameter)
	case "uint64":
		f, ok := parameter.(float64)
		if ok {
			return uint64(f), nil
		}
		return nil, fmt.Errorf("Unable to parse parameter. %v", parameter)
	case "uint8":
		f, ok := parameter.(float64)
		if ok {
			return uint8(f), nil
		}
		return nil, fmt.Errorf("Unable to parse parameter. %v", parameter)
	case "string":
		s, ok := parameter.(string)
		if ok {
			return s, nil
		}
		return nil, fmt.Errorf("Unable to parse parameter. %v", parameter)
	case "bytes":
		s, ok := parameter.(string)
		if ok {
			return []byte(s), nil
		}
		return nil, fmt.Errorf("Unable to parse parameter. %v", parameter)
	case "bytes32":
		s, ok := parameter.(string)
		if ok {
			var arr [32]byte
			copy(arr[:], []byte(s))
			return arr, nil
		}
		return nil, fmt.Errorf("Unable to parse parameter. %v", parameter)
	case "bool":
		b, ok := parameter.(bool)
		if ok {
			return b, nil
		}
		return nil, fmt.Errorf("Unable to parse parameter. %v", parameter)
	case "address":
		s, ok := parameter.(string)
		if ok {
			return common.HexToAddress(s), nil
		}
		return nil, fmt.Errorf("Unable to parse parameter. %v", parameter)
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
