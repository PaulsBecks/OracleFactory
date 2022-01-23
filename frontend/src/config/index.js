const local = {
  BASE_URL: "http://localhost:8080",
};

const prod = {
  BASE_URL: "https://oracles.work/api",
};

let config = prod;

console.log(process.env);
if (process.env.NODE_ENV === "PROD" || process.env.REACT_APP_ENV === "PROD") {
  let envUrl = process.env.REACT_APP_BASE_URL;
  if (envUrl && envUrl !== "") {
    prod["BASE_URL"] = envUrl;
  }
  config = prod;
}

export default config;
