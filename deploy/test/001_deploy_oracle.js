const { getNamedAccounts, deployments, network, run } = require("hardhat");
const { verify } = require("../../utils/verify");

module.exports = async ({ getNamedAccounts, deployments }) => {
  const { deploy, log } = deployments;
  const { deployer } = await getNamedAccounts();
  const args = [100000, 100];

  // Deploy the PriceOracle contract
  const oracle = await deploy('PriceOracle', {
    from: deployer,
    args: args,
    log: true,
    waitConfirmations: 6,
  });

  // Log the address at which the Oracle is deployed
  console.log('Oracle deployed at address ' + oracle.address)

  // Verify the deployment
  if (network.name == "testnet" && process.env.TESTNET_ETHERSCAN_KEY) {
    log("Verifying...");
    await verify(oracle.address, arguments);
  }
};

module.exports.tags = ["factory"];
