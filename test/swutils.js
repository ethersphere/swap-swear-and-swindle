const Swap = artifacts.require("./Swap.sol");

const { sign } = require('./testutils')

async function signCheque(signer, beneficiary, serial, amount) {
  const swap = await Swap.deployed();
  const hash = await swap.chequeHash(swap.address, beneficiary, serial, amount);
  return (await sign(signer, hash));
}

async function signNote(signer, beneficiary, serial, amount, witness, validFrom, validUntil, remark) {
  const swap = await Swap.deployed();
  const hash = await swap.noteHash(swap.address, beneficiary, serial, amount, witness, validFrom, validUntil, remark);
  return { ...await sign(signer, hash), hash };
}

async function signInvoice(signer, noteId, swapBalance, serial) {
  const swap = await Swap.deployed();
  const hash = await swap.invoiceHash(noteId, swapBalance, serial);
  return { ...await sign(signer, hash), hash };
}

module.exports = {
  signCheque,
  signNote,
  signInvoice
}
