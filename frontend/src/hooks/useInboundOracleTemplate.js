import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useInboundOracleTemplate(id) {
  const [inboundOracleTemplate, setInboundOracleTemplate] = useState();

  async function fetchInboundOracleTemplate() {
    const _inboundOracleTemplate = await getData(
      "/inboundOracleTemplates/" + id
    );
    setInboundOracleTemplate(_inboundOracleTemplate.inboundOracleTemplate);
  }

  useEffect(() => {
    fetchInboundOracleTemplate();
  }, []);

  return [inboundOracleTemplate];
}
