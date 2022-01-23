import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useSmartContractProvider(id) {
  const [smartContractProvider, setSmartContractProvider] = useState();

  async function fetchSmartContractProvider() {
    const data = await getData("/smartContractProviders/" + id);
    const _smartContractProvider = data.smartContractProvider;
    setSmartContractProvider(_smartContractProvider);
  }

  useEffect(() => {
    fetchSmartContractProvider();
  }, []); // eslint-disable-line

  return [smartContractProvider];
}
