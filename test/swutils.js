async function sign(hash, signer) {
  let signature = await web3.eth.sign(hash, signer)

  let rs = signature.substr(0,130);  
  let v = parseInt(signature.substr(130, 2), 16) + 27

  return rs + v.toString(16)
}

async function signCheque(swap, beneficiary, cumulativePayout, signee) {
  const hash = await swap.chequeHash(
    swap.address,
    beneficiary,
    cumulativePayout
  )

  return await sign(hash, signee)
}

async function signCashOut(swap, sender, cumulativePayout, beneficiaryAgent, calleePayout, signee) {
  const hash = await swap.cashOutHash(
    swap.address,
    sender,
    cumulativePayout,
    beneficiaryAgent,
    calleePayout
  )

  return await sign(hash, signee)
}

async function signCustomDecreaseTimeout(swap, beneficiary, decreaseTimeout, signee) {
  const hash = await swap.customDecreaseTimeoutHash(
    swap.address,
    beneficiary,
    decreaseTimeout
  )

  return await sign(hash, signee)
}

module.exports = {
  signCustomDecreaseTimeout,
  signCashOut,
  signCheque,
  sign
};
