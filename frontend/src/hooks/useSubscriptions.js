import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useSubscriptions() {
  const [subscriptions, setSubscriptions] = useState([]);
  async function fetchSubscriptions() {
    const _subscriptions = await getData("/subscriptions");
    setSubscriptions(_subscriptions.Subscriptions);
  }
  useEffect(() => {
    fetchSubscriptions();
  }, []);
  return [subscriptions];
}
