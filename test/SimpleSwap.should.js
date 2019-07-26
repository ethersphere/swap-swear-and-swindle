const {
  BN,
  balance,
  time,
  expectRevert,
  constants,
  expectEvent
} = require("openzeppelin-test-helpers");


const { signCheque } = require("./swutils");

const { expect } = require('chai');

function shouldReturnDEFAULT_HARDDEPPOSIT_DECREASE_TIMEOUT(expected) {
  it('should return the expected DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT', async function() {
    expect(await this.simpleSwap.DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT()).bignumber.to.be.equal(expected)
  })
}

function shouldReturnCheques(beneficiary, expectedSerial, expectedAmount, expectedPaidOut, expectedCashTimeout) {

}

function shouldReturnHarddeposits(beneficiary, expectedAmount, expectedDecreaseTimeout, expectedCanBeDecreasedAt) {

}

function shouldReturnTotalharddeposit(expectedHardDeposit) {

}

function shouldReturnIssuer(expectedIssuer) {
  it('should return the expected issuer', async function() {
    expect(await this.simpleSwap.issuer()).to.be.equal(expectedIssuer)
  })

}

function shouldReturnLiquidBalance(expectedLiquidBalance) {
  it('should return the expected liquidBalance', async function() {
    expect(await this.simpleSwap.liquidBalance()).bignumber.to.equal(expectedLiquidBalance)
  })
}

function shouldReturnLiquidBalanceFor(beneficiary, expectedLiquidBalance) {
  it('should return the expected liquidBalance', async function() {
    expect(await this.simpleSwap.liquidBalanceFor(beneficiary)).bignumber.to.equal(expectedLiquidBalance)
  })
}

function submitChequeInternal() {
  it('should update the cheque serial number', async function() {
    expect(this.postconditions.cheque.serial).bignumber.is.equal(this.signedCheque.serial, "serial was not updated")
  })

  it('should update the cheque amount', async function() {
    expect(this.postconditions.cheque.amount).bignumber.is.equal(this.signedCheque.amount, "amount was not updated")
  })

  it('should update the cheque timeout', async function() {
    expect(parseInt(this.postconditions.cheque.cashTimeout)).is.equal(parseInt(await time.latest()) + parseInt(this.signedCheque.timeout))
  })

  it('should emit a chequeSubmitted event', async function() {
    expectEvent.inLogs(this.logs, "ChequeSubmitted", {
      amount: this.signedCheque.amount,
      beneficiary: this.signedCheque.beneficiary,
      serial: this.signedCheque.serial
    })
  })
}

function shouldSubmitChequeIssuer(unsignedCheque, from) {
  beforeEach(async function() {
    this.preconditions = {
      cheque: await this.simpleSwap.cheques(unsignedCheque.beneficiary)
    }
    this.signedCheque = await signCheque(this.simpleSwap, unsignedCheque)
    const { logs } = await this.simpleSwap.submitChequeIssuer(this.signedCheque.beneficiary, this.signedCheque.serial, this.signedCheque.amount, this.signedCheque.timeout, this.signedCheque.signature, {from: from})
    this.logs = logs
    this.postconditions = {
      cheque: await this.simpleSwap.cheques(unsignedCheque.beneficiary)
    }
  })
  submitChequeInternal() 
}

function shouldNotSubmitChequeIssuer(toSignCheque, functionParams, from, value, revertMessage) {
  beforeEach(async function() {
    this.signedCheque = await signCheque(this.simpleSwap, toSignCheque)
  })
  it('reverts', async function() {
    await expectRevert(this.simpleSwap.submitChequeIssuer(
      functionParams.beneficiary,
      functionParams.serial, 
      functionParams.amount, 
      functionParams.timeout,
      this.signedCheque.signature, {from: from, value: value}), revertMessage)
  })
}

