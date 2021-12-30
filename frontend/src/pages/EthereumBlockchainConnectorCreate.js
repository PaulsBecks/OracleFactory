import { useState } from "react";
import { useHistory } from "react-router";
import { Button } from "semantic-ui-react";
import EthereumBlockchainConnectorForm from "../components/EthereumBlockchainConnectorForm";
import postData from "../services/postData";

export function EthereumBlockchainConnectorCreate() {
  const [ethereumBlockchainConnector, setEthereumBlockchainConnector] =
    useState({ IsOnChain: false });
  const history = useHistory();
  return (
    <div>
      <EthereumBlockchainConnectorForm
        ethereumBlockchainConnector={ethereumBlockchainConnector}
        setEthereumBlockchainConnector={setEthereumBlockchainConnector}
      />
      <br />
      <Button
        positive
        content="Create"
        onClick={async () => {
          const result = await postData(
            "/ethereumConnectors",
            ethereumBlockchainConnector
          );
          history.push("/");
        }}
      />
    </div>
  );
}
