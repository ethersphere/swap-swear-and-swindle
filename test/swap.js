const Swap = artifacts.require("./Swap.sol");
const OracleWitness = artifacts.require("./OracleWitness.sol");

require('chai')
    .use(require('chai-as-promised'))
    .use(require('bn-chai')(web3.utils.BN))
    .should();

const { getBalance, getTime, increaseTime, expectFail, matchLogs, matchStruct, sign, nulladdress, computeCost } = require('./testutils')
const { signCheque, signNote, signInvoice } = require('./swutils')

const epoch = 24 * 3600

contract('swap', function(accounts) {
  const [owner, bob, alice, carol] = accounts

  async function submitCheque(signer, beneficiary, serial, amount) {
    const swap = await Swap.deployed();
    const { sig } = await signCheque(signer, beneficiary, serial, amount);
    return swap.submitCheque(beneficiary, serial, amount, sig, { from: beneficiary });
  }

  const firstDeposit = 1000;

  it('should accept deposits', async() => {
    const swap = await Swap.deployed();
    const value = firstDeposit;

    const { logs } = await swap.send(value, { from: owner });

    matchLogs(logs, [
      { event: 'Deposit', args: { depositor: owner, amount: value }}
    ]);

    (await getBalance(swap.address)).should.eq.BN(value);
  })

  const firstCheque = 600;

  it('should not accept a cheque with serial 0', async() => {
    const swap = await Swap.deployed();
    await expectFail(submitCheque(owner, bob, 0, firstCheque));
  })

  it('should accept valid cheque (increasing value)', async() => {
    const swap = await Swap.deployed();

    const serial = 1;
    const amount = firstCheque;

    const { logs } = await submitCheque(owner, bob, serial, amount);

    matchLogs(logs, [
      { event: 'ChequeSubmitted', args: { amount, beneficiary: bob, serial } }
    ])

    const chequeInfo = await swap.cheques(bob);

    chequeInfo.serial.should.eq.BN(serial)
    chequeInfo.amount.should.eq.BN(amount)
    chequeInfo.timeout.should.gte.BN(await getTime() + 1 * epoch - 1)
  })

  it('should not allow cheque payout before timeout', async() => {
    const swap = await Swap.deployed();

    await expectFail(swap.cashCheque(bob));
  })

  it('should allow cheque payout after timeout (no hard deposit)', async() => {
    const swap = await Swap.deployed();

    await increaseTime(1 * epoch);

    let beneficiaryExpectedBalance = (await getBalance(bob)).addn(firstCheque);

    let { logs } = await swap.cashCheque(bob);

    matchLogs(logs, [
      { event: 'ChequeCashed', args: { beneficiary: bob, serial: 1, amount: firstCheque }}
    ]);

    (await getBalance(bob)).should.eq.BN(beneficiaryExpectedBalance);

    const chequeInfo = await swap.cheques(bob);

    chequeInfo.paidOut.should.eq.BN(firstCheque);
  })

  it('should not allow cheque payout if there is nothing to pay out', async() => {
    const swap = await Swap.deployed();
    await expectFail(swap.cashCheque(bob));
  })

  it('should not allow valid cheque if signed by owner and amount is not higher', async() => {
    const swap = await Swap.deployed();
    await expectFail(submitCheque(owner, bob, 2, firstCheque));
  })

  const secondCheque = firstCheque + 200;

  it('should accept valid cheque with higher amount', async() => {
    const swap = await Swap.deployed();

    const serial = 2;
    const amount = secondCheque;

    const { logs } = await submitCheque(owner, bob, serial, amount);

    matchLogs(logs, [
      { event: 'ChequeSubmitted', args: { amount, beneficiary: bob, serial } }
    ])

    const chequeInfo = await swap.cheques(bob);

    chequeInfo.serial.should.eq.BN(serial)
    chequeInfo.amount.should.eq.BN(amount)
    chequeInfo.timeout.should.gte.BN(getTime() + 1 * epoch - 1)
  })

  it('should not allow cheque payout before increased timeout', async() => {
    const swap = await Swap.deployed();
    await expectFail(swap.cashCheque(bob));
  })

  it('should allow cheque payout after timeout (with difference, no hard deposit)', async() => {
    const swap = await Swap.deployed();
    await increaseTime(1 * epoch);

    let beneficiaryExpectedBalance = (await getBalance(bob)).addn(secondCheque).subn(firstCheque);

    let { logs } = await swap.cashCheque(bob);

    matchLogs(logs, [
      { event: 'ChequeCashed', args: { beneficiary: bob, serial: 2, amount: secondCheque - firstCheque }}
    ]);

    (await getBalance(bob)).should.eq.BN(beneficiaryExpectedBalance);

    const chequeInfo = await swap.cheques(bob);

    chequeInfo.paidOut.should.eq.BN(secondCheque);
  })

  it('should accept a valid check with lower value if signed by beneficiary', async () => {
    const swap = await Swap.deployed();

    const serial = 3;
    const amount = firstCheque;

    const { sig: sigOwner } = await signCheque(owner, bob, serial, amount);
    const { sig: sigBob } = await signCheque(bob, bob, serial, amount);

    const { logs } = await swap.submitChequeLower(bob, serial, amount, sigOwner, sigBob);

    matchLogs(logs, [
      { event: 'ChequeSubmitted', args: { beneficiary: bob, serial, amount }}
    ]);

    const chequeInfo = await swap.cheques(bob);

    chequeInfo.serial.should.eq.BN(serial)
    chequeInfo.amount.should.eq.BN(firstCheque)
  })

  const thirdCheque = secondCheque + 300;

  it('should allow parital payment for a bouncing check', async () => {
    const swap = await Swap.deployed();

    const serial = 4;
    const amount = thirdCheque;

    await submitCheque(owner, bob, serial, amount);
    await increaseTime(1 * epoch);

    const paid = firstDeposit - secondCheque;
    const bounced = thirdCheque - firstDeposit;

    let beneficiaryExpectedBalance = (await getBalance(bob)).addn(paid);

    const { logs } = await swap.cashCheque(bob);

    matchLogs(logs, [
      { event: 'ChequeBounced', args: { paid, bounced, serial, beneficiary: bob } }
    ]);

    (await getBalance(bob)).should.eq.BN(beneficiaryExpectedBalance);
  })

  const refill = 500;

  it('should allow cheque to clear fully after refill', async () => {
    const swap = await Swap.deployed();

    await swap.send(refill);
    const bounced = thirdCheque - firstDeposit;

    let beneficiaryExpectedBalance = (await getBalance(bob)).addn(bounced);

    const { logs } = await swap.cashCheque(bob);

    matchLogs(logs, [
      { event: 'ChequeCashed', args: { amount: bounced, serial: 4, beneficiary: bob } }
    ]);

    (await getBalance(bob)).should.eq.BN(beneficiaryExpectedBalance);
  })

  const hardDepositBob1 = 200;

  it('should allow hard deposits if they do not exceed the global deposit', async() => {
    const swap = await Swap.deployed();

    const { logs } = await swap.increaseHardDeposit(bob, hardDepositBob1);

    matchLogs(logs, [
      { event: 'HardDepositChanged', args: { beneficiary: bob, amount: hardDepositBob1 } }
    ])

    const hardDeposit = await swap.hardDeposits(bob);
    hardDeposit.amount.should.eq.BN(hardDepositBob1);
    (await swap.totalDeposit()).should.eq.BN(hardDepositBob1);
  })

  it('should not allow hard deposits if they exceed the global deposit', async() => {
    const swap = await Swap.deployed();
    await expectFail(swap.increaseHardDeposit(bob, (await getBalance(swap.address)).subn(hardDepositBob1).addn(1)));
  })

  const fourthCheque = thirdCheque + 100;

  it('should use the hard deposit on valid cheque', async() => {
    const swap = await Swap.deployed();
    const serial = 5;

    const amount = fourthCheque;
    const diff = fourthCheque - thirdCheque;

    await submitCheque(owner, bob, serial, amount);
    await increaseTime(1 * epoch);

    let beneficiaryExpectedBalance = (await getBalance(bob)).addn(diff);

    const expectedHardDeposit = (await swap.hardDeposits(bob)).amount.subn(diff);

    const { logs } = await swap.cashCheque(bob);

    matchLogs(logs, [
      { event: 'ChequeCashed', args: { amount: diff, serial, beneficiary: bob } }
    ]);

    (await swap.hardDeposits(bob)).amount.should.eq.BN(expectedHardDeposit);
    (await getBalance(bob)).should.eq.BN(beneficiaryExpectedBalance);
  })

  const aliceCheque = 250;

  /* NOTE: at this point there are 300 wei in the contract with 100 wei locked for bob */
  it('should not spend ether locked away by hard deposit of another address', async() => {
    const swap = await Swap.deployed();

    await submitCheque(owner, alice, 1, aliceCheque);
    await increaseTime(1 * epoch);

    let swapBalance = await getBalance(swap.address)

    /* sanity check - make sure there is actually enough left so we know we're testing the right thing */
    swapBalance.should.gte.BN(aliceCheque);
    let paid = swapBalance.sub((await swap.hardDeposits(bob)).amount)
    let bounced = web3.utils.toBN(aliceCheque).sub(paid)


    const expectedBalanceAlice = (await getBalance(alice)).add(paid);

    const { logs } = await swap.cashCheque(alice);

    matchLogs(logs, [
      { event: 'ChequeBounced', args: { beneficiary: alice, paid, bounced, serial: 1 } }
    ]);

    (await getBalance(alice)).should.eq.BN(expectedBalanceAlice);
  })

  it('should not allow an instant decrease for hard deposits', async() => {
    const swap = await Swap.deployed();
    await expectFail(swap.decreaseHardDeposit(bob));
  })

  const hardDepositBobDecrease = 75;

  it('should allow to prepare a decrease for hard deposits', async() => {
    const swap = await Swap.deployed();

    const { logs } = await swap.prepareDecreaseHardDeposit(bob, hardDepositBobDecrease);

    matchLogs(logs, [
      { event: 'HardDepositDecreasePrepared', args: { beneficiary: bob, diff: hardDepositBobDecrease } }
    ]);

    const hardDeposit = await swap.hardDeposits(bob);

    hardDeposit.diff.should.eq.BN(hardDepositBobDecrease);
    hardDeposit.timeout.should.gte.BN(await getTime() + 2 * epoch - 1);
  })

  it('should not allow to decrease hard deposit before the timeout', async() => {
    const swap = await Swap.deployed();
    await expectFail(swap.decreaseHardDeposit(bob));
  })

  it('should allow to decrease hard deposit after the timeout', async() => {
    const swap = await Swap.deployed();

    let expectedHardDeposit = (await swap.hardDeposits(bob)).amount.subn(hardDepositBobDecrease);

    await increaseTime(2 * epoch);
    const { logs } = await swap.decreaseHardDeposit(bob);

    matchLogs(logs, [
      { event: 'HardDepositChanged', args: { beneficiary: bob, amount: expectedHardDeposit } }
    ]);

    (await swap.hardDeposits(bob)).amount.should.eq.BN(expectedHardDeposit);
    (await swap.totalDeposit()).should.eq.BN(expectedHardDeposit);
  })

  it('should not allow to do the same decrease twice', async() => {
    const swap = await Swap.deployed();
    await expectFail(swap.decreaseHardDeposit(bob));
  })

  it('should have enough liquid balance for the rest of the other accounts cheque now', async() => {
    const swap = await Swap.deployed();

    const amount = web3.utils.toBN(aliceCheque).sub((await swap.cheques(alice)).paidOut)

    const expectedBalanceAlice = (await getBalance(alice)).add(amount);

    const { logs } = await swap.cashCheque(alice);

    matchLogs(logs, [
      { event: 'ChequeCashed', args: { beneficiary: alice, serial: 1, amount } }
    ]);

    (await getBalance(alice)).should.be.eq.BN(expectedBalanceAlice);
  })

  let carolBond = 1000
  let carolBondValidTimeout = 3 * epoch

  it('should accept a valid note (bond)', async() => {
    const swap = await Swap.deployed();

    await swap.send(carolBond);

    let validity = await getTime() + carolBondValidTimeout

    let { sig, hash } = await signNote(owner, carol, 1, carolBond, nulladdress, validity, 0, "0x")

    await increaseTime(4 * epoch)
    let encoded = await swap.encodeNote(swap.address, carol, 1, carolBond, nulladdress, validity, 0, "0x")
    await swap.submitNote(encoded, sig, { from: carol });

    const { paidOut, timeout } = await swap.notes(hash)

    paidOut.should.eq.BN(0)
    timeout.should.gte.BN(await getTime() + 1 * epoch - 1)

    await increaseTime(1 * epoch)

    let expectedBalanceCarol = (await getBalance(carol)).addn(carolBond)

    let { receipt } = await swap.cashNote(encoded, carolBond, { from: carol });

    expectedBalanceCarol = expectedBalanceCarol.sub(await computeCost(receipt));

    (await getBalance(carol)).should.eq.BN(expectedBalanceCarol)

    /* already fully cashed out */
    await expectFail(swap.cashNote(encoded, carolBond, { from: carol }));
  })

  it('should accept a valid note (conditional bond)', async() => {
    const swap = await Swap.deployed();
    const oracle = await OracleWitness.deployed();

    await swap.send(carolBond);

    let bondTimeout = await getTime() + carolBondValidTimeout

    let { sig, hash } = await signNote(owner, carol, 1, carolBond, oracle.address, 0, bondTimeout, "0x")

    await oracle.testify(hash, 1)

    let encoded = await swap.encodeNote(swap.address, carol, 1, carolBond, oracle.address, 0, bondTimeout, "0x")
    await swap.submitNote(encoded, sig, { from: carol });

    const { paidOut, timeout } = await swap.notes(hash)

    paidOut.should.eq.BN(0)
    timeout.should.gte.BN(await getTime() + 1 * epoch - 1)

    /* cashout too soon */
    await expectFail(swap.cashNote(encoded, carolBond, { from: carol }))

    await increaseTime(1 * epoch)

    await oracle.testify(hash, 0)

    /* oracle says no */
    await expectFail(swap.cashNote(encoded, carolBond, { from: carol }))

    await oracle.testify(hash, 1)

    /* partial payment */
    let expectedBalanceCarol = (await getBalance(carol)).addn(carolBond / 4)
    var { receipt } = await swap.cashNote(encoded, carolBond / 4, { from: carol });
    expectedBalanceCarol = expectedBalanceCarol.sub(await computeCost(receipt));
    (await getBalance(carol)).should.eq.BN(expectedBalanceCarol)

    /* partial payment */
    expectedBalanceCarol = (await getBalance(carol)).addn(carolBond / 4)
    var { receipt } = await swap.cashNote(encoded, carolBond / 4, { from: carol });
    expectedBalanceCarol = expectedBalanceCarol.sub(await computeCost(receipt));
    (await getBalance(carol)).should.eq.BN(expectedBalanceCarol)

    await increaseTime(carolBondValidTimeout)

    /* too late for the rest */
    await expectFail(swap.cashNote(encoded, carolBond / 4, { from: carol }))
  })

  it('should allow to submit paid invoices', async() => {
    const swap = await Swap.deployed();

    await swap.send(carolBond + 200);

    await submitCheque(owner, carol, 1, 100)

    /* completely offchain cheque of 100 */

    /* owner issues note */
    let note = await signNote(owner, carol, 1, carolBond, nulladdress, 0, 0, "0x")

    /* carol issues invoice */
    let invoice = await signInvoice(carol, note.hash, 200, 2)

    /* owner issues cheque for invoice */
    let cheque = await signCheque(owner, carol, 3, carolBond + 200)

    let encoded = await swap.encodeNote(swap.address, carol, 1, carolBond, nulladdress, 0, 0, "0x")
    /* carol submits note anyway */
    await swap.submitNote(encoded, note.sig, { from: carol });

    /* owner presents paid invoice */
    await swap.submitPaidInvoice(encoded, 200, 2, invoice.sig, carolBond, cheque.sig)

    await increaseTime(2 * epoch)

    await expectFail(swap.cashNote(encoded, carolBond))

    let { logs } = await swap.cashCheque(carol)

    matchLogs(logs, [
      { event: 'ChequeCashed', args: { beneficiary: carol, serial: 3, amount: 1200 } }
    ])
  })

})
