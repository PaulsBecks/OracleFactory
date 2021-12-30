import React from "react";
import { Form, Segment } from "semantic-ui-react";

export default function EthereumBlockchainConnectorForm({
  ethereumBlockchainConnector,
  setEthereumBlockchainConnector,
}) {
  function updateEthereumBlockchainConnector(_, { name, value, checked }) {
    setEthereumBlockchainConnector({
      ...ethereumBlockchainConnector,
      [name]: value !== undefined ? value : checked,
    });
  }
  if (!ethereumBlockchainConnector) return "";

  return (
    <Form>
      <h2>Ethereum Connector</h2>
      <Segment>
        <Form.Input
          label="Ethereum Private Key"
          name="EthereumPrivateKey"
          value={ethereumBlockchainConnector.EthereumPrivateKey}
          onChange={updateEthereumBlockchainConnector}
          placeholder="Your private key"
        />
        <Form.Input
          label="Ethereum Gateway Address"
          name="EthereumAddress"
          value={ethereumBlockchainConnector.EthereumAddress}
          onChange={updateEthereumBlockchainConnector}
          placeholder="Your ethereum gateway"
        />
        <Form.Checkbox
          checked={ethereumBlockchainConnector.IsOnChain}
          label={
            ethereumBlockchainConnector.IsOnChain ? "on chain" : "off chain"
          }
          name="IsOnChain"
          toggle
          onChange={updateEthereumBlockchainConnector}
        />
      </Segment>
    </Form>
  );
}
