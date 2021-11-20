package forms

import (
	"fmt"
	"regexp"
)

type PubSubOracleBody struct {
	Oracle                 OracleBody
	ConsumerID             uint
	ProviderID             uint
	SubBlockchainEventID   uint
	UnsubBlockchainEventID uint
}

// TODO: create real validation
func (o *PubSubOracleBody) Valid() bool {
	ok, err := regexp.MatchString(`.+`, o.Oracle.Name)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ok && err == nil
}
