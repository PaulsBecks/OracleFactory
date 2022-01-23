import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useEventParameters(providerConsumerID) {
  const [parameters, setParameters] = useState([]);
  console.log(providerConsumerID);

  async function fetchParameters() {
    const data = await getData(
      "/providerConsumers/" + providerConsumerID + "/eventParameters"
    );
    setParameters(data.eventParameters);
  }

  useEffect(() => {
    fetchParameters();
  }, []); // eslint-disable-line

  return [parameters];
}
