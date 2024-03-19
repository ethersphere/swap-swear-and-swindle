import 'dotenv/config';
import 'hardhat-deploy';
import 'hardhat-deploy-ethers';
import '@nomicfoundation/hardhat-verify';
import 'hardhat-gas-reporter';
import { HardhatUserConfig } from 'hardhat/types';

const PRIVATE_RPC_MAINNET: string | undefined = process.env.PRIVATE_RPC_MAINNET;
const PRIVATE_RPC_TESTNET: string | undefined = process.env.PRIVATE_RPC_TESTNET;

const walletSecret: string = process.env.WALLET_SECRET === undefined ? 'undefined' : process.env.WALLET_SECRET;

if (walletSecret === 'undefined') {
  console.log('Please set your WALLET_SECRET in a .env file');
} else if (walletSecret.length !== 64) {
  console.log('WALLET_SECRET must be 64 characters long.');
}

const mainnetEtherscanKey: string | undefined = process.env.MAINNET_ETHERSCAN_KEY;
const testnetEtherscanKey: string | undefined = process.env.TESTNET_ETHERSCAN_KEY;
const accounts: string[] | { mnemonic: string } =
  walletSecret.length === 64 ? [walletSecret] : { mnemonic: walletSecret };

const config: HardhatUserConfig = {
  namedAccounts: {
    deployer: 0,
    admin: 1,
    stamper: 2,
    oracle: 3,
    redistributor: 4,
    pauser: 5,
    node_0: 6,
    node_1: 7,
    node_2: 8,
    node_3: 9,
    node_4: 10,
    node_5: 11,
    node_6: 12,
    node_7: 13,
  },
  defaultNetwork: 'hardhat',
  solidity: {
    compilers: [
      {
        version: '0.8.4',
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
        },
      },
      {
        version: '0.7.6',
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
        },
      },
    ],
  },
  networks: {
    hardhat: {
      chainId: 12345,
      accounts: [
        // deployer 0x3c8F39EE625fCF97cB6ee22bCe25BE1F1E5A5dE8
        {
          privateKey: '0x0d8f0a76e88539c4ceaa6ad01372cce44fb621b56b34b2cc614b4c77fb081f20',
          balance: '10000000000000000000000',
        },
        // admin 0x7E71bA1aB8AF3454a01CFafe358BEbb7691d02f8
        {
          privateKey: '0x8d56d322a1bb1e94c7d64ccd62aa2e5cc9760f59575eda0f7fd392bab8d6ba0d',
          balance: '10000000000000000000000',
        },
      ],
      deploy: ['deploy/local/'],
    },
    localhost: {
      url: 'http://localhost:8545',
      // accounts,  if not defined uses the same as above hardhat
      chainId: 12345,
      deploy: ['deploy/local/'],
    },
    testnet: {
      url: PRIVATE_RPC_TESTNET || 'https://1rpc.io/sepolia',
      accounts,
      chainId: 11155111,
      deploy: ['deploy/test/'],
    },
    mainnet: {
      url: PRIVATE_RPC_MAINNET || 'https://rpc.gnosischain.com',
      accounts,
      chainId: 100,
      deploy: ['deploy/main/'],
    },
  },
  etherscan: {
    apiKey: {
      mainnet: mainnetEtherscanKey || '',
      testnet: testnetEtherscanKey || '',
    },
    customChains: [
      {
        network: 'testnet',
        chainId: 11155111,
        urls: {
          apiURL: 'https://api-sepolia.etherscan.io/api',
          browserURL: 'https://sepolia.etherscan.io/address/',
        },
      },
      {
        network: 'mainnet',
        chainId: 100,
        urls: {
          apiURL: 'https://api.gnosisscan.io/api',
          browserURL: 'https://gnosisscan.io/address/',
        },
      },
    ],
  },
  namedAccounts: {
    deployer: 0,
    alice: 1,
    bob: 2,
    carol: 3,
    other: 4,
  },
  paths: {
    sources: 'contracts',
  },
};

export default config;
