import React from "react";
import { Form } from "semantic-ui-react";

export default function InboundOracleForm({ inboundOracle, setInboundOracle }) {
  function updateInboundOracle(_, { name, value }) {
    setInboundOracle({ ...inboundOracle, [name]: value });
  }
  if (!inboundOracle) return "";

  return (
    <Form>
      <Form.Input
        label="Name"
        name="Name"
        value={inboundOracle.Name}
        onChange={updateInboundOracle}
        placeholder="A name to recognize the oracle"
      />
    </Form>
  );
}
