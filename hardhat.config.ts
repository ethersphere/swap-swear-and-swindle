import 'dotenv/config';
import 'hardhat-deploy';
import 'hardhat-deploy-ethers';
import '@nomicfoundation/hardhat-verify';
import 'hardhat-gas-reporter';
import 'solidity-coverage';
import { HardhatUserConfig } from 'hardhat/types';

const PRIVATE_RPC_MAINNET: string | undefined = process.env.PRIVATE_RPC_MAINNET;
const PRIVATE_RPC_TESTNET: string | undefined = process.env.PRIVATE_RPC_TESTNET;

// We use WALLET_SECRET_2 because of collide with SI repo that uses WALLET_SECRET and we use both of them for cluster deployment
const walletSecret: string = process.env.WALLET_SECRET_2 === undefined ? 'undefined' : process.env.WALLET_SECRET_2;

if (walletSecret === 'undefined') {
  console.log('Please set your WALLET_SECRET_2 in a .env file');
} else if (walletSecret.length !== 64) {
  console.log('WALLET_SECRET_2 must be 64 characters long.');
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
  mocha: {
    timeout: Number.MAX_SAFE_INTEGER,
  },
  networks: {
    hardhat: {
      chainId: 12345,
      accounts: [
        // deployer 0x7E71bA1aB8AF3454a01CFafe358BEbb7691d02f8
        {
          privateKey: '0x8d56d322a1bb1e94c7d64ccd62aa2e5cc9760f59575eda0f7fd392bab8d6ba0d',
          balance: '10000000000000000000000',
        },
        // admin 0x77CbAdb1059dDC7334227e025fC940469f52FEd8
        {
          privateKey: '0xb65c0589ad60bc9985f0b6eafe5dd480b7ad63f073a7e9625dd23466a0d1947d',
          balance: '10000000000000000000000',
        },
        // named1 0xFCA295bC36F47A3Eb53F657b88f3f324374656C6
        {
          privateKey: '0x963893a36bd803209c07615b0650303706fb01158479a46fba4dea3fe8cf0734',
          balance: '10000000000000000000000',
        },
        // named2 0xB5963cAcF590909407433024cD3BA0319542E99D
        {
          privateKey: '0xee65b03b4dfdde207a44c6ff5da99201ee0642841ae9f2e07927e8d2ad523d55',
          balance: '10000000000000000000000',
        },
        // named3 0x9C8EEad79edDC16594489d63E5A9F7530b642079
        {
          privateKey: '0x34777daf03381f4666635bff0e03720a49f62ba28daa3ab6cabe0922e8574422',
          balance: '10000000000000000000000',
        },
        // other_1 0x626178434A88c3c8809D136d500b9707D749EA9B
        {
          privateKey: 'f09baf4a06da707abeb96568a1419b4eec094774eaa85ef85517457ffe25b515',
          balance: '10000000000000000000000',
        },
        // other_2 0xb22D48A49c0Aa99AC94072E229E52687E97da253
        {
          privateKey: '5d6172133423006770002831e395aca9d2dad3bcf9257e38c2f19224b4aef78b',
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
    localcluster: {
      url: 'http://geth-swap:8545',
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
  paths: {
    sources: 'contracts',
  },
  namedAccounts: {
    deployer: 0,
    admin: 1,
    alice: 2,
    bob: 3,
    carol: 4,
    other: 5,
  },
};

export default config;
