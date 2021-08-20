package forms

import (
	"fmt"
	"regexp"
)

type AuthBody struct {
	Email    string
	Password string
}

// TODO: create real validation
func (o *AuthBody) Valid() bool {
	ok, err := regexp.MatchString(`.+`, o.Email)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ok && err == nil
}
