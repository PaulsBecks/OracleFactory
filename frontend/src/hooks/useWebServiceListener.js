import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useWebServiceListener(id) {
  const [webServiceListener, setWebServiceListener] = useState();

  async function fetchWebServiceListener() {
    const data = await getData("/webServiceListeners/" + id);
    const _webServiceListener = data.webServiceListener;
    setWebServiceListener(_webServiceListener);
  }

  useEffect(() => {
    fetchWebServiceListener();
  }, []); // eslint-disable-line

  return [webServiceListener];
}
