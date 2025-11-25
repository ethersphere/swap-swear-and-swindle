export interface NetworkConfigItem {
  name: string;
  blockConfirmations: number;
}

export interface NetworkConfigMap {
  [key: string]: NetworkConfigItem;
}

export const networkConfig: NetworkConfigMap = {
  hardhat: {
    name: 'hardhat',
    blockConfirmations: 1,
  },
  localhost: {
    name: 'localhost',
    blockConfirmations: 1,
  },
  testnet: {
    name: 'testnet',
    blockConfirmations: 6,
  },
  mainnet: {
    name: 'mainnet',
    blockConfirmations: 6,
  },
  base: {
    name: 'base',
    blockConfirmations: 6,
  },
};

export const developmentChains = ['hardhat', 'localhost'];
