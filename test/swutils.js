const Swap = artifacts.require("./Swap.sol");

async function signCheque(swap, signer, beneficiary, serial, amount, timeout) {
  const hash = await swap.chequeHash(swap.address, beneficiary, serial, amount, timeout);
  return { sig: await web3.eth.sign(hash, signer) };
}

async function signNote(swap, signer, beneficiary, serial, amount, witness, validFrom, validUntil, remark) {
  const hash = await swap.noteHash(swap.address, beneficiary, serial, amount, witness, validFrom, validUntil, remark);
  return { sig: await web3.eth.sign(hash, signer), hash };
}

async function signInvoice(swap, signer, noteId, swapBalance, serial) {
  const hash = await swap.invoiceHash(noteId, swapBalance, serial);
  return { sig: await web3.eth.sign(hash, signer), hash };
}

module.exports = {
  signCheque,
  signNote,
  signInvoice
}
