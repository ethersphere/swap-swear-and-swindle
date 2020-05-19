const {
  BN,
  balance,
  time,
  expectEvent,
  expectRevert
} = require("@openzeppelin/test-helpers");

const ERC20SimpleSwap = artifacts.require('ERC20SimpleSwap')
const ERC20PresetMinterPauser = artifacts.require("ERC20PresetMinterPauser")

const { signCheque, signCashOut, signCustomDecreaseTimeout } = require("./swutils");
const { expect } = require('chai');

function shouldDeploy(issuer, defaultHardDepositTimeout, from, value) {
  beforeEach(async function() {
    this.ERC20PresetMinterPauser = await ERC20PresetMinterPauser.new("TestToken", "TEST", {from: issuer})
    await this.ERC20PresetMinterPauser.mint(issuer, 1000000000, {from: issuer});
    this.ERC20SimpleSwap = await ERC20SimpleSwap.new(issuer, this.ERC20PresetMinterPauser.address, defaultHardDepositTimeout, {from: from})   
    if(value != 0) {
      await this.ERC20PresetMinterPauser.transfer(this.ERC20SimpleSwap.address, value, {from: issuer});
    }
    this.postconditions = {
      issuer: await this.ERC20SimpleSwap.issuer(),
      defaultHardDepositTimeout: await this.ERC20SimpleSwap.defaultHardDepositTimeout()
    }
  })
  it('should set the issuer', function() {
    expect(this.postconditions.issuer).to.be.equal(issuer)
  })
  it('should set the defaultHardDepositTimeout', function() {
    expect(this.postconditions.defaultHardDepositTimeout).bignumber.to.be.equal(defaultHardDepositTimeout)
  })
}
function shouldReturnDefaultHardDepositTimeout(expected) {
  it('should return the expected defaultHardDepositTimeout', async function() {
    expect(await this.ERC20SimpleSwap.defaultHardDepositTimeout()).bignumber.to.be.equal(expected)
  })
}

function shouldReturnPaidOut(beneficiary, expectedAmount) {
  beforeEach(async function() {
    this.paidOut = await this.ERC20SimpleSwap.paidOut(beneficiary)
  })
  it('should return the expected amount', function() {
    expect(expectedAmount).bignumber.to.be.equal(this.paidOut)
  })
}

function shouldReturnTotalPaidOut(expectedAmount) {
  beforeEach(async function() {
    this.totalPaidOut = await this.ERC20SimpleSwap.totalPaidOut()
  })
  it('should return the expected amount', function() {
    expect(expectedAmount).bignumber.to.be.equal(this.totalPaidOut)
  })
}

function shouldReturnHardDeposits(beneficiary, expectedAmount, expectedDecreaseAmount,  expectedDecreaseTimeout, expectedCanBeDecreasedAt) {
  beforeEach(async function() {
    // If we expect this not to be the default value, we have to set the value here, as it depends on the most current time
    if(!expectedCanBeDecreasedAt.eq(new BN(0))) {
      this.expectedCanBeDecreasedAt = (await time.latest()).add(await this.ERC20SimpleSwap.defaultHardDepositTimeout())
    } else {
      this.expectedCanBeDecreasedAt = expectedCanBeDecreasedAt
    }
    this.exptectedCanBeDecreasedAt = (await time.latest()).add(await this.ERC20SimpleSwap.defaultHardDepositTimeout())
    this.hardDeposits = await this.ERC20SimpleSwap.hardDeposits(beneficiary)
  })
  it('should return the expected amount', function() {
    expect(expectedAmount).bignumber.to.be.equal(this.hardDeposits.amount)
  })
  it('should return the expected decreaseAmount', function() {
    expect(expectedDecreaseAmount).bignumber.to.be.equal(this.hardDeposits.decreaseAmount)
  })
  it('should return the expected timeout', function() {
    expect(expectedDecreaseTimeout).bignumber.to.be.equal(this.hardDeposits.timeout)
  })
  it('should return the exptected canBeDecreasedAt', function() {
    expect(this.expectedCanBeDecreasedAt.toNumber()).to.be.closeTo(this.hardDeposits.canBeDecreasedAt.toNumber(), 5)
  })
}

