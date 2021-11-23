import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useSubscription(id) {
  const [subscription, setSubscription] = useState();

  async function fetchSubscription() {
    const data = await getData("/subscriptions/" + id);
    const _subscription = data.subscription;
    setSubscription(_subscription);
  }

  useEffect(() => {
    fetchSubscription();
  }, []); // eslint-disable-line

  return [subscription];
}
