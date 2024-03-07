import 'dotenv/config';
import 'solidity-coverage';
import 'hardhat-deploy';
import 'hardhat-deploy-ethers';
import '@nomiclabs/hardhat-etherscan';
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
      deploy: ['deploy/local/'],
    },
    localhost: {
      url: 'http://localhost:8545',
      accounts,
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
