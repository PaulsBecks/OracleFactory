import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useWebServiceListeners() {
  const [webServiceListeners, setWebServiceListeners] = useState([]);
  async function fetchWebServiceListeners() {
    const _webServiceListeners = await getData("/webServiceListeners");
    setWebServiceListeners(_webServiceListeners.webServiceListeners);
  }
  useEffect(() => {
    fetchWebServiceListeners();
  }, []);
  return [webServiceListeners];
}
