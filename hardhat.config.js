require("@nomiclabs/hardhat-truffle5");
require("solidity-coverage")
require('dotenv/config');

const PRIVATE_RPC_MAINNET = !process.env.PRIVATE_RPC_MAINNET ? undefined : process.env.PRIVATE_RPC_MAINNET;
const PRIVATE_RPC_TESTNET = !process.env.PRIVATE_RPC_TESTNET ? undefined : process.env.PRIVATE_RPC_TESTNET;

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
    },
    localhost: {
      url: 'http://localhost:8545',
      accounts,
    },
    testnet: {
      url: PRIVATE_RPC_TESTNET ? PRIVATE_RPC_TESTNET : 'https://rpc2.sepolia.org',
      accounts,
      chainId: 11155111,
    },
    mainnet: {
      url: PRIVATE_RPC_MAINNET ? PRIVATE_RPC_MAINNET : 'https://rpc.gnosischain.com',
      accounts,
      chainId: 100,
    },
  },
  paths: {
    sources: 'contracts',
  },
};
