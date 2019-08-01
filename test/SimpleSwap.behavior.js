const {
  BN,
  balance,
  time,
  expectRevert,
  constants,
  expectEvent
} = require("openzeppelin-test-helpers");

const { expect } = require('chai');

const { sign } = require("./swutils");


const { computeCost } = require("./testutils");
const {
  shouldReturnDEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT,
  shouldReturnPaidOutCheques,
  shouldReturnHardDeposits,
  shouldReturnTotalHardDeposit,
  shouldReturnIssuer,
  shouldReturnLiquidBalance,
  shouldReturnBalanceFor,
  shouldSubmitChequeBeneficiary,
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
  shouldDeposit,
  shouldNotDeposit
} = require('./SimpleSwap.should.js')

// switch to false if you don't want to test the particular function
enabledTests = {
  DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT: true,
  cheques: true,
  hardDeposits: false,
  totalHardDeposit: false,
  issuer: false,
  liquidBalance: false,
  liquidBalanceFor: false,
  cashChequeBeneficiary: false,
  cashCheque: false,
  prepareDecreaseHardDeposit: false,
  decreaseHardDeposit: false,
  increaseHardDeposit: false,
  setCustomHardDepositDecreaseTimeout: false,
  withdraw: false, 
  deposit: false
}

// constants to make the test-log more readable
const describeFunction = 'FUNCTION: '
const describePreCondition = 'PRE-CONDITION: '
const describeTest = 'TEST: '

