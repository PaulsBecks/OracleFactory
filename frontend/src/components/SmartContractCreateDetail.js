import React, { useState } from "react";
import postData from "../services/postData";
import { ConsumerForm } from ".";
import { Button } from "semantic-ui-react";
import { Link, useHistory } from "react-router-dom";

export default function SmartContractCreate({ fromABI, pubSub, outbound }) {
  const history = useHistory();
  const [consumer, setConsumer] = useState({
    Description: "",
    Name: "",
    Private: true,
    BlockchainName: "Ethereum",
    ContractAddress: "",
    ContractAddressSynonym: "",
    consumers: fromABI
      ? []
      : [
          {
            EventName: "",
            Type: pubSub ? "function" : "event",
            inputs: [],
          },
        ],
  });

  const [loading, setLoading] = useState(false);
  return (
    <div>
      <ConsumerForm
        consumer={consumer}
        setConsumer={setConsumer}
        fromABI={fromABI}
        pubSub={pubSub}
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
          for (const element of consumer.consumers) {
            let result;
            if (element.Type === "function") {
              result = await postData(`/consumers`, {
                Name: element.Name ? element.Name : consumer.Name,
                Description: element.Description
                  ? element.Description
                  : consumer.Description,
                BlockchainAddress: consumer.BlockchainAddress,
                BlockchainName: consumer.BlockchainName,
                ContractAddress: consumer.ContractAddress,
                ContractAddressSynonym: consumer.ContractAddressSynonym,
                EventName: element.EventName,
                Private: consumer.Private,
              });
            } else {
              result = await postData(`/blockchainEvents`, {
                Name: element.Name ? element.Name : consumer.Name,
                Description: element.Description
                  ? element.Description
                  : consumer.Description,
                BlockchainAddress: consumer.BlockchainAddress,
                BlockchainName: consumer.BlockchainName,
                ContractAddress: consumer.ContractAddress,
                ContractAddressSynonym: consumer.ContractAddressSynonym,
                EventName: element.EventName,
                Private: consumer.Private,
              });
            }
            const smartContract =
              element.Type === "function"
                ? result.consumer
                : result.blockchainEvent;
            for (const input of element.inputs) {
              if (element.Type === "function") {
                await postData(
                  `/consumers/${smartContract.ID}/eventParameters`,
                  input
                );
              } else {
                await postData(
                  `/blockchainEvents/${smartContract.ID}/eventParameters`,
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
