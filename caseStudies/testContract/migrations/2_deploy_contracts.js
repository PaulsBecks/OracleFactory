const SimpleStorage = artifacts.require("SimpleStorage");
const SimpleStorage1 = artifacts.require("SimpleStorage1");
const SimpleStorage2 = artifacts.require("SimpleStorage2");

module.exports = function (deployer) {
  deployer.deploy(SimpleStorage);
  deployer.deploy(SimpleStorage1);
  deployer.deploy(SimpleStorage2);
};
