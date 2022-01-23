import React, { useState } from "react";
import postData from "../services/postData";
import { SmartContractConsumerForm } from ".";
import { Button } from "semantic-ui-react";
import { Link, useHistory } from "react-router-dom";

export default function SmartContractCreate({ fromABI, inbound, outbound }) {
  const history = useHistory();
  const [smartContractConsumer, setSmartContractConsumer] = useState({
    Description: "",
    Name: "",
    Private: true,
    BlockchainName: "Ethereum",
    ContractAddress: "",
    ContractAddressSynonym: "",
    smartContractConsumers: fromABI
      ? []
      : [
          {
            ContractName: "",
            Type: inbound ? "function" : "event",
            inputs: [],
          },
        ],
  });

  const [loading, setLoading] = useState(false);
  return (
    <div>
      <SmartContractConsumerForm
        smartContractConsumer={smartContractConsumer}
        setSmartContractConsumer={setSmartContractConsumer}
        fromABI={fromABI}
        inbound={inbound}
        outbound={outbound}
      />
      <br />
      <Button
        loading={loading}
        basic
        positive
        fluid
        floated="right"
        content="Create"
        onClick={async () => {
          setLoading(true);
          for (const element of smartContractConsumer.smartContractConsumers) {
            let result;
            if (element.Type === "function") {
              result = await postData(`/smartContractConsumers`, {
                Name: smartContractConsumer.Name,
                Description: smartContractConsumer.Description,
                BlockchainAddress: smartContractConsumer.BlockchainAddress,
                BlockchainName: smartContractConsumer.BlockchainName,
                ContractAddress: smartContractConsumer.ContractAddress,
                ContractAddressSynonym:
                  smartContractConsumer.ContractAddressSynonym,
                ContractName: element.ContractName,
                Private: smartContractConsumer.Private,
              });
            } else {
              result = await postData(`/smartContractProviders`, {
                Name: smartContractConsumer.Name,
                Description: smartContractConsumer.Description,
                BlockchainAddress: smartContractConsumer.BlockchainAddress,
                BlockchainName: smartContractConsumer.BlockchainName,
                ContractAddress: smartContractConsumer.ContractAddress,
                ContractAddressSynonym:
                  smartContractConsumer.ContractAddressSynonym,
                EventName: element.ContractName,
                Private: smartContractConsumer.Private,
              });
            }
            const smartContract =
              element.Type === "function"
                ? result.smartContractConsumer
                : result.smartContractProvider;
            for (const input of element.inputs) {
              if (element.Type === "function") {
                await postData(
                  `/smartContractConsumers/${smartContract.ID}/eventParameters`,
                  input
                );
              } else {
                await postData(
                  `/smartContractProviders/${smartContract.ID}/eventParameters`,
                  input
                );
              }
            }
          }
          setLoading(false);
          history.push("/smartContracts");
        }}
      />
      {/*<Button
        loading={loading}
        basic
        negative
        floated="right"
        content="Cancel"
        as={Link}
        to={"/"}
      />*/}
      <br />
      <br />
    </div>
  );
}
