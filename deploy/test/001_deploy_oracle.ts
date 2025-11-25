import verify from '../../utils/verify';
import { DeployFunction } from 'hardhat-deploy/types';
import { networkConfig } from '../../helper-hardhat-config';

const func: DeployFunction = async function ({ deployments, getNamedAccounts, network }) {
  const { deploy, log } = deployments;
  const { deployer } = await getNamedAccounts();
  const deployArgs = [100000, 100];

  // Get block confirmations for the current network
  const waitBlockConfirmations = networkConfig[network.name]?.blockConfirmations || 1;

  // Deploy the PriceOracle contract
  const oracle = await deploy('PriceOracle', {
    from: deployer,
    args: deployArgs,
    log: true,
    waitConfirmations: waitBlockConfirmations,
  });

  // Log the address at which the Oracle is deployed
  console.log('Oracle deployed at address ' + oracle.address);

  // Verify the deployment
  if (network.name === 'testnet' && process.env.ETHERSCAN_API_KEY) {
    log('Verifying...');
    await verify(oracle.address, deployArgs);
  }
};

func.tags = ['oracle'];
export default func;
