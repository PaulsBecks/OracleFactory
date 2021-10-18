import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useSmartContractPublisher(id) {
  const [smartContractPublisher, setSmartContractPublisher] = useState();

  async function fetchSmartContractPublisher() {
    const _smartContractPublisher = await getData(
      "/smartContractPublishers/" + id
    );
    setSmartContractPublisher(_smartContractPublisher.smartContractPublisher);
  }

  useEffect(() => {
    fetchSmartContractPublisher();
  }, []);

  return [smartContractPublisher];
}
