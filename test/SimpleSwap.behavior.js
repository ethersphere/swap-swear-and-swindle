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
  shouldReturnDEFAULT_HARDDEPPOSIT_DECREASE_TIMEOUT,
  shouldReturnCheques,
  shouldReturnHardDeposits,
  shouldReturnTotalHardDeposit,
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
  shouldDeposit,
  shouldNotDeposit
} = require('./SimpleSwap.should.js')

// switch to false if you don't want to test the particular function
enabledTests = {
  DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT: true,
  cheques: true,
  hardDeposits: true,
  totalHardDeposit: true,
  issuer: true,
  liquidBalance: true,
  liquidBalanceFor: true,
  submitChequeIssuer: true,
  submitChequeBeneficiary: true,
  submitCheque: true,
  cashChequeBeneficiary: true,
  cashCheque: true,
  prepareDecreaseHardDeposit: true,
  decreaseHardDeposit: true,
  increaseHardDeposit: true,
  setCustomHardDepositDecreaseTimeout: true,
  withdraw: true, 
  deposit: true
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
function shouldBehaveLikeSimpleSwap([issuer, alice, bob, agent], DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT) {
  const defaultCheque = {
    beneficiary: bob,
    serial: new BN(3),
    amount: new BN(500),
    timeout: new BN(86400),
    signee: issuer,
    signature: ""
  }
  context('as a simple swap', function () {
    describe(describeFunction + 'DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT', function () {
      if (enabledTests.DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT) {
        shouldReturnDEFAULT_HARDDEPPOSIT_DECREASE_TIMEOUT(DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT)
      }
    })

    describe(describeFunction + 'cheques', function () {
      if (enabledTests.cheques) {
        const beneficiary = defaultCheque.beneficiary
        context('when no cheque was ever submitted', function () {
          describe(describeTest + 'shouldReturnCheques', function () {
            const expectedSerial = new BN(0)
            const expectedAmount = new BN(0)
            const expectedPaidOut = new BN(0)
            const expectedCashTimeout = new BN(0)
            shouldReturnCheques(beneficiary, expectedSerial, expectedAmount, expectedPaidOut, expectedCashTimeout)
          })
        })
        context('when a cheque was submitted but not cashed', function () {
          const toSubmitCheque = defaultCheque
          describe(describePreCondition + 'shouldSubmitChequeBeneficiary', function () {
            shouldSubmitChequeBeneficiary(toSubmitCheque, defaultCheque.beneficiary)
            describe(describeTest + 'shouldReturnCheques', function () {
              const expectedSerial = toSubmitCheque.serial
              const expectedAmount = toSubmitCheque.amount
              const expectedPaidOut = new BN(0)
              const expectedCashTimeout = toSubmitCheque.timeout
              shouldReturnCheques(beneficiary, expectedSerial, expectedAmount, expectedPaidOut, expectedCashTimeout)
            })
            context('when a cheque was cashed', function () {
              describe(describePreCondition + 'shouldDeposit', function() {
                shouldDeposit(defaultCheque.amount, issuer)
                describe(describePreCondition + 'shouldCashChequeBeneficiary', function () {
                  beforeEach(async function() {
                    await time.increase(toSubmitCheque.timeout)
                  })
                  shouldCashChequeBeneficiary(defaultCheque.beneficiary, defaultCheque.amount, defaultCheque.beneficiary)
                  describe(describeTest + 'shouldReturnCheques', function () {
                    const expectedSerial = toSubmitCheque.serial
                    const expectedAmount = toSubmitCheque.amount
                    const expectedPaidOut = toSubmitCheque.amount
                    const expectedCashTimeout = toSubmitCheque.timeout
                    shouldReturnCheques(beneficiary, expectedSerial, expectedAmount, expectedPaidOut, expectedCashTimeout)
                  })
                })
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

    describe(describeFunction + 'liquidBalanceFor', function () {
      if (enabledTests.liquidBalanceFor) {
        const beneficiary = bob
        const depositAmount = new BN(50)
        context('when there is some balance', function () {
          describe(describePreCondition + 'shoulDeposit', function () {
            shouldDeposit(depositAmount, issuer)
            context('when there are no hard deposits', function () {
              describe(describeTest + 'shouldReturnLiquidBalanceFor', function () {
                shouldReturnLiquidBalanceFor(beneficiary, depositAmount)
              })
            })
            context('when there are no hard deposits', function () {
              const hardDeposit = new BN(10)
              describe('when these hard deposits are assigned to the beneficiary', function () {
                describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                  shouldIncreaseHardDeposit(beneficiary, hardDeposit, issuer)
                  describe(describeTest + 'shouldReturnLiquidBalanceFor', function () {
                    shouldReturnLiquidBalanceFor(beneficiary, depositAmount)
                  })
                })
              })
              describe('when these hard deposits are assigned to somebody else', function () {
                describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                  shouldIncreaseHardDeposit(alice, hardDeposit, issuer)
                  describe(describeTest + 'shouldReturnLiquidBalanceFor', function () {
                    shouldReturnLiquidBalanceFor(beneficiary, depositAmount.sub(hardDeposit))
                  })
                })
              })
            })
          })
        })
        describe('when there is no balance', function () {
          shouldReturnLiquidBalanceFor(beneficiary, new BN(0))
        })
      }
    })

    describe(describeFunction + 'submitChequeIssuer', function () {
      if (enabledTests.submitChequeIssuer) {
        context("when we don't send value along", function () {
          const value = new BN(0)
          context('when the sender is the issuer', function () {
            const sender = issuer
            context('when the first serial is higher than 0', function () {
              expect(defaultCheque.serial).bignumber.to.be.above(new BN(0), "Serial of defaultCheque not above 0")
              context('when the first serial is below MAX_UINT256', function () {
                expect(defaultCheque.serial).bignumber.to.be.below(constants.MAX_UINT256, "Serial of defaultCheque not above 0")
                context('when the beneficiary is a signee', function () {
                  const unsignedCheque = Object.assign({}, defaultCheque, { signee: defaultCheque.beneficiary })
                  expect(unsignedCheque.signee).to.be.equal(unsignedCheque.beneficiary, "Signee of unsignedCheque is not beneficiary")
                  context('when the signee signs the correct fields', function () {
                    context('when we send one cheque', function () {
                      context('when there is a liquidBalance to cover the cheque', function () {
                        describe(describePreCondition + 'shouldDeposit', function () {
                          describe(describePreCondition + 'shouldDeposit', function () {
                            shouldDeposit(unsignedCheque.amount.add(unsignedCheque.amount), issuer)
                            describe(describeTest + 'submitChequeIssuer', function () {
                              shouldSubmitChequeIssuer(unsignedCheque, sender)
                            })
                          })
                        })
                      })
                      context('when there is no liquidBalance to cover the cheque', function () {
                        describe(describeTest + 'shouldSubmitChequeIssuer', function () {
                          shouldSubmitChequeIssuer(unsignedCheque, sender)
                        })
                      })
                    })
                    context('when we send more than one cheque', async function () {
                      describe(describePreCondition + 'shouldSubmitChequeIssuer', function () {
                        shouldSubmitChequeIssuer(unsignedCheque, sender)
                        context('when the serial number is increasing', function () {
                          describe(describeTest + 'shouldSubmitChequeIssuer', function () {
                            const secondSerial = unsignedCheque.serial.add(new BN(1))
                            const increasing_serial_unsignedCheque = Object.assign({}, defaultCheque, { serial: secondSerial, signee: defaultCheque.beneficiary })
                            shouldSubmitChequeIssuer(increasing_serial_unsignedCheque, sender)
                          })
                        })
                        context('when the serial number stays the same', function () {
                          context('when the serial number is increasing', function () {
                            describe(describeTest + 'shouldSubmitChequeIssuer', function () {
                              const secondSerial = unsignedCheque.serial
                              const same_serial_unsignedCheque = Object.assign({}, defaultCheque, { serial: secondSerial, signee: defaultCheque.beneficiary })
                              shouldNotSubmitChequeIssuer(same_serial_unsignedCheque, same_serial_unsignedCheque, sender, value, "SimpleSwap: invalid serial")
                            })
                          })
                          context('when the serial number is decreasing', function () {
                            context('when the serial number is increasing', function () {
                              describe(describeTest + 'shouldSubmitChequeIssuer', function () {
                                const secondSerial = unsignedCheque.serial.sub(new BN(1))
                                const decreasing_serial_unsignedCheque = Object.assign({}, defaultCheque, { serial: secondSerial, signee: defaultCheque.beneficiary })
                                shouldNotSubmitChequeIssuer(decreasing_serial_unsignedCheque, decreasing_serial_unsignedCheque, sender, value, "SimpleSwap: invalid serial")
                              })
                            })
                          })
                        })
                      })
                    })
                  })
                  context('when the signee does not sign the correct fields', function () {
                    describe(describeTest + 'shouldSubmitChequeIssuer', function () {
                      const wrongBeneficiary = constants.ZERO_ADDRESS
                      const wrong_beneficiary_unsignedCheque = Object.assign({}, defaultCheque, { beneficiary: wrongBeneficiary, signee: defaultCheque.beneficiary })
                      shouldNotSubmitChequeIssuer(wrong_beneficiary_unsignedCheque, defaultCheque, sender, value, "SimpleSwap: invalid beneficiarySig")
                    })
                  })
                })
                context('when the beneficiary is not the signee', function () {
                  describe(describeTest + 'shouldSubmitChequeIssuer', function () {
                    const signee = alice
                    const wrong_signee_unsignedCheque = Object.assign({}, defaultCheque, { signee: signee })
                    shouldNotSubmitChequeIssuer(wrong_signee_unsignedCheque, defaultCheque, sender, value, "SimpleSwap: invalid beneficiarySig")
                  })
                })
              })
              context('when the first serial is at MAX_UINT256', function () {
                describe(describeTest + 'shouldSubmitChequeIssuer', function () {
                  const MAX_UINT256_unsignedCheque = Object.assign({}, defaultCheque, { serial: constants.MAX_UINT256, signee: defaultCheque.beneficiary })
                  shouldSubmitChequeIssuer(MAX_UINT256_unsignedCheque, issuer)
                  describe('when we want to submit a cheque afterwards', function () {
                    describe(describeTest + 'shouldNotSubmitChequeIssuer', function () {
                      const MAX_UINT256_wrap_unsignedCheque = Object.assign({}, defaultCheque, { serial: MAX_UINT256_unsignedCheque.serial.add(new BN(1)), signee: defaultCheque.beneficiary })
                      shouldNotSubmitChequeIssuer(MAX_UINT256_wrap_unsignedCheque, MAX_UINT256_wrap_unsignedCheque, sender, value, "SimpleSwap: invalid serial")
                    })
                  })
                })
              })
            })
            context('when the serial is 0', function () {
              describe(describeTest + 'shouldNotSubmitChequeIssuer', function () {
                const serial = new BN(0)
                const zero_serial_unsignedCheque = Object.assign({}, defaultCheque, { serial: serial, signee: defaultCheque.beneficiary })
                shouldNotSubmitChequeIssuer(zero_serial_unsignedCheque, zero_serial_unsignedCheque, sender, value, "SimpleSwap: invalid serial")
              })
            })
          })
          context('when the sender is not the issuer', function () {
            describe(describeTest + 'shouldNotSubmitChequeIssuer', function () {
              const sender = alice
              const unsignedCheque = Object.assign({}, defaultCheque, { signee: defaultCheque.beneficiary })
              shouldNotSubmitChequeIssuer(unsignedCheque, unsignedCheque, sender, value, "SimpleSwap: not issuer")
            })
          })
        })
        context("when we send value along", function () {
          describe(describeTest + 'shouldNotSubmitChequeIssuer', function () {
            const value = new BN(0)
            const sender = issuer
            const unsignedCheque = Object.assign({}, defaultCheque, { signee: defaultCheque.beneficiary })
            shouldNotSubmitChequeBeneficiary(unsignedCheque, unsignedCheque, sender, value, "revert")
          })
        })
      }
    })

    describe(describeFunction + 'submitChequeBeneficiary', function () {
      if (enabledTests.submitChequeBeneficiary) {
        context("when we don't send value along", function () {
          const value = new BN(0)
          context('when the sender is the beneficiary', function () {
            const sender = defaultCheque.beneficiary
            context('when the first serial is higher than 0', function () {
              expect(defaultCheque.serial).bignumber.to.be.above(new BN(0), "Serial of defaultCheque not above 0")
              context('when the first serial is below MAX_UINT256', function () {
                expect(defaultCheque.serial).bignumber.to.be.below(constants.MAX_UINT256, "Serial of defaultCheque not above 0")
                context('when the issuer is a signee', function () {
                  expect(defaultCheque.signee).to.be.equal(issuer, "Signee of defaultCheque is not issuer")
                  context('when the signee signs the correct fields', function () {
                    const unsignedCheque = Object.assign({}, defaultCheque)
                    context('when we send one cheque', function () {
                      context('when there is a liquidBalance to cover the cheque', function () {
                        describe(describePreCondition + 'shouldDeposit', function () {
                          shouldDeposit(unsignedCheque.amount.add(new BN(1)), issuer)
                          describe(describeTest + 'shouldSubmitChequeBeneficiary', function () {
                            shouldSubmitChequeBeneficiary(unsignedCheque, sender)
                          })
                        })
                      })
                      context('when there is no liquidBalance to cover the cheque', function () {
                        describe(describeTest + 'shouldSubmitChequeBeneficiary', function () {
                          shouldSubmitChequeBeneficiary(unsignedCheque, sender)
                        })
                      })
                    })
                    context('when we send more than one cheque', async function () {
                      describe(describePreCondition + 'shouldSubmitChequeBeneficiary', function () {
                        shouldSubmitChequeBeneficiary(unsignedCheque, sender)
                        context('when the serial number is increasing', function () {
                          describe(describeTest + 'shouldSubmitChequeBeneficiary', function () {
                            const secondSerial = unsignedCheque.serial.add(new BN(1))
                            const increasing_serial_unsignedCheque = Object.assign({}, defaultCheque, { serial: secondSerial })
                            shouldSubmitChequeBeneficiary(increasing_serial_unsignedCheque, sender)
                          })
                        })
                        context('when the serial number stays the same', function () {
                          describe(describeTest + 'shouldSubmitChequeBeneficiary', function () {
                            const secondSerial = unsignedCheque.serial
                            const same_serial_unsignedCheque = Object.assign({}, defaultCheque, { serial: secondSerial })
                            shouldNotSubmitChequeBeneficiary(same_serial_unsignedCheque, same_serial_unsignedCheque, sender, value, "SimpleSwap: invalid serial")
                          })
                          context('when the serial number is decreasing', function () {
                            describe(describeTest + 'shouldSubmitChequeBeneficiary', function () {
                              const secondSerial = unsignedCheque.serial.sub(new BN(1))
                              const decreasing_serial_unsignedCheque = Object.assign({}, defaultCheque, { serial: secondSerial })
                              shouldNotSubmitChequeBeneficiary(decreasing_serial_unsignedCheque, decreasing_serial_unsignedCheque, sender, value, "SimpleSwap: invalid serial")
                            })
                          })
                        })
                      })
                    })
                    context('when the signee does not sign the correct fields', function () {
                      describe(describeTest + 'shouldNotSubmitChequeBeneficiary', function () {
                        const wrongBeneficiary = constants.ZERO_ADDRESS
                        const wrong_beneficiary_unsignedCheque = Object.assign({}, defaultCheque, { beneficiary: wrongBeneficiary })
                        shouldNotSubmitChequeBeneficiary(wrong_beneficiary_unsignedCheque, defaultCheque, sender, value, "SimpleSwap: invalid issuerSig")
                      })
                    })
                  })
                  context('when the issuer is not the signee', function () {
                    describe(describeTest + 'shouldNotSubmitChequeBeneficiary', function () {
                      const signee = alice
                      const wrong_signee_unsignedCheque = Object.assign({}, defaultCheque, { signee: signee })
                      shouldNotSubmitChequeBeneficiary(wrong_signee_unsignedCheque, wrong_signee_unsignedCheque, sender, value, "SimpleSwap: invalid issuerSig")
                    })
                  })
                })
                context('when the first serial is at MAX_UINT256', function () {
                  describe(describeTest + 'shouldSubmitChequeBeneficiary', function () {
                    const MAX_UINT256_unsignedCheque = Object.assign({}, defaultCheque, { serial: constants.MAX_UINT256 })
                    shouldSubmitChequeBeneficiary(MAX_UINT256_unsignedCheque, defaultCheque.beneficiary)
                    context('when we want to submit a cheque afterwards', function () {
                      describe(describeTest + 'shouldNotSubmitChequeBeneficiary', function () {
                        // Solidity wraps integers
                        const MAX_UINT256_wrap_unsignedCheque = Object.assign({}, defaultCheque, { serial: MAX_UINT256_unsignedCheque.serial.add(new BN(1)) })
                        shouldNotSubmitChequeBeneficiary(MAX_UINT256_wrap_unsignedCheque, MAX_UINT256_wrap_unsignedCheque, sender, value, "SimpleSwap: invalid serial")
                      })
                    })
                  })
                })
              })
            })
            context('when the serial is 0', function () {
              describe(describeTest + 'shouldNotSubmitChequeBeneficiary', function () {
                const serial = new BN(0)
                const zero_serial_unsignedCheque = Object.assign({}, defaultCheque, { serial: serial })
                shouldNotSubmitChequeBeneficiary(zero_serial_unsignedCheque, zero_serial_unsignedCheque, sender, value, "SimpleSwap: invalid serial")
              })
            })
          })
          context('when the sender is not the beneficiary', function () {
            describe(describeTest + 'shouldNotSubmitChequeBeneficiary', function () {
              const sender = alice
              shouldNotSubmitChequeBeneficiary(defaultCheque, defaultCheque, sender, value, "SimpleSwap: invalid issuerSig")
            })
          })
        })
        context('when we send value along', function () {
          const value = new BN(1)
          const sender = defaultCheque.beneficiary
          shouldNotSubmitChequeBeneficiary(defaultCheque, defaultCheque, sender, value, "revert")
        })
      }
    })

    describe(describeFunction + 'submitCheque', function () {
      if (enabledTests.submitChequeBeneficiary) {
        context("when we don't send value along", function () {
          const value = new BN(0)
          context('when the first serial is higher than 0', function () {
            expect(defaultCheque.serial).bignumber.to.be.above(new BN(0), "Serial of defaultCheque not above 0")
            context('when the first serial is below MAX_UINT256', function () {
              expect(defaultCheque.serial).bignumber.to.be.below(constants.MAX_UINT256, "Serial of defaultCheque not above 0")
              context('when the beneficiary and issuer are a signee', function () {
                const signees = [issuer, defaultCheque.beneficiary]
                context('when the signee signs the correct fields', function () {
                  const unsignedCheque = Object.assign({}, defaultCheque, { signee: signees })
                  context('when we send one cheque', function () {
                    context('when there is a liquidBalance to cover the cheque', function () {
                      describe(describePreCondition + 'shouldDeposit', function () {
                        shouldDeposit(unsignedCheque.amount, issuer)
                        describe('when the sender is neither the beneficiary nor the issuer', function () {
                          describe(describeTest + 'shouldSubmitChequeBeneficiary', function () {
                            const sender = alice
                            shouldSubmitCheque(unsignedCheque, sender)
                          })
                        })
                        describe('when the sender is the beneficiary', function () {
                          describe(describeTest + 'shouldSubmitChequeBeneficiary', function () {
                            const sender = unsignedCheque.beneficiary
                            shouldSubmitCheque(unsignedCheque, sender)
                          })
                        })
                        describe('when the sender is the issuer', function () {
                          describe(describeTest + 'shouldSubmitChequeBeneficiary', function () {
                            const sender = issuer
                            shouldSubmitCheque(unsignedCheque, sender)
                          })
                        })

                      })
                    })
                    context('when there is no liquidBalance to cover the cheque', function () {
                      describe(describeTest + 'shouldSubmitCheque', function () {
                        const sender = alice
                        shouldSubmitCheque(unsignedCheque, sender)
                      })
                    })
                  })
                  context('when we send more than one cheque', async function () {
                    const sender = alice
                    shouldSubmitCheque(unsignedCheque, sender)
                    context('when the serial number is increasing', function () {
                      describe(describeTest + 'shouldSubmitCheque', function () {
                        const secondSerial = unsignedCheque.serial.add(new BN(1))
                        const increasing_serial_unsignedCheque = Object.assign({}, defaultCheque, { serial: secondSerial, signee: signees })
                        shouldSubmitCheque(increasing_serial_unsignedCheque, sender)
                      })
                    })
                    context('when the serial number stays the same', function () {
                      describe(describeTest + 'shouldNotSubmitCheque', function () {
                        const secondSerial = unsignedCheque.serial
                        const same_serial_unsignedCheque = Object.assign({}, defaultCheque, { serial: secondSerial, signee: signees })
                        shouldNotSubmitCheque(same_serial_unsignedCheque, same_serial_unsignedCheque, sender, value, "SimpleSwap: invalid serial")
                      })
                    })
                    context('when the serial number is decreasing', function () {
                      describe(describeTest + 'shouldNotSubmitCheque', function () {
                        const secondSerial = unsignedCheque.serial.sub(new BN(1))
                        const decreasing_serial_unsignedCheque = Object.assign({}, defaultCheque, { serial: secondSerial, signee: signees })
                        shouldNotSubmitCheque(decreasing_serial_unsignedCheque, decreasing_serial_unsignedCheque, sender, value, "SimpleSwap: invalid serial")
                      })
                    })
                  })
                })
                context("when the signees don't not sign the correct fields", function () {
                  describe(describeTest + 'shouldNotSubmitCheque', function () {
                    const sender = alice
                    const wrongBeneficiary = constants.ZERO_ADDRESS
                    const wrong_beneficiary_unsignedCheque = Object.assign({}, defaultCheque, { beneficiary: wrongBeneficiary, signee: signees })
                    const functionParams = defaultCheque
                    shouldNotSubmitCheque(wrong_beneficiary_unsignedCheque, functionParams, sender, value, "SimpleSwap: invalid issuerSig")
                  })
                })
              })
              context('when the issuer is not the signee', function () {
                describe(describeTest + 'shouldNotSubmitCheque', function () {
                  const sender = alice
                  const signees = [alice, defaultCheque.beneficiary]
                  const wrong_signee_unsignedCheque = Object.assign({}, defaultCheque, { signee: signees })
                  shouldNotSubmitCheque(wrong_signee_unsignedCheque, wrong_signee_unsignedCheque, sender, value, "SimpleSwap: invalid issuerSig")
                })
              })
              context('when the beneficiary is not the signee', function () {
                describe(describeTest + 'shouldNotSubmitCheque', function () {
                  const sender = alice
                  const signees = [issuer, alice]
                  const wrong_signee_unsignedCheque = Object.assign({}, defaultCheque, { signee: signees })
                  shouldNotSubmitCheque(wrong_signee_unsignedCheque, wrong_signee_unsignedCheque, sender, value, "SimpleSwap: invalid beneficiarySig")
                })
              })
              context('when neither the issuer nor the beneficiary are a signee', function () {
                describe(describeTest + 'shouldNotSubmitCheque', function () {
                  const sender = alice
                  const signees = [alice, alice]
                  const wrong_signee_unsignedCheque = Object.assign({}, defaultCheque, { signee: signees })
                  shouldNotSubmitCheque(wrong_signee_unsignedCheque, wrong_signee_unsignedCheque, sender, value, "SimpleSwap: invalid issuerSig")
                })
              })
            })
            context('when the first serial is at MAX_UINT256', function () {
              const sender = alice
              describe(describePreCondition + 'shouldSubmitCheque', function () {
                const signees = [issuer, defaultCheque.beneficiary]
                const MAX_UINT256_unsignedCheque = Object.assign({}, defaultCheque, { serial: constants.MAX_UINT256, signee: signees })
                shouldSubmitCheque(MAX_UINT256_unsignedCheque, defaultCheque.beneficiary)
                describe(describeTest + 'shouldNotSubmitCheque', function () {
                  // Solidity wraps integers
                  const MAX_UINT256_wrap_unsignedCheque = Object.assign({}, defaultCheque, { serial: MAX_UINT256_unsignedCheque.serial.add(new BN(1)), signee: signees })
                  shouldNotSubmitCheque(MAX_UINT256_wrap_unsignedCheque, MAX_UINT256_wrap_unsignedCheque, sender, value, "SimpleSwap: invalid serial")
                })
              })
            })
          })
          context('when the serial is 0', function () {
            const sender = alice
            const signees = [issuer, defaultCheque.beneficiary]
            const serial = new BN(0)
            const zero_serial_unsignedCheque = Object.assign({}, defaultCheque, { serial: serial, signee: signees })
            shouldNotSubmitCheque(zero_serial_unsignedCheque, zero_serial_unsignedCheque, sender, value, "SimpleSwap: invalid serial")
          })
        })
        context('when we send value along', function () {
          const sender = alice
          const value = new BN(1)
          let signees = [issuer, defaultCheque.beneficiary]
          const unsignedCheque = Object.assign({}, defaultCheque, { signee: signees })
          describe(describeTest, 'shouldNotSubmitCheque', function () {
            shouldNotSubmitCheque(unsignedCheque, defaultCheque, sender, value, "revert")
          })
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
                context('when the beneficiaryPrincipal is a signee', function () {
                  const signee = defaultCheque.beneficiary
                  context('when the beneficiaryPrincipal signs the correct fields', function () {
                    context('when the beneficiaryPrincipal and beneficiaryAgent are not the sender', function () {
                      const sender = alice
                      context("when the beneficiaryAgent is not the beneficiaryPrincipal", function () {
                        const beneficiaryAgent = agent
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
                                          describe('when the hardDeposits are assigned to the beneficiaryPrincipal', function () {
                                            const hardDepositReceiver = defaultCheque.beneficiary
                                            context('when the hardDeposit is more the requestPayout', function () {
                                              const hardDeposit = requestPayout.add(new BN(1))
                                              describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                                                shouldIncreaseHardDeposit(hardDepositReceiver, hardDeposit, issuer)
                                                describe(describeTest + 'shouldCashCheque', function () {
                                                  shouldCashCheque(defaultCheque.beneficiary, beneficiaryAgent, requestPayout, calleePayout, sender)
                                                })
                                              })
                                            })
                                            context('when the hardDeposit equals the requestPayout', function () {
                                              const hardDeposit = requestPayout
                                              describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                                                shouldIncreaseHardDeposit(hardDepositReceiver, hardDeposit, issuer)
                                                describe(describeTest + 'shouldCashCheque', function () {
                                                  shouldCashCheque(defaultCheque.beneficiary, beneficiaryAgent, requestPayout, calleePayout, sender)
                                                })
                                              })
                                            })
                                            context('when the hardDeposit is less than the requestPayout', function () {
                                              const hardDeposit = requestPayout.sub(new BN(1))
                                              describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                                                shouldIncreaseHardDeposit(hardDepositReceiver, hardDeposit, issuer)
                                                describe(describeTest + 'shouldCashCheque', function () {
                                                  shouldCashCheque(defaultCheque.beneficiary, beneficiaryAgent, requestPayout, calleePayout, sender)
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
                                                shouldCashCheque(defaultCheque.beneficiary, beneficiaryAgent, requestPayout, calleePayout, sender)
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

                                          shouldCashCheque(defaultCheque.beneficiary, beneficiaryAgent, requestPayout, calleePayout, sender)
                                        })
                                      })
                                    })
                                  })
                                  context('when there is no balance', function () {
                                    describe(describeTest + 'shouldNotCashCheque', function () {
                                      const toSignFields = {
                                        requestPayout,
                                        beneficiaryAgent,
                                        calleePayout,
                                      }
                                      const toSubmitFields = Object.assign({}, toSignFields, { beneficiaryPrincipal: defaultCheque.beneficiary })
                                      shouldNotCashCheque(toSignFields, toSubmitFields, value, sender, signee, "SimpleSwap: cannot pay callee")
                                    })
                                  })
                                })
                                context('when the requestPayout is less than the submitted value', function () {
                                  describe(describePreCondition + 'shouldDeposit', function () {
                                    shouldDeposit(defaultCheque.amount, issuer)
                                    const requestPayout = defaultCheque.amount.sub(new BN(1))
                                    describe(describeTest + 'shouldCashCheque', function () {
                                      shouldCashCheque(defaultCheque.beneficiary, beneficiaryAgent, requestPayout, calleePayout, sender)
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
                                  shouldCashCheque(defaultCheque.beneficiary, beneficiaryAgent, requestPayout, calleePayout, sender)
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
                                    shouldCashCheque(defaultCheque.beneficiary, beneficiaryAgent, requestPayout, calleePayout, sender)
                                    describe(describeTest + 'shouldCashCheque', function () {
                                      shouldCashCheque(defaultCheque.beneficiary, beneficiaryAgent, requestPayout, calleePayout, sender)
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
                                    shouldCashCheque(defaultCheque.beneficiary, beneficiaryAgent, requestPayout, calleePayout, sender)
                                    describe(describeTest + 'shouldNotCashCheque', function () {
                                      const toSignFields = {
                                        requestPayout,
                                        beneficiaryAgent,
                                        calleePayout,
                                      }
                                      const toSubmitFields = Object.assign({}, toSignFields, { beneficiaryPrincipal: defaultCheque.beneficiary })
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
                              beneficiaryAgent,
                              calleePayout,
                            }
                            const toSubmitFields = Object.assign({}, toSignFields, { beneficiaryPrincipal: defaultCheque.beneficiary })
                            shouldNotCashCheque(toSignFields, toSubmitFields, value, sender, signee, "SimpleSwap: not enough balance owed")
                          })
                        })
                      })
                      context('when the beneficiaryAgent is the beneficiaryPrincipal', function () {
                        const beneficiaryAgent = defaultCheque.beneficiary
                        describe(describePreCondition + 'shouldDeposit', function () {
                          shouldDeposit(defaultCheque.amount, issuer)
                          describe(describePreCondition + 'shouldSubmitChequeBeneficiary', function () {
                            shouldSubmitChequeBeneficiary(defaultCheque, defaultCheque.beneficiary)
                            describe(describeTest + 'shouldCashCheque', function () {
                              const waitTime = defaultCheque.timeout.add(new BN(100))

                              beforeEach(async function () {
                                await time.increase(waitTime)
                              })
                              shouldCashCheque(defaultCheque.beneficiary, beneficiaryAgent, defaultCheque.amount, new BN(1), alice)
                            })
                          })
                        })
                      })
                    })
                  })
                  context('when the beneficiary does not sign the correct fields', function () {
                    const sender = alice
                    const beneficiaryAgent = defaultCheque.beneficiary
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
                            beneficiaryAgent,
                            calleePayout,
                          }
                          const toSubmitFields = Object.assign({}, toSignFields, { beneficiaryPrincipal: defaultCheque.beneficiary, requestPayout: new BN(1) })
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
                  const beneficiaryAgent = defaultCheque.beneficiary
                  describe(describePreCondition + 'shouldDeposit', function () {
                    shouldDeposit(defaultCheque.amount, issuer)
                    describe(describePreCondition + 'shouldSubmitChequeBeneficiary', function () {
                      shouldSubmitChequeBeneficiary(defaultCheque, defaultCheque.beneficiary)
                      const signee = alice
                      describe(describeTest + 'shouldNotCashCheque', function () {
                        const toSignFields = {
                          requestPayout: new BN(0),
                          beneficiaryAgent,
                          calleePayout,
                        }
                        const toSubmitFields = Object.assign({}, toSignFields, { beneficiaryPrincipal: defaultCheque.beneficiary })
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
                    shouldCashCheque(defaultCheque.beneficiary, agent, defaultCheque.amount, calleePayout, alice)
                  })
                })
              })
            })
          })
          context('when the signature has expired', async function () {
            const beneficiaryAgent = agent
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
                    beneficiaryAgent,
                    calleePayout,
                  }
                  const toSubmitFields = Object.assign({}, toSignFields, { beneficiaryPrincipal: defaultCheque.beneficiary })
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
          const beneficiaryAgent = agent
          const calleePayout = new BN(0)
          const sender = alice
          const signee = defaultCheque.beneficiary
          describe(describeTest + 'shouldNotCashCheque', function () {
            const toSignFields = {
              requestPayout: new BN(0),
              beneficiaryAgent,
              calleePayout,
            }
            const toSubmitFields = Object.assign({}, toSignFields, { beneficiaryPrincipal: defaultCheque.beneficiary })
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
                context('when the beneficiaryAgent is not the beneficiaryPrincipal', function () {
                  const beneficiaryAgent = agent
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
                                        shouldCashChequeBeneficiary(beneficiaryAgent, requestPayout, sender)
                                      })
                                    })
                                  })
                                  context('when the hardDeposit equals the requestPayout', function () {
                                    const hardDeposit = requestPayout
                                    describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                                      shouldIncreaseHardDeposit(hardDepositReceiver, hardDeposit, issuer)
                                      describe(describeTest + 'shouldCashChequeBeneficiary', function () {
                                        shouldCashChequeBeneficiary(beneficiaryAgent, requestPayout, sender)
                                      })
                                    })
                                  })
                                  context('when the hardDeposit is less than the requestPayout', function () {
                                    const hardDeposit = requestPayout.sub(new BN(1))
                                    describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                                      shouldIncreaseHardDeposit(hardDepositReceiver, hardDeposit, issuer)
                                      describe(describeTest + 'shouldCashChequeBeneficiary', function () {
                                        shouldCashChequeBeneficiary(beneficiaryAgent, requestPayout, sender)
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
                                      shouldCashChequeBeneficiary(beneficiaryAgent, requestPayout, sender)
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
                                shouldCashChequeBeneficiary(beneficiaryAgent, requestPayout, sender)
                              })
                            })
                          })
                        })
                        context('when there is no balance', function () {
                          describe(describeTest + 'shouldCashChequeBeneficiary', function () {
                            shouldCashChequeBeneficiary(beneficiaryAgent, requestPayout, sender)
                          })
                        })
                      })
                      context('when the requestPayout is less than the submitted value', function () {
                        const requestPayout = defaultCheque.amount.sub(new BN(1))
                        describe(describeTest + 'shouldCashChequeBeneficiary', function () {
                          shouldCashChequeBeneficiary(beneficiaryAgent, requestPayout, sender)
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
                      shouldCashChequeBeneficiary(beneficiaryAgent, requestPayout, sender)
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
                        shouldCashChequeBeneficiary(beneficiaryAgent, requestPayout, sender)
                        describe(describeTest + 'shouldCashChequeBeneficiary', function () {
                          shouldCashChequeBeneficiary(beneficiaryAgent, requestPayout, sender)
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
                          shouldCashChequeBeneficiary(beneficiaryAgent, requestPayout, sender)
                          describe(describeTest + 'shouldNotCashChequeBeneficiary', function () {
                            shouldNotCashChequeBeneficiary(beneficiaryAgent, requestPayout, sender, value, "SimpleSwap: not enough balance owed")
                          })
                        })
                      })
                    })
                  })
                })
                context('when the beneficiaryAgent is the beneficiaryPrincipal', function () {
                  const beneficiaryAgent = defaultCheque.beneficiary
                  const waitTime = defaultCheque.timeout.add(new BN(100))
                  beforeEach(async function () {
                    await time.increase(waitTime)
                  })
                  const requestPayout = defaultCheque.amount
                  describe(describeTest + 'shouldCashChequeBeneficiary', function () {
                    shouldCashChequeBeneficiary(beneficiaryAgent, requestPayout, sender)
                  })
                })
              })
            })
            context("when we don't submit a cheque beforeHand", function () {
              describe(describeTest + 'shouldNotCashChequeBeneficiary', function () {
                shouldNotCashChequeBeneficiary(agent, defaultCheque.amount, sender, value, "SimpleSwap: not enough balance owed")
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
                  shouldNotCashChequeBeneficiary(agent, defaultCheque.amount, sender, value, "SimpleSwap: not enough balance owed")
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
                  shouldNotCashChequeBeneficiary(agent, defaultCheque.amount, sender, value, "revert")
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