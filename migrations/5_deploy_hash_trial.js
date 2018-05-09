var HashWitness = artifacts.require("./HashWitness.sol");
var SimpleTrial = artifacts.require("./SimpleTrial.sol");

module.exports = function(deployer, network, accounts) {
  deployer.deploy(SimpleTrial, HashWitness.address);
};
