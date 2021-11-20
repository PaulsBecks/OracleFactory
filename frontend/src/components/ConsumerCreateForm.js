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

export default function ConsumerForm({
  consumer,
  setConsumer,
  fromABI,
  pubSub,
  outbound,
}) {
  function updateConsumer(_, { name, value }) {
    setConsumer({ ...consumer, [name]: value });
  }

  const [abi, setAbi] = useState("[]");

  useEffect(() => {
    setAbi(formToAbi(consumer.consumers));
  }, [consumer]);

  if (!consumer) return "";

  return (
    <div>
      <Form>
        {!fromABI && (
          <>
            <Form.Group widths="equal">
              <Form.Input
                label={
                  <label>
                    {pubSub ? "Publisher" : "Listener"} Name
                    <Popup
                      content={`The name is displayed in the list and cards to identify the ${
                        pubSub ? "publisher" : "listener"
                      }`}
                      trigger={<Icon name="info circle" />}
                    />
                  </label>
                }
                name="Name"
                value={consumer.Name}
                onChange={updateConsumer}
                placeholder="The name"
              />
            </Form.Group>
            <Form.Group widths="equal">
              <Form.Input
                label={
                  <label>
                    Description{" "}
                    <Popup
                      content={`Describe what this ${
                        pubSub ? "publisher" : "listener"
                      } is doing.`}
                      trigger={<Icon name="info circle" />}
                    />
                  </label>
                }
                name="Description"
                value={consumer.Description}
                onChange={updateConsumer}
                placeholder="A description"
              />
            </Form.Group>
          </>
        )}
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
            value={consumer.BlockchainName}
            onChange={updateConsumer}
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
            value={consumer.ContractAddress}
            onChange={updateConsumer}
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
            value={consumer.ContractAddressSynonym}
            onChange={updateConsumer}
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
            <Identicon string={consumer.ContractAddress} size={50} />
          </Form.Field>
          <Form.Field>
            <label>
              Visibility{" "}
              <Popup
                content="Deactivate the toggle to share the oracle template with other users."
                trigger={<Icon name="info circle" />}
              />
            </label>
            <Form.Checkbox
              checked={consumer.Private}
              label={consumer.Private ? "private" : "public"}
              name="Private"
              toggle
              onChange={(event, { name, checked }) =>
                updateConsumer(event, { name, value: checked })
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
                setConsumer({
                  ...consumer,
                  consumers: parseAbi(abi),
                })
              }
            />
          </>
        )}
        <div>
          {consumer.consumers.map((oracle, oracleI) =>
            fromABI ? (
              <Segment color="green" style={{ margin: "1em 1em 0 0" }}>
                <h3>{oracle.Type === "function" ? "Publisher" : "Listener"}</h3>
                <Form.Group widths="equal">
                  <Form.Input
                    label={
                      <label>
                        Description{" "}
                        <Popup
                          content={`Describe what this ${
                            pubSub ? "publisher" : "listener"
                          } is doing.`}
                          trigger={<Icon name="info circle" />}
                        />
                      </label>
                    }
                    name="Description"
                    value={oracle.Description}
                    onChange={(event, { value }) => {
                      let _consumer = {
                        ...consumer,
                      };
                      _consumer.consumers[oracleI]["Description"] = value;
                      setConsumer(_consumer);
                    }}
                    placeholder="A description"
                  />
                </Form.Group>
                <Form.Group widths="equal">
                  <Form.Input
                    label="Contract Name"
                    name="ContractName"
                    value={oracle.ContractName}
                    placeholder="The name of the event"
                    onChange={(event, { value }) => {
                      let _consumer = {
                        ...consumer,
                      };
                      _consumer.consumers[oracleI]["ContractName"] = value;
                      setConsumer(_consumer);
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
                        let _consumer = {
                          ...consumer,
                        };
                        _consumer.consumers[oracleI].inputs[i]["Name"] = value;
                        setConsumer(_consumer);
                      }}
                    />
                    <Form.Input
                      label="Type"
                      value={input.Type}
                      onChange={(event, { value }) => {
                        let _consumer = {
                          ...consumer,
                        };
                        _consumer.consumers[oracleI].inputs[i]["Type"] = value;
                        setConsumer(_consumer);
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
                    label={pubSub ? "Contract Name" : "Event Name"}
                    name="ContractName"
                    value={oracle.ContractName}
                    placeholder="The name of the event"
                    onChange={(event, { value }) => {
                      let _consumer = {
                        ...consumer,
                      };
                      _consumer.consumers[oracleI]["ContractName"] = value;
                      setConsumer(_consumer);
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
                        let _consumer = {
                          ...consumer,
                        };
                        _consumer.consumers[oracleI].inputs[i]["Name"] = value;
                        setConsumer(_consumer);
                      }}
                    />
                    <Form.Input
                      label="Type"
                      value={input.Type}
                      onChange={(event, { value }) => {
                        let _consumer = {
                          ...consumer,
                        };
                        _consumer.consumers[oracleI].inputs[i]["Type"] = value;
                        setConsumer(_consumer);
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
                      setConsumer({
                        ...consumer,
                        consumers: [
                          {
                            ...consumer.consumers[0],
                            inputs: [
                              ...consumer.consumers[0].inputs,
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
