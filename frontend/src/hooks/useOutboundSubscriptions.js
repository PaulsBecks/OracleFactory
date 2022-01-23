import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useOutboundSubscriptions() {
  const [outboundSubscriptions, setOutboundSubscriptions] = useState([]);

  async function fetchOutboundSubscriptions() {
    const _outboundSubscriptions = await getData("/outboundSubscriptions");
    setOutboundSubscriptions(_outboundSubscriptions.outboundSubscriptions);
  }

  useEffect(() => {
    fetchOutboundSubscriptions();
  }, []);

  return [outboundSubscriptions];
}
