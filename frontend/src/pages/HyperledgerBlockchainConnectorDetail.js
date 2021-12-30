import { useParams } from "react-router";
import useHyperledgerBlockchainConnector from "../hooks/useHyperledgerBlockchainConnector";

export function HyperledgerBlockchainConnectorDetail() {
  const { hyperledgerBlockchainConnectorID } = useParams();
  const [hyperledgerConnector, start, stop] = useHyperledgerBlockchainConnector(
    hyperledgerBlockchainConnectorID
  );
  if (!hyperledgerConnector) {
    return "";
  }
  return <div>{JSON.stringify(hyperledgerConnector)}</div>;
}
