import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useInboundOracle() {
  const [inboundOracleTemplates, setInboundOracleTemplates] = useState([]);

  async function fetchInboundOracleTemplates() {
    const _inboundOracleTemplates = await getData("/inboundOracleTemplates");
    setInboundOracleTemplates(_inboundOracleTemplates.inboundOracleTemplates);
  }
  useEffect(() => {
    fetchInboundOracleTemplates();
  }, []);
  return [inboundOracleTemplates];
}
