import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useWebServicePublisher(id) {
  const [webServicePublisher, setWebServicePublisher] = useState();

  async function fetchWebServicePublisher() {
    const data = await getData("/webServicePublishers/" + id);
    const _webServicePublisher = data.webServicePublisher;
    setWebServicePublisher(_webServicePublisher);
  }

  useEffect(() => {
    fetchWebServicePublisher();
  }, []); // eslint-disable-line

  return [webServicePublisher];
}
