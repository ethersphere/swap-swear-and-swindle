const Swap = artifacts.require('./SimpleSwap')
const {
  constants
} = require("openzeppelin-test-helpers");

function getSwap() {  
  return Swap.new(constants.ZERO_ADDRESS)    
}

async function signCheque(swap, cheque) {
  const hash = await swap.chequeHash(
    swap.address,
    cheque.beneficiary,
    cheque.serial,
    cheque.amount,
    cheque.timeout
  );
  
  if(cheque.signee.length == 2) {
    cheque.signature = []
    cheque.signature[0] = await web3.eth.sign(hash, cheque.signee[0])
    cheque.signature[1] = await web3.eth.sign(hash, cheque.signee[1])
  } else {
    cheque.signature = await web3.eth.sign(hash, cheque.signee)
  }
  return cheque
}

module.exports = {
  signCheque
};
