import React from "react";
import { Form } from "semantic-ui-react";

export default function InboundOracleTemplateForm({
  inboundOracleTemplate,
  setInboundOracleTemplate,
}) {
  function updateInboundOracleTemplate(_, { name, value }) {
    setInboundOracleTemplate({ ...inboundOracleTemplate, [name]: value });
  }
  if (!inboundOracleTemplate) return "";

  return (
    <Form>
      <Form.Input
        label="Blockchain Name"
        name="BlockchainName"
        value={inboundOracleTemplate.BlockchainName}
        onChange={updateInboundOracleTemplate}
        disabled
        placeholder="The name of the blockchain"
      />
      <Form.Input
        label="Blockchain Address"
        name="BlockchainAddress"
        value={inboundOracleTemplate.BlockchainAddress}
        onChange={updateInboundOracleTemplate}
        placeholder="The address of the blockchain"
      />
      <Form.Input
        label="Contract Address"
        name="ContractAddress"
        value={inboundOracleTemplate.ContractAddress}
        onChange={updateInboundOracleTemplate}
        placeholder="The address of the contract"
      />
      <Form.Input
        label="Contract Name"
        name="ContractName"
        value={inboundOracleTemplate.EventName}
        onChange={updateInboundOracleTemplate}
        placeholder="The name of the event"
      />
    </Form>
  );
}
