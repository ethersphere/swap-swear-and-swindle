const { getNamedAccounts, deployments, network, run } = require("hardhat");
const { verify } = require("../../utils/verify");

module.exports = async ({ getNamedAccounts, deployments }) => {
  const { deploy, log } = deployments;
  const { deployer } = await getNamedAccounts();

  // This code is just used for Sepolia testnet deployment
  const waitBlockConfirmations = network.name != "mainnet" ? 1 : 6;

  log("----------------------------------------------------");
  // sBZZ token address
  // TODO this still needs to be done for the first time
};

module.exports.tags = ["factory"];
