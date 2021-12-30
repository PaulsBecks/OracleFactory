const local = {
  BASE_URL: "http://localhost:8080",
};

const prod = {
  BASE_URL: "http://pub-sub-oracle:8080",
};

let config = prod;

if (process.env.NODE_ENV === "development") {
  config = local;
}

export default config;
