const Swap = artifacts.require("./Swap.sol");
const SimpleSwap = artifacts.require("./SimpleSwap.sol");
const SoftSwap = artifacts.require("./SoftSwap.sol");
const OracleWitness = artifacts.require("./OracleWitness.sol");

const {
  BN,
  balance,
  time,
  shouldFail,
  constants,
  expectEvent
} = require("openzeppelin-test-helpers");

const { signCheque, signNote, signInvoice, encodeNote } = require("./swutils");
const { computeCost } = require("./testutils");

const epoch = 24 * 3600;

async function submitChequeBeneficiary(swap, cheque) {
  cheque = { timeout: epoch, ...cheque };
  const { owner, beneficiary, serial, amount, timeout } = cheque;
  const { sig } = await signCheque(swap, owner, cheque); // sig is the signature of the owner
  return swap.submitChequeBeneficiary(serial, amount, timeout, sig, {
    from: beneficiary
  });
}

const softSwapTests = (accounts, Swap) => {
  const [owner, bob, alice] = accounts;

  async function prepareSwap(prefilledAmount = 1000) {
    const swap = await Swap.new(owner);
    await swap.send(prefilledAmount, { from: owner });
    return { swap, prefilledAmount: new BN(prefilledAmount) };
  }

};

const swapTests = (accounts, Swap) => {
  const [owner, bob, alice, carol] = accounts;

  async function prepareSwap(prefilledAmount = 1000) {
    const swap = await Swap.new(owner);
    await swap.send(prefilledAmount, { from: owner });
    return { swap, prefilledAmount: new BN(prefilledAmount) };
  }

  it("should accept a valid note (bond)", async () => {
    const { swap } = await prepareSwap(1000);

    const noteTimeout = 5000;
    const noteAmount = 500;

    let validity = (await time.latest()).addn(noteTimeout);
    
    let note = {
      swap: swap.address,
      beneficiary: carol,
      serial: 1,
      amount: noteAmount,
      witness: constants.ZERO_ADDRESS,
      validFrom: validity,
      validUntil: 0,
      remark: '0x',
      timeout: epoch
    }    
    
    let { sig, hash } = await signNote(owner, note);

    await time.increase(4 * epoch);

    let encoded = await encodeNote(note);
    await swap.submitNote(encoded, sig, { from: carol });

    const { paidOut, timeout } = await swap.notes(hash);

    paidOut.should.bignumber.equal(new BN(0));
    timeout.should.bignumber.gte((await time.latest()).addn(1 * epoch - 1));

    await time.increase(1 * epoch);

    let expectedBalanceCarol = (await balance.current(carol)).addn(noteAmount);

    let { receipt } = await swap.cashNote(encoded, noteAmount, { from: carol });

    expectedBalanceCarol = expectedBalanceCarol.sub(await computeCost(receipt));

    (await balance.current(carol)).should.bignumber.equal(expectedBalanceCarol);

    // already fully cashed out
    await shouldFail.reverting(
      swap.cashNote(encoded, noteAmount, { from: carol })
    );
  });

  // TODO: split
  it("should accept a valid note (conditional bond)", async () => {    
    const { swap } = await prepareSwap(1000);
    const oracle = await OracleWitness.new();

    const noteAmount = 500;
    const noteTimeout = 2 * epoch;

    let bondTimeout = (await time.latest()).addn(noteTimeout);

    let remark = '0xaa'

    let note = {
      swap: swap.address,
      beneficiary: carol,
      serial: 1,
      amount: noteAmount,
      witness: oracle.address,
      validFrom: 0,
      validUntil: bondTimeout,
      remark: remark,
      timeout: epoch
    }

    let { sig, hash } = await signNote(owner, note)      
    let encoded = await encodeNote(note)

    await oracle.testify(remark, 1);

    await swap.submitNote(encoded, sig, { from: carol });

    const { paidOut, timeout } = await swap.notes(hash);

    paidOut.should.bignumber.equal(new BN(0));
    timeout.should.bignumber.gte((await time.latest()).addn(1 * epoch - 1));

    // cashout too soon
    await shouldFail.reverting(
      swap.cashNote(encoded, noteAmount, { from: carol })
    );

    await time.increase(1 * epoch);

    await oracle.testify(remark, 0);

    // oracle says no
    await shouldFail.reverting(
      swap.cashNote(encoded, noteAmount, { from: carol })
    );

    await oracle.testify(remark, 1);

    // partial payment
    let expectedBalanceCarol = (await balance.current(carol)).addn(
      noteAmount / 4
    );
    var { receipt } = await swap.cashNote(encoded, noteAmount / 4, {
      from: carol
    });
    expectedBalanceCarol = expectedBalanceCarol.sub(await computeCost(receipt));
    (await balance.current(carol)).should.bignumber.equal(expectedBalanceCarol);

    // partial payment
    expectedBalanceCarol = (await balance.current(carol)).addn(noteAmount / 4);
    var { receipt } = await swap.cashNote(encoded, noteAmount / 4, {
      from: carol
    });
    expectedBalanceCarol = expectedBalanceCarol.sub(await computeCost(receipt));
    (await balance.current(carol)).should.bignumber.equal(expectedBalanceCarol);

    await time.increase(2 * epoch);

    // too late for the rest
    await shouldFail.reverting(
      swap.cashNote(encoded, noteAmount / 4, { from: carol })
    );
  });

  // TODO: split
  it("should allow to submit paid invoices", async () => {    
    const { swap } = await prepareSwap(1000);
    const noteAmount = new BN(500);

    const beneficiary = carol;
    const cheques = [
      {
        owner,
        beneficiary,
        serial: new BN(1),
        amount: new BN(100),
        timeout: epoch
      },
      {
        serial: new BN(2),
        amount: new BN(200),
        timeout: epoch
      },
      {
        owner,
        beneficiary,
        serial: new BN(3),
        amount: noteAmount.addn(200),
        timeout: epoch
      }
    ];

    await submitChequeBeneficiary(swap, cheques[0]);

    // completely offchain cheque of 100

    let note = {
      swap: swap.address,
      beneficiary,
      serial: 1,
      amount: noteAmount,
      witness: constants.ZERO_ADDRESS,
      validFrom: 0,
      validUntil: 0,
      remark: '0x',
      timeout: epoch
    }

    // owner issues note
    let { sig, hash } = await signNote(owner, note);

    // carol issues invoice
    let invoice = await signInvoice(
      swap,
      beneficiary,
      hash,
      cheques[1].amount,
      cheques[1].serial
    );

    // owner issues cheque for invoice
    let cheque = await signCheque(swap, owner, cheques[2]);

    let encoded = await encodeNote(note);
    // carol submits note anyway
    await swap.submitNote(encoded, sig, { from: carol });
    //function submitPaidInvoice(bytes memory encoded, uint swapBalance, uint serial, bytes memory invoiceSig, uint amount, uint timeout, bytes memory chequeSig) public {
    // owner presents paid invoice
    await swap.submitPaidInvoice(
      encoded,
      cheques[1].amount,
      cheques[1].serial,
      invoice.sig,
      noteAmount,
      epoch,
      cheque.sig
    );

    await time.increase(2 * epoch);

    await shouldFail.reverting(swap.cashNote(encoded, noteAmount));

    let { logs } = await swap.cashCheque(carol, cheques[2].amount);

    expectEvent.inLogs(logs, "ChequeCashed", {
      beneficiary: carol,
      serial: cheques[2].serial,
      payout: cheques[2].amount,
      requestPayout: cheques[2].amount
    });
  });
};

contract("SoftSwap", function(accounts) {
  softSwapTests(accounts, SoftSwap);
});

contract("Swap", function(accounts) {
  softSwapTests(accounts, Swap);
  swapTests(accounts, Swap);
});
