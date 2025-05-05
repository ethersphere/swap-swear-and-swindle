import verify from '../../utils/verify';
import { DeployFunction } from 'hardhat-deploy/types';

const func: DeployFunction = async function ({ deployments, getNamedAccounts, network }) {
  const { deploy, log } = deployments;
  const { deployer } = await getNamedAccounts();
  const deployArgs = [100000, 100];

  // Deploy the PriceOracle contract
  const oracle = await deploy('PriceOracle', {
    from: deployer,
    args: deployArgs,
    log: true,
    waitConfirmations: 6,
  });

  // Log the address at which the Oracle is deployed
  console.log('Oracle deployed at address ' + oracle.address);

  // Verify the deployment
  if (network.name == 'testnet' && process.env.TESTNET_ETHERSCAN_KEY) {
    log('Verifying...');
    await verify(oracle.address, deployArgs);
  }
};

func.tags = ['oracle'];
export default func;
