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
        value={outboundOracle.Oracle.Name}
        onChange={(_, { value }) =>
          setOutboundOracle({
            ...outboundOracle,
            Oracle: { ...outboundOracle.Oracle, Name: value },
          })
        }
        placeholder="A name to recognize the oracle"
      />
    </Form>
  );
}