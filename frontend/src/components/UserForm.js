import React from "react";
import { Form, Segment } from "semantic-ui-react";

export default function UserForm({ user, setUser }) {
  function updateUser(_, { name, value }) {
    setUser({ ...user, [name]: value });
  }
  if (!user) return "";

  return (
    <Form>
      <h2>Ethereum</h2>
      <Segment>
        <Form.Input
          label="Ethereum Private Key"
          name="EthereumPrivateKey"
          value={user.EthereumPrivateKey}
          onChange={updateUser}
          placeholder="Your private key"
        />
        <Form.Input
          label="Ethereum Gateway Address"
          name="EthereumAddress"
          value={user.EthereumAddress}
          onChange={updateUser}
          placeholder="Your ethereum gateway"
        />
      </Segment>
      <h2>Hyperledger</h2>
      <Segment>
        <Form.Group widths="equal">
          <Form.Input
            label="Hyperledger Organization Name"
            name="HyperledgerOrganizationName"
            value={user.HyperledgerOrganizationName}
            onChange={updateUser}
            placeholder="Your ethereum gateway"
          />
          <Form.Input
            label="Hyperledger Channel"
            name="HyperledgerChannel"
            value={user.HyperledgerChannel}
            onChange={updateUser}
            placeholder="Your ethereum gateway"
          />
        </Form.Group>
        <Form.TextArea
          label="Hyperledger Config"
          name="HyperledgerConfig"
          value={user.HyperledgerConfig}
          onChange={updateUser}
          placeholder="Your ethereum gateway"
        />
        <Form.TextArea
          label="Hyperledger Certificate"
          name="HyperledgerCert"
          value={user.HyperledgerCert}
          onChange={updateUser}
          placeholder="Your ethereum gateway"
        />
        <Form.TextArea
          label="Hyperledger Key"
          name="HyperledgerKey"
          value={user.HyperledgerKey}
          onChange={updateUser}
          placeholder="Your ethereum gateway"
        />
      </Segment>
    </Form>
  );
}
