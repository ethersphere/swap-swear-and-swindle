require("@nomiclabs/hardhat-truffle5");
require("solidity-coverage");
require("dotenv/config");
require("hardhat-deploy");
require("@nomicfoundation/hardhat-verify");

const PRIVATE_RPC_MAINNET = !process.env.PRIVATE_RPC_MAINNET
  ? undefined
  : process.env.PRIVATE_RPC_MAINNET;
const PRIVATE_RPC_TESTNET = !process.env.PRIVATE_RPC_TESTNET
  ? undefined
  : process.env.PRIVATE_RPC_TESTNET;

const walletSecret =
  process.env.WALLET_SECRET === undefined
    ? "undefined"
    : process.env.WALLET_SECRET;
if (walletSecret === "undefined") {
  console.log("Please set your WALLET_SECRET in a .env file");
}

const mainnetEtherscanKey = process.env.MAINNET_ETHERSCAN_KEY;
const testnetEtherscanKey = process.env.TESTNET_ETHERSCAN_KEY;
const accounts =
  walletSecret.length === 64 ? [walletSecret] : { mnemonic: walletSecret };

// Config for hardhat.
module.exports = {
  defaultNetwork: "hardhat",
  solidity: {
    version: "0.8.19",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200,
      },
    },
  },
  networks: {
    localhost: {
      url: "http://localhost:8545",
      accounts,
    },
    testnet: {
      url: PRIVATE_RPC_TESTNET
        ? PRIVATE_RPC_TESTNET
        : "https://1rpc.io/sepolia",
      accounts,
      chainId: 11155111,
    },
    mainnet: {
      url: PRIVATE_RPC_MAINNET
        ? PRIVATE_RPC_MAINNET
        : "https://rpc.gnosischain.com",
      accounts,
      chainId: 100,
    },
  },
  etherscan: {
    apiKey: {
      mainnet: mainnetEtherscanKey || "",
      testnet: testnetEtherscanKey || "",
    },
    customChains: [
      {
        network: "testnet",
        chainId: 11155111,
        urls: {
          apiURL: "https://api-sepolia.etherscan.io/api",
          browserURL: "https://sepolia.etherscan.io/address/",
        },
      },
      {
        network: "mainnet",
        chainId: 100,
        urls: {
          apiURL: "https://api.gnosisscan.io/",
          browserURL: "https://gnosisscan.io/address/",
        },
      },
    ],
  },
  namedAccounts: {
    deployer: {
      default: 0, // here this will by default take the first account as deployer
      1: 0, // similarly on mainnet it will take the first account as deployer. Note though that depending on how hardhat network are configured, the account 0 on one network can be different than on another
    },
  },
  paths: {
    sources: "contracts",
  },
};
