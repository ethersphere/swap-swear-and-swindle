const abi = require('ethereumjs-abi')

const EIP712Domain = [
  { name: 'name', type: 'string' },
  { name: 'version', type: 'string' },
  { name: 'chainId', type: 'uint256' }
]

const ChequeType = [
  { name: 'chequebook', type: 'address' },
  { name: 'beneficiary', type: 'address' },
  { name: 'cumulativePayout', type: 'uint256' }
]

const CashoutType = [
  { name: 'chequebook', type: 'address' },
  { name: 'sender', type: 'address' },
  { name: 'requestPayout', type: 'uint256' },
  { name: 'recipient', type: 'address' },
  { name: 'callerPayout', type: 'uint256' }
]

const CustomDecreaseTimeoutType = [
  { name: 'chequebook', type: 'address' },
  { name: 'beneficiary', type: 'address' },
  { name: 'decreaseTimeout', type: 'uint256' }
]

async function sign(hash, signer) {
  let signature = await web3.eth.sign(hash, signer)

  let rs = signature.substr(0,130);  
  let v = parseInt(signature.substr(130, 2), 16) + 27

  return rs + v.toString(16)
}

function signTypedData(eip712data, signee) {
  return new Promise((resolve, reject) => 
    web3.currentProvider.send({
      method: 'eth_signTypedData',
      params: [signee, eip712data]
    },
    (err, result) => err == null ? resolve(result.result) : reject(err))
  )
}

// the chainId is set to 1 due to bug in ganache where the wrong id is reported via rpc
async function signCheque(swap, beneficiary, cumulativePayout, signee, chainId = 1) {
  const cheque = {
    chequebook: swap.address,
    beneficiary,
    cumulativePayout: cumulativePayout.toNumber()
  }

  const eip712data = {
    types: {
      EIP712Domain,
      Cheque: ChequeType
    },
    domain: {
      name: "Chequebook",
      version: "1.0",
      chainId
    },
    primaryType: 'Cheque',
    message: cheque
  }

  return signTypedData(eip712data, signee)
}

async function signCashOut(swap, sender, cumulativePayout, beneficiaryAgent, callerPayout, signee, chainId = 1) {
  const eip712data = {
    types: {
      EIP712Domain,
      Cashout: CashoutType
    },
    domain: {
      name: "Chequebook",
      version: "1.0",
      chainId
    },
    primaryType: 'Cashout',
    message: {
      chequebook: swap.address,
      sender,
      requestPayout: cumulativePayout.toNumber(),
      recipient: beneficiaryAgent,
      callerPayout: callerPayout.toNumber()
    }
  }

  return signTypedData(eip712data, signee)
}

async function signCustomDecreaseTimeout(swap, beneficiary, decreaseTimeout, signee, chainId = 1) {
  const eip712data = {
    types: {
      EIP712Domain,
      CustomDecreaseTimeout: CustomDecreaseTimeoutType
    },
    domain: {
      name: "Chequebook",
      version: "1.0",
      chainId
    },
    primaryType: 'CustomDecreaseTimeout',
    message: {
      chequebook: swap.address,
      beneficiary,
      decreaseTimeout: decreaseTimeout.toNumber()
    }
  }
  return signTypedData(eip712data, signee)
}

module.exports = {
  signCustomDecreaseTimeout,
  signCashOut,
  signCheque,
  sign
};
