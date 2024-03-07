import { ethers } from 'hardhat';

const EIP712Domain = [
  { name: 'name', type: 'string' },
  { name: 'version', type: 'string' },
  { name: 'chainId', type: 'uint256' },
];

const ChequeType = [
  { name: 'chequebook', type: 'address' },
  { name: 'beneficiary', type: 'address' },
  { name: 'cumulativePayout', type: 'uint256' },
];

const CashoutType = [
  { name: 'chequebook', type: 'address' },
  { name: 'sender', type: 'address' },
  { name: 'requestPayout', type: 'uint256' },
  { name: 'recipient', type: 'address' },
  { name: 'callerPayout', type: 'uint256' },
];

const CustomDecreaseTimeoutType = [
  { name: 'chequebook', type: 'address' },
  { name: 'beneficiary', type: 'address' },
  { name: 'decreaseTimeout', type: 'uint256' },
];

async function sign(hash: string, signer: string): Promise<string> {
  const signature = await ethers.provider.getSigner(signer).signMessage(ethers.utils.arrayify(hash));
  return signature;
}

function signTypedData(eip712data: any, signee: string): Promise<string> {
  const signer = ethers.provider.getSigner(signee);
  return signer._signTypedData(eip712data.domain, eip712data.types, eip712data.message);
}

// the chainId is set to 31337 which is the hardhat default
async function signCheque(
  swap: any,
  beneficiary: string,
  cumulativePayout: any,
  signee: string,
  chainId = 31337
): Promise<string> {
  const cheque = {
    chequebook: swap.address,
    beneficiary,
    cumulativePayout: cumulativePayout.toBigInt(),
  };

  const eip712data = {
    types: {
      EIP712Domain,
      Cheque: ChequeType,
    },
    domain: {
      name: 'Chequebook',
      version: '1.0',
      chainId,
    },
    primaryType: 'Cheque',
    message: cheque,
  };

  return signTypedData(eip712data, signee);
}

async function signCashOut(
  swap: any, // Changed from ethers.Contract to any
  sender: string,
  cumulativePayout: any,
  beneficiaryAgent: string,
  callerPayout: any,
  signee: string,
  chainId = 31337
): Promise<string> {
  const eip712data = {
    types: {
      EIP712Domain,
      Cashout: CashoutType,
    },
    domain: {
      name: 'Chequebook',
      version: '1.0',
      chainId,
    },
    primaryType: 'Cashout',
    message: {
      chequebook: swap.address,
      sender,
      requestPayout: cumulativePayout.toBigInt(),
      recipient: beneficiaryAgent,
      callerPayout: callerPayout.toBigInt(),
    },
  };

  return signTypedData(eip712data, signee);
}

async function signCustomDecreaseTimeout(
  swap: any,
  beneficiary: string,
  decreaseTimeout: any,
  signee: string,
  chainId = 31337
): Promise<string> {
  const eip712data = {
    types: {
      EIP712Domain,
      CustomDecreaseTimeout: CustomDecreaseTimeoutType,
    },
    domain: {
      name: 'Chequebook',
      version: '1.0',
      chainId,
    },
    primaryType: 'CustomDecreaseTimeout',
    message: {
      chequebook: swap.address,
      beneficiary,
      decreaseTimeout: decreaseTimeout.toBigInt(),
    },
  };
  return signTypedData(eip712data, signee);
}

export { signCustomDecreaseTimeout, signCashOut, signCheque, sign };
