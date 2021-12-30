import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useEthereumBlockchainConnectors() {
  const [ethereumBlockchainConnectors, setEthereumBlockchainConnectors] =
    useState([]);
  async function fetchEthereumBlockchainConnectors() {
    const _ethereumBlockchainConnectors = await getData("/ethereumConnectors");
    setEthereumBlockchainConnectors(
      _ethereumBlockchainConnectors.ethereumConnectors
    );
  }
  useEffect(() => {
    fetchEthereumBlockchainConnectors();
  }, []);

  return [ethereumBlockchainConnectors];
}
