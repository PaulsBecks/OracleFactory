import axios from "axios";
import config from "../config";
export default async function getData(url) {
  try {
    const response = await axios.get(config.BASE_URL + url);
    return response.data;
  } catch (err) {
    return [];
  }
}
