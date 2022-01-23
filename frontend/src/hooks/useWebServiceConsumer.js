import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useWebServiceConsumer(id) {
  const [webServiceConsumer, setWebServiceConsumer] = useState();

  async function fetchWebServiceConsumer() {
    const data = await getData("/webServiceConsumers/" + id);
    const _webServiceConsumer = data.webServiceConsumer;
    setWebServiceConsumer(_webServiceConsumer);
  }

  useEffect(() => {
    fetchWebServiceConsumer();
  }, []); // eslint-disable-line

  return [webServiceConsumer];
}
