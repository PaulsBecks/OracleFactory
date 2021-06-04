import axios from "axios";
import config from "../config";
export default async function postData(url, data) {
  try {
    const response = await axios.post(config.BASE_URL + url, data);
    return response.data;
  } catch (err) {
    return [];
  }
}
