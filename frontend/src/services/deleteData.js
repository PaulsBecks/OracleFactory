import axios from "axios";
import config from "../config";
import getHeaders from "./utils/getHeaders";
export default async function deleteData(url) {
  const headers = getHeaders();

  try {
    const response = await axios.delete(config.BASE_URL + url, { headers });
    return response.data;
  } catch (err) {
    if (err?.response?.status === 401) {
      localStorage.clear();
      window.location.reload();
    }
    return {};
  }
}