function shouldReturnTotalHardDeposit(expectedTotalHardDeposit) {
  beforeEach(async function() {
    this.totalHardDeposit = await this.ERC20SimpleSwap.totalHardDeposit()
  })

  it('should return the expectedTotalHardDeposit', function() {
    expect(expectedTotalHardDeposit).bignumber.to.be.equal(this.totalHardDeposit)
  })
}

function shouldReturnIssuer(expectedIssuer) {
  it('should return the expected issuer', async function() {
    expect(await this.ERC20SimpleSwap.issuer()).to.be.equal(expectedIssuer)
  })

}

function shouldReturnLiquidBalance(expectedLiquidBalance) {
  it('should return the expected liquidBalance', async function() {
    expect(await this.ERC20SimpleSwap.liquidBalance()).bignumber.to.equal(expectedLiquidBalance)
  })
}

function shouldReturnLiquidBalanceFor(beneficiary, expectedLiquidBalanceFor) {
  it('should return the expected liquidBalance', async function() {
    expect(await this.ERC20SimpleSwap.liquidBalanceFor(beneficiary)).bignumber.to.equal(expectedLiquidBalanceFor)
  })
}


function cashChequeInternal(beneficiary, recipient, cumulativePayout, callerPayout, from) {

  beforeEach(async function() {
    let requestPayout = cumulativePayout.sub(this.preconditions.paidOut)
    //if the requested payout is less than the liquidBalance available for beneficiary
    if(requestPayout.lt(this.preconditions.liquidBalanceFor)) {
      // full amount requested can be paid out
      this.totalPayout = requestPayout
    } else {
      // partial amount requested can be paid out (the liquid balance available to the node)
      this.totalPayout = this.preconditions.liquidBalanceFor
    }
    this.totalPaidOut = this.preconditions.totalPaidOut + this.totalPayout
  })
  
  it('should update the totalHardDeposit and hardDepositFor ', function() {
    let expectedDecreaseHardDeposit
    // if the hardDeposits can cover the totalPayout
    if(this.totalPayout.lt(this.preconditions.hardDepositFor.amount)) {
      // hardDeposit decreases by totalPayout
      expectedDecreaseHardDeposit = this.totalPayout
    } else {
      // hardDeposit decreases by the full amount (and rest is from global liquid balance)
      expectedDecreaseHardDeposit = this.preconditions.hardDepositFor.amount
    }
    // totalHarddeposit
    expect(this.postconditions.totalHardDeposit).bignumber.to.be.equal(this.preconditions.totalHardDeposit.sub(expectedDecreaseHardDeposit))    
    // hardDepositFor
    expect(this.postconditions.hardDepositFor.amount).bignumber.to.be.equal(this.preconditions.hardDepositFor.amount.sub(expectedDecreaseHardDeposit))    
  })
  
  it('should update paidOut', async function() {
    expect(this.postconditions.paidOut).bignumber.to.be.equal(this.preconditions.paidOut.add(this.totalPayout))
  })

  it('should update totalPaidOut', async function() {
    expect(this.postconditions.totalPaidOut).bignumber.to.be.equal(this.preconditions.paidOut.add(this.totalPayout))
  })

  it('should transfer the correct amount to the recipient', async function() {
    expect(this.postconditions.recipientBalance).bignumber.to.be.equal(this.preconditions.recipientBalance.add(this.totalPayout.sub(callerPayout)))
  })
  it('should transfer the correct amount to the caller', async function() {
    let expectedAmountCaller
    if(recipient == from) {
      expectedAmountCaller = this.totalPayout
    } else {
      expectedAmountCaller = callerPayout
    }
    expect(this.postconditions.callerBalance).bignumber.to.be.equal(this.preconditions.callerBalance.add(expectedAmountCaller))
  })
  
  it('should emit a ChequeCashed event', function() {
    expectEvent.inLogs(this.logs, "ChequeCashed", {
      beneficiary,
      recipient: recipient,
      caller: from,
      totalPayout: this.totalPayout,
      cumulativePayout,
      callerPayout,
    })
  })
  it('should only emit a chequeBounced event when insufficient funds', function() {
    if(this.totalPayout.lt(cumulativePayout.sub(this.preconditions.paidOut))) {
      expectEvent.inLogs(this.logs, "ChequeBounced", {})
    } else {
      const events = this.logs.filter(e => e.event === 'ChequeBounced');
      expect(events.length > 0).to.equal(false, `There is a ChequeBounced event`)
    }
  })

  it('should only set the bounced field when insufficient funds', function() {
    if(this.totalPayout.lt(cumulativePayout.sub(this.preconditions.paidOut))) {
      expect(this.postconditions.bounced).to.be.true
    } else {
      expect(this.postconditions.bounced).to.be.false
    }
  })
}

