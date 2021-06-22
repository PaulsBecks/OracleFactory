package forms

import (
	"fmt"
	"regexp"
)

type UserBody struct {
	EthereumPrivateKey          string
	EthereumPublicKey           string
	EthereumAddress             string
	HyperledgerConfig           string
	HyperledgerCert             string
	HyperledgerKey              string
	HyperledgerOrganizationName string
	HyperledgerChannel          string
}

// TODO: create real validation
func (o *UserBody) Valid() bool {
	ok, err := regexp.MatchString(`.+`, o.EthereumPrivateKey)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ok && err == nil
}
