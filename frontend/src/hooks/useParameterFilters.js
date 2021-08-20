import { useEffect, useState } from "react";
import getData from "../services/getData";
import deleteData from "../services/deleteData";
import postData from "../services/postData";

export default function useParameterFilters(oracleID) {
  const [parameterFilters, setParameterFilters] = useState([]);
  const [loading, setLoading] = useState(false);

  async function fetchParameterFilters() {
    setLoading(true);
    const data = await getData("/oracles/" + oracleID + "/parameterFilters");
    setParameterFilters(data.parameterFilters);
    setLoading(false);
  }

  async function createParameterFilter(parameterFilter) {
    setLoading(true);
    await postData(
      "/oracles/" + oracleID + "/parameterFilters",
      parameterFilter
    );
    fetchParameterFilters();
    setLoading(false);
  }

  async function deleteParameterFilter(parameterFilterId) {
    setLoading(true);
    await deleteData(
      "/oracles/" + oracleID + "/parameterFilters/" + parameterFilterId
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
