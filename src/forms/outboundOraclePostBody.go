package forms

import (
	"fmt"
	"regexp"
)

type OutboundOraclePostBody struct {
	URI  string
	Name string
}

// TODO: create real validation
func (o *OutboundOraclePostBody) Valid() bool {
	ok, err := regexp.MatchString(`.+`, o.URI)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ok && err == nil
}
