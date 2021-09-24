package ethereum

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ParseValues(event *models.Event) ([]interface{}, error) {
	var bodyData map[string]interface{}

	if e := json.Unmarshal(event.Body, &bodyData); e != nil {
		return nil, e
	}
	fmt.Println(event)

	var parameters []interface{}
	fmt.Println(event.GetEventValues())
	for _, eventValue := range event.GetEventValues() {
		eventParameter := eventValue.GetEventParameter()
		fmt.Println(eventValue, eventParameter)
		v := bodyData[eventParameter.Name]
		parameter, err := utils.TransformParameterType(v, eventParameter.Type)
		if err != nil {
			return nil, err
		}
		parameters = append(parameters, parameter)
	}
	return parameters, nil
}

func CreateTransaction(inboundOracle *models.InboundOracle, user *models.User, event *models.Event) error {
	client, err := ethclient.Dial(inboundOracle.GetOracle().GetUser().EthereumAddress)
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

	address := common.HexToAddress(inboundOracle.InboundOracleTemplate.OracleTemplate.ContractAddress)
	inputs := inboundOracle.InboundOracleTemplate.GetEventParameterJSON()
	name := inboundOracle.InboundOracleTemplate.OracleTemplate.EventName
	fmt.Println(address)
	fmt.Println(inputs, name)
	abi := "[{\"inputs\":" + inputs + ",\"name\":\"" + name + "\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	fmt.Println(abi)
	instance, err := NewStore(address, client, abi)
	if err != nil {
		return err
	}

	parameters, err := ParseValues(event)
	if err != nil {
		return err
	}

	tx, err := instance.StoreTransactor.contract.Transact(auth, name, parameters...)
	if err != nil {
		return err
	}

	fmt.Printf("INFO: tx sent %s\n", tx.Hash().Hex())
	return nil
}
