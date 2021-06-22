package forms

import (
	"fmt"
	"regexp"
)

type InboundOracleBody struct {
	Name string
}

// TODO: create real validation
func (o *InboundOracleBody) Valid() bool {
	ok, err := regexp.MatchString(`.+`, o.Name)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ok && err == nil
}
