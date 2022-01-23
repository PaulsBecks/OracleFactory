const local = {
  BASE_URL: "http://localhost:8080",
};

const prod = {
  BASE_URL: "http://54.194.28.109:8080",
};

let config = local;

console.log(process.env);
if (process.env.NODE_ENV === "PROD" || process.env.REACT_APP_ENV === "PROD") {
  let envUrl = process.env.REACT_APP_BASE_URL;
  if (envUrl && envUrl !== "") {
    prod["BASE_URL"] = envUrl;
  }
  config = prod;
}

export default config;
