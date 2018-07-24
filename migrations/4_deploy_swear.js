var Swindle = artifacts.require("./Swindle.sol");
var Swear = artifacts.require("./Swear.sol")
var SwearSwap = artifacts.require("./SwearSwap.sol")

module.exports = function(deployer, network, accounts) {
  deployer.deploy(Swear, Swindle.address)
  deployer.deploy(SwearSwap, Swindle.address)
};