// @param balance total ether deposited in checkbook
// @param liquidBalance totalDeposit - hardDeposits
// @param issuer the issuer of the checkbook
// @param alice a counterparty of the checkbook 
// @param bob a counterparty of the checkbook
function shouldBehaveLikeSimpleSwap([issuer, alice, bob, _recipient], DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT) {
  const defaultCheque = {
    beneficiary: bob,
    amount: new BN(500),
  }
  context('as a simple swap', function () {
    describe(describeFunction + 'DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT', function () {
      if (enabledTests.DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT) {
        shouldReturnDEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT(DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT)
      }
    })
    describe(describeFunction + 'paidOutCheques', function () {
      if (enabledTests.cheques) {
        const beneficiary = defaultCheque.beneficiary
        context('when no cheque was ever cashed', function () {
          describe(describeTest + 'shouldReturnPaidOutCheques', function () {
            const expectedAmount = new BN(0)
            shouldReturnPaidOutCheques(beneficiary, expectedAmount)
          })
        })
        context('when a cheque was cashed', function () {
          describe(describePreCondition + 'shouldDeposit', function() {
            shouldDeposit(defaultCheque.amount, issuer)
            describe(describePreCondition + 'shouldCashChequeBeneficiary', function () {
              shouldCashChequeBeneficiary(defaultCheque.beneficiary, defaultCheque.amount, issuer, defaultCheque.beneficiary)
              describe(describeTest + 'shouldReturnPaidOutCheques', function () {
                const expectedAmount = defaultCheque.amount
                shouldReturnPaidOutCheques(beneficiary, expectedAmount)
              })
            })
          })
        })
      }
    })

    describe(describeFunction + 'hardDeposits', function () {
      if (enabledTests.hardDeposits) {
        const beneficiary = defaultCheque.beneficiary
        context('when no hardDeposit was allocated', function() {
          const expectedAmount = new BN(0)
          const exptectedDecreaseAmount = new BN(0)
          const exptectedCanBeDecreasedAt = new BN(0)
          context('when no custom decreaseTimeout was set', function() {
            const expectedDecreaseTimeout = new BN(0)
            describe(describeTest + 'shouldReturnHardDeposits', function() {
              shouldReturnHardDeposits(beneficiary, expectedAmount, exptectedDecreaseAmount,  expectedDecreaseTimeout, exptectedCanBeDecreasedAt)
            })
          })
          context('when a custom decreaseTimeout was set', function() {
            const expectedDecreaseTimeout = new BN(60)
            describe(describePreCondition + 'shouldSetCustomDecreaseTimeout', function() {
              shouldSetCustomHardDepositDecreaseTimeout(beneficiary, expectedDecreaseTimeout, issuer)
              describe(describeTest + 'shouldReturnHardDeposits', function() {
                shouldReturnHardDeposits(beneficiary, expectedAmount, exptectedDecreaseAmount,  expectedDecreaseTimeout, exptectedCanBeDecreasedAt)
              })
            })
          })
        })
        context('when a hardDeposit was allocated', function() {
          describe(describePreCondition + 'shouldDeposit', function() {
            const depositAmount = new BN (50)
            shouldDeposit(depositAmount, issuer)
            describe(describePreCondition + 'shouldIncreaseHardDeposit', function() {
              shouldIncreaseHardDeposit(beneficiary, depositAmount, issuer)
              context('when the hardDeposit was not requested to decrease', function() {
                const expectedDecreaseAmount = new BN(0)
                const expectedCanBeDecreasedAt = new BN(0)
                const expectedDecreaseTimeout = new BN(0)
                describe(describeTest + 'shouldReturnHardDeposits', function() {
                  shouldReturnHardDeposits(beneficiary, depositAmount, expectedDecreaseAmount, expectedDecreaseTimeout, expectedCanBeDecreasedAt)
                })
              })
              context('when the hardDeposit was requested to decrease', function() {
                describe(describePreCondition + 'shouldPrepareDecreaseHardDeposit', function() {
                  const toDecrease = depositAmount.div(new BN(2))
                  shouldPrepareDecreaseHardDeposit(beneficiary, toDecrease, issuer)
                  describe(describeTest + 'shouldReturnHardDeposits', function() {
                    const expectedDecreaseTimeout = new BN(0)
      
                    shouldReturnHardDeposits(beneficiary, depositAmount, toDecrease, expectedDecreaseTimeout, new BN(42)) // 42 (not BN(0)) signifies that we have to define it later
                  })
                })
              })
            })
          })
        })
      }
    })

    describe(describeFunction + 'totalHardDeposits', function() {
      if(enabledTests.totalHardDeposit) {
        context('when there are no hardDeposits', function() {
          describe(describeTest + 'shouldReturnTotalHardDeposit', function() {
            shouldReturnTotalHardDeposit(new BN(0))
          })
        })
        context('when there are hardDeposits', function() {
          const depositAmount = new BN(50)
          describe(describePreCondition + 'shouldDeposit', function() {
            shouldDeposit(depositAmount, issuer)
            describe(describePreCondition + 'shouldIncreaseHardDeposit', function() {
              shouldIncreaseHardDeposit(defaultCheque.beneficiary, depositAmount, issuer)
              describe(describeTest + 'shouldReturnTotalHardDeposit', function() {
                shouldReturnTotalHardDeposit(depositAmount)
              })
            })
          })
        })
      }
    })

    describe(describeFunction + 'issuer', function () {
      if (enabledTests.issuer) {
        shouldReturnIssuer(issuer)
      }
    })

    describe(describeFunction + 'liquidBalance', function () {
      if (enabledTests.liquidBalance) {
        context('when there is some balance', function () {
          describe(describePreCondition + 'shouldDeposit', function () {
            const depositAmount = new BN(50)
            shouldDeposit(depositAmount, issuer)
            context('when there are hardDeposits', function () {
              describe('when the hardDeposits equal the depositAmount', function () {
                describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                  const hardDeposit = depositAmount
                  shouldIncreaseHardDeposit(defaultCheque.beneficiary, hardDeposit, issuer)
                  describe(describeTest + 'liquidBalance', function () {
                    shouldReturnLiquidBalance(depositAmount.sub(hardDeposit))
                  })
                })
                describe('when the hardDeposits are lower than the depositAmount', function () {
                  describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                    const hardDeposit = depositAmount.sub(new BN(40))
                    shouldIncreaseHardDeposit(defaultCheque.beneficiary, hardDeposit, issuer)
                    describe(describeTest + 'shouldReturnLiquidBalance', function () {
                      shouldReturnLiquidBalance(depositAmount.sub(hardDeposit))
                    })
                  })
                })
              })
              context('when there are no hardDeposits', function () {
                describe(describeTest + 'shouldReturnLiquidBalance', function () {
                  shouldReturnLiquidBalance(depositAmount)
                })
              })
            })
          })
          context('when there is no balance', function () {
            describe(describeTest + 'shouldReturnLiquidBalance', function () {
              shouldReturnLiquidBalance(new BN(0))
            })
          })
        })
      }
    })

    describe(describeFunction + 'shouldReturnBalanceFor', function () {
      if (enabledTests.liquidBalanceFor) {
        const beneficiary = bob
        const depositAmount = new BN(50)
        context('when there is some balance', function () {
          describe(describePreCondition + 'shoulDeposit', function () {
            shouldDeposit(depositAmount, issuer)
            context('when there are no hard deposits', function () {
              describe(describeTest + 'shouldReturnBalanceFor', function () {
                shouldReturnBalanceFor(beneficiary, depositAmount)
              })
            })
            context('when there are no hard deposits', function () {
              const hardDeposit = new BN(10)
              describe('when these hard deposits are assigned to the beneficiary', function () {
                describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                  shouldIncreaseHardDeposit(beneficiary, hardDeposit, issuer)
                  describe(describeTest + 'shouldReturnBalanceFor', function () {
                    shouldReturnBalanceFor(beneficiary, depositAmount)
                  })
                })
              })
              describe('when these hard deposits are assigned to somebody else', function () {
                describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                  shouldIncreaseHardDeposit(alice, hardDeposit, issuer)
                  describe(describeTest + 'shouldReturnBalanceFor', function () {
                    shouldReturnBalanceFor(beneficiary, depositAmount.sub(hardDeposit))
                  })
                })
              })
            })
          })
        })
        describe('when there is no balance', function () {
          shouldReturnBalanceFor(beneficiary, new BN(0))
        })
      }
    })

    describe(describeFunction + 'cashCheque', function () {
      if (enabledTests.cashCheque) {
        context("when we don't send value along", function () {
          const value = new BN(0)
          context('when the signature has not expired', function () {
            context('when the expiry is in the future', function () {
              beforeEach(async function () {
                this.expiry = (await time.latest()).add(time.duration.years(1))
              })
              context('when the calleePayout is non-zero', function () {
                const calleePayout = new BN(1)
                context('when the beneficiary is a signee', function () {
                  const signee = defaultCheque.beneficiary
                  context('when the beneficiary signs the correct fields', function () {
                    context('when the beneficiary and recipient are not the sender', function () {
                      const sender = alice
                      context("when the recipient is not the beneficiary", function () {
                        const recipient = _recipient
                        context('when we submit the cheque beforeHand', function () {
                          describe(describePreCondition + 'shouldSubmitChequeBeneficiary', function () {
                            shouldSubmitChequeBeneficiary(defaultCheque, defaultCheque.beneficiary)
                            context('when we have not cashed a cheque before', function () {
                              context('when we have waited more than timeoutDuration', function () {
                                const waitTime = defaultCheque.timeout.add(new BN(100))
                                beforeEach(async function () {
                                  await time.increase(waitTime)
                                })
                                context('when the requestPayout is equal to the submitted value', function () {
                                  const requestPayout = defaultCheque.amount
                                  context('when there is some balance', function () {
                                    context('when the balance is bigger than the requestPayout', function () {
                                      describe(describePreCondition + 'shouldDeposit', function () {
                                        const depositAmount = requestPayout.add(new BN(50))
                                        shouldDeposit(depositAmount, issuer)
                                        context('when there are hardDeposits', function () {
                                          describe('when the hardDeposits are assigned to the beneficiary', function () {
                                            const hardDepositReceiver = defaultCheque.beneficiary
                                            context('when the hardDeposit is more the requestPayout', function () {
                                              const hardDeposit = requestPayout.add(new BN(1))
                                              describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                                                shouldIncreaseHardDeposit(hardDepositReceiver, hardDeposit, issuer)
                                                describe(describeTest + 'shouldCashCheque', function () {
                                                  shouldCashCheque(defaultCheque.beneficiary, recipient, requestPayout, calleePayout, sender)
                                                })
                                              })
                                            })
                                            context('when the hardDeposit equals the requestPayout', function () {
                                              const hardDeposit = requestPayout
                                              describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                                                shouldIncreaseHardDeposit(hardDepositReceiver, hardDeposit, issuer)
                                                describe(describeTest + 'shouldCashCheque', function () {
                                                  shouldCashCheque(defaultCheque.beneficiary, recipient, requestPayout, calleePayout, sender)
                                                })
                                              })
                                            })
                                            context('when the hardDeposit is less than the requestPayout', function () {
                                              const hardDeposit = requestPayout.sub(new BN(1))
                                              describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                                                shouldIncreaseHardDeposit(hardDepositReceiver, hardDeposit, issuer)
                                                describe(describeTest + 'shouldCashCheque', function () {
                                                  shouldCashCheque(defaultCheque.beneficiary, recipient, requestPayout, calleePayout, sender)
                                                })
                                              })
                                            })
                                          })
                                          describe('when the hardDeposits are assigned to somebody else', function () {
                                            const hardDepositReceiver = alice
                                            const hardDeposit = requestPayout.add(new BN(1))
                                            describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                                              shouldIncreaseHardDeposit(hardDepositReceiver, hardDeposit, issuer)
                                              describe(describeTest + 'shouldCashCheque', function () {
                                                shouldCashCheque(defaultCheque.beneficiary, recipient, requestPayout, calleePayout, sender)
                                              })
                                            })
                                          })
                                        })
                                      })
                                    })
                                    context('when the balance equals the requestPayout', function () {
                                      describe(describePreCondition + 'shouldDeposit', function () {
                                        const depositAmount = requestPayout
                                        shouldDeposit(depositAmount, issuer)
                                        describe(describeTest + 'shouldCashCheque', function () {

                                          shouldCashCheque(defaultCheque.beneficiary, recipient, requestPayout, calleePayout, sender)
                                        })
                                      })
                                    })
                                  })
                                  context('when there is no balance', function () {
                                    describe(describeTest + 'shouldNotCashCheque', function () {
                                      const toSignFields = {
                                        requestPayout,
                                        recipient,
                                        calleePayout,
                                      }
                                      const toSubmitFields = Object.assign({}, toSignFields, { beneficiary: defaultCheque.beneficiary })
                                      shouldNotCashCheque(toSignFields, toSubmitFields, value, sender, signee, "SimpleSwap: cannot pay callee")
                                    })
                                  })
                                })
                                context('when the requestPayout is less than the submitted value', function () {
                                  describe(describePreCondition + 'shouldDeposit', function () {
                                    shouldDeposit(defaultCheque.amount, issuer)
                                    const requestPayout = defaultCheque.amount.sub(new BN(1))
                                    describe(describeTest + 'shouldCashCheque', function () {
                                      shouldCashCheque(defaultCheque.beneficiary, recipient, requestPayout, calleePayout, sender)
                                    })
                                  })
                                })
                              })
                            })
                            context('when we have waited timeoutDuration', function () {
                              const waitTime = defaultCheque.timeout
                              const requestPayout = defaultCheque.amount
                              beforeEach(async function () {
                                await time.increase(waitTime)
                              })
                              describe(describePreCondition + 'shouldDeposit', function () {
                                shouldDeposit(defaultCheque.amount, issuer)
                                describe(describeTest + 'shouldCashCheque', function () {
                                  shouldCashCheque(defaultCheque.beneficiary, recipient, requestPayout, calleePayout, sender)
                                })
                              })
                            })
                            context('when we have cashed a cheque before', function () {
                              describe('when we have cashed the partial amount before', function () {
                                const waitTime = defaultCheque.timeout.add(new BN(100))
                                beforeEach(async function () {
                                  await time.increase(waitTime)
                                })
                                describe(describePreCondition + 'shouldDeposit', function () {
                                  shouldDeposit(defaultCheque.amount, issuer)
                                  describe(describePreCondition + 'shouldCashChequeBeneficiary', function () {
                                    const requestPayout = defaultCheque.amount.div(new BN(2))
                                    shouldCashCheque(defaultCheque.beneficiary, recipient, requestPayout, calleePayout, sender)
                                    describe(describeTest + 'shouldCashCheque', function () {
                                      shouldCashCheque(defaultCheque.beneficiary, recipient, requestPayout, calleePayout, sender)
                                    })
                                  })
                                })
                              })
                              describe('when we have cashed the full amount before', function () {
                                const waitTime = defaultCheque.timeout.add(new BN(100))
                                beforeEach(async function () {
                                  await time.increase(waitTime)
                                })
                                describe(describePreCondition + 'shouldDeposit', function () {
                                  shouldDeposit(defaultCheque.amount, issuer)
                                  describe(describePreCondition + 'shouldCashCheque', function () {
                                    const requestPayout = defaultCheque.amount
                                    shouldCashCheque(defaultCheque.beneficiary, recipient, requestPayout, calleePayout, sender)
                                    describe(describeTest + 'shouldNotCashCheque', function () {
                                      const toSignFields = {
                                        requestPayout,
                                        recipient,
                                        calleePayout,
                                      }
                                      const toSubmitFields = Object.assign({}, toSignFields, { beneficiary: defaultCheque.beneficiary })
                                      shouldNotCashCheque(toSignFields, toSubmitFields, value, sender, signee, "SimpleSwap: not enough balance owed")
                                    })
                                  })
                                })
                              })
                            })
                          })
                        })
                        context("when we don't submit a cheque beforeHand", function () {
                          const requestPayout = defaultCheque.timeout.add(new BN(1))
                          describe(describeTest + 'shouldNotCashCheque', function () {
                            const toSignFields = {
                              requestPayout,
                              recipient,
                              calleePayout,
                            }
                            const toSubmitFields = Object.assign({}, toSignFields, { beneficiary: defaultCheque.beneficiary })
                            shouldNotCashCheque(toSignFields, toSubmitFields, value, sender, signee, "SimpleSwap: not enough balance owed")
                          })
                        })
                      })
                      context('when the recipient is the beneficiary', function () {
                        const recipient = defaultCheque.beneficiary
                        describe(describePreCondition + 'shouldDeposit', function () {
                          shouldDeposit(defaultCheque.amount, issuer)
                          describe(describePreCondition + 'shouldSubmitChequeBeneficiary', function () {
                            shouldSubmitChequeBeneficiary(defaultCheque, defaultCheque.beneficiary)
                            describe(describeTest + 'shouldCashCheque', function () {
                              const waitTime = defaultCheque.timeout.add(new BN(100))

                              beforeEach(async function () {
                                await time.increase(waitTime)
                              })
                              shouldCashCheque(defaultCheque.beneficiary, recipient, defaultCheque.amount, new BN(1), alice)
                            })
                          })
                        })
                      })
                    })
                  })
                  context('when the beneficiary does not sign the correct fields', function () {
                    const sender = alice
                    const recipient = defaultCheque.beneficiary
                    describe(describePreCondition + 'shouldDeposit', function () {
                      shouldDeposit(defaultCheque.amount, issuer)
                      describe(describePreCondition + 'shouldSubmitChequeBeneficiary', function () {
                        shouldSubmitChequeBeneficiary(defaultCheque, defaultCheque.beneficiary)

                        describe(describeTest + 'shouldNotCashCheque', function () {
                          const waitTime = defaultCheque.timeout.add(new BN(100))
                          beforeEach(async function () {
                            await time.increase(waitTime)
                          })
                          const toSignFields = {
                            requestPayout: new BN(0),
                            recipient,
                            calleePayout,
                          }
                          const toSubmitFields = Object.assign({}, toSignFields, { beneficiary: defaultCheque.beneficiary, requestPayout: new BN(1) })
                          shouldNotCashCheque(toSignFields, toSubmitFields, value, sender, signee, "SimpleSwap: invalid beneficiarySig")
                        })
                      })
                    })
                  })
                })
                context('when the beneficiary is not a signee', function () {
                  const sender = alice
                  const waitTime = defaultCheque.timeout.add(new BN(100))
                  beforeEach(async function () {
                    await time.increase(waitTime)
                  })
                  const recipient = defaultCheque.beneficiary
                  describe(describePreCondition + 'shouldDeposit', function () {
                    shouldDeposit(defaultCheque.amount, issuer)
                    describe(describePreCondition + 'shouldSubmitChequeBeneficiary', function () {
                      shouldSubmitChequeBeneficiary(defaultCheque, defaultCheque.beneficiary)
                      const signee = alice
                      describe(describeTest + 'shouldNotCashCheque', function () {
                        const toSignFields = {
                          requestPayout: new BN(0),
                          recipient,
                          calleePayout,
                        }
                        const toSubmitFields = Object.assign({}, toSignFields, { beneficiary: defaultCheque.beneficiary })
                        shouldNotCashCheque(toSignFields, toSubmitFields, value, sender, signee, "SimpleSwap: invalid beneficiarySig")
                      })
                    })
                  })
                })
              })
            })
            context('when the calleePayout is zero', function () {
              beforeEach(async function () {
                this.expiry = (await time.latest()).add(time.duration.years(1))
              })
              const calleePayout = new BN(0)
              describe(describePreCondition + 'shouldDeposit', function () {
                shouldDeposit(defaultCheque.amount, issuer)
                describe(describePreCondition + 'shouldSubmitChequeBeneficiary', function () {
                  shouldSubmitChequeBeneficiary(defaultCheque, defaultCheque.beneficiary)
                  describe(describeTest + 'shouldCashCheque', function () {
                    const waitTime = defaultCheque.timeout.add(new BN(100))

                    beforeEach(async function () {
                      await time.increase(waitTime)
                    })
                    shouldCashCheque(defaultCheque.beneficiary, _recipient, defaultCheque.amount, calleePayout, alice)
                  })
                })
              })
            })
          })
          context('when the signature has expired', async function () {
            const recipient = _recipient
            const calleePayout = new BN(0)
            const sender = alice
            const signee = defaultCheque.beneficiary

            beforeEach(async function () {
              this.expiry = (await time.latest()).sub(time.duration.days(new BN(1)))
            })
            describe(describePreCondition + 'shouldDeposit', function () {
              shouldDeposit(defaultCheque.amount, issuer)
              describe(describePreCondition + 'shouldSubmitChequeBeneficiary', function () {
                shouldSubmitChequeBeneficiary(defaultCheque, defaultCheque.beneficiary)
                describe(describeTest + 'shouldNotCashCheque', function () {
                  const toSignFields = {
                    requestPayout: new BN(0),
                    recipient,
                    calleePayout,
                  }
                  const toSubmitFields = Object.assign({}, toSignFields, { beneficiary: defaultCheque.beneficiary })
                  shouldNotCashCheque(toSignFields, toSubmitFields, value, sender, signee, "SimpleSwap: beneficiarySig expired")
                })
              })
            })
          })
        })
        context('when we send value along', function () {
          beforeEach(async function () {
            this.expiry = (await time.latest()).sub(time.duration.days(new BN(1)))
          })
          const value = new BN(1)
          const recipient = _recipient
          const calleePayout = new BN(0)
          const sender = alice
          const signee = defaultCheque.beneficiary
          describe(describeTest + 'shouldNotCashCheque', function () {
            const toSignFields = {
              requestPayout: new BN(0),
              recipient,
              calleePayout,
            }
            const toSubmitFields = Object.assign({}, toSignFields, { beneficiary: defaultCheque.beneficiary })
            shouldNotCashCheque(toSignFields, toSubmitFields, value, sender, signee, "revert")
          })
        })
      }
    })

    describe(describeFunction + 'cashChequeBeneficiary', function () {
      if (enabledTests.cashChequeBeneficiary) {
        context("when we don't send value along", function () {
          const value = new BN(0)
          context('when the sender is the beneficiary', function () {
            const sender = defaultCheque.beneficiary
            context('when we submit the cheque beforeHand', function () {
              describe(describePreCondition + 'shouldSubmitChequeBeneficiary', function () {
                shouldSubmitChequeBeneficiary(defaultCheque, defaultCheque.beneficiary)
                context('when the recipient is not the beneficiary', function () {
                  const recipient = _recipient
                  context('when we have not cashed a cheque before', function () {
                    context('when we have waited more than timeoutDuration', function () {
                      const waitTime = defaultCheque.timeout.add(new BN(100))
                      beforeEach(async function () {
                        await time.increase(waitTime)
                      })
                      context('when the requestPayout is equal to the submitted value', function () {
                        const requestPayout = defaultCheque.amount
                        context('when there is some balance', function () {
                          context('when the balance is bigger than the requestPayout', function () {
                            describe(describePreCondition + 'shouldDeposit', function () {
                              const depositAmount = requestPayout.add(new BN(50))
                              shouldDeposit(depositAmount, issuer)
                              context('when there are hardDeposits', function () {
                                describe('when the hardDeposits are assigned to the sender', function () {
                                  const hardDepositReceiver = sender
                                  context('when the hardDeposit is more the requestPayout', function () {
                                    const hardDeposit = requestPayout.add(new BN(1))
                                    describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                                      shouldIncreaseHardDeposit(hardDepositReceiver, hardDeposit, issuer)
                                      describe(describeTest + 'shouldCashChequeBeneficiary', function () {
                                        shouldCashChequeBeneficiary(recipient, requestPayout, sender)
                                      })
                                    })
                                  })
                                  context('when the hardDeposit equals the requestPayout', function () {
                                    const hardDeposit = requestPayout
                                    describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                                      shouldIncreaseHardDeposit(hardDepositReceiver, hardDeposit, issuer)
                                      describe(describeTest + 'shouldCashChequeBeneficiary', function () {
                                        shouldCashChequeBeneficiary(recipient, requestPayout, sender)
                                      })
                                    })
                                  })
                                  context('when the hardDeposit is less than the requestPayout', function () {
                                    const hardDeposit = requestPayout.sub(new BN(1))
                                    describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                                      shouldIncreaseHardDeposit(hardDepositReceiver, hardDeposit, issuer)
                                      describe(describeTest + 'shouldCashChequeBeneficiary', function () {
                                        shouldCashChequeBeneficiary(recipient, requestPayout, sender)
                                      })
                                    })
                                  })
                                })
                                describe('when the hardDeposits are assigned to somebody else', function () {
                                  const hardDepositReceiver = alice
                                  const hardDeposit = requestPayout.add(new BN(1))
                                  describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                                    shouldIncreaseHardDeposit(hardDepositReceiver, hardDeposit, issuer)
                                    describe(describeTest + 'shouldCashChequeBeneficiary', function () {
                                      shouldCashChequeBeneficiary(recipient, requestPayout, sender)
                                    })
                                  })
                                })
                              })
                            })
                          })
                          context('when the balance equals the requestPayout', function () {
                            describe(describePreCondition + 'shouldDeposit', function () {
                              const depositAmount = requestPayout
                              shouldDeposit(depositAmount, issuer)
                              describe(describeTest + 'shouldCashChequeBeneficiary', function () {
                                shouldCashChequeBeneficiary(recipient, requestPayout, sender)
                              })
                            })
                          })
                        })
                        context('when there is no balance', function () {
                          describe(describeTest + 'shouldCashChequeBeneficiary', function () {
                            shouldCashChequeBeneficiary(recipient, requestPayout, sender)
                          })
                        })
                      })
                      context('when the requestPayout is less than the submitted value', function () {
                        const requestPayout = defaultCheque.amount.sub(new BN(1))
                        describe(describeTest + 'shouldCashChequeBeneficiary', function () {
                          shouldCashChequeBeneficiary(recipient, requestPayout, sender)
                        })
                      })
                    })
                  })
                  context('when we have waited timeoutDuration', function () {
                    const waitTime = defaultCheque.timeout
                    const requestPayout = defaultCheque.amount
                    beforeEach(async function () {
                      await time.increase(waitTime)
                    })
                    describe(describeTest + 'shouldCashChequeBeneficiary', function () {
                      shouldCashChequeBeneficiary(recipient, requestPayout, sender)
                    })
                  })
                  context('when we have cashed a cheque before', function () {
                    describe('when we have cashed the partial amount before', function () {
                      const waitTime = defaultCheque.timeout.add(new BN(100))
                      beforeEach(async function () {
                        await time.increase(waitTime)
                      })
                      describe(describePreCondition + 'shouldCashChequeBeneficiary', function () {
                        const requestPayout = defaultCheque.amount.div(new BN(2))
                        shouldCashChequeBeneficiary(recipient, requestPayout, sender)
                        describe(describeTest + 'shouldCashChequeBeneficiary', function () {
                          shouldCashChequeBeneficiary(recipient, requestPayout, sender)
                        })
                      })
                    })
                    describe('when we have cashed the full amount before', function () {
                      const waitTime = defaultCheque.timeout.add(new BN(100))
                      beforeEach(async function () {
                        await time.increase(waitTime)
                      })
                      describe(describePreCondition + 'shouldDeposit', function () {
                        shouldDeposit(defaultCheque.amount, issuer)
                        describe(describePreCondition + 'shouldCashChequeBeneficiary', function () {
                          const requestPayout = defaultCheque.amount
                          shouldCashChequeBeneficiary(recipient, requestPayout, sender)
                          describe(describeTest + 'shouldNotCashChequeBeneficiary', function () {
                            shouldNotCashChequeBeneficiary(recipient, requestPayout, sender, value, "SimpleSwap: not enough balance owed")
                          })
                        })
                      })
                    })
                  })
                })
                context('when the recipient is the beneficiary', function () {
                  const recipient = defaultCheque.beneficiary
                  const waitTime = defaultCheque.timeout.add(new BN(100))
                  beforeEach(async function () {
                    await time.increase(waitTime)
                  })
                  const requestPayout = defaultCheque.amount
                  describe(describeTest + 'shouldCashChequeBeneficiary', function () {
                    shouldCashChequeBeneficiary(recipient, requestPayout, sender)
                  })
                })
              })
            })
            context("when we don't submit a cheque beforeHand", function () {
              describe(describeTest + 'shouldNotCashChequeBeneficiary', function () {
                shouldNotCashChequeBeneficiary(_recipient, defaultCheque.amount, sender, value, "SimpleSwap: not enough balance owed")
              })
            })
          })
          context('when the sender is not the beneficiary', function () {
            const sender = alice
            describe(describePreCondition, 'shouldDeposit', function () {
              shouldDeposit(defaultCheque.amount, issuer)
              describe(describePreCondition + 'shouldSubmitChequeBeneficiary', function () {
                shouldSubmitChequeBeneficiary(defaultCheque, defaultCheque.beneficiary)
                describe(describeTest + 'shouldNotCashChequeBeneficiary', function () {
                  shouldNotCashChequeBeneficiary(recipient, defaultCheque.amount, sender, value, "SimpleSwap: not enough balance owed")
                })
              })
            })
          })
          context('when we send value along', function () {
            const sender = alice
            const value = new BN(1)
            describe(describePreCondition + 'shouldDeposit', function () {
              shouldDeposit(defaultCheque.amount, issuer)
              describe(describePreCondition + 'shouldSubmitChequeBeneficiary', function () {
                shouldSubmitChequeBeneficiary(defaultCheque, defaultCheque.beneficiary)
                describe(describeTest + 'shouldNotCashChequeBeneficiary', function () {
                  shouldNotCashChequeBeneficiary(_recipient, defaultCheque.amount, sender, value, "revert")
                })
              })
            })
          })
        })
      }
    })

    describe(describeFunction + 'prepareDecreaseHardDeposit', function () {
      if (enabledTests.prepareDecreaseHardDeposit) {
        const beneficiary = defaultCheque.beneficiary
        context("when we don't send value along", function () {
          const value = new BN(0)
          context('when there are hardDeposits', function () {
            const hardDepositAmount = new BN(50)
            describe(describePreCondition + 'shouldDeposit', function () {
              shouldDeposit(hardDepositAmount, issuer)
              describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                shouldIncreaseHardDeposit(beneficiary, hardDepositAmount, issuer)
                context('when the sender is the issuer', function () {
                  const sender = issuer
                  context('when the decreaseAmount is the hardDepositAmount', function () {
                    const decreaseAmount = hardDepositAmount
                    context('when we have set a custom decreaseTimeout', function () {
                      describe(describePreCondition + 'shouldSetCustomHardDepositDecreaseTimeout', function () {
                        const customTimeout = new BN(10)
                        shouldSetCustomHardDepositDecreaseTimeout(beneficiary, customTimeout, issuer)
                        context('when we have not set a custom decreaseTimeout', function () {
                          describe(describeTest + 'prepareDecreaseHardDeposit', function () {
                            shouldPrepareDecreaseHardDeposit(beneficiary, decreaseAmount, sender)
                          })
                        })
                      })
                    })
                    context('when we have not set a custom decreaseTimeout', function () {
                      describe(describeTest + 'prepareDecreaseHardDeposit', function () {
                        shouldPrepareDecreaseHardDeposit(beneficiary, decreaseAmount, sender)
                      })
                    })
                  })
                  context('when the decreaseAmount is less than the hardDepositAmount', function () {
                    const decreaseAmount = hardDepositAmount.div(new BN(2))
                    describe(describeTest + 'prepareDecreaseHardDeposit', function () {
                      shouldPrepareDecreaseHardDeposit(beneficiary, decreaseAmount, sender)
                    })
                  })
                  context('when the decreaseAmount is higher than the hardDepositAmount', function () {
                    const decreaseAmount = hardDepositAmount.add(new BN(1))
                    const revertMessage = "SimpleSwap: hard deposit not sufficient"
                    describe(describeTest + 'shouldNotPrepareDecreaseHardDeposit', function () {
                      shouldNotPrepareDecreaseHardDeposit(beneficiary, decreaseAmount, sender, value, revertMessage)
                    })
                  })
                })
                context('when the sender is the issuer', function () {
                  const sender = alice
                  const revertMessage = "SimpleSwap: not issuer"
                  const decreaseAmount = hardDepositAmount
                  describe(describeTest + 'shouldNotPrepareDecreaseHardDeposit', function () {
                    shouldNotPrepareDecreaseHardDeposit(beneficiary, decreaseAmount, sender, value, revertMessage)
                  })
                })
              })
              context('when there are no hardDeposits', function () {
                const sender = issuer
                const revertMessage = "SimpleSwap: hard deposit not sufficient"
                const decreaseAmount = new BN(50)
                describe(describeTest + 'shouldNotPrepareDecreaseHardDeposit', function () {
                  shouldNotPrepareDecreaseHardDeposit(beneficiary, decreaseAmount, sender, value, revertMessage)
                })
              })
            })
          })
        })
        context('when we send value along', function () {
          const value = new BN(1)
          const sender = issuer
          const revertMessage = "revert"
          const decreaseAmount = new BN(50)
          describe(describeTest + 'shouldNotPrepareDecreaseHardDeposit', function () {
            shouldNotPrepareDecreaseHardDeposit(beneficiary, decreaseAmount, sender, value, revertMessage)
          })
        })
      }
    })

    describe(describeFunction + 'decreaseHardDeposit', function () {
      if (enabledTests.decreaseHardDeposit) {
        const beneficiary = defaultCheque.beneficiary
        context("when we don't send value along", function () {
          const value = new BN(0)
          context('when the sender is the issuer', function () {
            const sender = issuer
            context("when we have prepared the decreaseHardDeposit", function () {
              const hardDeposit = new BN(50)
              describe(describePreCondition + "shouldDeposit", function () {
                shouldDeposit(hardDeposit, issuer)
                describe(describePreCondition + "shouldIncreaseHardDeposit", function () {
                  shouldIncreaseHardDeposit(beneficiary, hardDeposit, issuer)
                  describe(describePreCondition + "shouldPrepareDecreaseHardDeposit", function () {
                    shouldPrepareDecreaseHardDeposit(beneficiary, hardDeposit, issuer)
                    context('when we have waited more than decreaseTimeout time', function () {
                      beforeEach(async function () {
                        await time.increase(await this.simpleSwap.DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT())
                      })
                      describe(describeTest + 'shouldDecreaseHardDeposit', function () {
                        shouldDecreaseHardDeposit(beneficiary, sender)
                      })
                    })
                    context('when we have not waited more than decreaseTimeout time', function () {
                      describe(describeTest + 'shouldNotDecreaseHardDeposit', function () {
                        const revertMessage = "SimpleSwap: deposit not yet timed out"
                        shouldNotDecreaseHardDeposit(beneficiary, sender, value, revertMessage)
                      })
                    })
                  })
                })
              })
            })
            context('when we have not prepared the decreaseHardDeposit', function () {
              describe(describeTest + 'shouldNotDecreaseHardDeposit', function () {
                const revertMessage = "SimpleSwap: deposit not yet timed out"
                shouldNotDecreaseHardDeposit(beneficiary, sender, value, revertMessage)
              })
            })
          })
        })
        context("when we send value along", function () {
          const value = new BN(1)
          const sender = issuer
          describe(describeTest + 'shouldNotDecreaseHardDeposit', function () {
            const revertMessage = "revert"
            shouldNotDecreaseHardDeposit(beneficiary, sender, value, revertMessage)
          })

        })
      }
    })

    describe(describeFunction + 'increaseHardDeposit', function () {
      if (enabledTests.increaseHardDeposit) {
        const hardDepositIncrease = new BN(50)
        const beneficiary = defaultCheque.beneficiary
        context("when we don't send value along", function () {
          const value = new BN(0)
          context('when the sender is the issuer', function () {
            const sender = issuer
            context('when there is more liquidBalance than the requested hardDepositIncrease', function () {
              const deposit = hardDepositIncrease.mul(new BN(2))
              describe(describePreCondition + 'shouldDeposit', function () {
                shouldDeposit(deposit, issuer)
                context('when we have set a customHardDepositDecreaseTimeout', function () {
                  const customHardDepositDecreaseTimeout = new BN(60)
                  describe(describePreCondition + 'shouldSetCustomHardDepositDecreaseTimeout', function () {
                    shouldSetCustomHardDepositDecreaseTimeout(beneficiary, customHardDepositDecreaseTimeout, issuer)
                    describe(describeTest + 'shouldIncreaseHardDeposit', function () {
                      shouldIncreaseHardDeposit(beneficiary, hardDepositIncrease, sender)
                    })
                  })
                })
              })
            })
            context('when there is as much liquidBalance as the requested hardDepositIncrease', function () {
              const deposit = hardDepositIncrease
              describe(describePreCondition + 'shouldDeposit', function () {
                shouldDeposit(deposit, issuer)
                describe(describeTest + 'shouldIncreaseHardDeposit', function () {
                  shouldIncreaseHardDeposit(beneficiary, hardDepositIncrease, sender)
                })
              })
            })
            context('when the liquidBalance is less than the requested hardDepositIncrease', function () {
              describe(describeTest + 'shouldNotIncreaseHardDeposit', function () {
                const revertMessage = "SimpleSwap: hard deposit cannot be more than balance"
                shouldNotIncreaseHardDeposit(beneficiary, hardDepositIncrease, sender, value, revertMessage)
              })
            })
          })
          context('when the sender is not the issuer', function () {
            const sender = alice
            describe(describeTest + 'shouldNotIncreaseHardDeposit', function () {
              const revertMessage = "SimpleSwap: not issuer"
              shouldNotIncreaseHardDeposit(beneficiary, hardDepositIncrease, sender, value, revertMessage)
            })
          })
        })
        context('when we send value along', function () {
          const value = new BN(1)
          const hardDepositIncrease = new BN(50)
          const beneficiary = defaultCheque.beneficiary
          const sender = issuer
          describe(describeTest + 'shouldNotIncreaseHardDeposit', function () {
            const revertMessage = "revert"
            shouldNotIncreaseHardDeposit(beneficiary, hardDepositIncrease, sender, value, revertMessage)
          })
        })
      }
    })

    describe(describeFunction + 'setCustomHardDepositDecreaseTimeout', function () {
      if (enabledTests.setCustomHardDepositDecreaseTimeout) {
        const beneficiary = defaultCheque.beneficiary
        const decreaseTimeout = new BN(60)
        context("when we don't send value along", function () {
          const value = new BN(0)
          context('when the sender is the issuer', function () {
            const sender = issuer
            context('when the beneficiary is a signee', function () {
              const signee = beneficiary
              context('when the beneficiary signs the correct fields', function () {
                describe(describeTest + 'shouldSetCustomHardDepositDecreaseTimeout', function () {
                  shouldSetCustomHardDepositDecreaseTimeout(beneficiary, decreaseTimeout, sender)
                })
              })
              context('when the beneficiary does not sign the correct fields', function () {
                describe(describeTest + 'shouldNotSetCustomHardDepositDecreaseTimeout', function () {
                  const toSubmit = { beneficiary, decreaseTimeout }
                  const toSign = { beneficiary, decreaseTimeout: decreaseTimeout.sub(new BN(1)) }
                  const revertMessage = "SimpleSwap: invalid beneficiarySig"
                  shouldNotSetCustomHardDepositDecreaseTimeout(toSubmit, toSign, signee, sender, value, revertMessage)
                })
              })
            })
            context('when the beneficiary is not a signee', function () {
              const signee = alice
              describe(describeTest + 'shouldNotSetCustomHardDepositDecreaseTimeout', function () {
                const toSubmit = { beneficiary, decreaseTimeout }
                const toSign = toSubmit
                const revertMessage = "SimpleSwap: invalid beneficiarySig"
                shouldNotSetCustomHardDepositDecreaseTimeout(toSubmit, toSign, signee, sender, value, revertMessage)
              })
            })
          })
          context('when the sender is not the issuer', function () {
            const sender = alice
            describe(describeTest + 'shouldNotSetCustomHardDepositDecreaseTimeout', function () {
              const toSubmit = { beneficiary, decreaseTimeout }
              const toSign = toSubmit
              const signee = beneficiary
              const revertMessage = "SimpleSwap: not issuer"
              shouldNotSetCustomHardDepositDecreaseTimeout(toSubmit, toSign, signee, sender, value, revertMessage)
            })
          })
        })
        context('when we send value along', function () {
          const value = new BN(1)
          const sender = issuer
          const signee = beneficiary
          describe(describeTest + 'shouldNotSetCustomHardDepositDecreaseTimeout', function () {
            const toSubmit = { beneficiary, decreaseTimeout }
            const toSign = toSubmit
            const revertMessage = "revert"
            shouldNotSetCustomHardDepositDecreaseTimeout(toSubmit, toSign, signee, sender, value, revertMessage)
          })
        })
      }
    })

    describe(describeFunction + 'withdraw', function () {
      if (enabledTests.withdraw) {
        const withdrawAmount = new BN(50)
        context("when we don't send value along", function () {
          const value = new BN(0)
          context('when the sender is the issuer', function () {
            const sender = issuer
            context('when the liquidBalance is more than the withdrawAmount', function () {
              const depositAmount = withdrawAmount.mul(new BN(2))
              describe(describePreCondition + 'shouldDeposit', function () {
                shouldDeposit(depositAmount, issuer)
                describe(describeTest + 'shouldWithdraw', function () {
                  shouldWithdraw(withdrawAmount, sender)
                })
              })
            })
            context('when the liquidBalance equals the withdrawAount', function () {
              const depositAmount = withdrawAmount
              describe(describePreCondition + 'shouldDeposit', function () {
                shouldDeposit(depositAmount, issuer)
                describe(describeTest + 'shouldWithdraw', function () {
                  shouldWithdraw(withdrawAmount, sender)
                })
              })
            })
            context('when the liquidBalance is less than the withdrawAmount', function () {
              const revertMessage = "SimpleSwap: liquidBalance not sufficient"
              shouldNotWithdraw(withdrawAmount, sender, value, revertMessage)
            })
          })
          context('when the sender is not the issuer', function () {
            const sender = alice
            const revertMessage = "SimpleSwap: not issuer"
            shouldNotWithdraw(withdrawAmount, sender, value, revertMessage)
          })
        })
        context('when we send value along', function () {
          const value = new BN(1)
          const sender = alice
          const revertMessage = "revert"
          shouldNotWithdraw(withdrawAmount, sender, value, revertMessage)
        })
      }
    })

    describe(describeFunction + 'deposit', function () {
      if (enabledTests.deposit) {
        context('when the depositAmount is not zero', function() {
          const depositAmount = new BN(1)
          describe(describeTest + 'shouldDeposit', function() {
            shouldDeposit(depositAmount, issuer)
          })
        })
        context('when the depositAmount is zero', function() {
          const depositAmount = new BN(0)
          describe(describeTest + 'shouldDeposit', function() {
            shouldNotDeposit(depositAmount, issuer)
          })
        })
      }
    })
  })
}

module.exports = {
  shouldBehaveLikeSimpleSwap
};