function shouldCashChequeBeneficiary(recipient, cumulativePayout, signee, from) {
  beforeEach(async function() {
    this.preconditions = {
      callerBalance: await this.ERC20PresetMinterPauser.balanceOf(from),
      recipientBalance: await this.ERC20PresetMinterPauser.balanceOf(recipient),
      totalHardDeposit: await this.ERC20SimpleSwap.totalHardDeposit(),
      hardDepositFor: await this.ERC20SimpleSwap.hardDeposits(from),
      liquidBalance: await this.ERC20SimpleSwap.liquidBalance(),
      liquidBalanceFor: await this.ERC20SimpleSwap.liquidBalanceFor(from),
      chequebookBalance: await this.ERC20SimpleSwap.balance(),
      paidOut: await this.ERC20SimpleSwap.paidOut(from),
      totalPaidOut: await this.ERC20SimpleSwap.totalPaidOut()
    }

    const issuerSig = await signCheque(this.ERC20SimpleSwap, from, cumulativePayout, signee)
  
    const { logs, receipt } = await this.ERC20SimpleSwap.cashChequeBeneficiary(recipient, cumulativePayout, issuerSig, {from: from})
    this.logs = logs
    this.receipt = receipt
  
    this.postconditions = {
      callerBalance: await this.ERC20PresetMinterPauser.balanceOf(from),
      recipientBalance: await this.ERC20PresetMinterPauser.balanceOf(recipient),
      totalHardDeposit: await this.ERC20SimpleSwap.totalHardDeposit(),
      hardDepositFor: await this.ERC20SimpleSwap.hardDeposits(from),
      liquidBalance: await this.ERC20SimpleSwap.liquidBalance(),
      liquidBalanceFor: await this.ERC20SimpleSwap.liquidBalanceFor(from),
      chequebookBalance: await this.ERC20SimpleSwap.balance(),
      paidOut: await this.ERC20SimpleSwap.paidOut(from),
      totalPaidOut: await this.ERC20SimpleSwap.totalPaidOut(),
      bounced: await this.ERC20SimpleSwap.bounced()
    }
  })
  cashChequeInternal(from, recipient, cumulativePayout, new BN(0), from)
}
function shouldNotCashChequeBeneficiary(recipient, toSubmitCumulativePayout, toSignCumulativePayout, signee, from, value, revertMessage) {
  beforeEach(async function() {
    this.issuerSig = await signCheque(this.ERC20SimpleSwap, from, toSignCumulativePayout, signee)
  })
  it('reverts', async function() {
    await expectRevert(this.ERC20SimpleSwap.cashChequeBeneficiary(
      recipient,
      toSubmitCumulativePayout,
      this.issuerSig,
     {from: from, value: value}), 
     revertMessage
    )
  })
}
function shouldCashCheque(beneficiary, recipient, cumulativePayout, callerPayout, from, beneficiarySignee, issuerSignee) {
  beforeEach(async function() {
    const beneficiarySig = await signCashOut(this.ERC20SimpleSwap, from, cumulativePayout, recipient, callerPayout, beneficiarySignee)
    const issuerSig = await signCheque(this.ERC20SimpleSwap, beneficiary, cumulativePayout, issuerSignee)
    this.preconditions = {
      callerBalance: await this.ERC20PresetMinterPauser.balanceOf(from),
      recipientBalance: await this.ERC20PresetMinterPauser.balanceOf(recipient),
      totalHardDeposit: await this.ERC20SimpleSwap.totalHardDeposit(),
      hardDepositFor: await this.ERC20SimpleSwap.hardDeposits(beneficiary),
      liquidBalance: await this.ERC20SimpleSwap.liquidBalance(),
      liquidBalanceFor: await this.ERC20SimpleSwap.liquidBalanceFor(beneficiary),
      chequebookBalance: await this.ERC20SimpleSwap.balance(),
      paidOut: await this.ERC20SimpleSwap.paidOut(beneficiary),
      totalPaidOut: await this.ERC20SimpleSwap.totalPaidOut()
    }
    const { logs, receipt } = await this.ERC20SimpleSwap.cashCheque(beneficiary, recipient, cumulativePayout, beneficiarySig, callerPayout, issuerSig, {from: from})
    this.logs = logs
    this.receipt = receipt
  
    this.postconditions = {
      callerBalance: await this.ERC20PresetMinterPauser.balanceOf(from),
      recipientBalance: await this.ERC20PresetMinterPauser.balanceOf(recipient),
      totalHardDeposit: await this.ERC20SimpleSwap.totalHardDeposit(),
      hardDepositFor: await this.ERC20SimpleSwap.hardDeposits(beneficiary),
      liquidBalance: await this.ERC20SimpleSwap.liquidBalance(),
      liquidBalanceFor: await this.ERC20SimpleSwap.liquidBalanceFor(beneficiary),
      chequebookBalance: await this.ERC20SimpleSwap.balance(),
      paidOut: await this.ERC20SimpleSwap.paidOut(beneficiary),
      totalPaidOut: await this.ERC20SimpleSwap.totalPaidOut(),
      bounced: await this.ERC20SimpleSwap.bounced()
    }
  })
  cashChequeInternal(beneficiary, recipient, cumulativePayout, callerPayout, from)
}
function shouldNotCashCheque(beneficiaryToSign, issuerToSign, toSubmitFields, value, from, beneficiarySignee, issuerSignee, revertMessage) {
  beforeEach(async function() {
    this.beneficiarySig = await signCashOut(this.ERC20SimpleSwap, from, beneficiaryToSign.cumulativePayout, beneficiaryToSign.recipient, beneficiaryToSign.callerPayout, beneficiarySignee)
    this.issuerSig = await signCheque(this.ERC20SimpleSwap, issuerToSign.beneficiary, issuerToSign.cumulativePayout, issuerSignee)
  })
  it('reverts', async function() {
    await expectRevert(this.ERC20SimpleSwap.cashCheque(
      toSubmitFields.beneficiary, 
      toSubmitFields.recipient, 
      toSubmitFields.cumulativePayout, 
      this.beneficiarySig, 
      toSubmitFields.callerPayout,
      this.issuerSig,
      {from: from, value: value}), 
      revertMessage
    )
  })
}
function shouldPrepareDecreaseHardDeposit(beneficiary, decreaseAmount, from) {
  beforeEach(async function() {
    this.preconditions = {
      hardDepositFor: await this.ERC20SimpleSwap.hardDeposits(beneficiary)
    }
    const { logs } = await this.ERC20SimpleSwap.prepareDecreaseHardDeposit(beneficiary, decreaseAmount , {from: from})
    this.logs = logs

    this.postconditions = {
      hardDepositFor: await this.ERC20SimpleSwap.hardDeposits(beneficiary)
    }
  })

  it("should update the canBeDecreasedAt", async function() {
    let expectedCanBeDecreasedAt
    let personalDecreaseTimeout = (await this.ERC20SimpleSwap.hardDeposits(beneficiary)).timeout
    // if personalDecreaseTimeout is zero
    if(personalDecreaseTimeout.eq(new BN(0))) {
      // use the contract's default
      expectedCanBeDecreasedAt = await this.ERC20SimpleSwap.defaultHardDepositTimeout()
    } else {
      // use the value that was set
      expectedCanBeDecreasedAt = personalDecreaseTimeout
    }
    expect(this.postconditions.hardDepositFor.canBeDecreasedAt.toNumber()).to.be.closeTo(((await time.latest()).add(expectedCanBeDecreasedAt)).toNumber(), 5)
  })

  it('should update the decreaseAmount', function() {
    expect(this.postconditions.hardDepositFor.decreaseAmount).bignumber.to.be.equal(decreaseAmount)
  })

  it('should emit a HardDepositDecreasePrepared event', function() {
    expectEvent.inLogs(this.logs, 'HardDepositDecreasePrepared', {
      beneficiary,
      decreaseAmount
    })
  })
}
function shouldNotPrepareDecreaseHardDeposit(beneficiary, decreaseAmount, from, value, revertMessage) {
  it('reverts', async function() {
    await expectRevert(this.ERC20SimpleSwap.prepareDecreaseHardDeposit(
      beneficiary,
      decreaseAmount,
      {from: from, value: value}), 
      revertMessage
    )
  })
}
function shouldDecreaseHardDeposit(beneficiary, from) {
  beforeEach(async function() {
    this.preconditions = {
      hardDeposit: await this.ERC20SimpleSwap.totalHardDeposit(),
      hardDepositFor: await this.ERC20SimpleSwap.hardDeposits(beneficiary)
    }

    const { logs } = await this.ERC20SimpleSwap.decreaseHardDeposit(beneficiary, {from: from})
    this.logs = logs
    
    this.postconditions = {
      hardDeposit: await this.ERC20SimpleSwap.totalHardDeposit(),
      hardDepositFor: await this.ERC20SimpleSwap.hardDeposits(beneficiary)
    }
  })

  it('decreases the hardDeposit amount for the beneficiary', function() { 
    expect(this.postconditions.hardDepositFor.amount).bignumber.to.be.equal(this.preconditions.hardDepositFor.amount.sub(this.preconditions.hardDepositFor.decreaseAmount))
  })

 
  it('decreases the total hardDeposits', function() {
    expect(this.postconditions.hardDeposit).bignumber.to.be.equal(this.preconditions.hardDeposit.sub(this.preconditions.hardDepositFor.decreaseAmount))
  })

  it('resets the canBeDecreased at', function() {
    expect(this.postconditions.hardDepositFor.canBeDecreasedAt).bignumber.to.be.equal(new BN(0))
  })

  it('emits a hardDepositAmountChanged event', function() {
    expectEvent.inLogs(this.logs, 'HardDepositAmountChanged', {
      beneficiary,
      amount: this.postconditions.hardDepositFor.amount
    })
  })
}
function shouldNotDecreaseHardDeposit(beneficiary, from, value, revertMessage) {
  it('reverts', async function() {
    await expectRevert(this.ERC20SimpleSwap.decreaseHardDeposit(
      beneficiary,
      {from: from, value: value}), 
      revertMessage
    )
  })

}
function shouldIncreaseHardDeposit(beneficiary, amount, from) {
  beforeEach(async function () {
    this.preconditions = {
      balance: await this.ERC20SimpleSwap.balance(),
      liquidBalance: await this.ERC20SimpleSwap.liquidBalance(),
      liquidBalanceFor: await this.ERC20SimpleSwap.liquidBalanceFor(beneficiary),
      totalHardDeposit: await this.ERC20SimpleSwap.totalHardDeposit(),
      hardDepositFor: await this.ERC20SimpleSwap.hardDeposits(beneficiary),
    }
    const { logs } = await this.ERC20SimpleSwap.increaseHardDeposit(beneficiary, amount, { from: from })
    this.logs = logs
    this.postconditions = {
      balance: await this.ERC20SimpleSwap.balance(),
      liquidBalance: await this.ERC20SimpleSwap.liquidBalance(),
      liquidBalanceFor: await this.ERC20SimpleSwap.liquidBalanceFor(beneficiary),
      totalHardDeposit: await this.ERC20SimpleSwap.totalHardDeposit(),
      hardDepositFor: await this.ERC20SimpleSwap.hardDeposits(beneficiary)
    }
  })

  it('should decrease the liquidBalance', function () {
    expect(this.postconditions.liquidBalance).bignumber.to.be.equal(this.preconditions.liquidBalance.sub(amount))
  })

  it('should not affect the liquidBalanceFor', function () {
    expect(this.postconditions.liquidBalanceFor).bignumber.to.be.equal(this.preconditions.liquidBalanceFor)
  })

  it('should not affect the balance', function () {
    expect(this.postconditions.balance).bignumber.to.be.equal(this.preconditions.balance)
  })

  it('should increase the totalHardDeposit', function () {
    expect(this.postconditions.totalHardDeposit).bignumber.to.be.equal(this.preconditions.totalHardDeposit.add(amount))
  })

  it('should increase the hardDepositFor', function() {
    expect(this.postconditions.hardDepositFor.amount).bignumber.to.be.equal(this.preconditions.hardDepositFor.amount.add(amount))
  })

  it('should not influence the timeout', function() {
    expect(this.postconditions.hardDepositFor.timeout).bignumber.to.be.equal(this.preconditions.hardDepositFor.timeout)
  })

  it('should set canBeDecreasedAt to zero', function() {
    expect(this.postconditions.hardDepositFor.canBeDecreasedAt).bignumber.to.be.equal(new BN(0))
  })

  it('emits a hardDepositAmountChanged event', function() {
    expectEvent.inLogs(this.logs, 'HardDepositAmountChanged', {
      beneficiary,
      amount
    })
  })
}
function shouldNotIncreaseHardDeposit(beneficiary, amount, from, value, revertMessage) {
  it('reverts', async function() {
    await expectRevert(this.ERC20SimpleSwap.increaseHardDeposit(
      beneficiary,
      amount,
      {from: from, value: value}), 
      revertMessage
    )
  })
}
function shouldSetCustomHardDepositTimeout(beneficiary, timeout, from) {
  beforeEach(async function() {
    const beneficiarySig = await signCustomDecreaseTimeout(this.ERC20SimpleSwap, beneficiary, timeout, beneficiary)

    const { logs } = await this.ERC20SimpleSwap.setCustomHardDepositTimeout(beneficiary, timeout, beneficiarySig, {from: from})
    this.logs = logs

    this.postconditions = {
      hardDepositFor: await this.ERC20SimpleSwap.hardDeposits(beneficiary)
    }
  })

  it('should have set the timeout', async function() {
    expect(this.postconditions.hardDepositFor.timeout).bignumber.to.be.equal(timeout)
  })

  it('emits a HardDepositTimeoutChanged event', function() {
    expectEvent.inLogs(this.logs, 'HardDepositTimeoutChanged', {
      beneficiary,
      timeout
    })
  })
}
function shouldNotSetCustomHardDepositTimeout(toSubmit, toSign, signee, from, value, revertMessage) {
  beforeEach(async function() {
    this.beneficiarySig = await signCustomDecreaseTimeout(this.ERC20SimpleSwap, toSign.beneficiary, toSign.timeout, signee)
  })
  it('reverts', async function() {
    await expectRevert(this.ERC20SimpleSwap.setCustomHardDepositTimeout(
      toSubmit.beneficiary,
      toSubmit.timeout,
      this.beneficiarySig,
      {from: from, value: value}), 
      revertMessage
    )
  })
}

