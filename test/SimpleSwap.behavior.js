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
} = require('./SimpleSwap.should.js')

// switch to false if you don't want to test the particular function
enabledTests = {
  DEFAULT_HARDDEPPOSIT_DECREASE_TIMEOUT: true,
  cheques: true,
  harddeposits: true,
  totalharddeposit: true,
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
  withdraw: true
}

// constants to make the test-log more readable
const describeFunction = 'FUNCTION: '
const describePreCondition = 'PRE-CONDITION: '
const describeTest = 'TEST: '

// @param balance total ether deposited in checkbook
// @param liquidBalance totalDeposit - harddeposits
// @param issuer the issuer of the checkbook
// @param alice a counterparty of the checkbook 
// @param bob a counterparty of the checkbook
function shouldBehaveLikeSimpleSwap([issuer, alice, bob], DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT) {
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
        //TODO
      }
    })

    describe(describeFunction + 'harddeposits', function () {
      if (enabledTests.harddeposits) {
        //TODO
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
            context('when there are harddeposits', function () {
              describe('when the harddeposits equal the depositAmount', function () {
                describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                  const hardDeposit = depositAmount
                  shouldIncreaseHardDeposit(defaultCheque.beneficiary, hardDeposit, issuer)
                  describe(describeTest + 'liquidBalance', function () {
                    shouldReturnLiquidBalance(depositAmount.sub(hardDeposit))
                  })
                })
                describe('when the harddeposits are lower than the depositAmount', function () {
                  describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                    const hardDeposit = depositAmount.sub(new BN(40))
                    shouldIncreaseHardDeposit(defaultCheque.beneficiary, hardDeposit, issuer)
                    describe(describeTest + 'shouldReturnLiquidBalance', function () {
                      shouldReturnLiquidBalance(depositAmount.sub(hardDeposit))
                    })
                  })
                })
              })
              context('when there are no harddeposits', function () {
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
                            const secondSerial = new BN(parseInt(unsignedCheque.serial) + 1)
                            const increasing_serial_unsignedCheque = Object.assign({}, defaultCheque, { serial: secondSerial, signee: defaultCheque.beneficiary })
                            shouldSubmitChequeIssuer(increasing_serial_unsignedCheque, sender)
                          })
                        })
                        context('when the serial number stays the same', function () {
                          context('when the serial number is increasing', function () {
                            describe(describeTest + 'shouldSubmitChequeIssuer', function () {
                              const secondSerial = new BN(parseInt(unsignedCheque.serial))
                              const same_serial_unsignedCheque = Object.assign({}, defaultCheque, { serial: secondSerial, signee: defaultCheque.beneficiary })
                              shouldNotSubmitChequeIssuer(same_serial_unsignedCheque, same_serial_unsignedCheque, sender, value, "SimpleSwap: invalid serial")
                            })
                          })
                          context('when the serial number is decreasing', function () {
                            context('when the serial number is increasing', function () {
                              describe(describeTest + 'shouldSubmitChequeIssuer', function () {
                                const secondSerial = new BN(parseInt(unsignedCheque.serial) + -1)
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
                            const secondSerial = new BN(parseInt(unsignedCheque.serial) + 1)
                            const increasing_serial_unsignedCheque = Object.assign({}, defaultCheque, { serial: secondSerial })
                            shouldSubmitChequeBeneficiary(increasing_serial_unsignedCheque, sender)
                          })
                        })
                        context('when the serial number stays the same', function () {
                          describe(describeTest + 'shouldSubmitChequeBeneficiary', function () {
                            const secondSerial = new BN(parseInt(unsignedCheque.serial))
                            const same_serial_unsignedCheque = Object.assign({}, defaultCheque, { serial: secondSerial })
                            shouldNotSubmitChequeBeneficiary(same_serial_unsignedCheque, same_serial_unsignedCheque, sender, value, "SimpleSwap: invalid serial")
                          })
                          context('when the serial number is decreasing', function () {
                            describe(describeTest + 'shouldSubmitChequeBeneficiary', function () {
                              const secondSerial = new BN(parseInt(unsignedCheque.serial) + -1)
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
                      describe(describeTest + 'shouldSubmitCheque', function() {
                        const sender = alice
                        shouldSubmitCheque(unsignedCheque, sender)
                      })
                    })
                  })
                  context('when we send more than one cheque', async function () {
                    const sender = alice
                    shouldSubmitCheque(unsignedCheque, sender)
                    context('when the serial number is increasing', function () {
                      describe(describeTest + 'shouldSubmitCheque', function() {
                        const secondSerial = new BN(parseInt(unsignedCheque.serial) + 1)
                        const increasing_serial_unsignedCheque = Object.assign({}, defaultCheque, { serial: secondSerial, signee: signees })
                        shouldSubmitCheque(increasing_serial_unsignedCheque, sender)
                      })
                    })
                    context('when the serial number stays the same', function () {
                      describe(describeTest + 'shouldNotSubmitCheque', function() {
                        const secondSerial = new BN(parseInt(unsignedCheque.serial))
                        const same_serial_unsignedCheque = Object.assign({}, defaultCheque, { serial: secondSerial, signee: signees })
                        shouldNotSubmitCheque(same_serial_unsignedCheque, same_serial_unsignedCheque, sender, value, "SimpleSwap: invalid serial")
                      })
                    })
                    context('when the serial number is decreasing', function () {
                      describe(describeTest + 'shouldNotSubmitCheque', function() {
                        const secondSerial = new BN(parseInt(unsignedCheque.serial) + -1)
                        const decreasing_serial_unsignedCheque = Object.assign({}, defaultCheque, { serial: secondSerial, signee: signees })
                        shouldNotSubmitCheque(decreasing_serial_unsignedCheque, decreasing_serial_unsignedCheque, sender, value, "SimpleSwap: invalid serial")
                      })
                    })
                  })
                })
                context("when the signees don't not sign the correct fields", function () {
                  describe(describeTest + 'shouldNotSubmitCheque', function() {
                    const sender = alice
                    const wrongBeneficiary = constants.ZERO_ADDRESS
                    const wrong_beneficiary_unsignedCheque = Object.assign({}, defaultCheque, { beneficiary: wrongBeneficiary, signee: signees })
                    const functionParams = defaultCheque
                    shouldNotSubmitCheque(wrong_beneficiary_unsignedCheque, functionParams, sender, value, "SimpleSwap: invalid issuerSig")
                  })
                })
              })
              context('when the issuer is not the signee', function () {
                describe(describeTest + 'shouldNotSubmitCheque', function() {
                  const sender = alice
                  const signees = [alice, defaultCheque.beneficiary]
                  const wrong_signee_unsignedCheque = Object.assign({}, defaultCheque, { signee: signees })
                  shouldNotSubmitCheque(wrong_signee_unsignedCheque, wrong_signee_unsignedCheque, sender, value, "SimpleSwap: invalid issuerSig")
                })
              })
              context('when the beneficiary is not the signee', function () {
                describe(describeTest + 'shouldNotSubmitCheque', function() {
                  const sender = alice
                  const signees = [issuer, alice]
                  const wrong_signee_unsignedCheque = Object.assign({}, defaultCheque, { signee: signees })
                  shouldNotSubmitCheque(wrong_signee_unsignedCheque, wrong_signee_unsignedCheque, sender, value, "SimpleSwap: invalid beneficiarySig")
                })
              })
              context('when neither the issuer nor the beneficiary are a signee', function () {
                describe(describeTest + 'shouldNotSubmitCheque', function() {
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

    describe(describeFunction + 'cashChequeBeneficiary', function () {
      if (enabledTests.cashChequeBeneficiary) {
        let cheque = defaultCheque
        context('when there is sufficient balance in the chequebook', function () {
          shouldDeposit(defaultCheque.amount, issuer)
          context('when the beneficiary has submitted a cheque', function () {
            shouldSubmitChequeBeneficiary(cheque, cheque.beneficiary)
            context('when sufficient time has passed', function () {
              beforeEach(async function () {
                time.increase(new BN(86400))
                this.currentBeneficiaryBalance = await balance.current(cheque.beneficiary)
                this.currentCheque = await this.simpleSwap.cheques(cheque.beneficiary)
                const { logs, receipt } = await this.simpleSwap.cashChequeBeneficiary(cheque.beneficiary, cheque.amount, { from: cheque.beneficiary })
                this.logs = logs
                this.receipt = receipt
              })
              // _shouldCashChequeInternal(cheque.beneficiary, cheque.beneficiary, defaultCheque.amount, new BN(0))

            })
          })
        })
      }
    })

    describe(describeFunction + 'cashCheque', function () {
      if (enabledTests.cashCheque) {
        let cheque = defaultCheque
        context('when there is sufficient balance in the chequebook', function () {
          shouldDeposit(defaultCheque.amount, issuer)
          context('when the beneficiary has submitted a cheque', function () {
            shouldSubmitChequeBeneficiary(cheque, cheque.beneficiary)
            context('when sufficient time has passed', function () {
              before(function () { time.increase(new BN(86400)) })
              const sender = alice
              it('is here', function () {

              })
            })
          })
        })
      }
    })

    describe(describeFunction + 'prepareDecreaseHardDeposit', function () {
      if (enabledTests.prepareDecreaseHardDeposit) {
        let amount = new BN(50)
        let beneficiary = bob
        context('when the sender is the issuer', function () {
          context('when the hard deposit is high enough', function () {
            context('when no custom decreaseTimeout is set', function () {
              shouldDecreaseHardDeposit()
            })
            context('when a custom decreaseTimeout is set', function () {
              let decreaseTimeout = new BN(100)
              beforeEach(async function () {
                const data = web3.utils.keccak256(web3.eth.abi.encodeParameters(['address', 'address', 'uint256'], [this.simpleSwap.address, beneficiary, decreaseTimeout.toString()]))
                await this.simpleSwap.setCustomHardDepositDecreaseTimeout(
                  beneficiary,
                  decreaseTimeout,
                  await sign(data, beneficiary), {
                    from: issuer
                  })
              })
              shouldDecreaseHardDeposit()
            })
          })
          context('when the hard deposit is not high enough', function () {
            beforeEach(async function () {
              await this.simpleSwap.send(amount)
              await this.simpleSwap.increaseHardDeposit(beneficiary, amount.divn(2))
            })
            it('reverts', async function () {
              await expectRevert(this.simpleSwap.prepareDecreaseHardDeposit(
                beneficiary,
                amount, {
                  from: issuer
                }), "SimpleSwap: hard deposit not sufficient")
            })

          })
        })
        context('when the sender is not the issuer', function () {
          let sender = bob
          it('reverts', async function () {
            await expectRevert(this.simpleSwap.prepareDecreaseHardDeposit(
              beneficiary,
              amount,
              { from: sender }), "SimpleSwap: not issuer")
          })
        })
      }
    })

    describe(describeFunction + 'decreaseHardDeposit', function () {
      if (enabledTests.decreaseHardDeposit) {
        let beneficiary = bob
        let amount = new BN(500)
        let decrease = new BN(400)
        beforeEach(async function () {
          await this.simpleSwap.send(amount)
          await this.simpleSwap.increaseHardDeposit(beneficiary, amount)
        })
        context('when decrease is ready', function () {
          context('when there is enough hard deposit left', function () {
            beforeEach(async function () {
              await this.simpleSwap.prepareDecreaseHardDeposit(beneficiary, decrease)
              await time.increase(await this.simpleSwap.DEFAULT_HARDDEPPOSIT_DECREASE_TIMEOUT())
              let { logs } = await this.simpleSwap.decreaseHardDeposit(beneficiary)
              this.logs = logs
            })

            it('should fire the HardDepositAmountChanged event', async function () {
              expectEvent.inLogs(this.logs, 'HardDepositAmountChanged', {
                beneficiary,
                amount: amount.sub(decrease)
              })
            })

            it('should set the new amount', async function () {
              expect((await this.simpleSwap.hardDeposits(beneficiary))[0]).bignumber.is.equal(amount.sub(decrease))
            })
          })
          // TODO: when there is not enough left
        })

        context('when timeout not yet expired', function () {
          beforeEach(async function () {
            await this.simpleSwap.prepareDecreaseHardDeposit(beneficiary, amount)
          })
          it('reverts', async function () {
            await expectRevert(
              this.simpleSwap.decreaseHardDeposit(beneficiary, { from: issuer }),
              "SimpleSwap: deposit not yet timed out"
            )
          })
        })

        context('when no decrease prepared', async function () {
          it('reverts', async function () {
            await expectRevert(
              this.simpleSwap.decreaseHardDeposit(beneficiary, { from: issuer }),
              "SimpleSwap: deposit not yet timed out"
            )
          })
        })

      }
    })

    describe(describeFunction + 'increaseHardDeposit', function () {
      if (enabledTests.increaseHardDeposit) {
        let amount = new BN(50)
        let beneficiary = bob
        context('when the sender is the issuer', function () {
          let sender = issuer
          context('when the totalHardDeposit is below the swap balance', function () {
            shouldDeposit(amount.muln(2), issuer)
            describe('when there is no prior deposit', function () {
              shouldIncreaseHardDeposit(sender, amount)
            })
            context('when there is a prior deposit', function () {
              shouldIncreaseHardDeposit(sender, amount)
              describe('when the totalHardDeposit is below the swap balance', function () {
                shouldIncreaseHardDeposit(sender, amount)
              })
            })
          })
          context('when the totalHardDeposit exceeds the swap balance', function () {
            it('reverts', async function () {
              await expectRevert(this.simpleSwap.increaseHardDeposit(
                bob,
                new BN(amount),
                { from: sender }), "SimpleSwap: hard deposit cannot be more than balance ")
            })
          })
          function shouldIncreaseHardDeposit(sender, amount) {
            beforeEach(async function () {
              this.previousTotalHardDeposit = await this.simpleSwap.totalHardDeposit()
              this.previousHardDeposit = (await this.simpleSwap.hardDeposits(beneficiary))[0]
              let { logs } = await this.simpleSwap.increaseHardDeposit(
                beneficiary,
                amount,
                { from: sender }
              )
              this.logs = logs
            })

            it('should fire the HardDepositAmountChanged event', async function () {
              expectEvent.inLogs(this.logs, 'HardDepositAmountChanged', {
                beneficiary,
                amount: this.previousHardDeposit.add(amount)
              })
            })
            it('increases the totalHardDeposit', async function () {
              expect(await this.simpleSwap.totalHardDeposit()).bignumber.is.equal(this.previousTotalHardDeposit.add(amount))
            })
            it('increases the hardDeposit amount', async function () {
              expect((await this.simpleSwap.hardDeposits(beneficiary))[0]).bignumber.is.equal(this.previousHardDeposit.add(amount))
            })
            it('reset the canBeDecreasedAt  value', async function () {
              expect((await this.simpleSwap.hardDeposits(beneficiary))[3]).bignumber.is.equal(new BN(0))
            })
          }
        })
        context('when the sender is not the issuer', function () {
          let sender = bob
          it('reverts', async function () {
            await expectRevert(this.simpleSwap.increaseHardDeposit(
              bob,
              new BN(amount),
              { from: sender }), "SimpleSwap: not issuer")
          })
        })
      }
    })

    describe(describeFunction + 'setCustomHardDepositDecreaseTimeout', function () {
      if (enabledTests.setCustomHardDepositDecreaseTimeout) {
        let beneficiary = bob
        let decreaseTimeout = new BN(10)
        beforeEach(function () {
          this.data = web3.utils.keccak256(web3.eth.abi.encodeParameters(['address', 'address', 'uint256'], [this.simpleSwap.address, beneficiary, decreaseTimeout.toString()]))
        })
        describe('when the sender is the issuer', function () {
          let sender = issuer
          describe('when the beneficiarySig is valid', function () {
            beforeEach(async function () {
              let { logs } = await this.simpleSwap.setCustomHardDepositDecreaseTimeout(
                beneficiary,
                decreaseTimeout,
                await sign(this.data, beneficiary),
                { from: sender }
              )
              this.logs = logs
            })

            it('should set the decreaseTimeout', async function () {
              expect((await this.simpleSwap.hardDeposits(beneficiary))[2]).bignumber.is.equal(decreaseTimeout)
            })

            it('should fire the HardDepositDecreaseTimeoutChanged', async function () {
              expectEvent.inLogs(this.logs, 'HardDepositDecreaseTimeoutChanged', {
                beneficiary,
                decreaseTimeout
              })
            })
          })
          context('when the beneficiarySig is invalid', function () {
            it('reverts', async function () {
              await expectRevert.unspecified(this.simpleSwap.setCustomHardDepositDecreaseTimeout(
                beneficiary,
                decreaseTimeout,
                '0x',
                { from: sender }
              ))
            })
          })
        })
        context('when the sender is not the issuer', function () {
          let sender = alice
          it('reverts', async function () {
            await expectRevert.unspecified(this.simpleSwap.setCustomHardDepositDecreaseTimeout(
              beneficiary,
              decreaseTimeout,
              await sign(this.data, beneficiary),
              { from: sender }
            ))
          })
        })
      }
    })

    describe(describeFunction + 'withdraw', function () {
      if (enabledTests.withdraw) {
        let amount = new BN(100)
        beforeEach(async function () {
          await this.simpleSwap.send(amount)
        })

        context('when the sender is the issuer', function () {
          let sender = issuer
          context('when the liquid balance is high enough', function () {
            beforeEach(async function () {
              let issuerBalancePrior = await balance.current(issuer)
              let { logs, receipt } = await this.simpleSwap.withdraw(
                amount,
                { from: sender }
              )

              this.logs = logs
              this.expectedBalance = issuerBalancePrior.add(amount).sub(await computeCost(receipt))
            })

            it('should change the issuer balance correctly', async function () {
              expect(await balance.current(issuer)).bignumber.is.equal(this.expectedBalance)
            })
          })
          context('when the liquid balance is too low', function () {
            beforeEach(async function () {
              await this.simpleSwap.increaseHardDeposit(bob, new BN(1), { from: sender })
            })

            it('reverts', async function () {
              await expectRevert(this.simpleSwap.withdraw(amount, {
                from: sender
              }), "SimpleSwap: liquidBalance not sufficient")
            })
          })
        })

        context('when the sender is not the issuer', function () {
          let sender = bob
          it('reverts', async function () {
            await expectRevert(this.simpleSwap.withdraw(amount, {
              from: sender
            }), 'SimpleSwap: not issuer')
          })
        })
      }
    })

    describe(describeFunction + 'deposit', function () {
      if (enabledTests.deposit) {
        shouldDeposit(new BN(1), issuer)
      }
    })
  })
}

module.exports = {
  shouldBehaveLikeSimpleSwap
};