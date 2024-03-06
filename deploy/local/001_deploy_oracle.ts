import { DeployFunction } from 'hardhat-deploy/types';

const func: DeployFunction = async function ({ deployments, getNamedAccounts }) {
  const { deploy, log } = deployments;
  const { deployer } = await getNamedAccounts();
  const waitBlockConfirmations = 1;

  log('----------------------------------------------------');
  const deployArgs: [number, number] = [100, 200];

  // Deploy the PriceOracle contract
  const oracle = await deploy('PriceOracle', {
    from: deployer,
    args: deployArgs,
    log: true,
    waitConfirmations: waitBlockConfirmations,
  });

  // Log the address at which the Oracle is deployed
  log(`Oracle deployed at address ${oracle.address}`);
};

export default func;
func.tags = ['factory'];
