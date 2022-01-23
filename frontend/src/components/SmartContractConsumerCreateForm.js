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

export default function SmartContractConsumerForm({
  smartContractConsumer,
  setSmartContractConsumer,
  fromABI,
  inbound,
  outbound,
}) {
  function updateSmartContractConsumer(_, { name, value }) {
    setSmartContractConsumer({ ...smartContractConsumer, [name]: value });
  }

  const [abi, setAbi] = useState("[]");

  useEffect(() => {
    setAbi(formToAbi(smartContractConsumer.smartContractConsumers));
  }, [smartContractConsumer]);

  if (!smartContractConsumer) return "";

  return (
    <div>
      <Form>
        <Form.Group widths="equal">
          <Form.Input
            label={
              <label>
                {inbound ? "Consumer" : "Provider"} Name
                <Popup
                  content={`The name is displayed in the list and cards to identify the ${
                    inbound ? "consumer" : "provider"
                  }`}
                  trigger={<Icon name="info circle" />}
                />
              </label>
            }
            name="Name"
            value={smartContractConsumer.Name}
            onChange={updateSmartContractConsumer}
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
                    inbound ? "consumer" : "provider"
                  } is doing.`}
                  trigger={<Icon name="info circle" />}
                />
              </label>
            }
            name="Description"
            value={smartContractConsumer.Description}
            onChange={updateSmartContractConsumer}
            placeholder="A description"
          />
        </Form.Group>
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
            value={smartContractConsumer.BlockchainName}
            onChange={updateSmartContractConsumer}
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
            value={smartContractConsumer.ContractAddress}
            onChange={updateSmartContractConsumer}
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
            value={smartContractConsumer.ContractAddressSynonym}
            onChange={updateSmartContractConsumer}
            placeholder="The Subscription Name"
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
              string={smartContractConsumer.ContractAddress}
              size={50}
            />
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
              checked={smartContractConsumer.Private}
              label={smartContractConsumer.Private ? "private" : "public"}
              name="Private"
              toggle
              onChange={(event, { name, checked }) =>
                updateSmartContractConsumer(event, { name, value: checked })
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
                setSmartContractConsumer({
                  ...smartContractConsumer,
                  smartContractConsumers: parseAbi(abi),
                })
              }
            />
          </>
        )}
        <div>
          {smartContractConsumer.smartContractConsumers.map(
            (subscription, subscriptionI) =>
              fromABI ? (
                <Segment color="green" style={{ margin: "1em 1em 0 0" }}>
                  <h3>
                    {subscription.Type === "function" ? "Consumer" : "Provider"}
                  </h3>
                  <Form.Group widths="equal">
                    <Form.Input
                      label="Contract Name"
                      name="ContractName"
                      value={subscription.ContractName}
                      placeholder="The name of the event"
                      onChange={(event, { value }) => {
                        let _smartContractConsumer = {
                          ...smartContractConsumer,
                        };
                        _smartContractConsumer.smartContractConsumers[
                          subscriptionI
                        ]["ContractName"] = value;
                        setSmartContractConsumer(_smartContractConsumer);
                      }}
                    />
                  </Form.Group>
                  <Header>Input Parameters</Header>
                  {subscription.inputs.map((input, i) => (
                    <Form.Group widths="equal">
                      <Form.Input
                        label="Name"
                        value={input.Name}
                        onChange={(_, { value }) => {
                          let _smartContractConsumer = {
                            ...smartContractConsumer,
                          };
                          _smartContractConsumer.smartContractConsumers[
                            subscriptionI
                          ].inputs[i]["Name"] = value;
                          setSmartContractConsumer(_smartContractConsumer);
                        }}
                      />
                      <Form.Input
                        label="Type"
                        value={input.Type}
                        onChange={(event, { value }) => {
                          let _smartContractConsumer = {
                            ...smartContractConsumer,
                          };
                          _smartContractConsumer.smartContractConsumers[
                            subscriptionI
                          ].inputs[i]["Type"] = value;
                          setSmartContractConsumer(_smartContractConsumer);
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
                      content="Cancel Subscription"
                    />
                  </div>
                </Segment>
              ) : (
                <>
                  <Form.Group widths="equal">
                    <Form.Input
                      label={inbound ? "Contract Name" : "Event Name"}
                      name="ContractName"
                      value={subscription.ContractName}
                      placeholder="The name of the event"
                      onChange={(event, { value }) => {
                        let _smartContractConsumer = {
                          ...smartContractConsumer,
                        };
                        _smartContractConsumer.smartContractConsumers[
                          subscriptionI
                        ]["ContractName"] = value;
                        setSmartContractConsumer(_smartContractConsumer);
                      }}
                    />
                  </Form.Group>
                  <label>
                    <b>Input Parameters</b>
                  </label>
                  {subscription.inputs.length === 0 && (
                    <p>No input parameters added yet</p>
                  )}
                  {subscription.inputs.map((input, i) => (
                    <Form.Group widths="equal">
                      <Form.Input
                        label="Name"
                        value={input.Name}
                        onChange={(_, { value }) => {
                          let _smartContractConsumer = {
                            ...smartContractConsumer,
                          };
                          _smartContractConsumer.smartContractConsumers[
                            subscriptionI
                          ].inputs[i]["Name"] = value;
                          setSmartContractConsumer(_smartContractConsumer);
                        }}
                      />
                      <Form.Input
                        label="Type"
                        value={input.Type}
                        onChange={(event, { value }) => {
                          let _smartContractConsumer = {
                            ...smartContractConsumer,
                          };
                          _smartContractConsumer.smartContractConsumers[
                            subscriptionI
                          ].inputs[i]["Type"] = value;
                          setSmartContractConsumer(_smartContractConsumer);
                        }}
                      />
                      <Form.Checkbox
                        checked={input.Indexed}
                        label="Indexed"
                        name="Indexed"
                        toggle
                        onChange={(event, { name, checked }) => {
                          let _smartContractConsumer = {
                            ...smartContractConsumer,
                          };
                          _smartContractConsumer.smartContractConsumers[
                            subscriptionI
                          ].inputs[i]["Indexed"] = checked;
                          setSmartContractConsumer(_smartContractConsumer);
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
                        setSmartContractConsumer({
                          ...smartContractConsumer,
                          smartContractConsumers: [
                            {
                              ...smartContractConsumer
                                .smartContractConsumers[0],
                              inputs: [
                                ...smartContractConsumer
                                  .smartContractConsumers[0].inputs,
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
