import { useEffect, useState } from "react";
import getData from "../services/getData";
import postData from "../services/postData";

export default function useHyperledgerBlockchainConnector(id) {
  const [hyperledgerBlockchainConnector, setHyperledgerBlockchainConnector] =
    useState();
  async function fetchHyperledgerBlockchainConnector() {
    const data = await getData("/hyperledgerConnectors/" + id);
    setHyperledgerBlockchainConnector(data.hyperledgerConnector);
  }
  useEffect(() => {
    fetchHyperledgerBlockchainConnector();
  }, []);

  const start = async () => {
    await postData(
      "/outboundOracles/" +
        hyperledgerBlockchainConnector.OutboundOracleID +
        "/start"
    );
    fetchHyperledgerBlockchainConnector();
  };

  const stop = async () => {
    await postData(
      "/outboundOracles/" +
        hyperledgerBlockchainConnector.OutboundOracleID +
        "/stop"
    );
    fetchHyperledgerBlockchainConnector();
  };

  return [hyperledgerBlockchainConnector, start, stop];
}
