import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useSubscriptions() {
  const [subscriptions, setSubscriptions] = useState([]);
  async function fetchSubscriptions() {
    const _subscriptions = await getData("/subscriptions");
    setSubscriptions(_subscriptions.subscriptions);
  }
  useEffect(() => {
    fetchSubscriptions();
  }, []);
  return [subscriptions];
}
