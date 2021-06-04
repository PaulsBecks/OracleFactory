import axios from "axios";
import config from "../config";
export default async function putData(url, data) {
  try {
    const response = await axios.put(config.BASE_URL + url, data);
    return response.data;
  } catch (err) {
    return [];
  }
}
