import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useEventParameters(listenerPublisherID) {
  const [parameters, setParameters] = useState([]);
  console.log(listenerPublisherID);

  async function fetchParameters() {
    const data = await getData(
      "/listenerPublishers/" + listenerPublisherID + "/eventParameters"
    );
    setParameters(data.eventParameters);
  }

  useEffect(() => {
    fetchParameters();
  }, []); // eslint-disable-line

  return [parameters];
}
