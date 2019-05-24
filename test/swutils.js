const Swap = artifacts.require('./Swap')
const {
  constants
} = require("openzeppelin-test-helpers");

function getSwap() {  
  return Swap.new(constants.ZERO_ADDRESS)    
}

async function signCheque(swap, signer, cheque) {
  const { owner, beneficiary, serial, amount, timeout } = cheque;
  const hash = await swap.chequeHash(
    swap.address,
    beneficiary,
    serial,
    amount,
    timeout
  );
  return { sig: await web3.eth.sign(hash, signer) };
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
  signInvoice
};
