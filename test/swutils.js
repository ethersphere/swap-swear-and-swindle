const Swap = artifacts.require('./Swap')
const {
  constants
} = require("openzeppelin-test-helpers");

function getSwap() {  
  return Swap.new(constants.ZERO_ADDRESS)    
}

async function sign(hash, signer) {
  let signature = await web3.eth.sign(hash, signer)

  let rs = signature.substr(0,130);  
  let v = parseInt(signature.substr(130, 2), 16) + 27

  return rs + v.toString(16)
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
    cheque.signature.issuer = await sign(hash, cheque.signee[0])
    cheque.signature.beneficiary = await sign(hash, cheque.signee[1])
  } else {
    cheque.signature = await sign(hash, cheque.signee)
  }
  return cheque
}

async function signCashOut(swap, sender, requestPayout, beneficiaryAgent, expiry, calleePayout, signee) {
  const hash = await swap.cashOutHash(
    swap.address,
    sender,
    requestPayout,
    beneficiaryAgent,
    expiry,
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

async function signNote(signer, note) {  
  const hash = await (await getSwap()).noteHash([
    note.swap,
    note.serial,
    note.amount,
    note.beneficiary,
    note.witness,
    note.validFrom,
    note.validUntil,
    note.remark,
    note.timeout
  ]);
  return { sig: await web3.eth.sign(hash, signer), hash };
}

async function encodeNote(note) {  
  return (await getSwap()).encodeNote([
    note.swap,
    note.serial,
    note.amount,
    note.beneficiary,
    note.witness,
    note.validFrom,
    note.validUntil,
    note.remark,
    note.timeout
  ]);  
}

async function signInvoice(swap, signer, noteId, swapBalance, serial) {
  const hash = await swap.invoiceHash(noteId, swapBalance, serial);
  return { sig: await web3.eth.sign(hash, signer), hash };
}

module.exports = {
  signCheque,
  encodeNote,
  signNote,
  signInvoice,
  signCustomDecreaseTimeout,
  signCashOut,
  signCheque,
  sign
};
