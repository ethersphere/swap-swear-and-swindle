import { DeployFunction } from 'hardhat-deploy/types';

const func: DeployFunction = async function ({ deployments, getNamedAccounts, network, ethers }) {
  const { deploy, log } = deployments;
  const { deployer } = await getNamedAccounts();

  log('----------------------------------------------------');
  log('Deployer address at ', deployer);
  log('----------------------------------------------------');

  // Check if this is run on Cluster and is already using deployed token from Storage Incentive
  const deployedToken = '0x942C6684eB9874C63d4ed26Ab0623F951D253081';
  const code = await ethers.provider.getCode(deployedToken);
  let tokenAddress;

  // If there's code, then there's a contract deployed
  if (code !== '0x') {
    console.log('A contract is deployed at this address.');
    tokenAddress = deployedToken;
  } else {
    const token = await deploy('TestToken', {
      from: deployer,
      log: true,
    });
    tokenAddress = token.address;
  }

  log(`Token deployed at address ${tokenAddress}`);

  const deployArgs: string[] = [tokenAddress];
  const factory = await deploy('SimpleSwapFactory', {
    from: deployer,
    args: deployArgs,
    log: true,
    waitConfirmations: 1,
  });

  log(`Factory deployed at address ${factory.address}`);
};

export default func;
func.tags = ['factory'];