function shouldWithdraw(amount, from) {
  beforeEach(async function() {
    this.preconditions = {
      callerBalance: await this.ERC20PresetMinterPauser.balanceOf(from),
      liquidBalance: await this.ERC20SimpleSwap.liquidBalance()
    }

    await this.ERC20SimpleSwap.withdraw(amount, {from: from})

    this.postconditions = {
      callerBalance: await  this.ERC20PresetMinterPauser.balanceOf(from),
      liquidBalance: await this.ERC20SimpleSwap.liquidBalance()
    }
  })

  it('should have updated the liquidBalance', function() {
    expect(this.postconditions.liquidBalance).bignumber.to.be.equal(this.preconditions.liquidBalance.sub(amount))
  })

  it('should have updated the callerBalance', function() {
    expect(this.postconditions.callerBalance).bignumber.to.be.equal(this.preconditions.callerBalance.add(amount))
  })
}
function shouldNotWithdraw(amount, from, value, revertMessage) {
  it('reverts', async function() {
    await expectRevert(this.ERC20SimpleSwap.withdraw(
      amount,
      {from: from, value: value}), 
      revertMessage
    )
  })
}

function shouldDeposit(amount, from) {
  beforeEach(async function() {
    this.preconditions = {
      balance: await this.ERC20SimpleSwap.balance(),
      totalHardDeposit: await this.ERC20SimpleSwap.totalHardDeposit(),
      liquidBalance: await this.ERC20SimpleSwap.liquidBalance()
    }
    const { logs } = await this.ERC20PresetMinterPauser.transfer(this.ERC20SimpleSwap.address, amount, {from: from})
    this.logs = logs
  })
  it('should update the liquidBalance of the checkbook', async function() {
    expect(await this.ERC20SimpleSwap.liquidBalance()).bignumber.to.equal(this.preconditions.liquidBalance.add(amount))
  })
  it('should update the balance of the checkbook', async function() {
    expect(await this.ERC20SimpleSwap.balance()).bignumber.to.equal(this.preconditions.balance.add(amount))
  })
  it('should not afect the totalHardDeposit', async function() {
    expect(await this.ERC20SimpleSwap.totalHardDeposit()).bignumber.to.equal(this.preconditions.totalHardDeposit)
  })
  it('should emit a transfer event', async function() {
    expectEvent.inLogs(this.logs, "Transfer", {
      from: from,
      to: this.ERC20SimpleSwap.address,
      value: amount
    })
  })
}
module.exports = {
  shouldDeploy,
  shouldReturnDefaultHardDepositTimeout,
  shouldReturnPaidOut,
  shouldReturnTotalPaidOut,
  shouldReturnHardDeposits,
  shouldReturnTotalHardDeposit,
  shouldReturnIssuer,
  shouldReturnLiquidBalance,
  shouldReturnLiquidBalanceFor,
  shouldCashChequeBeneficiary,
  shouldNotCashChequeBeneficiary,
  shouldCashCheque,
  shouldNotCashCheque,
  shouldPrepareDecreaseHardDeposit,
  shouldNotPrepareDecreaseHardDeposit,
  shouldDecreaseHardDeposit,
  shouldNotDecreaseHardDeposit,
  shouldIncreaseHardDeposit,
  shouldNotIncreaseHardDeposit,
  shouldSetCustomHardDepositTimeout,
  shouldNotSetCustomHardDepositTimeout,
  shouldWithdraw,
  shouldNotWithdraw,
  shouldDeposit,
}

