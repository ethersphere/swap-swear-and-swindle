const SimpleSwap = artifacts.require("./SimpleSwap.sol");

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

const epoch = new BN(24 * 3600);

async function submitChequeBeneficiary(swap, cheque, sender) {
  cheque = { timeout: epoch, ...cheque };
  const { owner, beneficiary, serial, amount, timeout } = cheque;
  const { sig: ownerSig } = await signCheque(swap, owner, cheque); 
  return swap.submitChequeBeneficiary(serial, amount, timeout, ownerSig, {
    from: sender
  });
}

async function submitChequeOwner(swap, cheque, sender) {
    cheque = { timeout: epoch, ...cheque };
    const { owner, beneficiary, serial, amount, timeout } = cheque;
    const { sig: beneficiarySig } = await signCheque(swap, beneficiary, cheque);
    return swap.submitChequeOwner(beneficiary, serial, amount, timeout, beneficiarySig, {
      from: sender
    });
}

async function submitCheque(swap, cheque, sender) {
    cheque = { timeout: epoch, ...cheque };
    const { owner, beneficiary, serial, amount, timeout } = cheque;
    const { sig: ownerSig } = await signCheque(swap, owner, cheque);
    const { sig: beneficiarySig } = await signCheque(swap, beneficiary, cheque)
    return swap.submitCheque(beneficiary, serial, amount, timeout, ownerSig, beneficiarySig, {
      from: sender
    })
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
    const sender = bob
    const { swap, prefilledAmount } = await prepareSwap();
    await shouldFail.reverting(
      submitChequeBeneficiary(swap, {
        owner,
        beneficiary: bob,
        serial: new BN(0),
        amount: prefilledAmount
      }, sender)
    );
  });

  it("should accept a valid first cheque by owner", async () => {
    const { swap, prefilledAmount } = await prepareSwap();

    const cheque = {
      owner,
      beneficiary: bob,
      serial: new BN(1),
      amount: prefilledAmount,
    };

    const sender = cheque.owner

    var { logs } = await submitChequeOwner(swap, cheque, sender);

    expectEvent.inLogs(logs, "ChequeSubmitted", {
      amount: cheque.amount,
      beneficiary: cheque.beneficiary,
      serial: cheque.serial,
      timeout: epoch
    });

    var chequeInfo = await swap.cheques(cheque.beneficiary);

    chequeInfo.serial.should.bignumber.equal(cheque.serial);
    chequeInfo.amount.should.bignumber.equal(cheque.amount);
    chequeInfo.timeout.should.bignumber.gte(
      (await time.latest()).addn(1 * epoch - 1)
    );
  });

  it("should accept a valid first cheque by the beneficiary", async () => {
    const { swap, prefilledAmount } = await prepareSwap();

    const cheque = {
      owner,
      beneficiary: bob,
      serial: new BN(1),
      amount: prefilledAmount,
    };
    const sender = cheque.beneficiary

    var { logs } = await submitChequeBeneficiary(swap, cheque, sender);

    expectEvent.inLogs(logs, "ChequeSubmitted", {
      amount: cheque.amount,
      beneficiary: cheque.beneficiary,
      serial: cheque.serial,
      timeout: epoch
    });

    var chequeInfo = await swap.cheques(cheque.beneficiary);
    chequeInfo.serial.should.bignumber.equal(cheque.serial);
    chequeInfo.amount.should.bignumber.equal(cheque.amount);
    chequeInfo.timeout.should.bignumber.gte(
      (await time.latest()).addn(1 * epoch - 1)
    );
  })

  it("should accept a valid first cheque by another party", async () => {
    const { swap, prefilledAmount } = await prepareSwap();

    const cheque = {
      owner,
      beneficiary: bob,
      serial: new BN(1),
      amount: prefilledAmount,
    };

    const sender = alice

    var { logs } = await submitCheque(swap, cheque, sender);

    expectEvent.inLogs(logs, "ChequeSubmitted", {
      amount: cheque.amount,
      beneficiary: cheque.beneficiary,
      serial: cheque.serial,
      timeout: epoch
    });

    var chequeInfo = await swap.cheques(cheque.beneficiary);
    chequeInfo.serial.should.bignumber.equal(cheque.serial);
    chequeInfo.amount.should.bignumber.equal(cheque.amount);
    chequeInfo.timeout.should.bignumber.gte(
      (await time.latest()).addn(1 * epoch - 1)
    );
  })

  it("should accept a valid second cheque by owner (higher amount)", async () => {
    const { swap, prefilledAmount } = await prepareSwap();

    const amounts = [new BN(0.4 * prefilledAmount), new BN(0.6 * prefilledAmount)]

    const cheques = [{
      owner,
      beneficiary: bob,
      serial: new BN(1),
      amount: amounts[0],
    }, {
      owner,
      beneficiary: bob,
      serial: new BN(2),
      amount: amounts[1],
    }];

    const sender = cheques[0].owner

    var { logs } = await submitChequeOwner(swap, cheques[0], sender);

    expectEvent.inLogs(logs, "ChequeSubmitted", {
      amount: cheques[0].amount,
      beneficiary: cheques[0].beneficiary,
      serial: cheques[0].serial,
      timeout: epoch
    });

    var chequeInfo = await swap.cheques(cheques[0].beneficiary);

    chequeInfo.serial.should.bignumber.equal(cheques[0].serial);
    chequeInfo.amount.should.bignumber.equal(cheques[0].amount);
    chequeInfo.timeout.should.bignumber.gte(
      (await time.latest()).addn(1 * epoch - 1)
    );

    var { logs } = await submitChequeOwner(swap, cheques[1], sender);

    expectEvent.inLogs(logs, "ChequeSubmitted", {
      amount: amounts[1],
      beneficiary: cheques[1].beneficiary,
      serial: cheques[1].serial,
      timeout: epoch
    });

    var chequeInfo = await swap.cheques(cheques[1].beneficiary);

    chequeInfo.serial.should.bignumber.equal(cheques[1].serial);
    chequeInfo.amount.should.bignumber.equal(cheques[1].amount);
    chequeInfo.timeout.should.bignumber.gte(
      (await time.latest()).addn(1 * epoch - 1)
    );
  });

  it("should accept a valid second cheque by owner (lower amount)", async () => {
    const { swap, prefilledAmount } = await prepareSwap();

    const amounts = [new BN(0.6 * prefilledAmount), new BN(0.4 * prefilledAmount)]

    const cheques = [{
      owner,
      beneficiary: bob,
      serial: new BN(1),
      amount: amounts[0],
    }, {
      owner,
      beneficiary: bob,
      serial: new BN(2),
      amount: amounts[1],
    }];

    const sender = cheques[0].owner

    var { logs } = await submitChequeOwner(swap, cheques[0], sender);

    expectEvent.inLogs(logs, "ChequeSubmitted", {
      amount: cheques[0].amount,
      beneficiary: cheques[0].beneficiary,
      serial: cheques[0].serial,
      timeout: epoch
    });

    var chequeInfo = await swap.cheques(cheques[0].beneficiary);

    chequeInfo.serial.should.bignumber.equal(cheques[0].serial);
    chequeInfo.amount.should.bignumber.equal(cheques[0].amount);
    chequeInfo.timeout.should.bignumber.gte(
      (await time.latest()).addn(1 * epoch - 1)
    );

    var { logs } = await submitChequeOwner(swap, cheques[1], sender);

    expectEvent.inLogs(logs, "ChequeSubmitted", {
      amount: amounts[1],
      beneficiary: cheques[1].beneficiary,
      serial: cheques[1].serial,
      timeout: epoch
    });

    var chequeInfo = await swap.cheques(cheques[1].beneficiary);

    chequeInfo.serial.should.bignumber.equal(cheques[1].serial);
    chequeInfo.amount.should.bignumber.equal(cheques[1].amount);
    chequeInfo.timeout.should.bignumber.gte(
      (await time.latest()).addn(1 * epoch - 1)
    );
  });

  it("should accept a valid second cheque by the beneficiary (higher amount)", async () => {
    const { swap, prefilledAmount } = await prepareSwap();

    const amounts = [new BN(0.4 * prefilledAmount), new BN(0.6 * prefilledAmount)]

    const cheques = [{
      owner,
      beneficiary: bob,
      serial: new BN(1),
      amount: amounts[0],
    }, {
      owner,
      beneficiary: bob,
      serial: new BN(2),
      amount: amounts[1],
    }];

    const sender = cheques[0].beneficiary

    var { logs } = await submitChequeBeneficiary(swap, cheques[0], sender);

    expectEvent.inLogs(logs, "ChequeSubmitted", {
      amount: cheques[0].amount,
      beneficiary: cheques[0].beneficiary,
      serial: cheques[0].serial,
      timeout: epoch
    });

    var chequeInfo = await swap.cheques(cheques[0].beneficiary);

    chequeInfo.serial.should.bignumber.equal(cheques[0].serial);
    chequeInfo.amount.should.bignumber.equal(cheques[0].amount);
    chequeInfo.timeout.should.bignumber.gte(
      (await time.latest()).addn(1 * epoch - 1)
    );

    var { logs } = await submitChequeBeneficiary(swap, cheques[1], sender);

    expectEvent.inLogs(logs, "ChequeSubmitted", {
      amount: amounts[1],
      beneficiary: cheques[0].beneficiary,
      serial: cheques[1].serial,
      timeout: epoch
    });

    var chequeInfo = await swap.cheques(cheques[1].beneficiary);

    chequeInfo.serial.should.bignumber.equal(cheques[1].serial);
    chequeInfo.amount.should.bignumber.equal(cheques[1].amount);
    chequeInfo.timeout.should.bignumber.gte(
      (await time.latest()).addn(1 * epoch - 1)
    );
  })

  it("should accept a valid second cheque by the beneficiary (lower amount)", async () => {
    const { swap, prefilledAmount } = await prepareSwap();

    const amounts = [new BN(0.6 * prefilledAmount), new BN(0.4 * prefilledAmount)]

    const cheques = [{
      owner,
      beneficiary: bob,
      serial: new BN(1),
      amount: amounts[0],
    }, {
      owner,
      beneficiary: bob,
      serial: new BN(2),
      amount: amounts[1],
    }];

    const sender = cheques[0].beneficiary

    var { logs } = await submitChequeBeneficiary(swap, cheques[0], sender);

    expectEvent.inLogs(logs, "ChequeSubmitted", {
      amount: cheques[0].amount,
      beneficiary: cheques[0].beneficiary,
      serial: cheques[0].serial,
      timeout: epoch
    });

    var chequeInfo = await swap.cheques(cheques[0].beneficiary);

    chequeInfo.serial.should.bignumber.equal(cheques[0].serial);
    chequeInfo.amount.should.bignumber.equal(cheques[0].amount);
    chequeInfo.timeout.should.bignumber.gte(
      (await time.latest()).addn(1 * epoch - 1)
    );

    var { logs } = await submitChequeBeneficiary(swap, cheques[1], sender);

    expectEvent.inLogs(logs, "ChequeSubmitted", {
      amount: amounts[1],
      beneficiary: cheques[1].beneficiary,
      serial: cheques[1].serial,
      timeout: epoch
    });

    var chequeInfo = await swap.cheques(cheques[1].beneficiary);

    chequeInfo.serial.should.bignumber.equal(cheques[1].serial);
    chequeInfo.amount.should.bignumber.equal(cheques[1].amount);
    chequeInfo.timeout.should.bignumber.gte(
      (await time.latest()).addn(1 * epoch - 1)
    );
  })

  it("should accept a valid second cheque by another party (higher amount)", async () => {
    const { swap, prefilledAmount } = await prepareSwap();

    const amounts = [new BN(0.4 * prefilledAmount), new BN(0.6 * prefilledAmount)]

    const cheques = [{
      owner,
      beneficiary: bob,
      serial: new BN(1),
      amount: amounts[0],
    }, {
      owner,
      beneficiary: bob,
      serial: new BN(2),
      amount: amounts[1],
    }];

    const sender = alice

    var { logs } = await submitCheque(swap, cheques[0], sender);

    expectEvent.inLogs(logs, "ChequeSubmitted", {
      amount: cheques[0].amount,
      beneficiary: cheques[0].beneficiary,
      serial: cheques[0].serial,
      timeout: epoch
    });

    var chequeInfo = await swap.cheques(cheques[0].beneficiary);

    chequeInfo.serial.should.bignumber.equal(cheques[0].serial);
    chequeInfo.amount.should.bignumber.equal(cheques[0].amount);
    chequeInfo.timeout.should.bignumber.gte(
      (await time.latest()).addn(1 * epoch - 1)
    );

    var { logs } = await submitCheque(swap, cheques[1], sender);

    expectEvent.inLogs(logs, "ChequeSubmitted", {
      amount: amounts[1],
      beneficiary: cheques[0].beneficiary,
      serial: cheques[1].serial,
      timeout: epoch
    });

    var chequeInfo = await swap.cheques(cheques[1].beneficiary);

    chequeInfo.serial.should.bignumber.equal(cheques[1].serial);
    chequeInfo.amount.should.bignumber.equal(cheques[1].amount);
    chequeInfo.timeout.should.bignumber.gte(
      (await time.latest()).addn(1 * epoch - 1)
    );
  })

  it("should accept a valid second cheque by another party (lower amount)", async () => {
    const { swap, prefilledAmount } = await prepareSwap();

    const amounts = [new BN(0.6 * prefilledAmount), new BN(0.4 * prefilledAmount)]

    const cheques = [{
      owner,
      beneficiary: bob,
      serial: new BN(1),
      amount: amounts[0],
    }, {
      owner,
      beneficiary: bob,
      serial: new BN(2),
      amount: amounts[1],
    }];

    const sender = alice

    var { logs } = await submitCheque(swap, cheques[0], sender);

    expectEvent.inLogs(logs, "ChequeSubmitted", {
      amount: cheques[0].amount,
      beneficiary: cheques[0].beneficiary,
      serial: cheques[0].serial,
      timeout: epoch
    });

    var chequeInfo = await swap.cheques(cheques[0].beneficiary);

    chequeInfo.serial.should.bignumber.equal(cheques[0].serial);
    chequeInfo.amount.should.bignumber.equal(cheques[0].amount);
    chequeInfo.timeout.should.bignumber.gte(
      (await time.latest()).addn(1 * epoch - 1)
    );

    var { logs } = await submitCheque(swap, cheques[1], sender);

    expectEvent.inLogs(logs, "ChequeSubmitted", {
      amount: amounts[1],
      beneficiary: cheques[0].beneficiary,
      serial: cheques[1].serial,
      timeout: epoch
    });

    var chequeInfo = await swap.cheques(cheques[1].beneficiary);

    chequeInfo.serial.should.bignumber.equal(cheques[1].serial);
    chequeInfo.amount.should.bignumber.equal(cheques[1].amount);
    chequeInfo.timeout.should.bignumber.gte(
      (await time.latest()).addn(1 * epoch - 1)
    );
  })

  it("should not allow the owner to submit a cheque via submitChequeBeneficiary", async() => {
    const { swap, prefilledAmount } = await prepareSwap();

    const cheque = {
      owner,
      beneficiary: bob,
      serial: new BN(1),
      amount: prefilledAmount,
    };
    const sender = cheque.owner

    await shouldFail.reverting(submitChequeBeneficiary(swap, cheque, sender))

  })

  it("should not allow the beneficiary to submit a cheque via submitChequeOwner", async() => {
    const { swap, prefilledAmount } = await prepareSwap();

    const cheque = {
      owner,
      beneficiary: bob,
      serial: new BN(1),
      amount: prefilledAmount,
    };
    const sender = cheque.beneficiary

    await shouldFail.reverting(submitChequeOwner(swap, cheque, sender))
  })

  it("should not allow a cheque with the same or decreasing serial number", async () => {
    const { swap, prefilledAmount } = await prepareSwap();

    const serials = [new BN(1), new BN(0), new BN(1)]

    const cheques = [{
      owner,
      beneficiary: bob,
      serial: serials[0],
      amount: prefilledAmount,
    }, {
      owner,
      beneficiary: bob,
      serial: serials[1],
      amount: prefilledAmount,
    }, {
      owner,
      beneficiary: bob,
      serial: serials[2],
      amount: prefilledAmount,
    }
  ]
    const sender = cheques[0].beneficiary
    await submitChequeBeneficiary(swap, cheques[0], sender)
    await shouldFail.reverting(submitChequeBeneficiary(swap, cheques[1], sender))
    await shouldFail.reverting(submitChequeBeneficiary(swap, cheques[2], sender))
  })

  it("should be cheaper in gas to submit the same cheque via submitChequeOwner or submitChequeBeneficiary than via submitCheque", async() => {
    const { swap: swapAlice, prefilledAmount } = await prepareSwap();
    const { swap: swapOwner, prefilledAmount1 } = await prepareSwap();
    const { swap: swapBeneficiary, prefilledAmount2 } = await prepareSwap();

    const cheque = {
      owner,
      beneficiary: bob,
      serial: new BN(1),
      amount: prefilledAmount,
    }

    const { receipt: receiptOne} = await submitChequeOwner(swapOwner, cheque, cheque.owner)
    const { receipt: receiptTwo } = await submitChequeBeneficiary(swapBeneficiary, cheque, cheque.beneficiary)
    const { receipt: receiptThree } = await submitCheque(swapAlice, cheque, alice)

    let gasCostOwner = await computeCost(receiptOne)
    let gasCostBeneficiary = await computeCost(receiptTwo)
    let gasCostAlice = await computeCost(receiptThree)

    gasCostOwner.should.bignumber.be.below(gasCostAlice, "submitChequeOwner should cost less than submitCheque")
    gasCostBeneficiary.should.bignumber.be.below(gasCostAlice, "submitChequeBeneficiary should cost less than submitCheque")
  })

  it("should allow cheque payout after timeout", async () => {
    const { swap, prefilledAmount } = await prepareSwap();

    const cheque = {
      owner,
      beneficiary: bob,
      serial: new BN(1),
      amount: prefilledAmount,
    };

    const sender = cheque.owner

    var { logs } = await submitChequeOwner(swap, cheque, sender);

    expectEvent.inLogs(logs, "ChequeSubmitted", {
      amount: cheque.amount,
      beneficiary: cheque.beneficiary,
      serial: cheque.serial,
      timeout: epoch
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

    var { logs } = await swap.cashCheque(cheque.beneficiary, cheque.amount);

    expectEvent.inLogs(logs, "ChequeCashed", {
      beneficiary: cheque.beneficiary,
      serial: cheque.serial,
      payout: cheque.amount,
      requestPayout: cheque.amount
    });

    (await balance.current(cheque.beneficiary)).should.bignumber.equal(
      beneficiaryExpectedBalance
    );

    var chequeInfo = await swap.cheques(cheque.beneficiary);

    chequeInfo.paidOut.should.bignumber.equal(cheque.amount);

  })

  it("should not allow cheque payout before timeout", async () => {
    const { swap, prefilledAmount } = await prepareSwap();
    await submitChequeBeneficiary(swap, {
      owner,
      beneficiary: bob,
      serial: new BN(1),
      amount: prefilledAmount
    }, bob);
    await shouldFail.reverting(swap.cashCheque(bob, prefilledAmount));
  });

  it("should not allow cheque payout if there is nothing to pay out", async () => {
    const { swap, prefilledAmount } = await prepareSwap();
    await submitChequeBeneficiary(swap, {
      owner,
      beneficiary: bob,
      serial: new BN(1),
      amount: prefilledAmount
    }, bob);
    await time.increase(1 * epoch);
    await swap.cashCheque(bob, prefilledAmount);
    await shouldFail.reverting(swap.cashCheque(bob, prefilledAmount));
  });

  it("should not allow cheque payout before increased timeout", async () => {
    const { swap } = await prepareSwap(1000);

    await submitChequeBeneficiary(swap, {
      owner,
      beneficiary: bob,
      serial: new BN(1),
      amount: new BN(500)
    }, bob);
    await time.increase(1 * epoch);
    await swap.cashCheque(bob, new BN(500));

    await shouldFail.reverting(swap.cashCheque(bob, new BN(500)));
  });

  it("should allow partial payments for a bouncing check", async () => {
    const { swap, prefilledAmount } = await prepareSwap(1000);

    const beneficiary = bob;

    const cheque = {
      owner,
      beneficiary, 
      serial: new BN(1),
      amount: new BN(1500),
    };
    const sender = cheque.beneficiary

    await submitChequeBeneficiary(swap, cheque, sender);
    await time.increase(1 * epoch);

    var beneficiaryExpectedBalance = (await balance.current(beneficiary)).add(
      prefilledAmount
    );
    var { logs } = await swap.cashCheque(beneficiary, cheque.amount);

    const bounced = cheque.amount.sub(prefilledAmount);

    expectEvent.inLogs(logs, "ChequeBounced", { });

    (await balance.current(beneficiary)).should.bignumber.equal(
      beneficiaryExpectedBalance
    );

    await swap.send(bounced);

    var beneficiaryExpectedBalance = (await balance.current(beneficiary)).add(
      bounced
    );

    var { logs } = await swap.cashCheque(beneficiary, bounced);

    expectEvent.inLogs(logs, "ChequeCashed", {
      payout: bounced,
      requestPayout: bounced,
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
      amount: new BN(400),
    };

    const sender = cheque.beneficiary

    await swap.increaseHardDeposit(beneficiary, hardDepositAmount);

    await submitChequeBeneficiary(swap, cheque, sender);
    await time.increase(1 * epoch);

    let beneficiaryExpectedBalance = (await balance.current(beneficiary)).add(
      cheque.amount
    );

    var { logs } = await swap.cashCheque(beneficiary, cheque.amount);

    expectEvent.inLogs(logs, "ChequeCashed", {
      payout: cheque.amount,
      requestPayout: cheque.amount,
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
      amount: new BN(600),
    };

    const sender = cheque.beneficiary

    const hardDepositAmount = new BN(500);
    const available = prefilledAmount.sub(hardDepositAmount);

    await swap.increaseHardDeposit(bob, hardDepositAmount);

    await submitChequeBeneficiary(swap, cheque, sender);
    await time.increase(1 * epoch);

    const expectedBalanceAlice = (await balance.current(beneficiary)).add(
      available
    );

    var { logs } = await swap.cashCheque(beneficiary, cheque.amount);

    expectEvent.inLogs(logs, "ChequeBounced", { });
    expectEvent.inLogs(logs, "ChequeCashed", { 
      payout: available,
      requestPayout: cheque.amount,
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

contract("SimpleSwap", function(accounts) {
  simpleSwapTests(accounts, SimpleSwap);
});
