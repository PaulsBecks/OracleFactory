package forms

import (
	"fmt"
	"regexp"
)

type InboundOracleTemplateBody struct {
	BlockchainName    string
	BlockchainAddress string
	ContractName      string
	ContractAddress   string
	Private           bool
}

// TODO: create real validation
func (o *InboundOracleTemplateBody) Valid() bool {
	ok, err := regexp.MatchString(`.+`, o.BlockchainName)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ok && err == nil
}
