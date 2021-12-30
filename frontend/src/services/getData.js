import axios from "axios";
import config from "../config";
import getHeaders from "./utils/getHeaders";
export default async function getData(url) {
  const headers = getHeaders();
  console.log(config, process.env.NODE_ENV);
  try {
    const response = await axios.get(config.BASE_URL + url, { headers });
    return response.data;
  } catch (err) {
    if (err?.response?.status === 401) {
      localStorage.clear();
      window.location.reload();
    }
    return {};
  }
}
