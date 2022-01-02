package chaincode

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const oracleName = "on-chain-oracle"
const contractName = "test-contract"
const topic = "test-topic"

type SmartContract struct {
	contractapi.Contract
}

// StartOnChainSubscription subscribes to a topic "test-topic" at the local on-chain oracle.
func (s *SmartContract) StartOnChainSubscription(ctx contractapi.TransactionContextInterface) {
	params := [][]byte{[]byte("Subscribe"), []byte(topic), []byte(contractName)}
	stub := ctx.GetStub()
	stub.InvokeChaincode(oracleName, params, stub.GetChannelID())
}

// StartOffChainSubscription subscribes to the tpoic "test-topic" at an off chain oracle.
func (s *SmartContract) StartOffChainSubscription(ctx contractapi.TransactionContextInterface) {
	payload := map[string]interface{}{
		"topic":         "test-topic",
		"smartContract": contractName,
	}
	json, _ := json.Marshal(payload) // payload is fine. We don't need to check the error
	ctx.GetStub().SetEvent("Subscribe", json)
}

func (s *SmartContract) Callback(ctx contractapi.TransactionContextInterface, number int) {
	payload := map[string]interface{}{
		"number": number,
	}
	json, _ := json.Marshal(payload) // payload is fine. We don't need to check the error
	ctx.GetStub().SetEvent("Published", json)
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	return nil
}
