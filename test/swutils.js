const abi = require('ethereumjs-abi')

const EIP712Domain = [
  { name: 'name', type: 'string' },
  { name: 'version', type: 'string' },
  { name: 'chainId', type: 'uint256' }
]

const ChequeType = [
  { name: 'swap', type: 'address' },
  { name: 'beneficiary', type: 'address' },
  { name: 'cumulativePayout', type: 'uint256' }
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

async function signCheque(swap, beneficiary, cumulativePayout, signee, chainId = 1) {
  const cheque = {
    swap: swap.address,
    beneficiary,
    cumulativePayout: cumulativePayout.toNumber()
  }

  const eip712data = {
    types: {
      EIP712Domain,
      Cheque: ChequeType
    },
    domain: {
      name: "ERC20SimpleSwap",
      version: "1.0",
      chainId
    },
    primaryType: 'Cheque',
    message: cheque
  }

  return signTypedData(eip712data, signee)
}

async function signCashOut(swap, sender, cumulativePayout, beneficiaryAgent, calleePayout, signee) {
  const encodedCashOut = abi.solidityPack(
    ['address', 'address', 'uint256', 'address', 'uint256'],
    [swap.address, sender, cumulativePayout, beneficiaryAgent, calleePayout]
  )
  const hash = web3.utils.keccak256(encodedCashOut)
  return await sign(hash, signee)
}

async function signCustomDecreaseTimeout(swap, beneficiary, decreaseTimeout, signee) {
  const encodedCustomDecreaseTimeout = abi.solidityPack(
    ['address', 'address', 'uint256'],
    [swap.address, beneficiary, decreaseTimeout]
  )
  const hash = web3.utils.keccak256(encodedCustomDecreaseTimeout)

  return await sign(hash, signee)
}

module.exports = {
  signCustomDecreaseTimeout,
  signCashOut,
  signCheque,
  sign
};
