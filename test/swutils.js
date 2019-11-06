const abi = require('ethereumjs-abi')

async function sign(hash, signer) {
  let signature = await web3.eth.sign(hash, signer)

  let rs = signature.substr(0,130);  
  let v = parseInt(signature.substr(130, 2), 16) + 27

  return rs + v.toString(16)
}

async function signCheque(swap, beneficiary, cumulativePayout, signee) {
  const encodedCheque = abi.solidityPack(
    ['address', 'address', 'uint256'],
    [swap.address, beneficiary, cumulativePayout]
  )
  const hash = web3.utils.keccak256(encodedCheque)
  return await sign(hash, signee)
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
