import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useHyperledgerBlockchainConnectors() {
  const [hyperledgerBlockchainConnectors, setHyperledgerBlockchainConnectors] =
    useState([]);
  async function fetchHyperledgerBlockchainConnectors() {
    const _hyperledgerBlockchainConnectors = await getData(
      "/hyperledgerConnectors"
    );
    setHyperledgerBlockchainConnectors(
      _hyperledgerBlockchainConnectors.hyperledgerConnectors
    );
  }
  useEffect(() => {
    fetchHyperledgerBlockchainConnectors();
  }, []);
  return [hyperledgerBlockchainConnectors];
}
