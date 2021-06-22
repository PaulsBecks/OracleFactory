package forms

import (
	"fmt"
	"regexp"
)

type OutboundOracleTemplateBody struct {
	BlockchainName    string
	BlockchainAddress string
	EventName         string
	ContractAddress   string
}

// TODO: create real validation
func (o *OutboundOracleTemplateBody) Valid() bool {
	ok, err := regexp.MatchString(`.+`, o.BlockchainName)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ok && err == nil
}
