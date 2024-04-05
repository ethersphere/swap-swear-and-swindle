import { DeployFunction } from 'hardhat-deploy/types';

const func: DeployFunction = async function ({ deployments, config, network }) {
  const { get, log } = deployments;

  const SimpleSwapFactory = await get('SimpleSwapFactory');
  const PriceOracle = await get('PriceOracle');
  const networkURL = config.networks[network.name].url;

  // Generate content for the environment file
  let content = '';

  content += `echo "----- USE THE COMMANDS BELOW TO SETUP YOUR TERMINALS -----" >&2\n\n`;
  content += `export BEE_SWAP_FACTORY_ADDRESS=${SimpleSwapFactory.address}\n`;
  content += `export BEE_SWAP_LEGACY_FACTORY_ADDRESSES=${SimpleSwapFactory.address}\n`;
  content += `export BEE_SWAP_PRICE_ORACLE_ADDRESS=${PriceOracle.address}\n`;
  content += `export BEE_SWAP_ENDPOINT=${networkURL}\n`;

  // Output the content to the terminal
  console.log(content);
  log(`Exported contract addresses to console`);

  log('----------------------------------------------------');
};

export default func;
func.tags = ['variables'];
