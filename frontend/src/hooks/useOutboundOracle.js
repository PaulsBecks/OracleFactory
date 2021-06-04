import { useEffect, useState } from "react";
import getData from "../services/getData";
import putData from "../services/putData";

export default function useOutboundOracle(id) {
  const [outboundOracle, setOutboundOracle] = useState();
  const [loading, setLoading] = useState(false);

  async function fetchOutboundOracle() {
    const _outboundOracle = await getData("/outboundOracles/" + id);
    setOutboundOracle(_outboundOracle.outboundOracle);
  }

  async function updateOutboundOracle(data) {
    setLoading(true);
    await putData("/outboundOracles/" + id, data);
    await fetchOutboundOracle();
    setLoading(false);
  }

  useEffect(() => {
    fetchOutboundOracle();
  }, []);

  return [outboundOracle, updateOutboundOracle, loading];
}
