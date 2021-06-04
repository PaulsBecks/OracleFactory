import React from "react";
import { Form } from "semantic-ui-react";

export default function OutboundOracleForm({
  outboundOracle,
  setOutboundOracle,
}) {
  function updateOutboundOracle(_, { name, value }) {
    setOutboundOracle({ ...outboundOracle, [name]: value });
  }
  if (!outboundOracle) return "";

  return (
    <Form>
      <Form.Input
        label="Name"
        name="Name"
        value={outboundOracle.Name}
        onChange={updateOutboundOracle}
        placeholder="A name to recognize the oracle"
      />
      <Form.Input
        label="Forward to"
        name="URI"
        value={outboundOracle.URI}
        onChange={updateOutboundOracle}
        placeholder="http://your.domain/here"
      />
    </Form>
  );
}
