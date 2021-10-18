import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useInboundOracle() {
  const [smartContractPublishers, setSmartContractPublishers] = useState([]);

  async function fetchSmartContractPublishers() {
    const _smartContractPublishers = await getData("/smartContractPublishers");
    setSmartContractPublishers(
      _smartContractPublishers.smartContractPublishers
    );
  }
  useEffect(() => {
    fetchSmartContractPublishers();
  }, []);
  return [smartContractPublishers];
}
