import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useOutboundOracleTemplate(id) {
  const [outboundOracleTemplate, setOutboundOracleTemplate] = useState();

  async function fetchOutboundOracleTemplate() {
    const _outboundOracleTemplate = await getData(
      "/outboundOracleTemplates/" + id
    );
    setOutboundOracleTemplate(_outboundOracleTemplate.outboundOracleTemplate);
  }

  useEffect(() => {
    fetchOutboundOracleTemplate();
  }, []);

  return [outboundOracleTemplate];
}
