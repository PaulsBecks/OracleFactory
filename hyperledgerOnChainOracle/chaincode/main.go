package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type OracleContract struct {
	contractapi.Contract
}

var subscriptionsMap = make(map[string][]string)
var subscriptionsMapKey = "subscriptionsmap"

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func findPosOfElement(s []string, e string) (int, error) {
	for i, a := range s {
		if a == e {
			return i, nil
		}
	}
	return 0, fmt.Errorf("Element not in list.")
}

func remove(s []string, a string) ([]string, error) {
	pos, err := findPosOfElement(s, a)
	if err != nil {
		return s, err
	}
	s[pos] = s[len(s)-1]
	return s[:len(s)-1], nil
}

func readState(ctx contractapi.TransactionContextInterface) {
	assetJSON, err := ctx.GetStub().GetState(subscriptionsMapKey)
	if err != nil {
		// should not happen
	}
	if assetJSON == nil {
		// should not happen
	}
	err = json.Unmarshal(assetJSON, &subscriptionsMap)
}

func saveState(ctx contractapi.TransactionContextInterface) {
	assetJSON, err := json.Marshal(subscriptionsMap)
	if err != nil {
		// This should not happen
	}
	ctx.GetStub().PutState(subscriptionsMapKey, assetJSON)
}

func (o *OracleContract) Subscribe(ctx contractapi.TransactionContextInterface, topic string, smartContract string) error {
	readState(ctx)
	subscriptions := subscriptionsMap[topic]
	if subscriptions == nil {
		subscriptions = []string{smartContract}
	} else {
		if !contains(subscriptions, smartContract) {
			subscriptions = append(subscriptions, smartContract)
		} else {
			return fmt.Errorf("Already subscribed")
		}
	}
	subscriptionsMap[topic] = subscriptions
	saveState(ctx)
	return nil
}

func (o *OracleContract) Unsubscribe(ctx contractapi.TransactionContextInterface, topic string, smartContract string) error {
	readState(ctx)
	subscriptions, ok := subscriptionsMap[topic]
	var err error
	if !ok {
		return fmt.Errorf("No topic %s", topic)
	} else {
		subscriptions, err = remove(subscriptions, smartContract)
		if err != nil {
			return err
		}
	}
	subscriptionsMap[topic] = subscriptions
	saveState(ctx)
	return nil
}

func (o *OracleContract) Publish(ctx contractapi.TransactionContextInterface, topic string, content string) error {
	readState(ctx)
	subscriptions, ok := subscriptionsMap[topic]
	var err error
	if !ok {
		return err
	}
	for _, sub := range subscriptions {
		ctx.GetStub().InvokeChaincode(sub, ctx.GetStub().GetArgs(), ctx.GetStub().GetChannelID())
	}
	return err
}

func (o *OracleContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	return nil
}
