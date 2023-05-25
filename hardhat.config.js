require("@nomiclabs/hardhat-truffle5");
require("solidity-coverage")


const walletSecret = process.env.WALLET_SECRET === undefined ? 'undefined' : process.env.WALLET_SECRET;
if (walletSecret === 'undefined') {
  console.log('Please set your WALLET_SECRET in a .env file');
}
const accounts = walletSecret.length === 64 ? [walletSecret] : { mnemonic: walletSecret };


// Config for hardhat.
module.exports = {
  solidity: { version: '0.8.10',
    settings: {
      optimizer: {
        enabled: true,
        runs: 200
      },
    }
  },
  networks: {
    hardhat: {
      accounts,
    },
    localhost: {
      url: 'http://localhost:8545',
      accounts,
    },
    staging: {
      url: 'https://goerli.infura.io/v3/' + process.env.INFURA_TOKEN,
      accounts,
    },
  },
  paths: {
    sources: 'contracts',
  },
};
