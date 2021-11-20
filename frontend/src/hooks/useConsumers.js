import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function usePubSubOracle() {
  const [consumers, setConsumers] = useState([]);

  async function fetchConsumers() {
    const _consumers = await getData("/consumers");
    setConsumers(_consumers.consumers);
  }
  useEffect(() => {
    fetchConsumers();
  }, []);
  return [consumers];
}
