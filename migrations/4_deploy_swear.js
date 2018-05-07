var Swindle = artifacts.require("./Swindle.sol");
var Swear = artifacts.require("./Swear.sol")

module.exports = function(deployer, network, accounts) {
  deployer.deploy(Swear, Swindle.address)  
};
