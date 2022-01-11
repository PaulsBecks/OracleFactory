package forms

import (
	"fmt"
	"regexp"
)

type EventParameterBody struct {
	Name    string
	Type    string
	Indexed bool
}

// TODO: create real validation
func (o *EventParameterBody) Valid() bool {
	ok, err := regexp.MatchString(`.+`, o.Name)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ok && err == nil
}
