package hyperledger

import (
	"fmt"

	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

func CreateTransaction(inboundOracle *models.InboundOracle, user *models.User, event *models.Event) error {
	//_ = os.Setenv("DISCOVERY_AS_LOCALHOST", "false")
	organizationName := user.HyperledgerOrganizationName
	cert := user.HyperledgerCert
	key := user.HyperledgerKey
	gatewayConfig := user.HyperledgerConfig
	channel := user.HyperledgerChannel

	contractAddress := inboundOracle.InboundOracleTemplate.OracleTemplate.ContractAddress
	contractName := inboundOracle.InboundOracleTemplate.OracleTemplate.EventName
	parameters := []string{}
	for _, eventValue := range event.EventValues {
		parameters = append(parameters, eventValue.Value)
	}
	fmt.Println(parameters)

	wallet := gateway.NewInMemoryWallet()
	wallet.Put("appUser", gateway.NewX509Identity(organizationName, string(cert), string(key)))
	config := config.FromRaw([]byte(gatewayConfig), "yaml")
	gw, err := gateway.Connect(
		gateway.WithConfig(config),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		return err
	}
	defer gw.Close()

	network, err := gw.GetNetwork(channel)
	if err != nil {
		return fmt.Errorf("Failed to get network: %v", err)
	}

	contract := network.GetContract(contractAddress)

	_, err = contract.SubmitTransaction(contractName, parameters...)
	if err != nil {
		return fmt.Errorf("Failed to Submit transaction: %v", err)
	}

	return nil
}
