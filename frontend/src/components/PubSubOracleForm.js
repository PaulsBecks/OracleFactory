import React from "react";
import { Form } from "semantic-ui-react";

export default function PubSubOracleForm({ pubSubOracle, setPubSubOracle }) {
  function updatePubSubOracle(_, { name, value }) {
    setPubSubOracle({ ...pubSubOracle, [name]: value });
  }
  if (!pubSubOracle) return "";

  return (
    <Form>
      <Form.Input
        label="Name"
        name="Name"
        value={pubSubOracle.Oracle.Name}
        onChange={(_, { value }) =>
          setPubSubOracle({
            ...pubSubOracle,
            Oracle: { ...pubSubOracle.Oracle, Name: value },
          })
        }
        placeholder="A name to recognize the oracle"
      />
    </Form>
  );
}
