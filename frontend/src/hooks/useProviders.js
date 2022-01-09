import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useProviders() {
  const [providers, setProviders] = useState([]);
  async function fetchProviders() {
    const _providers = await getData("/providers");
    setProviders(_providers.Providers);
  }
  useEffect(() => {
    fetchProviders();
  }, []);
  return [providers];
}
