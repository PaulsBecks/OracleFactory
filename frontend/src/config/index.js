const local = {
  BASE_URL: "http://localhost:8080",
};

const prod = {
  BASE_URL: "http://3.249.80.172:8080",
};

let config = prod;

if (process.env.NODE_ENV === "PROD") {
  config = prod;
}

export default config;
