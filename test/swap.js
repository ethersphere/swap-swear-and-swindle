const Swap = artifacts.require("./Swap.sol");
const OracleWitness = artifacts.require("./OracleWitness.sol");

require('chai')
    .use(require('chai-as-promised'))
    .use(require('bn-chai')(web3.utils.BN))
    .should();

const { matchLogs, matchStruct, computeCost } = require('./testutils')
const { signCheque, signNote, signInvoice } = require('./swutils')
const { balance, time, shouldFail, constants } = require('openzeppelin-test-helpers')

const epoch = 24 * 3600

contract('swap', function(accounts) {
  const [owner, bob, alice, carol] = accounts

  async function submitCheque(swap, signer, beneficiary, serial, amount, timeout = epoch) {
    const { sig } = await signCheque(swap, signer, beneficiary, serial, amount, timeout);
    return swap.submitCheque(beneficiary, serial, amount, timeout, sig, { from: beneficiary });
  }

  async function prepareSwap(amount = 1000) {
    const swap = await Swap.new(owner);
    await swap.send(amount, { from: owner });
    return { swap, amount }
  }

  it('should accept deposits', async() => {
    const { swap } = await prepareSwap(0)
    const amount = 1000
    const { logs } = await swap.send(amount, { from: owner });

    matchLogs(logs, [
      { event: 'Deposit', args: { depositor: owner, amount }}
    ]);

    (await balance.current(swap.address)).should.eq.BN(amount);
  })

  it('should not accept a cheque with serial 0', async() => {
    const { swap, amount } = await prepareSwap()
    await shouldFail.reverting(submitCheque(swap, owner, bob, 0, amount));
  })

  it('should accept valid cheque (increasing value, no hard deposit)', async() => {
    const { swap, amount } = await prepareSwap()
    const serial = 1;

    var { logs } = await submitCheque(swap, owner, bob, serial, amount);

    matchLogs(logs, [
      { event: 'ChequeSubmitted', args: { amount, beneficiary: bob, serial } }
    ])

    var chequeInfo = await swap.cheques(bob);

    chequeInfo.serial.should.eq.BN(serial)
    chequeInfo.amount.should.eq.BN(amount)
    chequeInfo.timeout.should.gte.BN((await time.latest()).addn(1 * epoch - 1))

    await time.increase(1 * epoch);
    let beneficiaryExpectedBalance = (await balance.current(bob)).addn(amount);

    var { logs } = await swap.cashCheque(bob);

    matchLogs(logs, [
      { event: 'ChequeCashed', args: { beneficiary: bob, serial: 1, amount }}
    ]);

    (await balance.current(bob)).should.eq.BN(beneficiaryExpectedBalance);

    var chequeInfo = await swap.cheques(bob);

    chequeInfo.paidOut.should.eq.BN(amount);
  })

  it('should not allow cheque payout before timeout', async() => {
    const { swap, amount } = await prepareSwap()
    await submitCheque(swap, owner, bob, 1, amount);
    await shouldFail.reverting(swap.cashCheque(bob));
  })

  it('should not allow cheque payout if there is nothing to pay out', async() => {
    const { swap, amount } = await prepareSwap()
    await submitCheque(swap, owner, bob, 1, amount);
    await time.increase(1 * epoch);
    await swap.cashCheque(bob);
    await shouldFail.reverting(swap.cashCheque(bob));
  })

  it('should not allow valid cheque if signed by owner and amount is not higher', async() => {
    const { swap, amount } = await prepareSwap()
    await submitCheque(swap, owner, bob, 1, amount);
    await shouldFail.reverting(submitCheque(swap, owner, bob, 2, amount));
  })

  it('should accept valid cheque with higher amount', async() => {
    const { swap, amount } = await prepareSwap(1000)

    await submitCheque(swap, owner, bob, 1, 500)
    await time.increase(1 * epoch);
    await swap.cashCheque(bob);

    var { logs } = await submitCheque(swap, owner, bob, 2, amount)

    matchLogs(logs, [
      { event: 'ChequeSubmitted', args: { amount, beneficiary: bob, serial: 2 } }
    ])

    var chequeInfo = await swap.cheques(bob);

    chequeInfo.serial.should.eq.BN(2)
    chequeInfo.amount.should.eq.BN(amount)
    chequeInfo.timeout.should.gte.BN((await time.latest()).addn(1 * epoch - 1))

    await time.increase(1 * epoch);

    const beneficiaryExpectedBalance = (await balance.current(bob)).addn(500)

    var { logs } = await swap.cashCheque(bob);

    matchLogs(logs, [
      { event: 'ChequeCashed', args: { beneficiary: bob, serial: 2, amount: 500 }}
    ]);

    (await balance.current(bob)).should.eq.BN(beneficiaryExpectedBalance);

    var chequeInfo = await swap.cheques(bob);

    chequeInfo.paidOut.should.eq.BN(amount);
  })

  it('should not allow cheque payout before increased timeout', async() => {
    const { swap, amount } = await prepareSwap(1000)

    await submitCheque(swap, owner, bob, 1, 500)
    await time.increase(1 * epoch);
    await swap.cashCheque(bob);

    await shouldFail.reverting(swap.cashCheque(bob));
  })

  it('should accept a valid check with lower value if signed by beneficiary', async () => {
    const { swap } = await prepareSwap(1000)

    await submitCheque(swap, owner, bob, 1, 500)

    const { sig: sigOwner } = await signCheque(swap, owner, bob, 2, 400, epoch);
    const { sig: sigBob } = await signCheque(swap, bob, bob, 2, 400, epoch);

    const { logs } = await swap.submitChequeLower(bob, 2, 400, epoch, sigOwner, sigBob);

    matchLogs(logs, [
      { event: 'ChequeSubmitted', args: { beneficiary: bob, serial: 2, amount: 400 }}
    ]);

    const chequeInfo = await swap.cheques(bob);

    chequeInfo.serial.should.eq.BN(2)
    chequeInfo.amount.should.eq.BN(400)
  })

  it('should allow partial payments for a bouncing check', async () => {
    const { swap } = await prepareSwap(1000)

    await submitCheque(swap, owner, bob, 1, 1500)
    await time.increase(1 * epoch)

    var beneficiaryExpectedBalance = (await balance.current(bob)).addn(1000);
    var { logs } = await swap.cashCheque(bob)

    matchLogs(logs, [
      { event: 'ChequeBounced', args: { paid: 1000, bounced: 500, serial: 1, beneficiary: bob } }
    ]);

    (await balance.current(bob)).should.eq.BN(beneficiaryExpectedBalance);

    await swap.send(500)

    var beneficiaryExpectedBalance = (await balance.current(bob)).addn(500);

    var { logs } = await swap.cashCheque(bob)

    matchLogs(logs, [
      { event: 'ChequeCashed', args: { amount: 500, serial: 1, beneficiary: bob } }
    ]);
  })

  it('should allow hard deposits if they do not exceed the global deposit', async() => {
    const { swap } = await prepareSwap(1000)
    var { logs } = await swap.increaseHardDeposit(bob, 500);

    matchLogs(logs, [
      { event: 'HardDepositChanged', args: { beneficiary: bob, amount: 500 } }
    ]);

    var { logs } = await swap.increaseHardDeposit(alice, 500);

    matchLogs(logs, [
      { event: 'HardDepositChanged', args: { beneficiary: alice, amount: 500 } }
    ]);

    (await swap.hardDeposits(bob)).amount.should.eq.BN(500);
    (await swap.hardDeposits(alice)).amount.should.eq.BN(500);
    (await swap.totalDeposit()).should.eq.BN(1000);
  })


  it('should not allow hard deposits if they exceed the global deposit', async() => {
    const { swap } = await prepareSwap(1000)
    await shouldFail.reverting(swap.increaseHardDeposit(bob, 1001));
  })

  it('should use the hard deposit on valid cheque', async() => {
    const { swap } = await prepareSwap(1000)
    await swap.increaseHardDeposit(bob, 500)

    await submitCheque(swap, owner, bob, 1, 400)
    await time.increase(1 * epoch);

    let beneficiaryExpectedBalance = (await balance.current(bob)).addn(400);

    var { logs } = await swap.cashCheque(bob);

    matchLogs(logs, [
      { event: 'ChequeCashed', args: { amount: 400, serial: 1, beneficiary: bob } }
    ]);

    (await swap.hardDeposits(bob)).amount.should.eq.BN(100);
    (await balance.current(bob)).should.eq.BN(beneficiaryExpectedBalance);
  })

  // TODO: only part covered by HD, but enough
  // TODO: only part covered by HD, but not enough
  // TODO: many cheques test

  it('should not spend ether locked away by hard deposit of another address', async() => {
    const { swap } = await prepareSwap(1000)
    await swap.increaseHardDeposit(bob, 500)

    await submitCheque(swap, owner, alice, 1, 600)
    await time.increase(1 * epoch)

    const expectedBalanceAlice = (await balance.current(alice)).addn(500);

    var { logs } = await swap.cashCheque(alice)

    matchLogs(logs, [
      { event: 'ChequeBounced', args: { paid: 500, bounced:100, serial: 1, beneficiary: alice } }
    ]);

    (await balance.current(alice)).should.eq.BN(expectedBalanceAlice);
  })

  it('should not allow an instant decrease for hard deposits', async() => {
    const { swap } = await prepareSwap(1000)
    await swap.increaseHardDeposit(bob, 500)
    await shouldFail.reverting(swap.decreaseHardDeposit(bob));
  })

  it('should allow to prepare a decrease for hard deposits', async() => {
    const { swap } = await prepareSwap(1000)
    await swap.increaseHardDeposit(bob, 500)

    const { logs } = await swap.prepareDecreaseHardDeposit(bob, 200);

    matchLogs(logs, [
      { event: 'HardDepositDecreasePrepared', args: { beneficiary: bob, diff: 200 } }
    ]);

    const hardDeposit = await swap.hardDeposits(bob);

    hardDeposit.diff.should.eq.BN(200);
    hardDeposit.timeout.should.gte.BN((await time.latest()).addn(2 * epoch - 1));
  })

  it('should not allow to decrease hard deposit before the timeout', async() => {
    const { swap } = await prepareSwap(1000)
    await swap.increaseHardDeposit(bob, 500)
    await swap.prepareDecreaseHardDeposit(bob, 200);
    await shouldFail.reverting(swap.decreaseHardDeposit(bob));
  })

  it('should allow to decrease hard deposit after the timeout', async() => {
    const { swap } = await prepareSwap(1000)
    await swap.increaseHardDeposit(bob, 500)
    await swap.prepareDecreaseHardDeposit(bob, 200);

    let expectedHardDeposit = (await swap.hardDeposits(bob)).amount.subn(200);

    await time.increase(2 * epoch);
    const { logs } = await swap.decreaseHardDeposit(bob);

    matchLogs(logs, [
      { event: 'HardDepositChanged', args: { beneficiary: bob, amount: 300 } }
    ]);

    (await swap.hardDeposits(bob)).amount.should.eq.BN(expectedHardDeposit);
    (await swap.totalDeposit()).should.eq.BN(expectedHardDeposit);
  })

  it('should not allow to do the same decrease twice', async() => {
    const { swap } = await prepareSwap(1000)
    await swap.increaseHardDeposit(bob, 500)
    await swap.prepareDecreaseHardDeposit(bob, 200);
    await time.increase(2 * epoch);
    await swap.decreaseHardDeposit(bob);
    await shouldFail.reverting(swap.decreaseHardDeposit(bob));
  })

  // TODO: split
  it('should accept a valid note (bond)', async() => {
    const { swap } = await prepareSwap(1000)

    const noteTimeout = 5000;
    const noteAmount = 500

    let validity = (await time.latest()).addn(noteTimeout)

    let { sig, hash } = await signNote(swap, owner, carol, 1, noteAmount, constants.ZERO_ADDRESS, validity, 0, "0x")

    await time.increase(4 * epoch)

    let encoded = await swap.encodeNote(swap.address, carol, 1, noteAmount, constants.ZERO_ADDRESS, validity, 0, "0x")

    await swap.submitNote(encoded, sig, { from: carol });

    const { paidOut, timeout } = await swap.notes(hash)

    paidOut.should.eq.BN(0)
    timeout.should.gte.BN((await time.latest()).addn(1 * epoch - 1))

    await time.increase(1 * epoch)

    let expectedBalanceCarol = (await balance.current(carol)).addn(noteAmount)

    let { receipt } = await swap.cashNote(encoded, noteAmount, { from: carol });

    expectedBalanceCarol = expectedBalanceCarol.sub(await computeCost(receipt));

    (await balance.current(carol)).should.eq.BN(expectedBalanceCarol)

    // already fully cashed out
    await shouldFail.reverting(swap.cashNote(encoded, noteAmount, { from: carol }));
  })

  // TODO: split
  it('should accept a valid note (conditional bond)', async() => {
    const { swap } = await prepareSwap(1000)
    const oracle = await OracleWitness.new();

    const noteAmount = 500
    const noteTimeout = 2 * epoch

    let bondTimeout = (await time.latest()).addn(noteTimeout)

    let { sig, hash } = await signNote(swap, owner, carol, 1, noteAmount, oracle.address, 0, bondTimeout, "0x")

    await oracle.testify(hash, 1)

    let encoded = await swap.encodeNote(swap.address, carol, 1, noteAmount, oracle.address, 0, bondTimeout, "0x")
    await swap.submitNote(encoded, sig, { from: carol });

    const { paidOut, timeout } = await swap.notes(hash)

    paidOut.should.eq.BN(0)
    timeout.should.gte.BN((await time.latest()).addn(1 * epoch - 1))

    // cashout too soon
    await shouldFail.reverting(swap.cashNote(encoded, noteAmount, { from: carol }))

    await time.increase(1 * epoch)

    await oracle.testify(hash, 0)

    // oracle says no
    await shouldFail.reverting(swap.cashNote(encoded, noteAmount, { from: carol }))

    await oracle.testify(hash, 1)

    // partial payment
    let expectedBalanceCarol = (await balance.current(carol)).addn(noteAmount / 4)
    var { receipt } = await swap.cashNote(encoded, noteAmount / 4, { from: carol });
    expectedBalanceCarol = expectedBalanceCarol.sub(await computeCost(receipt));
    (await balance.current(carol)).should.eq.BN(expectedBalanceCarol)

    // partial payment
    expectedBalanceCarol = (await balance.current(carol)).addn(noteAmount / 4)
    var { receipt } = await swap.cashNote(encoded, noteAmount / 4, { from: carol });
    expectedBalanceCarol = expectedBalanceCarol.sub(await computeCost(receipt));
    (await balance.current(carol)).should.eq.BN(expectedBalanceCarol)

    await time.increase(2 * epoch)

    // too late for the rest
    await shouldFail.reverting(swap.cashNote(encoded, noteAmount / 4, { from: carol }))
  })

  // TODO: split
  it('should allow to submit paid invoices', async() => {
    const { swap } = await prepareSwap(1000)
    const noteAmount = 500

    await submitCheque(swap, owner, carol, 1, 100)

    // completely offchain cheque of 100

    // owner issues note
    let note = await signNote(swap, owner, carol, 1, noteAmount, constants.ZERO_ADDRESS, 0, 0, "0x")

    // carol issues invoice
    let invoice = await signInvoice(swap, carol, note.hash, 200, 2)

    // owner issues cheque for invoice
    let cheque = await signCheque(swap, owner, carol, 3, noteAmount + 200, epoch)

    let encoded = await swap.encodeNote(swap.address, carol, 1, noteAmount, constants.ZERO_ADDRESS, 0, 0, "0x")
    // carol submits note anyway
    await swap.submitNote(encoded, note.sig, { from: carol });

    // owner presents paid invoice
    await swap.submitPaidInvoice(encoded, 200, 2, invoice.sig, noteAmount, epoch, cheque.sig)

    await time.increase(2 * epoch)

    await shouldFail.reverting(swap.cashNote(encoded, noteAmount))

    let { logs } = await swap.cashCheque(carol)

    matchLogs(logs, [
      { event: 'ChequeCashed', args: { beneficiary: carol, serial: 3, amount: 700 } }
    ])
  })

})
