import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useWebServiceProviders() {
  const [webServiceProviders, setWebServiceProviders] = useState([]);
  async function fetchWebServiceProviders() {
    const _webServiceProviders = await getData("/webServiceProviders");
    setWebServiceProviders(_webServiceProviders.webServiceProviders);
  }
  useEffect(() => {
    fetchWebServiceProviders();
  }, []);
  return [webServiceProviders];
}
