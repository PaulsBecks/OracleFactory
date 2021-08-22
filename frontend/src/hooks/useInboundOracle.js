import { useEffect, useState } from "react";
import getData from "../services/getData";
import postData from "../services/postData";
import putData from "../services/putData";

export default function useInboundOracle(id) {
  const [inboundOracle, setInboundOracle] = useState();
  const [loading, setLoading] = useState(false);

  async function fetchInboundOracle() {
    const _inboundOracle = await getData("/inboundOracles/" + id);
    setInboundOracle(_inboundOracle.inboundOracle);
  }
  async function fetchWrapper(callback) {
    setLoading(true);
    await callback();
    await fetchInboundOracle();
    setLoading(false);
  }

  async function updateInboundOracle(data) {
    await fetchWrapper(
      async () => await putData("/inboundOracles/" + id, data)
    );
  }

  async function startInboundOracle() {
    await fetchWrapper(
      async () => await postData("/inboundOracles/" + id + "/start")
    );
  }

  async function stopInboundOracle() {
    await fetchWrapper(
      async () => await postData("/inboundOracles/" + id + "/stop")
    );
  }

  useEffect(() => {
    fetchInboundOracle();
  }, []);

  return [
    inboundOracle,
    updateInboundOracle,
    loading,
    startInboundOracle,
    stopInboundOracle,
  ];
}
