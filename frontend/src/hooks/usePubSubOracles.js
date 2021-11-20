import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function usePubSubOracles() {
  const [pubSubOracles, setPubSubOracles] = useState([]);

  async function fetchPubSubOracles() {
    const _pubSubOracles = await getData("/pubSubOracles");
    setPubSubOracles(_pubSubOracles.pubSubOracles);
  }

  useEffect(() => {
    fetchPubSubOracles();
  }, []);

  return [pubSubOracles];
}
