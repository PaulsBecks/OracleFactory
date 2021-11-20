import { useEffect, useState } from "react";
import getData from "../services/getData";
import postData from "../services/postData";
import putData from "../services/putData";

export default function usePubSubOracle(id) {
  const [pubSubOracle, setPubSubOracle] = useState();
  const [loading, setLoading] = useState(false);

  async function fetchPubSubOracle() {
    const _pubSubOracle = await getData("/pubSubOracles/" + id);
    setPubSubOracle(_pubSubOracle.pubSubOracle);
  }
  async function fetchWrapper(callback) {
    setLoading(true);
    await callback();
    await fetchPubSubOracle();
    setLoading(false);
  }

  async function updatePubSubOracle(data) {
    await fetchWrapper(async () => await putData("/pubSubOracles/" + id, data));
  }

  async function startPubSubOracle() {
    await fetchWrapper(
      async () => await postData("/pubSubOracles/" + id + "/start")
    );
  }

  async function stopPubSubOracle() {
    await fetchWrapper(
      async () => await postData("/pubSubOracles/" + id + "/stop")
    );
  }

  useEffect(() => {
    fetchPubSubOracle();
  }, []);

  return [
    pubSubOracle,
    updatePubSubOracle,
    loading,
    startPubSubOracle,
    stopPubSubOracle,
  ];
}
