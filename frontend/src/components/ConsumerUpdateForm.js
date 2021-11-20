import React from "react";
import { Form } from "semantic-ui-react";

export default function ConsumerForm({ consumer, setConsumer }) {
  function updateConsumer(_, { name, value }) {
    setConsumer({ ...consumer, [name]: value });
  }
  if (!consumer) return "";

  return (
    <Form>
      <Form.Input
        label="Blockchain Name"
        name="BlockchainName"
        value={consumer.BlockchainName}
        onChange={updateConsumer}
        disabled
        placeholder="The name of the blockchain"
      />
      <Form.Input
        label="Blockchain Address"
        name="BlockchainAddress"
        value={consumer.BlockchainAddress}
        onChange={updateConsumer}
        placeholder="The address of the blockchain"
      />
      <Form.Input
        label="Contract Address"
        name="ContractAddress"
        value={consumer.ContractAddress}
        onChange={updateConsumer}
        placeholder="The address of the contract"
      />
      <Form.Input
        label="Contract Name"
        name="ContractName"
        value={consumer.EventName}
        onChange={updateConsumer}
        placeholder="The name of the event"
      />
    </Form>
  );
}
