import { DeployFunction } from 'hardhat-deploy/types';
import { networkConfig } from '../../helper-hardhat-config';

const func: DeployFunction = async function ({ deployments, getNamedAccounts, network }) {
  const { deploy, log } = deployments;
  const { deployer } = await getNamedAccounts();

  // Get block confirmations for the current network
  const waitBlockConfirmations = networkConfig[network.name]?.blockConfirmations || 1;

  log('----------------------------------------------------');
  // sBZZ token address
  // TODO this still needs to be done for the first time
};

func.tags = ['factory'];
export default func;
