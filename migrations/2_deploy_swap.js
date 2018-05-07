var Migrations = artifacts.require("./Migrations.sol");
var Swap = artifacts.require("./Swap.sol");
var OracleWitness = artifacts.require("./OracleWitness.sol");

module.exports = function(deployer, network, accounts) {
  deployer.deploy(Migrations);
  deployer.deploy(Swap, accounts[0]);
  deployer.deploy(OracleWitness);
};
