import { DeployFunction } from 'hardhat-deploy/types';

const func: DeployFunction = async function ({ deployments, getNamedAccounts, network, ethers }) {
  const { deploy, log } = deployments;
  const { deployer, deployer_si } = await getNamedAccounts();

  log('----------------------------------------------------');
  log('Deployer address at ', deployer);
  log('----------------------------------------------------');

  // Send ETH to S3 deployer if using localcluster and geth
  const amountEth = ethers.utils.parseEther('10'); // 10 ETH
  if (network.name == 'localcluster') {
    await deployments.rawTx({
      from: ethers.utils.getAddress(deployer_si),
      to: ethers.utils.getAddress(deployer),
      value: amountEth,
    });
    log('Sent ETH to S3 deployer from SI deployer');
    log('----------------------------------------------------');
  }

  // Check if this is run on Cluster and is already using deployed token from Storage Incentive
  const deployedToken = '0x6AAB14FE9cccd64A502d23842d916eB5321c26E7';
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
