package forms

import (
	"fmt"
	"regexp"
)

type OutboundOraclePostBody struct {
	Oracle                  OracleBody
	WebServicePublisherID   uint
	SmartContractListenerID uint
}

// TODO: create real validation
func (o *OutboundOraclePostBody) Valid() bool {
	ok, err := regexp.MatchString(`.+`, o.Oracle.Name)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ok && err == nil
}
