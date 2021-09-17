const BlindAuction = artifacts.require("BlindAuction");

module.exports = function (deployer) {
  deployer.deploy(
    BlindAuction,
    3,
    0,
    "0xe4EFfB267484Cd790143484de3Bae7fDfbE31F00"
  );
};
