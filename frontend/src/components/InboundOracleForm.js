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
        value={inboundOracle.Oracle.Name}
        onChange={(_, { value }) =>
          setInboundOracle({
            ...inboundOracle,
            Oracle: { ...inboundOracle.Oracle, Name: value },
          })
        }
        placeholder="A name to recognize the oracle"
      />
    </Form>
  );
}
