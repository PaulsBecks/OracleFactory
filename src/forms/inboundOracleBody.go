package forms

import (
	"fmt"
	"regexp"
)

type InboundOracleBody struct {
	Oracle                   OracleBody
	SmartContractPublisherID uint
	WebServiceListenerID     uint
}

// TODO: create real validation
func (o *InboundOracleBody) Valid() bool {
	ok, err := regexp.MatchString(`.+`, o.Oracle.Name)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ok && err == nil
}
