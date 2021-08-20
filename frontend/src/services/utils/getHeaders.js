export default function getHeaders() {
  try {
    const token = JSON.parse(window.localStorage.getItem("authToken"));

    return { authorization: "Bearer " + token.token };
  } catch (err) {
    return null;
  }
}
