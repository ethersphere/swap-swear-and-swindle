const { getNamedAccounts, deployments, network, run } = require("hardhat");
const { verify } = require("../utils/verify");

module.exports = async ({ getNamedAccounts, deployments }) => {
  const { deploy, log } = deployments;
  const { deployer } = await getNamedAccounts();

  const waitBlockConfirmations = network.name != "testnet" ? 1 : 6;

  log("----------------------------------------------------");
  const arguments = ["0xa66be4A7De4DfA5478Cb2308469D90115C45aA23"];
  const factory = await deploy("SimpleSwapFactory", {
    from: deployer,
    args: arguments,
    log: true,
    waitConfirmations: waitBlockConfirmations,
  });

  log("Factory deployed at address " + factory.address);

  // Verify the deployment
  if (network.name == "testnet" && process.env.ETHERSCAN_API_KEY) {
    log("Verifying...");
    await verify(factory.address, arguments);
  }
};

module.exports.tags = ["all", "factory"];
