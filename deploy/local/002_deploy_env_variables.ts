import fs from 'fs';
import path from 'path';
import { DeployFunction } from 'hardhat-deploy/types';

const func: DeployFunction = async function ({ deployments, config }) {
  const { get, log } = deployments;

  const SimpleSwapFactory = await get('SimpleSwapFactory');
  const PriceOracle = await get('PriceOracle');

  // Generate content for the environment file
  let content = '';

  content += `echo "----- USE THE COMMANDS BELOW TO SETUP YOUR TERMINALS -----" >&2\n\n`;
  content += `export BEE_SWAP_FACTORY_ADDRESS=${SimpleSwapFactory.address}\n`;
  content += `export BEE_SWAP_LEGACY_FACTORY_ADDRESSES=${SimpleSwapFactory.address}\n`;
  content += `export BEE_SWAP_PRICE_ORACLE_ADDRESS=${PriceOracle.address}\n`;
  content += `export BEE_SWAP_ENDPOINT=${config.networks.localhost.url}\n`;

  const envFilePath = path.join(__dirname, '../../deployedContracts.sh');

  // Write the content to the file
  fs.writeFileSync(envFilePath, content, { flag: 'a' });
  log(`Exported contract addresses to ${envFilePath}`);

  log('----------------------------------------------------');
};

export default func;
func.tags = ['variables'];
