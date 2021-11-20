import { useEffect, useState } from "react";
import getData from "../services/getData";
import postData from "../services/postData";
import putData from "../services/putData";

export default function useOutboundOracle(id) {
  const [outboundOracle, setOutboundOracle] = useState();
  const [pubSubOracle, setPubSubOracle] = useState();
  const [loading, setLoading] = useState(false);

  async function fetchOutboundOracle() {
    const _outboundOracle = await getData("/outboundOracles/" + id);
    setOutboundOracle(_outboundOracle.outboundOracle);
    setPubSubOracle(_outboundOracle.pubSubOracle);
  }

  async function fetchWrapper(callback) {
    setLoading(true);
    await callback();
    await fetchOutboundOracle();
    setLoading(false);
  }

  async function updateOutboundOracle(data) {
    await fetchWrapper(
      async () => await putData("/outboundOracles/" + id, data)
    );
  }

  async function startOutboundOracle() {
    await fetchWrapper(
      async () => await postData("/outboundOracles/" + id + "/start")
    );
  }

  async function stopOutboundOracle() {
    await fetchWrapper(
      async () => await postData("/outboundOracles/" + id + "/stop")
    );
  }

  useEffect(() => {
    fetchOutboundOracle();
  }, []);

  return [
    outboundOracle,
    updateOutboundOracle,
    loading,
    startOutboundOracle,
    stopOutboundOracle,
    pubSubOracle,
  ];
}
