package forms

import (
	"fmt"
	"regexp"
)

type OutboundOraclePostBody struct {
	URI string "json:uri"
}

func (o *OutboundOraclePostBody) Valid() bool {
	ok, err := regexp.MatchString(`.+`, o.URI)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ok && err == nil
}
