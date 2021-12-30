import { useState } from "react";
import { useHistory } from "react-router";
import { Button } from "semantic-ui-react";
import HyperledgerBlockchainConnectorForm from "../components/HyperledgerBlockchainConnectorForm";
import postData from "../services/postData";

export function HyperledgerBlockchainConnectorCreate() {
  const [hyperledgerBlockchainConnector, setHyperledgerBlockchainConnector] =
    useState({ IsOnChain: false });
  const history = useHistory();
  return (
    <div>
      <HyperledgerBlockchainConnectorForm
        hyperledgerBlockchainConnector={hyperledgerBlockchainConnector}
        setHyperledgerBlockchainConnector={setHyperledgerBlockchainConnector}
      />
      <br />
      <Button
        content="Create"
        positive
        basic
        onClick={async () => {
          const result = await postData(
            "/hyperledgerConnectors",
            hyperledgerBlockchainConnector
          );
          history.push("/");
        }}
      />
    </div>
  );
}
