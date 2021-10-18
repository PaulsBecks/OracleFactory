import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useSmartContractListener(id) {
  const [smartContractListener, setSmartContractListener] = useState();

  async function fetchSmartContractListener() {
    const data = await getData("/smartContractListeners/" + id);
    const _smartContractListener = data.smartContractListener;
    setSmartContractListener(_smartContractListener);
  }

  useEffect(() => {
    fetchSmartContractListener();
  }, []); // eslint-disable-line

  return [smartContractListener];
}
