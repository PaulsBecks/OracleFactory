import React, { useState } from "react";
import postData from "../services/postData";
import { SmartContractPublisherForm } from ".";
import { Button } from "semantic-ui-react";
import { Link, useHistory } from "react-router-dom";

export default function SmartContractCreate({ fromABI, inbound, outbound }) {
  const history = useHistory();
  const [smartContractPublisher, setSmartContractPublisher] = useState({
    Description: "",
    Name: "",
    Private: true,
    BlockchainName: "Ethereum",
    ContractAddress: "",
    ContractAddressSynonym: "",
    smartContractPublishers: fromABI
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
      <SmartContractPublisherForm
        smartContractPublisher={smartContractPublisher}
        setSmartContractPublisher={setSmartContractPublisher}
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
          for (const element of smartContractPublisher.smartContractPublishers) {
            let result;
            if (element.Type === "function") {
              result = await postData(`/smartContractPublishers`, {
                Name: smartContractPublisher.Name,
                Description: smartContractPublisher.Description,
                BlockchainAddress: smartContractPublisher.BlockchainAddress,
                BlockchainName: smartContractPublisher.BlockchainName,
                ContractAddress: smartContractPublisher.ContractAddress,
                ContractAddressSynonym:
                  smartContractPublisher.ContractAddressSynonym,
                ContractName: element.ContractName,
                Private: smartContractPublisher.Private,
              });
            } else {
              result = await postData(`/smartContractListeners`, {
                Name: smartContractPublisher.Name,
                Description: smartContractPublisher.Description,
                BlockchainAddress: smartContractPublisher.BlockchainAddress,
                BlockchainName: smartContractPublisher.BlockchainName,
                ContractAddress: smartContractPublisher.ContractAddress,
                ContractAddressSynonym:
                  smartContractPublisher.ContractAddressSynonym,
                EventName: element.ContractName,
                Private: smartContractPublisher.Private,
              });
            }
            const smartContract =
              element.Type === "function"
                ? result.smartContractPublisher
                : result.smartContractListener;
            for (const input of element.inputs) {
              if (element.Type === "function") {
                await postData(
                  `/smartContractPublishers/${smartContract.ID}/eventParameters`,
                  input
                );
              } else {
                await postData(
                  `/smartContractListeners/${smartContract.ID}/eventParameters`,
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
