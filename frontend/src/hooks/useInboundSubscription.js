import { useEffect, useState } from "react";
import getData from "../services/getData";
import postData from "../services/postData";
import putData from "../services/putData";

export default function useInboundSubscription(id) {
  const [inboundSubscription, setInboundSubscription] = useState();
  const [loading, setLoading] = useState(false);

  async function fetchInboundSubscription() {
    const _inboundSubscription = await getData("/inboundSubscriptions/" + id);
    setInboundSubscription(_inboundSubscription.inboundSubscription);
  }
  async function fetchWrapper(callback) {
    setLoading(true);
    await callback();
    await fetchInboundSubscription();
    setLoading(false);
  }

  async function updateInboundSubscription(data) {
    await fetchWrapper(
      async () => await putData("/inboundSubscriptions/" + id, data)
    );
  }

  async function startInboundSubscription() {
    await fetchWrapper(
      async () => await postData("/inboundSubscriptions/" + id + "/start")
    );
  }

  async function stopInboundSubscription() {
    await fetchWrapper(
      async () => await postData("/inboundSubscriptions/" + id + "/stop")
    );
  }

  useEffect(() => {
    fetchInboundSubscription();
  }, []);

  return [
    inboundSubscription,
    updateInboundSubscription,
    loading,
    startInboundSubscription,
    stopInboundSubscription,
  ];
}
