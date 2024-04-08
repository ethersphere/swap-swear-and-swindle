import { DeployFunction } from 'hardhat-deploy/types';

const func: DeployFunction = async function ({ deployments, getNamedAccounts, network }) {
  const { deploy, log } = deployments;
  const { deployer } = await getNamedAccounts();

  // This code is just used for Sepolia testnet deployment
  const waitBlockConfirmations = network.name !== 'mainnet' ? 1 : 6;

  log('----------------------------------------------------');
  // sBZZ token address
  // TODO this still needs to be done for the first time
};

func.tags = ['factory'];
export default func;
