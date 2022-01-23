import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useWebServiceProvider(id) {
  const [webServiceProvider, setWebServiceProvider] = useState();

  async function fetchWebServiceProvider() {
    const data = await getData("/webServiceProviders/" + id);
    const _webServiceProvider = data.webServiceProvider;
    setWebServiceProvider(_webServiceProvider);
  }

  useEffect(() => {
    fetchWebServiceProvider();
  }, []); // eslint-disable-line

  return [webServiceProvider];
}
