import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useOutboundOracles() {
  const [outboundOracles, setOutboundOracles] = useState([]);

  async function fetchOutboundOracles() {
    const _outboundOracles = await getData("/outboundOracles");
    setOutboundOracles(_outboundOracles.outboundOracles);
  }

  useEffect(() => {
    fetchOutboundOracles();
  }, []);

  return [outboundOracles];
}
