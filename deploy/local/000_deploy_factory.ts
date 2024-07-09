import { DeployFunction } from 'hardhat-deploy/types';
import { rm } from 'fs';
import { promisify } from 'util';
const rmAsync = promisify(rm);

async function deleteDirectory(directoryPath: string) {
  try {
    await rmAsync(directoryPath, { recursive: true, force: true });
    console.log(`Deleted directory and all its contents: ${directoryPath}`);
  } catch (error) {
    console.error('Error deleting directory:', error);
  }
}

const func: DeployFunction = async function ({ deployments, getNamedAccounts, ethers }) {
  const { deploy, log } = deployments;
  const { deployer } = await getNamedAccounts();

  log('----------------------------------------------------');
  log('Deployer address at ', deployer);
  log('----------------------------------------------------');

  // Check if this is run on Cluster and is already using deployed token from Storage Incentive
  const deployedToken = '0x6AAB14FE9cccd64A502d23842d916eB5321c26E7';
  const code = await ethers.provider.getCode(deployedToken);
  let tokenAddress;

  // If there's code, then there's a contract deployed and we are deploying on running hardhat node
  if (code !== '0x') {
    tokenAddress = deployedToken;

    // Do cleanup of previous deployment
    const directoryToDelete = 'deployments/localhost/';
    //await deleteDirectory(directoryToDelete);
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
