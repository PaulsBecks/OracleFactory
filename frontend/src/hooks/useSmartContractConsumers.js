import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useInboundSubscription() {
  const [smartContractConsumers, setSmartContractConsumers] = useState([]);

  async function fetchSmartContractConsumers() {
    const _smartContractConsumers = await getData("/smartContractConsumers");
    setSmartContractConsumers(_smartContractConsumers.smartContractConsumers);
  }
  useEffect(() => {
    fetchSmartContractConsumers();
  }, []);
  return [smartContractConsumers];
}
