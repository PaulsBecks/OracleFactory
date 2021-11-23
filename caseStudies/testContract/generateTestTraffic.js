const SimpleStorage = artifacts.require("SimpleStorage");

async function generateTestTraffic() {
  try {
    let instance = await SimpleStorage.deployed();

    instance.start();
  } catch (err) {
    console.log(err);
  }
}

module.exports = generateTestTraffic;
