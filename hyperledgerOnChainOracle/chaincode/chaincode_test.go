package chaincode_test

import (
	"testing"

	"github.com/paulsbecks/hyperledgerOnChainPubSubOracle/chaincode"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/paulsbecks/hyperledgerOnChainPubSubOracle/mocks"
	"github.com/stretchr/testify/require"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o mocks/transaction.go -fake-name TransactionContext . transactionContext
type transactionContext interface {
	contractapi.TransactionContextInterface
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o mocks/chaincodestub.go -fake-name ChaincodeStub . chaincodeStub
type chaincodeStub interface {
	shim.ChaincodeStubInterface
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o mocks/statequeryiterator.go -fake-name StateQueryIterator . stateQueryIterator
type stateQueryIterator interface {
	shim.StateQueryIteratorInterface
}

func TestSubscription(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)
	oracle := new(chaincode.OracleContract)
	err := oracle.InitLedger(transactionContext)
	require.NoError(t, err)

	err = oracle.Subscribe(transactionContext, "test-topic", "TestContract")
	require.NoError(t, err)

	err = oracle.Subscribe(transactionContext, "test-topic", "TestContract")
	require.EqualError(t, err, "Already subscribed")

}

func TestUnsubscription(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)
	oracle := new(chaincode.OracleContract)
	err := oracle.InitLedger(transactionContext)
	require.NoError(t, err)

	err = oracle.Unsubscribe(transactionContext, "test-topic", "TestContract2")
	require.EqualError(t, err, "Element not in list.")

	err = oracle.Unsubscribe(transactionContext, "test-topic", "TestContract")
	require.NoError(t, err)

	err = oracle.Unsubscribe(transactionContext, "test-topic", "TestContract")
	require.EqualError(t, err, "Element not in list.")

	err = oracle.Unsubscribe(transactionContext, "test-topic2", "TestContract")
	require.EqualError(t, err, "No topic test-topic2")
}

func TestPublish(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)
	oracle := new(chaincode.OracleContract)
	err := oracle.InitLedger(transactionContext)
	require.NoError(t, err)

	err = oracle.Subscribe(transactionContext, "test-topic", "TestContract")
	require.NoError(t, err)

	err = oracle.Publish(transactionContext, "test-topic", "content")
	require.NoError(t, err)
}
