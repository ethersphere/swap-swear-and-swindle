import { DeployFunction } from 'hardhat-deploy/types';

const func: DeployFunction = async function ({ deployments, getNamedAccounts, network }) {

  const { deploy, log } = deployments;
  const { deployer } = await getNamedAccounts();
  const waitBlockConfirmations = network.name != "testnet" ? 1 : 6;

  log("----------------------------------------------------");

  const token = await deploy("TestToken", {
    from: deployer,
    log: true,
  });

  log(`Token deployed at address ${token.address}`);

  const deployArgs: string[] = [token.address];
  const factory = await deploy("SimpleSwapFactory", {
    from: deployer,
    args: deployArgs,
    log: true,
    waitConfirmations: waitBlockConfirmations,
  });

  log(`Factory deployed at address ${factory.address}`);
};

export default func;
func.tags = ["factory"];
