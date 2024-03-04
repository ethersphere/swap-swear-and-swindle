const fs = require("fs");
const path = require("path");

const func = async function({ deployments }) {
  const { get, log } = deployments;

  const SimpleSwapFactory = await get("SimpleSwapFactory");

  // Generate content for the environment file
  let content = "";

  content += `echo "----- USE THE COMMANDS BELOW TO SETUP YOUR TERMINALS -----" >&2\n\n`;
  content += `export BEE_SWAP_FACTORY_ADDRESS=${SimpleSwapFactory.address}\n`;
  content += `export BEE_SWAP_LEGACY_FACTORY_ADDRESSES=${SimpleSwapFactory.address}\n`;
  content += `export BEE_SWAP_ENDPOINT=${networks.localhost.url}\n`;

  const envFilePath = path.join(__dirname, "../../deployedContracts.sh");

  // Write the content to the file
  fs.writeFileSync(envFilePath, content, { flag: "a" });
  console.log(`Exported contract addresses to ${envFilePath}`);

  log("----------------------------------------------------");
};

module.exports = func;
func.tags = ["variables"];
