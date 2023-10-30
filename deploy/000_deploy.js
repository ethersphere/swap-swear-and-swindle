const { getNamedAccounts, deployments, network, run } = require("hardhat");
const { verify } = require("../utils/verify");

module.exports = async ({ getNamedAccounts, deployments }) => {
  const { deploy, log } = deployments;
  const { deployer } = await getNamedAccounts();

  // This code is just used for Sepolia testnet deployment
  const waitBlockConfirmations = network.name != "testnet" ? 1 : 6;

  log("----------------------------------------------------");
  // sBZZ token address
  const arguments = ["0x543dDb01Ba47acB11de34891cD86B675F04840db"];
  const factory = await deploy("SimpleSwapFactory", {
    from: deployer,
    args: arguments,
    log: true,
    waitConfirmations: waitBlockConfirmations,
  });

  log("Factory deployed at address " + factory.address);

  // Verify the deployment
  if (network.name == "testnet" && process.env.TESTNET_ETHERSCAN_KEY) {
    log("Verifying...");
    await verify(factory.address, arguments);
  }
};

module.exports.tags = ["factory"];
