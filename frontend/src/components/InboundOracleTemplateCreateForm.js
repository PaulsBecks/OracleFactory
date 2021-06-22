import React, { useEffect, useState } from "react";
import { Button, Form, Segment, Header } from "semantic-ui-react";

function formToAbi(oracles) {
  return JSON.stringify(
    oracles.map((oracle) => {
      delete oracle.ContractAddress;
      return oracle;
    })
  )
    .replace(/"ContractName":/g, '"name":')
    .replace(/"Name":/g, '"name":')
    .replace(/"Type":/g, '"type":');
}

function parseAbi(abi) {
  let oracles = [];
  try {
    JSON.parse(abi);
    const parsedAbi = JSON.parse(
      abi.replace(/"name":/g, '"Name":').replace(/"type":/g, '"Type":')
    );
    oracles = parsedAbi.map((oracle) => {
      const _oracle = { ...oracle, ContractName: oracle.Name };
      delete _oracle.Name;
      return _oracle;
    });
    return oracles;
  } catch (err) {
    console.log(err, "No valid abi");
  }
  return [];
}

export default function InboundOracleTemplateForm({
  inboundOracleTemplate,
  setInboundOracleTemplate,
}) {
  function updateInboundOracleTemplate(_, { name, value }) {
    setInboundOracleTemplate({ ...inboundOracleTemplate, [name]: value });
  }

  const [abi, setAbi] = useState("[]");

  useEffect(() => {
    setAbi(formToAbi(inboundOracleTemplate.inboundOracleTemplates));
  }, [inboundOracleTemplate]);

  if (!inboundOracleTemplate) return "";

  return (
    <Form>
      <Form.Group widths="equal">
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
      </Form.Group>

      <Form.TextArea
        label="ABI"
        value={abi}
        onChange={(_, { value }) => setAbi(value)}
        onBlur={() =>
          setInboundOracleTemplate({
            ...inboundOracleTemplate,
            inboundOracleTemplates: parseAbi(abi),
          })
        }
      />
      {/*<Button
        positive
        basic
        onClick={() =>
          setInboundOracleTemplate({
            ...inboundOracleTemplate,
            inboundOracleTemplates: parseAbi(abi),
          })
        }
        content="Save Abi"
      />*/}
      <div
        style={{ display: "flex", flexWrap: "wrap", alignItems: "flex-start" }}
      >
        {inboundOracleTemplate.inboundOracleTemplates.map((oracle) => (
          <Segment
            color="green"
            style={{ maxWidth: "30em", margin: "1em 1em 0 0" }}
          >
            <h3>
              {oracle.Type === "function"
                ? "Inbound Oracle"
                : "Outbound Oracle"}
            </h3>
            <Form.Group widths="equal">
              <Form.Input
                label="Contract Name"
                name="ContractName"
                value={oracle.ContractName}
                onChange={updateInboundOracleTemplate}
                placeholder="The name of the event"
              />
            </Form.Group>
            <Header>Inputs</Header>
            {oracle.inputs.map((input) => (
              <Form.Group widths="equal">
                <Form.Input label="Name" value={input.Name} />
                <Form.Input label="Type" value={input.Type} />
              </Form.Group>
            ))}
            <div>
              {/*<Button
                icon="plus"
                basic
                positive
                size="tiny"
                content="Add Input"
              />
              <Button
                icon="close"
                basic
                negative
                size="tiny"
                content="Cancel Oracle"
              />*/}
            </div>
          </Segment>
        ))}
      </div>
    </Form>
  );
}
