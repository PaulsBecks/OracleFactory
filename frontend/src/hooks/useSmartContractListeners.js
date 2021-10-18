import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useSmartContractListeners() {
  const [smartContractListeners, setSmartContractListeners] = useState([]);
  async function fetchSmartContractListeners() {
    const _smartContractListeners = await getData("/smartContractListeners");
    setSmartContractListeners(_smartContractListeners.smartContractListeners);
  }
  useEffect(() => {
    fetchSmartContractListeners();
  }, []);
  return [smartContractListeners];
}
