import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useSmartContractProviders() {
  const [smartContractProviders, setSmartContractProviders] = useState([]);
  async function fetchSmartContractProviders() {
    const _smartContractProviders = await getData("/smartContractProviders");
    setSmartContractProviders(_smartContractProviders.smartContractProviders);
  }
  useEffect(() => {
    fetchSmartContractProviders();
  }, []);
  return [smartContractProviders];
}
