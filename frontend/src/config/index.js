const local = {
  BASE_URL: "http://localhost:8080",
};

const prod = {
  BASE_URL: "https://oracles.work/api",
};

let config = prod;

if (process.env.NODE_ENV === "PROD") {
  config = prod;
}

export default config;
