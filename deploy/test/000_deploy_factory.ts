
import { verify } from "../../utils/verify";
import { DeployFunction } from 'hardhat-deploy/types';

const func: DeployFunction = async function ({ deployments, getNamedAccounts, network }) {
  const { deploy, log } = deployments;
  const { deployer } = await getNamedAccounts();

  // This code is just used for Sepolia testnet deployment
  const waitBlockConfirmations = network.name !== "testnet" ? 1 : 6;

  log("----------------------------------------------------");
  const deployArgs: string[] = ["0x543ddb01ba47acb11de34891cd86b675f04840db"];
  const factory = await deploy("SimpleSwapFactory", {
    from: deployer,
    args: deployArgs,
    log: true,
    waitConfirmations: waitBlockConfirmations,
  });

  log(`Factory deployed at address ${factory.address}`);

  // Verify the deployment
  if (network.name === "testnet" && process.env.TESTNET_ETHERSCAN_KEY) {
    log("Verifying...");
    await verify(factory.address, arguments);
  }
};

func.tags = ["factory"];
export default func;
