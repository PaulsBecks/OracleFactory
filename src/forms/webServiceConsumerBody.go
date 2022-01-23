package forms

import (
	"fmt"
	"regexp"
)

type WebServiceConsumerBody struct {
	Name        string
	Description string
	URL         string
	Private     bool
}

// TODO: create real validation
func (o *WebServiceConsumerBody) Valid() bool {
	ok, err := regexp.MatchString(`.+`, o.Name)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ok && err == nil
}
