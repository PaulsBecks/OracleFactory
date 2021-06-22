import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useInboundOracles() {
  const [inboundOracles, setInboundOracles] = useState([]);

  async function fetchInboundOracles() {
    const _inboundOracles = await getData("/inboundOracles");
    setInboundOracles(_inboundOracles.inboundOracles);
  }

  useEffect(() => {
    fetchInboundOracles();
  }, []);

  return [inboundOracles];
}
