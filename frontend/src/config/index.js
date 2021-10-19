const local = {
  BASE_URL: "http://localhost:8080",
};

const prod = {
  BASE_URL: "http://34.245.60.212:8080",
};

let config = prod;

if (process.env.NODE_ENV === "PROD") {
  config = prod;
}

export default config;
