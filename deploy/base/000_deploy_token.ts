import verify from '../../utils/verify';
import { DeployFunction } from 'hardhat-deploy/types';
import { networkConfig } from '../../helper-hardhat-config';

const func: DeployFunction = async function ({ deployments, getNamedAccounts, network }) {
  const { deploy, log } = deployments;
  const { deployer } = await getNamedAccounts();

  // Get block confirmations for the current network
  const waitBlockConfirmations = networkConfig[network.name]?.blockConfirmations || 1;

  log('----------------------------------------------------');
  log('Deploying TestToken...');

  const token = await deploy('TestToken', {
    from: deployer,
    args: [],
    log: true,
    waitConfirmations: waitBlockConfirmations,
  });

  log(`TestToken deployed at address ${token.address}`);

  // Verify the deployment
  if (network.name === 'base' && process.env.ETHERSCAN_API_KEY) {
    log('Verifying TestToken...');
    await verify(token.address, []);
  }
};

func.tags = ['token'];
export default func;
