const local = {
  BASE_URL: "http://localhost:8080",
};

const prod = {
  BASE_URL: "http://54.194.28.109:8080",
};

let config = local;

if (process.env.NODE_ENV === "PROD") {
  config = prod;
}

export default config;
