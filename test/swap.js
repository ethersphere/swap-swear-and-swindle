const Swap = artifacts.require("./Swap.sol");
require('chai')
    .use(require('chai-as-promised'))
    .use(require('chai-bignumber')(web3.BigNumber))
    .should();

const { getBalance, getTime, increaseTime, expectFail, matchLogs, sign } = require('./testutils')

const epoch = 24 * 3600

contract('swap', function(accounts) {
  const [owner, bob, alice] = accounts

  async function submitCheque(signer, beneficiary, serial, amount) {
    const swap = await Swap.deployed();
    const hash = await swap.chequeHash(beneficiary, serial, amount);
    const { r, s, v } = sign(signer, hash);

    return swap.submitCheque(beneficiary, serial, amount, r, s, v);
  }

  const firstDeposit = 1000;
  const firstCheque = 600;
  const secondCheque = firstCheque + 200;
  const thirdCheque = secondCheque + 300;
  const refill = 500;
  const hardDepositBob1 = 200;
  const fourthCheque = thirdCheque + 100;
  const aliceCheque = 250;
  const hardDepositBob2 = 50;

  it('should accept deposits', async() => {
    const swap = await Swap.deployed();
    const value = firstDeposit;

    const { logs } = await swap.send(value, { from: owner });

    matchLogs(logs, [
      { event: 'Deposit', args: { depositor: owner, amount: value }}
    ]);

    (await getBalance(swap.address)).should.bignumber.equal(value);
  })

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

    const [storedSerial, storedAmount,, timeout] = await swap.infos(bob);

    storedSerial.should.bignumber.equal(serial)
    storedAmount.should.bignumber.equal(amount)
    timeout.should.bignumber.gte(getTime() + 1 * epoch) /* TODO: make sure there are no edge cases that could cause this to fail */
  })

  it('should not allow cheque payout before timeout', async() => {
    const swap = await Swap.deployed();

    await expectFail(swap.cashCheque(bob));
  })

  it('should allow cheque payout after timeout (no hard deposit)', async() => {
    const swap = await Swap.deployed();

    await increaseTime(1 * epoch);

    let beneficiaryExpectedBalance = (await getBalance(bob)).plus(firstCheque);

    let { logs } = await swap.cashCheque(bob);

    matchLogs(logs, [
      { event: 'ChequeCashed', args: { beneficiary: bob, serial: 1, amount: firstCheque }}
    ]);

    (await getBalance(bob)).should.be.bignumber.equal(beneficiaryExpectedBalance);

    const [,, paidOut,] = await swap.infos(bob);

    paidOut.should.bignumber.equal(firstCheque);
  })

  it('should not allow cheque payout if there is nothing to pay out', async() => {
    const swap = await Swap.deployed();

    await expectFail(swap.cashCheque(bob));
  })

  it('should not allow valid cheque if signed by owner and amount is not higher', async() => {
    const swap = await Swap.deployed();

    await expectFail(submitCheque(owner, bob, 2, firstCheque));
  })

  it('should accept valid cheque with higher amount', async() => {
    const swap = await Swap.deployed();

    const serial = 2;
    const amount = secondCheque;

    const { logs } = await submitCheque(owner, bob, serial, amount);

    matchLogs(logs, [
      { event: 'ChequeSubmitted', args: { amount, beneficiary: bob, serial } }
    ])

    const [storedSerial, storedAmount,, timeout] = await swap.infos(bob);

    storedSerial.should.bignumber.equal(serial)
    storedAmount.should.bignumber.equal(amount)
    timeout.should.bignumber.gte(getTime() + 1 * epoch) /* TODO: make sure there are no edge cases that could cause this to fail */
  })

  it('should not allow cheque payout before increased timeout', async() => {
    const swap = await Swap.deployed();

    await expectFail(swap.cashCheque(bob));
  })

  it('should allow cheque payout after timeout (with difference, no hard deposit)', async() => {
    const swap = await Swap.deployed();

    await increaseTime(1 * epoch);

    let beneficiaryExpectedBalance = (await getBalance(bob)).plus(secondCheque).minus(firstCheque);

    let { logs } = await swap.cashCheque(bob);

    matchLogs(logs, [
      { event: 'ChequeCashed', args: { beneficiary: bob, serial: 2, amount: secondCheque - firstCheque }}
    ]);

    (await getBalance(bob)).should.be.bignumber.equal(beneficiaryExpectedBalance);

    const [,, paidOut,] = await swap.infos(bob);

    paidOut.should.bignumber.equal(secondCheque);
  })

  it('should accept a valid check with lower value if signed by beneficiary', async () => {
    const swap = await Swap.deployed();

    const serial = 3;
    const amount = firstCheque;

    const { logs } = await submitCheque(bob, bob, serial, amount);

    matchLogs(logs, [
      { event: 'ChequeSubmitted', args: { beneficiary: bob, serial, amount }}
    ]);

    const [storedSerial, storedAmount,,] = await swap.infos(bob);

    storedSerial.should.bignumber.equal(serial)
    storedAmount.should.bignumber.equal(firstCheque)
  })

  it('should allow parital payment for a bouncing check', async () => {
    const swap = await Swap.deployed();

    const serial = 4;
    const amount = thirdCheque;

    await submitCheque(owner, bob, serial, amount);
    await increaseTime(1 * epoch);

    const paid = firstDeposit - secondCheque;
    const bounced = thirdCheque - firstDeposit;

    let beneficiaryExpectedBalance = (await getBalance(bob)).plus(paid);

    const { logs } = await swap.cashCheque(bob);

    matchLogs(logs, [
      { event: 'ChequeBounced', args: { paid, bounced, serial, beneficiary: bob } }
    ]);

    (await getBalance(bob)).should.be.bignumber.equal(beneficiaryExpectedBalance);
  })

  it('should allow cheque to clear fully after refill', async () => {
    const swap = await Swap.deployed();

    await swap.send(refill);
    const bounced = thirdCheque - firstDeposit;

    let beneficiaryExpectedBalance = (await getBalance(bob)).plus(bounced);

    const { logs } = await swap.cashCheque(bob);

    matchLogs(logs, [
      { event: 'ChequeCashed', args: { amount: bounced, serial: 4, beneficiary: bob } }
    ]);

    (await getBalance(bob)).should.be.bignumber.equal(beneficiaryExpectedBalance);
  })

  it('should allow hard deposits if they do not exceed the global deposit', async() => {
    const swap = await Swap.deployed();

    const { logs } = await swap.increaseHardDeposit(bob, hardDepositBob1);

    matchLogs(logs, [
      { event: 'HardDepositChanged', args: { beneficiary: bob, amount: hardDepositBob1 } }
    ])

    const [deposit] = await swap.hardDeposits(bob);

    deposit.should.bignumber.equal(hardDepositBob1);
    (await swap.totalDeposit()).should.bignumber.equal(hardDepositBob1);
  })

  it('should not allow hard deposits if they exceed the global deposit', async() => {
    const swap = await Swap.deployed();

    await expectFail(swap.increaseHardDeposit(bob, await getBalance(swap.address) - hardDepositBob1 + 1));
  })

  it('should use the hard deposit on valid cheque', async() => {
    const swap = await Swap.deployed();
    const serial = 5;

    const amount = fourthCheque;
    const diff = fourthCheque - thirdCheque;

    await submitCheque(owner, bob, serial, amount);
    await increaseTime(1 * epoch);

    let beneficiaryExpectedBalance = (await getBalance(bob)).plus(diff);
    const expectedHardDeposit = (await swap.hardDeposits(bob))[0].minus(diff);

    const { logs } = await swap.cashCheque(bob);

    matchLogs(logs, [
      { event: 'ChequeCashed', args: { amount: diff, serial, beneficiary: bob } }
    ]);

    const [hardDeposit] = await swap.hardDeposits(bob);

    hardDeposit.should.bignumber.equal(expectedHardDeposit);
    (await getBalance(bob)).should.be.bignumber.equal(beneficiaryExpectedBalance);
  })

  /* NOTE: at this point there are 300 wei in the contract with 100 wei locked for bob */
  it('should not spend ether locked away by hard deposit of another address', async() => {
    const swap = await Swap.deployed();

    await submitCheque(owner, alice, 1, aliceCheque);
    await increaseTime(1 * epoch);

    /* sanity check - make sure there is actually enough left so we know we're testing the right thing */
    (await getBalance(swap.address)).should.be.bignumber.gte(aliceCheque);
    const expectedBalanceAlice = (await getBalance(alice)).plus(200);

    const { logs } = await swap.cashCheque(alice);

    matchLogs(logs, [
      { event: 'ChequeBounced', args: { beneficiary: alice, paid: 200, bounced: 50, serial: 1 } }
    ]);

    (await getBalance(alice)).should.bignumber.equal(expectedBalanceAlice);
  })

  it('should not allow an instant decrease for hard deposits', async() => {
    const swap = await Swap.deployed();

    await expectFail(swap.decreaseHardDeposit(bob));
  })

  it('should allow to prepare a decrease for hard deposits', async() => {
    const swap = await Swap.deployed();

    const { logs } = await swap.prepareDecreaseHardDeposit(bob, hardDepositBob2);

    matchLogs(logs, [
      { event: 'HardDepositDecreasePrepared', args: { beneficiary: bob, amount: hardDepositBob2 } }
    ]);

    const [, timeout, next] = await swap.hardDeposits(bob);

    next.should.bignumber.equal(hardDepositBob2);
    timeout.should.bignumber.gte(getTime() + 2 * epoch);
  })

  it('should not allow to decrease hard deposit before the timeout', async() => {
    const swap = await Swap.deployed();

    await expectFail(swap.decreaseHardDeposit(bob));
  })

  it('should allow to decrease hard deposit after the timeout', async() => {
    const swap = await Swap.deployed();

    await increaseTime(2 * epoch);
    const { logs } = await swap.decreaseHardDeposit(bob);

    matchLogs(logs, [
      { event: 'HardDepositChanged', args: { beneficiary: bob, amount: hardDepositBob2 } }
    ]);

    const [deposit] = await swap.hardDeposits(bob);
    deposit.should.bignumber.equal(hardDepositBob2);
    (await swap.totalDeposit()).should.bignumber.equal(hardDepositBob2);
  })

  it('should not allow to do the same decrease twice', async() => {
    const swap = await Swap.deployed();

    await expectFail(swap.decreaseHardDeposit(bob));
  })

  it('should have enough liquid balance for the rest of the other accounts cheque now', async() => {
    const swap = await Swap.deployed();

    const expectedBalanceAlice = (await getBalance(alice)).plus(50);

    const { logs } = await swap.cashCheque(alice);

    matchLogs(logs, [
      { event: 'ChequeCashed', args: { beneficiary: alice, serial: 1, amount: 50 } }
    ]);

    (await getBalance(alice)).should.be.bignumber.equal(expectedBalanceAlice);
  })

})
