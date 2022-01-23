import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useSmartContractConsumer(id) {
  const [smartContractConsumer, setSmartContractConsumer] = useState();

  async function fetchSmartContractConsumer() {
    const _smartContractConsumer = await getData(
      "/smartContractConsumers/" + id
    );
    setSmartContractConsumer(_smartContractConsumer.smartContractConsumer);
  }

  useEffect(() => {
    fetchSmartContractConsumer();
  }, []);

  return [smartContractConsumer];
}
