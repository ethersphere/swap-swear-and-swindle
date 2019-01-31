const Swap = artifacts.require("./Swap.sol");

const { sign } = require('./testutils')

async function signCheque(swap, signer, beneficiary, serial, amount, timeout) {
  const hash = await swap.chequeHash(swap.address, beneficiary, serial, amount, timeout);
  return (await sign(signer, hash));
}

async function signNote(swap, signer, beneficiary, serial, amount, witness, validFrom, validUntil, remark) {
  const hash = await swap.noteHash(swap.address, beneficiary, serial, amount, witness, validFrom, validUntil, remark);
  return { ...await sign(signer, hash), hash };
}

async function signInvoice(swap, signer, noteId, swapBalance, serial) {
  const hash = await swap.invoiceHash(noteId, swapBalance, serial);
  return { ...await sign(signer, hash), hash };
}

module.exports = {
  signCheque,
  signNote,
  signInvoice
}
