const Oracle = artifacts.require("Oracle");

module.exports = async function (deployer) {
  await deployer.deploy(Oracle);
  let contract = await Oracle.deployed();
  console.log("PUBSUB_ORACLE_ADDRESS: " + contract.address);
};
