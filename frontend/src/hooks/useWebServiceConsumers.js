import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useWebServiceConsumers() {
  const [webServiceConsumers, setWebServiceConsumers] = useState([]);
  async function fetchWebServiceConsumers() {
    const _webServiceConsumers = await getData("/webServiceConsumers");
    setWebServiceConsumers(_webServiceConsumers.webServiceConsumers);
  }
  useEffect(() => {
    fetchWebServiceConsumers();
  }, []);
  return [webServiceConsumers];
}
