package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username                    string
	Password                    string
	EthereumPrivateKey          string
	EthereumPublicKey           string
	EthereumAddress             string
	HyperledgerConfig           string
	HyperledgerCert             string
	HyperledgerKey              string
	HyperledgerOrganizationName string
	HyperledgerChannel          string
}
