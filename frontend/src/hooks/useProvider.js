import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useProvider(id) {
  const [provider, setProvider] = useState();

  async function fetchProvider() {
    const data = await getData("/providers/" + id);
    const _provider = data.provider;
    setProvider(_provider);
  }

  useEffect(() => {
    fetchProvider();
  }, []); // eslint-disable-line

  return [provider];
}
