var Swindle = artifacts.require("./Swindle.sol");
var OracleTrial = artifacts.require("./OracleTrial.sol");

module.exports = function(deployer, network, accounts) {
  deployer.deploy(Swindle);
  deployer.deploy(OracleTrial);
};
