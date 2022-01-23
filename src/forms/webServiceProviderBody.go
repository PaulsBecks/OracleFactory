package forms

import (
	"fmt"
	"regexp"
)

type WebServiceProviderBody struct {
	Name        string
	Description string
	Private     bool
}

// TODO: create real validation
func (o *WebServiceProviderBody) Valid() bool {
	ok, err := regexp.MatchString(`.+`, o.Name)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ok && err == nil
}
