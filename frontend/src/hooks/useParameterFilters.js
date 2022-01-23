import { useEffect, useState } from "react";
import getData from "../services/getData";
import deleteData from "../services/deleteData";
import postData from "../services/postData";

export default function useParameterFilters(subscriptionID) {
  const [parameterFilters, setParameterFilters] = useState([]);
  const [loading, setLoading] = useState(false);

  async function fetchParameterFilters() {
    setLoading(true);
    const data = await getData(
      "/subscriptions/" + subscriptionID + "/parameterFilters"
    );
    setParameterFilters(data.parameterFilters);
    setLoading(false);
  }

  async function createParameterFilter(parameterFilter) {
    setLoading(true);
    await postData(
      "/subscriptions/" + subscriptionID + "/parameterFilters",
      parameterFilter
    );
    fetchParameterFilters();
    setLoading(false);
  }

  async function deleteParameterFilter(parameterFilterId) {
    setLoading(true);
    await deleteData(
      "/subscriptions/" +
        subscriptionID +
        "/parameterFilters/" +
        parameterFilterId
    );
    fetchParameterFilters();
    setLoading(false);
  }

  useEffect(() => {
    fetchParameterFilters();
  }, []); // eslint-disable-line

  return [
    parameterFilters,
    createParameterFilter,
    deleteParameterFilter,
    loading,
  ];
}
