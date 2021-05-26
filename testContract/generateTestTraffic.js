const SimpleStorage = artifacts.require("SimpleStorage");

async function generateTestTraffic() {
  try {
    let instance = await SimpleStorage.deployed();
    while (true) {
      instance.set(42);
    }
  } catch (err) {
    console.log(err);
  }
}

module.exports = generateTestTraffic;
