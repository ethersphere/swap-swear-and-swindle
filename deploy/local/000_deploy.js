const { getNamedAccounts, deployments, network, run } = require("hardhat");
const { verify } = require("../../utils/verify");

module.exports = async ({ getNamedAccounts, deployments }) => {
  const { deploy, log } = deployments;
  const { deployer } = await getNamedAccounts();
  const waitBlockConfirmations = network.name != "testnet" ? 1 : 6;

  log("----------------------------------------------------");

  token = await deploy("TestToken", {
    from: deployer,
    log: true,
  });

  log("Token deployed at address " + token.address);

  const arguments = [token];
  const factory = await deploy("SimpleSwapFactory", {
    from: deployer,
    args: arguments,
    log: true,
    waitConfirmations: waitBlockConfirmations,
  });

  log("Factory deployed at address " + factory.address);
};

module.exports.tags = ["factory"];
