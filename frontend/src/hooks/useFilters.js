import { useEffect, useState } from "react";
import getData from "../services/getData";

export default function useFilters() {
  const [filters, setFilters] = useState([]);

  async function fetchFilters() {
    const data = await getData("/filters");
    setFilters(data.filters);
  }

  useEffect(() => {
    fetchFilters();
  }, []);

  return [filters];
}
