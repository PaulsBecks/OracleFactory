package forms

type EthereumConnectorBody struct {
	EthereumPrivateKey string
	EthereumAddress    string
}

// TODO: create real validation
func (o *EthereumConnectorBody) Valid() bool {
	return true
}

type HyperledgerConnectorBody struct {
	HyperledgerConfig           string
	HyperledgerCert             string
	HyperledgerKey              string
	HyperledgerOrganizationName string
	HyperledgerChannel          string
}

// TODO: create real validation
func (o *HyperledgerConnectorBody) Valid() bool {
	return true
}
