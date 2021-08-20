import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useEventParameters(OracleTemplateID) {
  const [parameters, setParameters] = useState([]);

  async function fetchParameters() {
    const data = await getData(
      "/oracleTemplates/" + OracleTemplateID + "/eventParameters"
    );
    setParameters(data.eventParameters);
  }

  useEffect(() => {
    fetchParameters();
  }, []); // eslint-disable-line

  return [parameters];
}
