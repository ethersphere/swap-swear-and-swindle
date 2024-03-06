const { network } = require("hardhat");

module.exports = async ({ getNamedAccounts, deployments }) => {
  const { deploy, log } = deployments;
  const { deployer } = await getNamedAccounts();
  const waitBlockConfirmations = 1;

  log("----------------------------------------------------");
  const args = [100, 200];

  // Deploy the PriceOracle contract
  const oracle = await deploy('PriceOracle', {
    from: deployer,
    args: args,
    log: true,
    waitConfirmations: waitBlockConfirmations,
  });

  // Log the address at which the Oracle is deployed
  console.log('Oracle deployed at address ' + oracle.address);

};

module.exports.tags = ["factory"];
