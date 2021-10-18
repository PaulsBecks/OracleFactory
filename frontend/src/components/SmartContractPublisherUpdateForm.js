import React from "react";
import { Form } from "semantic-ui-react";

export default function SmartContractPublisherForm({
  smartContractPublisher,
  setSmartContractPublisher,
}) {
  function updateSmartContractPublisher(_, { name, value }) {
    setSmartContractPublisher({ ...smartContractPublisher, [name]: value });
  }
  if (!smartContractPublisher) return "";

  return (
    <Form>
      <Form.Input
        label="Blockchain Name"
        name="BlockchainName"
        value={smartContractPublisher.BlockchainName}
        onChange={updateSmartContractPublisher}
        disabled
        placeholder="The name of the blockchain"
      />
      <Form.Input
        label="Blockchain Address"
        name="BlockchainAddress"
        value={smartContractPublisher.BlockchainAddress}
        onChange={updateSmartContractPublisher}
        placeholder="The address of the blockchain"
      />
      <Form.Input
        label="Contract Address"
        name="ContractAddress"
        value={smartContractPublisher.ContractAddress}
        onChange={updateSmartContractPublisher}
        placeholder="The address of the contract"
      />
      <Form.Input
        label="Contract Name"
        name="ContractName"
        value={smartContractPublisher.EventName}
        onChange={updateSmartContractPublisher}
        placeholder="The name of the event"
      />
    </Form>
  );
}
