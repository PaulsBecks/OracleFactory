package forms

import (
	"fmt"
	"regexp"
)

type SmartContractProviderBody struct {
	Name                   string
	Description            string
	BlockchainName         string
	BlockchainAddress      string
	EventName              string
	ContractAddress        string
	ContractAddressSynonym string
	Private                bool
}

// TODO: create real validation
func (o *SmartContractProviderBody) Valid() bool {
	ok, err := regexp.MatchString(`.+`, o.BlockchainName)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ok && err == nil
}
