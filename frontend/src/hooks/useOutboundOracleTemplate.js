import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useOutboundOracleTemplate(id) {
  const [outboundOracleTemplate, setOutboundOracleTemplate] = useState();

  async function fetchOutboundOracleTemplate() {
    const data = await getData("/outboundOracleTemplates/" + id);
    const _outboundOracleTemplate = data.outboundOracleTemplate;
    setOutboundOracleTemplate(_outboundOracleTemplate);
  }

  useEffect(() => {
    fetchOutboundOracleTemplate();
  }, []); // eslint-disable-line

  return [outboundOracleTemplate];
}
