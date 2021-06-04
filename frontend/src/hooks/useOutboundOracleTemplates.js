import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useOutboundOracleTemplates() {
  const [outboundOracleTemplates, setOutboundOracleTemplates] = useState([]);

  async function fetchOutboundOracleTemplates() {
    const _outboundOracleTemplates = await getData("/outboundOracleTemplates");
    setOutboundOracleTemplates(
      _outboundOracleTemplates.outboundOracleTemplates
    );
  }
  useEffect(() => {
    fetchOutboundOracleTemplates();
  }, []);
  return [outboundOracleTemplates];
}
