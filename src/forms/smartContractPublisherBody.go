package forms

import (
	"fmt"
	"regexp"
)

type SmartContractPublisherBody struct {
	Name                   string
	Description            string
	BlockchainName         string
	BlockchainAddress      string
	ContractName           string
	ContractAddress        string
	ContractAddressSynonym string
	Private                bool
}

// TODO: create real validation
func (o *SmartContractPublisherBody) Valid() bool {
	ok, err := regexp.MatchString(`.+`, o.BlockchainName)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ok && err == nil
}