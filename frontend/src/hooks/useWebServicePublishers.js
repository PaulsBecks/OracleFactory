import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useWebServicePublishers() {
  const [webServicePublishers, setWebServicePublishers] = useState([]);
  async function fetchWebServicePublishers() {
    const _webServicePublishers = await getData("/webServicePublishers");
    setWebServicePublishers(_webServicePublishers.webServicePublishers);
  }
  useEffect(() => {
    fetchWebServicePublishers();
  }, []);
  return [webServicePublishers];
}