function shouldSubmitChequeBeneficiary(unsignedCheque, from) {
  beforeEach(async function() {
    this.preconditions = {
      cheque: await this.simpleSwap.cheques(unsignedCheque.beneficiary)
    }
    this.signedCheque = await signCheque(this.simpleSwap, unsignedCheque)
    const { logs } = await this.simpleSwap.submitChequeBeneficiary(this.signedCheque.serial, this.signedCheque.amount, this.signedCheque.timeout, this.signedCheque.signature, {from: from})
    this.logs = logs
    this.postconditions = {
      cheque: await this.simpleSwap.cheques(unsignedCheque.beneficiary)
    }
  })
  submitChequeInternal() 
}
function shouldNotSubmitChequeBeneficiary(toSignCheque, functionParams, from, value, revertMessage) {
  beforeEach(async function() {
    this.signedCheque = await signCheque(this.simpleSwap, toSignCheque)
  })
  it('reverts', async function() {
    await expectRevert(this.simpleSwap.submitChequeBeneficiary(
      functionParams.serial, 
      functionParams.amount, 
      functionParams.timeout,
      this.signedCheque.signature, {from: from, value: value}), revertMessage)
  })

}
function shouldSubmitCheque(unsignedCheque, from) {
  beforeEach(async function() {
    this.preconditions = {
      cheque: await this.simpleSwap.cheques(unsignedCheque.beneficiary)
    }
    this.signedCheque = await signCheque(this.simpleSwap, unsignedCheque)
    const { logs } = await this.simpleSwap.submitCheque(
      this.signedCheque.beneficiary, 
      this.signedCheque.serial, 
      this.signedCheque.amount, 
      this.signedCheque.timeout, 
      this.signedCheque.signature.issuer, 
      this.signedCheque.signature.beneficiary, 
      {from: from}
    )
    this.logs = logs
    this.postconditions = {
      cheque: await this.simpleSwap.cheques(unsignedCheque.beneficiary)
    }
  })
  submitChequeInternal() 
}
function shouldNotSubmitCheque(unsignedCheque, functionParams, from, value, revertMessage) {
  beforeEach(async function() {
    this.signedCheque = await signCheque(this.simpleSwap, unsignedCheque)
  })
  it('reverts', async function() {
    await expectRevert(this.simpleSwap.submitCheque(
      functionParams.beneficiary, 
      functionParams.serial, 
      functionParams.amount, 
      functionParams.timeout, 
      this.signedCheque.signature.issuer, 
      this.signedCheque.signature.beneficiary, {from: from, value: value}), revertMessage)
  })
}
function cashChequeInternal(beneficiaryPrincipal, beneficiaryAgent, requestPayout, calleePayout, from) {
  let totalPayout 
  // if the requested payout is less than the liquidBalance available for beneficiary
  if(requestPayout.lt(this.preconditions.liquidBalanceFor)) {
    // full amount requested can be paid out
    totalPayout = requestPayout
  } else {
    // partial amount requested can be paid out (the liquid balance available to the node)
    totalPayout = liquidBalanceFor
  }
  it('should update the totalHardDeposit and hardDepositFor ', function() {
    let expectedDecreaseHardDeposit
    // if the harddeposits can cover the totalPayout
    if(totalPayout.lt(this.preconditions.hardDepositFor)) {
      // harddeposit decreases by totalPayout
      expectedDecreaseHardDeposit = totalPayout
    } else {
      // harddeposit decreases by the full amount (and rest is from global liquid balance)
      expectedDecreaseHardDeposit = this.preconditions.hardDepositFor
    }
    // totalHarddeposit
    expect(this.postconditions.totalHardDeposit).bignumber.to.be.equal(this.preconditions.totalHardDeposit.sub(expectedDecreaseHardDeposit))    
    // hardDepositFor
    expect(this.postconditions.hardDepositFor).bignumber.to.be.equal(this.preconditions.hardDepositFor.sub(expectedDecreaseHardDeposit))    
  })
  
  it('should update paidOut', async function() {
    expect(this.postconditions.cheque.paidOut).bignumber.to.be.equal(this.preconditions.cheque.paidOut.add(totalPayout))
  })

  it('should transfer the correct amount to the beneficiaryAgent', async function() {
    let beneficiaryAgentTransactionCosts
    // if the beneficiary agent equal the sender
    if(beneficiaryAgent == from) {
      // the beneficiaryAgent bears the transaction costs
      beneficiaryAgentTransactionCosts = await computeCost(this.receipt)
    } else {
      // somebody else pays for the transaction costs
      beneficiaryAgentTransactionCosts = new BN(0)
    }
    expect(this.postconditions.beneficiaryAgentBalance).bignumber.to.be.equal(
      this.postconditions.beneficiaryAgentBalance
        .add(totalPayout)
        .sub(calleePayout)
        .sub(beneficiaryAgentTransactionCosts)
      )
  })
  it('should transfer the correct amount to the callee', async function() {
    let expectedCalleeTransactionCosts = await computeCost(this.receipt)
    let expectedAmountCallee
    // if the beneficiary agent equal the sender
    if(beneficiaryAgent == from) {
      // the callee gets the totalPayout
      expectedAmountCallee = totalPayout.sub(expectedCalleeTransactionCosts)
    }   else {
      // the callee get's a part of the totalPayout
      expectedAmountCallee = calleePayout.sub(expectedCalleeTransactionCosts)
    }
    expect(this.postconditions.calleeBalance).bignumber.to.be.equal(this.preconditions.calleeBalance.add(expectedAmountCallee))
  })
  it('should emit a ChequeCashed event', function() {
    expectEvent.inLogs(this.logs, "ChequeCashed", {
      beneficiaryPrincipal: beneficiaryPrincipal,
      beneficiaryAgent: beneficiaryAgent,
      callee: from,
      serial: this.postconditions.cheque.serial,
      totalPayout: totalPayout,
      requestPayout: requestPayout,
      calleePayout: calleePayout
    })
  })
  if(totalPayout < requestPayout) {
    it('should emit a ChequeBounced event', function() {
      expectEvent.inLogs(this.logs, "ChequeBounced", {})
    })
  }
}
function shouldCashChequeBeneficiary(beneficiaryAgent, requestPayout, from) {
  beforeEach(async function() {
    this.preconditions = {
      calleeBalance: await balance.current(from),
      beneficiaryAgentBalance: await balance.current(beneficiaryAgent),
      beneficiaryPrincipalBalance: calleeBalance,
      totalHarddeposit: await this.simpleSwap.totalHardDeposit(),
      hardDepositFor: await this.simpleSwap.hardDepositFor(from),
      liquidBalance: await this.simpleSwap.liquidBalance(),
      liquidBalanceFor: await this.simpleSwap.liquidBalanceFor(from),
      chequebookBalance: await balance.current(this.simpleSwap.address),
      beneficiaryBalance: await balance.current(beneficiaryAgent),
      cheque: await this.simpleSwap.cheques(from)
    }
  
    const { logs, receipt } = this.simpleSwap.cashChequeBeneficiary(beneficiaryAgent, requestPayout, {from: from})
    this.logs = logs
    this.receipt = receipt
  
    this.postconditions = {
      calleeBalance: await balance.current(from),
      beneficiaryAgentBalance: await balance.current(beneficiaryAgent),
      beneficiaryPrincipalBalance: calleeBalance,
      totalHarddeposit: await this.simpleSwap.totalHardDeposit(),
      hardDepositFor: await this.simpleSwap.hardDepositFor(from),
      liquidBalance: await this.simpleSwap.liquidBalance(),
      liquidBalanceFor: await this.simpleSwap.liquidBalanceFor(from),
      chequebookBalance: await balance.current(this.simpleSwap.address),
      beneficiaryBalance: await balance.current(beneficiaryAgent),
      cheque: await this.simpleSwap.cheques(from)
    }
  })
  cashChequeInternal(from, beneficiaryAgent, requestPayout, 0, from)
}
function shouldNotCashChequeBeneficiary(beneficiaryAgent, requestPayout, from, value, revertMessage) {

}
function shouldCashCheque(beneficiaryPrincipal, beneficiaryAgent, requestPayout, beneficiarySig, expiry, calleePayout, from) {

}
function shouldNotCashCheque(beneficiaryPrincipal, beneficiaryAgent, requestPayout, value,  beneficiarySig, expiry, calleePayout, from, revertMessage) {

}
function shouldPrepareDecreaseHardDeposit(beneficiary, decreaseAmount, from) {
  beforeEach(async function() {
    await this.simpleSwap.send(amount)
    await this.simpleSwap.increaseHardDeposit(beneficiary, amount)

    let { logs } = await this.simpleSwap.prepareDecreaseHardDeposit(
      beneficiary,
      amount, {
        from: issuer
      }
    )

    this.logs = logs
    this.timeout = (await this.simpleSwap.hardDeposits(beneficiary))[2]
  })

  it('should fire the HardDepositDecreasePrepared event', function() {
    expectEvent.inLogs(this.logs, 'HardDepositDecreasePrepared', {
      beneficiary,
      decreaseAmount: amount
    })
  })

  it('should set the decreaseAmount', async function() {
    expect((await this.simpleSwap.hardDeposits(beneficiary))[1]).bignumber.is.equal(amount)
  })

  it('should set the canBeDecreasedAt', async function() {
    expect((await this.simpleSwap.hardDeposits(beneficiary))[3]).bignumber.is.gte((await time.latest()).add(this.timeout))
  })

}
function shouldNotPrepareDecreaseHardDeposit(beneficiary, decreaseAmount, from, value, revertMessage) {

}
function shouldDecreaseHardDeposit(beneficiary, from) {

}
function shouldNotDecreaseHardDeposit(beneficiary, from, value, revertMessage) {

}
function shouldIncreaseHardDeposit(beneficiary, amount, from) {
  beforeEach(async function () {
    this.preconditions = {
      balance: await balance.current(this.simpleSwap.address),
      liquidBalance: await this.simpleSwap.liquidBalance(),
      liquidBalanceFor: await this.simpleSwap.liquidBalanceFor(beneficiary),
      totalHardDeposit: await this.simpleSwap.totalHardDeposit(),
      hardDepositFor: await this.simpleSwap.hardDeposits(beneficiary),
    }
    const { logs } = this.simpleSwap.increaseHardDeposit(beneficiary, amount, { from: from })
    this.logs = logs
    this.postconditions = {
      balance: await balance.current(this.simpleSwap.address),
      liquidBalance: await this.simpleSwap.liquidBalance(),
      liquidBalanceFor: await this.simpleSwap.liquidBalanceFor(beneficiary),
      totalHardDeposit: await this.simpleSwap.totalHardDeposit(),
      hardDepositFor: await this.simpleSwap.hardDeposits(beneficiary)
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

  it('should increase the harddepositFor', function() {
    expect(this.postconditions.hardDepositFor.amount).bignumber.to.be.equal(this.preconditions.hardDepositFor.amount.add(amount))
  })

  it('should set the decreaseTimeout to the default value', function() {
    expect(this.postconditions.hardDepositFor.decreaseTimeout).bignumber.to.be.equal(new BN(0))
  })

  it('should set canBeDecreasedAt to zero', function() {
    expect(this.postconditions.hardDepositFor.decreaseTimeout).bignumber.to.be.equal(new BN(0))
  })



}
function shouldNotIncreaseHardDeposit(beneficiary, amount, from, value, revertMessage) {

}
function shouldSetCustomHardDepositDecreaseTimeout(beneficiary, decreaseTimeout, beneficiarySig, from) {

}
function shouldNotSetCustomHardDepositDecreaseTimeout(beneficiary, decreaseTimeout, beneficiarySig, from,value, revertMessage) {

}

function shouldWithdraw(amount, from) {

}
function shouldNotWithdraw(amount, from, value, revertMessage) {

}

function shouldDeposit(amount, from) {
  beforeEach(async function() {
    this.preconditions = {
      balance: await balance.current(this.simpleSwap.address),
      totalHardDeposit: await this.simpleSwap.totalHardDeposit(),
      liquidBalance: await this.simpleSwap.liquidBalance()
    }
    const { logs } = await this.simpleSwap.send(amount, {from: from})
    this.logs = logs
  })
  it('should update the liquidBalance of the checkbook', async function() {
    expect(await this.simpleSwap.liquidBalance()).bignumber.to.equal(this.preconditions.liquidBalance.add(amount))
  })
  it('should update the balance of the checkbook', async function() {
    expect(await balance.current(this.simpleSwap.address)).bignumber.to.equal(this.preconditions.balance.add(amount))
  })
  it('should not afect the totalHardDeposit', async function() {
    expect(await this.simpleSwap.totalHardDeposit()).bignumber.to.equal(this.preconditions.totalHardDeposit)
  })
  it('should emit a deposit event', async function() {
    expectEvent.inLogs(this.logs, "Deposit", {
      depositor: from,
      amount: amount
    })
  })
}

module.exports = {
  shouldReturnDEFAULT_HARDDEPPOSIT_DECREASE_TIMEOUT,
  shouldReturnCheques,
  shouldReturnHarddeposits,
  shouldReturnTotalharddeposit,
  shouldReturnIssuer,
  shouldReturnLiquidBalance,
  shouldReturnLiquidBalanceFor,
  shouldSubmitChequeIssuer,
  shouldNotSubmitChequeIssuer,
  shouldSubmitChequeBeneficiary,
  shouldNotSubmitChequeBeneficiary,
  shouldSubmitCheque,
  shouldNotSubmitCheque,
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
  shouldSetCustomHardDepositDecreaseTimeout,
  shouldNotSetCustomHardDepositDecreaseTimeout,
  shouldWithdraw,
  shouldNotWithdraw,
  shouldDeposit
}

