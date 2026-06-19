import verify from '../../utils/verify';
import { DeployFunction } from 'hardhat-deploy/types';
import { networkConfig } from '../../helper-hardhat-config';

const func: DeployFunction = async function ({ deployments, getNamedAccounts, network }) {
  const { deploy, log, get } = deployments;
  const { deployer } = await getNamedAccounts();

  // Get block confirmations for the current network
  const waitBlockConfirmations = networkConfig[network.name]?.blockConfirmations || 1;

  log('----------------------------------------------------');
  log('Deploying SimpleSwapFactory...');

  // Use existing TestToken deployed from other repo
  const existingTokenAddress = '0x239Db952bde69A15962436C6CD86FDd3b45342e4';
  const deployArgs: string[] = [existingTokenAddress];

  const factory = await deploy('SimpleSwapFactory', {
    from: deployer,
    args: deployArgs,
    log: true,
    waitConfirmations: waitBlockConfirmations,
  });

  log(`Factory deployed at address ${factory.address}`);
  log(`Factory is using existing token at ${existingTokenAddress}`);

  // Verify the deployment
  if (network.name === 'base' && process.env.ETHERSCAN_API_KEY) {
    log('Verifying SimpleSwapFactory...');
    await verify(factory.address, deployArgs);
  }
};

func.tags = ['factory'];
func.dependencies = ['oracle'];
export default func;
