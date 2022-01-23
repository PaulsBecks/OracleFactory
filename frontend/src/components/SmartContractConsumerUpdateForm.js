import React from "react";
import { Form } from "semantic-ui-react";

export default function SmartContractConsumerForm({
  smartContractConsumer,
  setSmartContractConsumer,
}) {
  function updateSmartContractConsumer(_, { name, value }) {
    setSmartContractConsumer({ ...smartContractConsumer, [name]: value });
  }
  if (!smartContractConsumer) return "";

  return (
    <Form>
      <Form.Input
        label="Blockchain Name"
        name="BlockchainName"
        value={smartContractConsumer.BlockchainName}
        onChange={updateSmartContractConsumer}
        disabled
        placeholder="The name of the blockchain"
      />
      <Form.Input
        label="Blockchain Address"
        name="BlockchainAddress"
        value={smartContractConsumer.BlockchainAddress}
        onChange={updateSmartContractConsumer}
        placeholder="The address of the blockchain"
      />
      <Form.Input
        label="Contract Address"
        name="ContractAddress"
        value={smartContractConsumer.ContractAddress}
        onChange={updateSmartContractConsumer}
        placeholder="The address of the contract"
      />
      <Form.Input
        label="Contract Name"
        name="ContractName"
        value={smartContractConsumer.EventName}
        onChange={updateSmartContractConsumer}
        placeholder="The name of the event"
      />
    </Form>
  );
}
