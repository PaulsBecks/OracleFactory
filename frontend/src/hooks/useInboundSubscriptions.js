import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useInboundSubscriptions() {
  const [inboundSubscriptions, setInboundSubscriptions] = useState([]);

  async function fetchInboundSubscriptions() {
    const _inboundSubscriptions = await getData("/inboundSubscriptions");
    setInboundSubscriptions(_inboundSubscriptions.inboundSubscriptions);
  }

  useEffect(() => {
    fetchInboundSubscriptions();
  }, []);

  return [inboundSubscriptions];
}
