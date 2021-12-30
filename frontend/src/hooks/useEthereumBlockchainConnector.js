import { useEffect, useState } from "react";
import getData from "../services/getData";
import postData from "../services/postData";

export default function useEthereumBlockchainConnector(id) {
  const [ethereumBlockchainConnector, setEthereumBlockchainConnector] =
    useState();
  async function fetchEthereumBlockchainConnector() {
    const _ethereumBlockchainConnector = await getData(
      "/ethereumConnectors/" + id
    );
    setEthereumBlockchainConnector(
      _ethereumBlockchainConnector.ethereumConnector
    );
  }
  useEffect(() => {
    fetchEthereumBlockchainConnector();
  }, []);

  const start = async () => {
    await postData(
      "/outboundOracles/" +
        ethereumBlockchainConnector.OutboundOracleID +
        "/start"
    );
    fetchEthereumBlockchainConnector();
  };

  const stop = async () => {
    await postData(
      "/outboundOracles/" +
        ethereumBlockchainConnector.OutboundOracleID +
        "/stop"
    );
    fetchEthereumBlockchainConnector();
  };

  return [ethereumBlockchainConnector, start, stop];
}
