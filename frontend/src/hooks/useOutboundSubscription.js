import { useEffect, useState } from "react";
import getData from "../services/getData";
import postData from "../services/postData";
import putData from "../services/putData";

export default function useOutboundSubscription(id) {
  const [outboundSubscription, setOutboundSubscription] = useState();
  const [loading, setLoading] = useState(false);

  async function fetchOutboundSubscription() {
    const _outboundSubscription = await getData("/outboundSubscriptions/" + id);
    setOutboundSubscription(_outboundSubscription.outboundSubscription);
  }

  async function fetchWrapper(callback) {
    setLoading(true);
    await callback();
    await fetchOutboundSubscription();
    setLoading(false);
  }

  async function updateOutboundSubscription(data) {
    await fetchWrapper(
      async () => await putData("/outboundSubscriptions/" + id, data)
    );
  }

  async function startOutboundSubscription() {
    await fetchWrapper(
      async () => await postData("/outboundSubscriptions/" + id + "/start")
    );
  }

  async function stopOutboundSubscription() {
    await fetchWrapper(
      async () => await postData("/outboundSubscriptions/" + id + "/stop")
    );
  }

  useEffect(() => {
    fetchOutboundSubscription();
  }, []);

  return [
    outboundSubscription,
    updateOutboundSubscription,
    loading,
    startOutboundSubscription,
    stopOutboundSubscription,
  ];
}
