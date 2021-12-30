import React from "react";
import { Form, Segment } from "semantic-ui-react";

export default function HyperledgerBlockchainConnectorForm({
  hyperledgerBlockchainConnector,
  setHyperledgerBlockchainConnector,
}) {
  function updateHyperledgerBlockchainConnector(_, { name, value, checked }) {
    setHyperledgerBlockchainConnector({
      ...hyperledgerBlockchainConnector,
      [name]: value !== undefined ? value : checked,
    });
  }
  if (!hyperledgerBlockchainConnector) return "";

  return (
    <Form>
      <h2>Hyperledger Connector</h2>
      <Segment>
        <Form.Group widths="equal">
          <Form.Input
            label="Hyperledger Organization Name"
            name="HyperledgerOrganizationName"
            value={hyperledgerBlockchainConnector.HyperledgerOrganizationName}
            onChange={updateHyperledgerBlockchainConnector}
            placeholder="Your ethereum gateway"
          />
          <Form.Input
            label="Hyperledger Channel"
            name="HyperledgerChannel"
            value={hyperledgerBlockchainConnector.HyperledgerChannel}
            onChange={updateHyperledgerBlockchainConnector}
            placeholder="Your ethereum gateway"
          />
        </Form.Group>
        <Form.TextArea
          label="Hyperledger Config"
          name="HyperledgerConfig"
          value={hyperledgerBlockchainConnector.HyperledgerConfig}
          onChange={updateHyperledgerBlockchainConnector}
          placeholder="Your ethereum gateway"
        />
        <Form.TextArea
          label="Hyperledger Certificate"
          name="HyperledgerCert"
          value={hyperledgerBlockchainConnector.HyperledgerCert}
          onChange={updateHyperledgerBlockchainConnector}
          placeholder="Your ethereum gateway"
        />
        <Form.TextArea
          label="Hyperledger Key"
          name="HyperledgerKey"
          value={hyperledgerBlockchainConnector.HyperledgerKey}
          onChange={updateHyperledgerBlockchainConnector}
          placeholder="Your ethereum gateway"
        />
        <Form.Checkbox
          checked={hyperledgerBlockchainConnector.IsOnChain}
          label={
            hyperledgerBlockchainConnector.IsOnChain ? "on chain" : "off chain"
          }
          name="IsOnChain"
          toggle
          onChange={updateHyperledgerBlockchainConnector}
        />
      </Segment>
    </Form>
  );
}
