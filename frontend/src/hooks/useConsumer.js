import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useConsumer(id) {
  const [consumer, setConsumer] = useState();

  async function fetchConsumer() {
    const _consumer = await getData("/consumers/" + id);
    setConsumer(_consumer.consumer);
  }

  useEffect(() => {
    fetchConsumer();
  }, []);

  return [consumer];
}
