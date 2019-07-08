const Swap = artifacts.require('./SimpleSwap')
const {
  constants
} = require("openzeppelin-test-helpers");

function getSwap() {  
  return Swap.new(constants.ZERO_ADDRESS)    
}

async function signCheque(swap, signer, cheque) {
  const { beneficiary, serial, amount, timeout } = cheque;
  const hash = await swap.chequeHash(
    swap.address,
    beneficiary,
    serial,
    amount,
    timeout
  );
  return { sig: await web3.eth.sign(hash, signer)} ;
}

module.exports = {
  signCheque
};
