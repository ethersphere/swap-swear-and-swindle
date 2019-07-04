
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
  cheque.signature = await web3.eth.sign(hash, cheque.signee)
  return cheque
}

module.exports = {
  signCheque
};
