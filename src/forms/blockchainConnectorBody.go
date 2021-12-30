package forms

type EthereumConnectorBody struct {
	IsOnChain          bool
	EthereumPrivateKey string
	EthereumAddress    string
}

// TODO: create real validation
func (o *EthereumConnectorBody) Valid() bool {
	return true
}

type HyperledgerConnectorBody struct {
	IsOnChain                   bool
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
