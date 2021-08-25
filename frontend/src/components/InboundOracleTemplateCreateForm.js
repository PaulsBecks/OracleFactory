import React, { useEffect, useState } from "react";
import {
  Checkbox,
  Form,
  Segment,
  Header,
  Icon,
  Popup,
  Button,
} from "semantic-ui-react";
import Identicon from "react-identicons";
import { parseAbi, formToAbi } from "../services/utils/abiTransformation";

export default function InboundOracleTemplateForm({
  inboundOracleTemplate,
  setInboundOracleTemplate,
  fromABI,
  inbound,
  outbound,
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
    <div>
      <Form>
        <Form.Group widths="equal">
          <Form.Select
            label={
              <label>
                Blockchain Name{" "}
                <Popup
                  content="The blockchain name you want to interact with. You can choose from the blockchains implemented."
                  trigger={<Icon name="info circle" />}
                />
              </label>
            }
            name="BlockchainName"
            value={inboundOracleTemplate.BlockchainName}
            onChange={updateInboundOracleTemplate}
            options={[
              { key: "Ethereum", text: "Ethereum", value: "Ethereum" },
              { key: "Hyperledger", text: "Hyperledger", value: "Hyperledger" },
            ]}
            placeholder="The name of the blockchain"
          />
        </Form.Group>
        <Form.Group widths="equal">
          <Form.Input
            label={
              <label>
                Contract Address{" "}
                <Popup
                  content="The blockchain address of the smart contract. For ethereum the address looks like 0x07A93d6C2D964b702662971Efaca43499fEB198c. For Hyperledger enter the channel name."
                  trigger={<Icon name="info circle" />}
                />
              </label>
            }
            name="ContractAddress"
            value={inboundOracleTemplate.ContractAddress}
            onChange={updateInboundOracleTemplate}
            placeholder="The address of the contract"
          />
        </Form.Group>
        <Form.Group widths="equal">
          <Form.Input
            label={
              <label>
                Contract Address Synonym{" "}
                <Popup
                  content="This synonym will be part of the name of the oracle template to identify it. The full name will be a combination of the Contract Address Synonym and the Contract Name."
                  trigger={<Icon name="info circle" />}
                />
              </label>
            }
            name="ContractAddressSynonym"
            value={inboundOracleTemplate.ContractAddressSynonym}
            onChange={updateInboundOracleTemplate}
            placeholder="The Oracle Name"
          />
        </Form.Group>
        <Form.Group widths="2">
          <Form.Field>
            <label>
              Icon{" "}
              <Popup
                style={{ marginLeft: "40px" }}
                content="This icon is generated from the contract address. It helps to easily identify the oracle template."
                trigger={<Icon name="info circle" />}
              />
            </label>
            <Identicon
              string={inboundOracleTemplate.ContractAddress}
              size={50}
            />
          </Form.Field>
          <Form.Field>
            <label>
              Oracle Template Visibility{" "}
              <Popup
                content="Deactive the toggle to share the oracle template with other users."
                trigger={<Icon name="info circle" />}
              />
            </label>
            <Form.Checkbox
              checked={inboundOracleTemplate.Private}
              label={inboundOracleTemplate.Private ? "private" : "public"}
              name="Private"
              toggle
              onChange={(event, { name, checked }) =>
                updateInboundOracleTemplate(event, { name, value: checked })
              }
            />
          </Form.Field>
        </Form.Group>
        <Form.Group></Form.Group>
        {fromABI && (
          <>
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
          </>
        )}
        <div>
          {inboundOracleTemplate.inboundOracleTemplates.map((oracle, oracleI) =>
            fromABI ? (
              <Segment color="green" style={{ margin: "1em 1em 0 0" }}>
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
                    placeholder="The name of the event"
                    onChange={(event, { value }) => {
                      let _inboundOracleTemplate = {
                        ...inboundOracleTemplate,
                      };
                      _inboundOracleTemplate.inboundOracleTemplates[oracleI][
                        "ContractName"
                      ] = value;
                      setInboundOracleTemplate(_inboundOracleTemplate);
                    }}
                  />
                </Form.Group>
                <Header>Input Parameters</Header>
                {oracle.inputs.map((input, i) => (
                  <Form.Group widths="equal">
                    <Form.Input
                      label="Name"
                      value={input.Name}
                      onChange={(_, { value }) => {
                        let _inboundOracleTemplate = {
                          ...inboundOracleTemplate,
                        };
                        _inboundOracleTemplate.inboundOracleTemplates[
                          oracleI
                        ].inputs[i]["Name"] = value;
                        setInboundOracleTemplate(_inboundOracleTemplate);
                      }}
                    />
                    <Form.Input
                      label="Type"
                      value={input.Type}
                      onChange={(event, { value }) => {
                        let _inboundOracleTemplate = {
                          ...inboundOracleTemplate,
                        };
                        _inboundOracleTemplate.inboundOracleTemplates[
                          oracleI
                        ].inputs[i]["Type"] = value;
                        setInboundOracleTemplate(_inboundOracleTemplate);
                      }}
                    />
                  </Form.Group>
                ))}
                <div>
                  <Button
                    icon="close"
                    basic
                    negative
                    size="tiny"
                    content="Cancel Oracle"
                  />
                </div>
              </Segment>
            ) : (
              <>
                <Form.Group widths="equal">
                  <Form.Input
                    label={inbound ? "Contract Name" : "Event Name"}
                    name="ContractName"
                    value={oracle.ContractName}
                    placeholder="The name of the event"
                    onChange={(event, { value }) => {
                      let _inboundOracleTemplate = {
                        ...inboundOracleTemplate,
                      };
                      _inboundOracleTemplate.inboundOracleTemplates[oracleI][
                        "ContractName"
                      ] = value;
                      setInboundOracleTemplate(_inboundOracleTemplate);
                    }}
                  />
                </Form.Group>
                <label>
                  <b>Input Parameters</b>
                </label>
                {oracle.inputs.length === 0 && (
                  <p>No input parameters added yet</p>
                )}
                {oracle.inputs.map((input, i) => (
                  <Form.Group widths="equal">
                    <Form.Input
                      label="Name"
                      value={input.Name}
                      onChange={(_, { value }) => {
                        let _inboundOracleTemplate = {
                          ...inboundOracleTemplate,
                        };
                        _inboundOracleTemplate.inboundOracleTemplates[
                          oracleI
                        ].inputs[i]["Name"] = value;
                        setInboundOracleTemplate(_inboundOracleTemplate);
                      }}
                    />
                    <Form.Input
                      label="Type"
                      value={input.Type}
                      onChange={(event, { value }) => {
                        let _inboundOracleTemplate = {
                          ...inboundOracleTemplate,
                        };
                        _inboundOracleTemplate.inboundOracleTemplates[
                          oracleI
                        ].inputs[i]["Type"] = value;
                        setInboundOracleTemplate(_inboundOracleTemplate);
                      }}
                    />
                  </Form.Group>
                ))}
                <div>
                  <Button
                    icon="plus"
                    basic
                    primary
                    size="tiny"
                    content="Add Input"
                    onClick={() =>
                      setInboundOracleTemplate({
                        ...inboundOracleTemplate,
                        inboundOracleTemplates: [
                          {
                            ...inboundOracleTemplate.inboundOracleTemplates[0],
                            inputs: [
                              ...inboundOracleTemplate.inboundOracleTemplates[0]
                                .inputs,
                              { Name: "", Type: "" },
                            ],
                          },
                        ],
                      })
                    }
                  />
                </div>
              </>
            )
          )}
        </div>
      </Form>
    </div>
  );
}
