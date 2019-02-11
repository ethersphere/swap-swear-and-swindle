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

const { signCheque, signNote, signInvoice } = require("./swutils");
const { computeCost } = require("./testutils");

const epoch = 24 * 3600;

async function submitCheque(swap, cheque) {
  cheque = { timeout: epoch, ...cheque };
  const { owner, beneficiary, serial, amount, timeout } = cheque;
  const { sig } = await signCheque(swap, owner, cheque);
  return swap.submitCheque(beneficiary, serial, amount, timeout, sig, {
    from: beneficiary
  });
}

const simpleSwapTests = (accounts, Swap) => {
  const [owner, bob, alice] = accounts;

  async function prepareSwap(prefilledAmount = 1000) {
    const swap = await Swap.new(owner);
    await swap.send(prefilledAmount, { from: owner });
    return { swap, prefilledAmount: new BN(prefilledAmount) };
  }

  it("should accept deposits", async () => {
    const { swap } = await prepareSwap(0);
    const amount = new BN(1000);
    const { logs } = await swap.send(amount, { from: owner });

    expectEvent.inLogs(logs, "Deposit", {
      depositor: owner,
      amount
    });

    (await balance.current(swap.address)).should.bignumber.equal(amount);
  });

  it("should not accept a cheque with serial 0", async () => {
    const { swap, prefilledAmount } = await prepareSwap();
    await shouldFail.reverting(
      submitCheque(swap, {
        owner,
        beneficiary: bob,
        serial: new BN(0),
        amount: prefilledAmount
      })
    );
  });

  it("should accept valid cheque (increasing value, no hard deposit)", async () => {
    const { swap, prefilledAmount } = await prepareSwap();

    const cheque = {
      owner,
      beneficiary: bob,
      serial: new BN(1),
      amount: prefilledAmount
    };

    var { logs } = await submitCheque(swap, cheque);

    expectEvent.inLogs(logs, "ChequeSubmitted", {
      amount: cheque.amount,
      beneficiary: cheque.beneficiary,
      serial: cheque.serial
    });

    var chequeInfo = await swap.cheques(cheque.beneficiary);

    chequeInfo.serial.should.bignumber.equal(cheque.serial);
    chequeInfo.amount.should.bignumber.equal(cheque.amount);
    chequeInfo.timeout.should.bignumber.gte(
      (await time.latest()).addn(1 * epoch - 1)
    );

    await time.increase(1 * epoch);
    let beneficiaryExpectedBalance = (await balance.current(
      cheque.beneficiary
    )).add(cheque.amount);

    var { logs } = await swap.cashCheque(cheque.beneficiary);

    expectEvent.inLogs(logs, "ChequeCashed", {
      beneficiary: cheque.beneficiary,
      serial: cheque.serial,
      amount: cheque.amount
    });

    (await balance.current(cheque.beneficiary)).should.bignumber.equal(
      beneficiaryExpectedBalance
    );

    var chequeInfo = await swap.cheques(cheque.beneficiary);

    chequeInfo.paidOut.should.bignumber.equal(cheque.amount);
  });

  it("should not allow cheque payout before timeout", async () => {
    const { swap, prefilledAmount } = await prepareSwap();
    await submitCheque(swap, {
      owner,
      beneficiary: bob,
      serial: new BN(1),
      amount: prefilledAmount
    });
    await shouldFail.reverting(swap.cashCheque(bob));
  });

  it("should not allow cheque payout if there is nothing to pay out", async () => {
    const { swap, prefilledAmount } = await prepareSwap();
    await submitCheque(swap, {
      owner,
      beneficiary: bob,
      serial: new BN(1),
      amount: prefilledAmount
    });
    await time.increase(1 * epoch);
    await swap.cashCheque(bob);
    await shouldFail.reverting(swap.cashCheque(bob));
  });

  it("should not allow valid cheque if signed by owner and amount is not higher", async () => {
    const { swap, prefilledAmount } = await prepareSwap();
    await submitCheque(swap, {
      owner,
      beneficiary: bob,
      serial: new BN(1),
      amount: prefilledAmount
    });
    await shouldFail.reverting(
      submitCheque(swap, {
        owner,
        beneficiary: bob,
        serial: new BN(2),
        amount: prefilledAmount
      })
    );
  });

  it("should accept valid cheque with higher amount", async () => {
    const { swap } = await prepareSwap(new BN(1000));

    const beneficiary = bob;

    const cheques = [
      {
        owner,
        beneficiary,
        serial: new BN(1),
        amount: new BN(500)
      },
      {
        owner,
        beneficiary,
        serial: new BN(2),
        amount: new BN(1000)
      }
    ];

    await submitCheque(swap, cheques[0]);
    await time.increase(1 * epoch);
    await swap.cashCheque(beneficiary);

    var { logs } = await submitCheque(swap, cheques[1]);

    expectEvent.inLogs(logs, "ChequeSubmitted", {
      amount: cheques[1].amount,
      beneficiary: cheques[1].beneficiary,
      serial: cheques[1].serial
    });

    var chequeInfo = await swap.cheques(beneficiary);

    chequeInfo.serial.should.bignumber.equal(cheques[1].serial);
    chequeInfo.amount.should.bignumber.equal(cheques[1].amount);
    chequeInfo.timeout.should.bignumber.gte(
      (await time.latest()).addn(1 * epoch - 1)
    );

    await time.increase(1 * epoch);

    const payoutAmount = cheques[1].amount.sub(cheques[0].amount);

    const beneficiaryExpectedBalance = (await balance.current(beneficiary)).add(
      payoutAmount
    );

    var { logs } = await swap.cashCheque(beneficiary);

    expectEvent.inLogs(logs, "ChequeCashed", {
      beneficiary,
      serial: cheques[1].serial,
      amount: payoutAmount
    });

    (await balance.current(beneficiary)).should.bignumber.equal(
      beneficiaryExpectedBalance
    );

    var chequeInfo = await swap.cheques(beneficiary);

    chequeInfo.paidOut.should.bignumber.equal(cheques[1].amount);
  });

  it("should not allow cheque payout before increased timeout", async () => {
    const { swap } = await prepareSwap(1000);

    await submitCheque(swap, {
      owner,
      beneficiary: bob,
      serial: new BN(1),
      amount: new BN(500)
    });
    await time.increase(1 * epoch);
    await swap.cashCheque(bob);

    await shouldFail.reverting(swap.cashCheque(bob));
  });

  it("should accept a valid check with lower value if signed by beneficiary", async () => {
    const { swap } = await prepareSwap(1000);

    const beneficiary = bob;

    const cheques = [
      {
        owner,
        beneficiary,
        serial: new BN(1),
        amount: new BN(500),
        timeout: epoch
      },
      {
        owner,
        beneficiary,
        serial: new BN(2),
        amount: new BN(400),
        timeout: epoch
      }
    ];

    await submitCheque(swap, cheques[0]);

    const { sig: sigOwner } = await signCheque(swap, owner, cheques[1]);
    const { sig: sigBeneficiary } = await signCheque(
      swap,
      beneficiary,
      cheques[1]
    );

    const { logs } = await swap.submitChequeLower(
      beneficiary,
      cheques[1].serial,
      cheques[1].amount,
      epoch,
      sigOwner,
      sigBeneficiary
    );

    expectEvent.inLogs(logs, "ChequeSubmitted", {
      beneficiary,
      serial: cheques[1].serial,
      amount: cheques[1].amount
    });

    const chequeInfo = await swap.cheques(beneficiary);

    chequeInfo.serial.should.bignumber.equal(cheques[1].serial);
    chequeInfo.amount.should.bignumber.equal(cheques[1].amount);
  });

  it("should allow partial payments for a bouncing check", async () => {
    const { swap, prefilledAmount } = await prepareSwap(1000);

    const beneficiary = bob;

    const cheque = {
      owner,
      beneficiary,
      serial: new BN(1),
      amount: new BN(1500)
    };

    await submitCheque(swap, cheque);
    await time.increase(1 * epoch);

    var beneficiaryExpectedBalance = (await balance.current(beneficiary)).add(
      prefilledAmount
    );
    var { logs } = await swap.cashCheque(beneficiary);

    const bounced = cheque.amount.sub(prefilledAmount);

    expectEvent.inLogs(logs, "ChequeBounced", {
      paid: prefilledAmount,
      bounced,
      serial: cheque.serial,
      beneficiary
    });

    (await balance.current(beneficiary)).should.bignumber.equal(
      beneficiaryExpectedBalance
    );

    await swap.send(bounced);

    var beneficiaryExpectedBalance = (await balance.current(beneficiary)).add(
      bounced
    );

    var { logs } = await swap.cashCheque(beneficiary);

    expectEvent.inLogs(logs, "ChequeCashed", {
      amount: bounced,
      serial: cheque.serial,
      beneficiary
    });
  });

  it("should allow hard deposits if they do not exceed the global deposit", async () => {
    const { swap } = await prepareSwap(1000);
    const amount = new BN(500);
    var { logs } = await swap.increaseHardDeposit(bob, amount);

    expectEvent.inLogs(logs, "HardDepositChanged", {
      beneficiary: bob,
      amount
    });

    var { logs } = await swap.increaseHardDeposit(alice, amount);

    expectEvent.inLogs(logs, "HardDepositChanged", {
      beneficiary: alice,
      amount
    });

    (await swap.hardDeposits(bob)).amount.should.bignumber.equal(amount);
    (await swap.hardDeposits(alice)).amount.should.bignumber.equal(amount);
    (await swap.totalDeposit()).should.bignumber.equal(amount.muln(2));
  });

  it("should not allow hard deposits if they exceed the global deposit", async () => {
    const { swap } = await prepareSwap(1000);
    await shouldFail.reverting(swap.increaseHardDeposit(bob, 1001));
  });

  it("should use the hard deposit on valid cheque", async () => {
    const { swap } = await prepareSwap(1000);

    const beneficiary = bob;
    const hardDepositAmount = new BN(500);

    const cheque = {
      owner,
      beneficiary,
      serial: new BN(1),
      amount: new BN(400)
    };

    await swap.increaseHardDeposit(beneficiary, hardDepositAmount);

    await submitCheque(swap, cheque);
    await time.increase(1 * epoch);

    let beneficiaryExpectedBalance = (await balance.current(beneficiary)).add(
      cheque.amount
    );

    var { logs } = await swap.cashCheque(beneficiary);

    expectEvent.inLogs(logs, "ChequeCashed", {
      amount: cheque.amount,
      serial: cheque.serial,
      beneficiary
    });

    (await swap.hardDeposits(beneficiary)).amount.should.bignumber.equal(
      hardDepositAmount.sub(cheque.amount)
    );
    (await balance.current(beneficiary)).should.bignumber.equal(
      beneficiaryExpectedBalance
    );
  });

  // TODO: only part covered by HD, but enough
  // TODO: only part covered by HD, but not enough
  // TODO: many cheques test

  it("should not spend ether locked away by hard deposit of another address", async () => {
    const { swap, prefilledAmount } = await prepareSwap(1000);

    const beneficiary = alice;
    const cheque = {
      owner,
      beneficiary,
      serial: new BN(1),
      amount: new BN(600)
    };

    const hardDepositAmount = new BN(500);
    const available = prefilledAmount.sub(hardDepositAmount);

    await swap.increaseHardDeposit(bob, hardDepositAmount);

    await submitCheque(swap, cheque);
    await time.increase(1 * epoch);

    const expectedBalanceAlice = (await balance.current(beneficiary)).add(
      available
    );

    var { logs } = await swap.cashCheque(beneficiary);

    expectEvent.inLogs(logs, "ChequeBounced", {
      paid: available,
      bounced: cheque.amount.sub(available),
      serial: cheque.serial,
      beneficiary
    });

    (await balance.current(alice)).should.bignumber.equal(expectedBalanceAlice);
  });

  it("should not allow an instant decrease for hard deposits", async () => {
    const { swap } = await prepareSwap(1000);
    await swap.increaseHardDeposit(bob, 500);
    await shouldFail.reverting(swap.decreaseHardDeposit(bob));
  });

  it("should allow to prepare a decrease for hard deposits", async () => {
    const { swap } = await prepareSwap(1000);
    await swap.increaseHardDeposit(bob, 500);

    const diff = new BN(200);

    const { logs } = await swap.prepareDecreaseHardDeposit(bob, diff);

    expectEvent.inLogs(logs, "HardDepositDecreasePrepared", {
      beneficiary: bob,
      diff
    });

    const hardDeposit = await swap.hardDeposits(bob);

    hardDeposit.diff.should.bignumber.equal(diff);
    hardDeposit.timeout.should.bignumber.gte(
      (await time.latest()).addn(epoch - 1)
    );
  });

  it("should not allow to decrease hard deposit before the timeout", async () => {
    const { swap } = await prepareSwap(1000);
    await swap.increaseHardDeposit(bob, 500);
    await swap.prepareDecreaseHardDeposit(bob, 200);
    await shouldFail.reverting(swap.decreaseHardDeposit(bob));
  });

  it("should allow to decrease hard deposit after the timeout", async () => {
    const { swap } = await prepareSwap(1000);
    const inital = new BN(500);
    const diff = new BN(200);
    await swap.increaseHardDeposit(bob, inital);
    await swap.prepareDecreaseHardDeposit(bob, diff);

    let expectedHardDeposit = (await swap.hardDeposits(bob)).amount.sub(diff);

    await time.increase(2 * epoch);
    const { logs } = await swap.decreaseHardDeposit(bob);

    expectEvent.inLogs(logs, "HardDepositChanged", {
      beneficiary: bob,
      amount: inital.sub(diff)
    });

    (await swap.hardDeposits(bob)).amount.should.bignumber.equal(
      expectedHardDeposit
    );
    (await swap.totalDeposit()).should.bignumber.equal(expectedHardDeposit);
  });

  it("should not allow to do the same decrease twice", async () => {
    const { swap } = await prepareSwap(1000);
    await swap.increaseHardDeposit(bob, 500);
    await swap.prepareDecreaseHardDeposit(bob, 200);
    await time.increase(2 * epoch);
    await swap.decreaseHardDeposit(bob);
    await shouldFail.reverting(swap.decreaseHardDeposit(bob));
  });
};

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

    let { sig, hash } = await signNote(
      swap,
      owner,
      carol,
      1,
      noteAmount,
      constants.ZERO_ADDRESS,
      validity,
      0,
      "0x",
      epoch
    );

    await time.increase(4 * epoch);

    let encoded = await swap.encodeNote([
      swap.address,      
      1,
      noteAmount,
      carol,
      constants.ZERO_ADDRESS,
      validity,
      0,
      "0x",
      epoch
    ]);

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

    let { sig, hash } = await signNote(
      swap,
      owner,
      carol,
      1,
      noteAmount,
      oracle.address,
      0,
      bondTimeout,
      "0x",
      epoch
    );

    await oracle.testify(hash, 1);

    let encoded = await swap.encodeNote([
      swap.address,      
      1,
      noteAmount,
      carol,
      oracle.address,
      0,
      bondTimeout,
      "0x",
      epoch
    ]);
    await swap.submitNote(encoded, sig, { from: carol });

    const { paidOut, timeout } = await swap.notes(hash);

    paidOut.should.bignumber.equal(new BN(0));
    timeout.should.bignumber.gte((await time.latest()).addn(1 * epoch - 1));

    // cashout too soon
    await shouldFail.reverting(
      swap.cashNote(encoded, noteAmount, { from: carol })
    );

    await time.increase(1 * epoch);

    await oracle.testify(hash, 0);

    // oracle says no
    await shouldFail.reverting(
      swap.cashNote(encoded, noteAmount, { from: carol })
    );

    await oracle.testify(hash, 1);

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
        amount: new BN(100)
      },
      {
        serial: new BN(2),
        amount: new BN(200)
      },
      {
        owner,
        beneficiary,
        serial: new BN(3),
        amount: noteAmount.addn(200),
        timeout: epoch
      }
    ];

    await submitCheque(swap, cheques[0]);

    // completely offchain cheque of 100

    // owner issues note
    let note = await signNote(
      swap,
      owner,
      beneficiary,
      1,
      noteAmount,
      constants.ZERO_ADDRESS,
      0,
      0,
      "0x",
      epoch
    );

    // carol issues invoice
    let invoice = await signInvoice(
      swap,
      beneficiary,
      note.hash,
      cheques[1].amount,
      cheques[1].serial
    );

    // owner issues cheque for invoice
    let cheque = await signCheque(swap, owner, cheques[2]);

    let encoded = await swap.encodeNote([
      swap.address,      
      1,
      noteAmount,
      beneficiary,
      constants.ZERO_ADDRESS,
      0,
      0,
      "0x",
      epoch
    ]);
    // carol submits note anyway
    await swap.submitNote(encoded, note.sig, { from: carol });

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

    let { logs } = await swap.cashCheque(carol);

    expectEvent.inLogs(logs, "ChequeCashed", {
      beneficiary: carol,
      serial: cheques[2].serial,
      amount: cheques[2].amount
    });
  });
};

contract("SimpleSwap", function(accounts) {
  simpleSwapTests(accounts, SimpleSwap);
});

contract("SoftSwap", function(accounts) {
  simpleSwapTests(accounts, SoftSwap);
  softSwapTests(accounts, SoftSwap);
});

contract("Swap", function(accounts) {
  simpleSwapTests(accounts, Swap);
  softSwapTests(accounts, Swap);
  swapTests(accounts, Swap);
});
