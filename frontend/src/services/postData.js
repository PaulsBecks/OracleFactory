import axios from "axios";
import config from "../config";
import getHeaders from "./utils/getHeaders";

export default async function postData(url, data) {
  const headers = getHeaders();
  try {
    const response = await axios.post(config.BASE_URL + url, data, { headers });
    return response.data;
  } catch (err) {
    return [];
  }
}
