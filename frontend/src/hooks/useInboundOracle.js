import { useEffect, useState } from "react";
import getData from "../services/getData";
import putData from "../services/putData";

export default function useInboundOracle(id) {
  const [inboundOracle, setInboundOracle] = useState();
  const [loading, setLoading] = useState(false);

  async function fetchInboundOracle() {
    const _inboundOracle = await getData("/inboundOracles/" + id);
    setInboundOracle(_inboundOracle.inboundOracle);
  }

  async function updateInboundOracle(data) {
    setLoading(true);
    await putData("/inboundOracles/" + id, data);
    await fetchInboundOracle();
    setLoading(false);
  }

  useEffect(() => {
    fetchInboundOracle();
  }, []);

  return [inboundOracle, updateInboundOracle, loading];
}
