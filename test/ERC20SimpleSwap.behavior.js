const {
  BN,
  time
} = require("@openzeppelin/test-helpers");

const {
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
} = require('./ERC20SimpleSwap.should.js')

// switch to false if you don't want to test the particular function
enabledTests = {
  defaultHardDepositTimeout: true,
  cheques: true,
  hardDeposits: true,
  totalHardDeposit: true,
  issuer: true,
  liquidBalance: true,
  liquidBalanceFor: true,
  cashChequeBeneficiary: true,
  cashCheque: true,
  prepareDecreaseHardDeposit: true,
  decreaseHardDeposit: true,
  increaseHardDeposit: true,
  setCustomHardDepositTimeout: true,
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
function shouldBehaveLikeERC20SimpleSwap([issuer, alice, bob, carol], defaultHardDepositTimeout) {
  // defaults used throught the tests
  const defaults = {
    beneficiary: bob,
    recipient: carol,
    firstCumulativePayout: new BN(500),
    secondCumulativePayout: new BN(1000),
    deposit: new BN(10000),
  }

  context('as a simple swap', function () {
    describe(describeFunction + 'defaultHardDepositTimeout', function () {
      if (enabledTests.defaultHardDepositTimeout) {
        shouldReturnDefaultHardDepositTimeout(defaultHardDepositTimeout)
      }
    })
    describe(describeFunction + 'paidOutCheques', function () {
      if (enabledTests.cheques) {
        const beneficiary = defaults.beneficiary
        context('when no cheque was ever cashed', function () {
          describe(describeTest + 'shouldReturnPaidOut', function () {
            const expectedAmount = new BN(0)
            shouldReturnPaidOut(beneficiary, expectedAmount)
            shouldReturnTotalPaidOut(expectedAmount)
          })
        })
        context('when a cheque was cashed', function () {
          describe(describePreCondition + 'shouldDeposit', function () {
            shouldDeposit(defaults.deposit, issuer)
            describe(describePreCondition + 'shouldCashChequeBeneficiary', function () {
              shouldCashChequeBeneficiary(defaults.recipient, defaults.firstCumulativePayout, issuer, defaults.beneficiary)
              describe(describeTest + 'shouldReturnPaidOut', function () {
                const expectedAmount = defaults.firstCumulativePayout
                shouldReturnPaidOut(beneficiary, expectedAmount)
                shouldReturnTotalPaidOut(expectedAmount)
              })
            })
          })
        })
      }
    })

    describe(describeFunction + 'hardDeposits', function () {
      if (enabledTests.hardDeposits) {
        const beneficiary = defaults.beneficiary
        context('when no hardDeposit was allocated', function () {
          const expectedAmount = new BN(0)
          const exptectedDecreaseAmount = new BN(0)
          const exptectedCanBeDecreasedAt = new BN(0)
          context('when no custom timeout was set', function () {
            const expectedDecreaseTimeout = new BN(0)
            describe(describeTest + 'shouldReturnHardDeposits', function () {
              shouldReturnHardDeposits(beneficiary, expectedAmount, exptectedDecreaseAmount, expectedDecreaseTimeout, exptectedCanBeDecreasedAt)
            })
          })
          context('when a custom timeout was set', function () {
            const expectedDecreaseTimeout = new BN(60)
            describe(describePreCondition + 'shouldSetCustomDecreaseTimeout', function () {
              shouldSetCustomHardDepositTimeout(beneficiary, expectedDecreaseTimeout, issuer)
              describe(describeTest + 'shouldReturnHardDeposits', function () {
                shouldReturnHardDeposits(beneficiary, expectedAmount, exptectedDecreaseAmount, expectedDecreaseTimeout, exptectedCanBeDecreasedAt)
              })
            })
          })
        })
        context('when a hardDeposit was allocated', function () {
          describe(describePreCondition + 'shouldDeposit', function () {
            const depositAmount = new BN(50)
            shouldDeposit(depositAmount, issuer)
            describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
              shouldIncreaseHardDeposit(beneficiary, depositAmount, issuer)
              context('when the hardDeposit was not requested to decrease', function () {
                const expectedDecreaseAmount = new BN(0)
                const expectedCanBeDecreasedAt = new BN(0)
                const expectedDecreaseTimeout = new BN(0)
                describe(describeTest + 'shouldReturnHardDeposits', function () {
                  shouldReturnHardDeposits(beneficiary, depositAmount, expectedDecreaseAmount, expectedDecreaseTimeout, expectedCanBeDecreasedAt)
                })
              })
              context('when the hardDeposit was requested to decrease', function () {
                describe(describePreCondition + 'shouldPrepareDecreaseHardDeposit', function () {
                  const toDecrease = depositAmount.div(new BN(2))
                  shouldPrepareDecreaseHardDeposit(beneficiary, toDecrease, issuer)
                  describe(describeTest + 'shouldReturnHardDeposits', function () {
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

    describe(describeFunction + 'totalHardDeposits', function () {
      if (enabledTests.totalHardDeposit) {
        context('when there are no hardDeposits', function () {
          describe(describeTest + 'shouldReturnTotalHardDeposit', function () {
            shouldReturnTotalHardDeposit(new BN(0))
          })
        })
        context('when there are hardDeposits', function () {
          const depositAmount = defaults.deposit
          describe(describePreCondition + 'shouldDeposit', function () {
            shouldDeposit(depositAmount, issuer)
            describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
              const hardDeposit = defaults.deposit
              shouldIncreaseHardDeposit(defaults.beneficiary, hardDeposit, issuer)
              describe(describeTest + 'shouldReturnTotalHardDeposit', function () {
                shouldReturnTotalHardDeposit(hardDeposit)
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
          const depositAmount = defaults.deposit
          describe(describePreCondition + 'shouldDeposit', function () {
            shouldDeposit(depositAmount, issuer)
            context('when there are hardDeposits', function () {
              describe('when the hardDeposits equal the depositAmount', function () {
                describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                  const hardDeposit = depositAmount
                  shouldIncreaseHardDeposit(defaults.beneficiary, hardDeposit, issuer)
                  describe(describeTest + 'liquidBalance', function () {
                    shouldReturnLiquidBalance(new BN(0))
                  })
                })
                describe('when the hardDeposits are lower than the depositAmount', function () {
                  const hardDeposit = defaults.deposit.div(new BN(2))
                  describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                    shouldIncreaseHardDeposit(defaults.beneficiary, hardDeposit, issuer)
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

    describe(describeFunction + 'shouldReturnLiquidBalanceFor', function () {
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

    describe(describeFunction + 'cashCheque', function () {
      if (enabledTests.cashCheque) {
        const beneficiary = defaults.beneficiary
        const firstCumulativePayout = defaults.firstCumulativePayout
        const recipient = defaults.recipient
        context('when the sender is not the issuer', function() {
          const caller = alice
          context("when we don't send value along", function () {
            const value = new BN(0)
            context('when the beneficiary provides the beneficiarySig', function () {
              const beneficiarySignee = beneficiary
              context('when the issuer provides the issuerSig', function () {
                const issuerSignee = issuer
                context('when the callerPayout is non-zero', function () {
                  const callerPayout = defaults.firstCumulativePayout.div(new BN(100))
                  context('when there is some money deposited', function () {
                    context('when the money fully covers the cheque', function() {
                      const depositAmount = firstCumulativePayout.add(defaults.secondCumulativePayout)
                      describe(describePreCondition + 'shouldDeposit', function () {
                        shouldDeposit(depositAmount, issuer)
                        context('when there are hardDeposits assigned to the beneficiary', function() {
                          context('when the hardDeposits fully cover the cheque', function() {
                            describe(describePreCondition + 'shouldIncreaseHardDeposit', function() {
                              shouldIncreaseHardDeposit(beneficiary, firstCumulativePayout, issuer)
                              context('when we submit one cheque', function() {
                                describe(describeTest + 'shouldCashCheque', function() {
                                  shouldCashCheque(beneficiary, recipient, firstCumulativePayout, callerPayout, caller, beneficiarySignee, issuerSignee)
                                })
                              })
                              context('when we attempt to submit two cheques', function() {
                                describe(describePreCondition + 'shouldCashCheque', function() {
                                  shouldCashCheque(beneficiary, recipient, firstCumulativePayout, callerPayout, caller, beneficiarySignee, issuerSignee)
                                  context('when the second cumulativePayout is higher than the first cumulativePayout', function() {
                                    const secondCumulativePayout = defaults.secondCumulativePayout
                                    describe(describeTest + 'shouldCashCheque', function() {
                                      shouldCashCheque(beneficiary, recipient, secondCumulativePayout, callerPayout, caller, beneficiarySignee, issuerSignee)
                                    })
                                  })
                                  context('when the second cumulativePayout is lower than the first cumulativePayout', function() {
                                    const secondCumulativePayout = firstCumulativePayout.sub(new BN(1))
                                    const revertMessage = 'SafeMath: subtraction overflow.'
                                    const beneficiaryToSign = {
                                      cumulativePayout: secondCumulativePayout,
                                      recipient,
                                      callerPayout
                                    }
                                    const issuerToSign = {
                                      beneficiary,
                                      cumulativePayout: secondCumulativePayout,
                                    }
                                    const toSubmit = Object.assign({}, beneficiaryToSign, issuerToSign)
                                    describe(describeTest + 'shouldNotCashCheque', function() {
                                      shouldNotCashCheque(beneficiaryToSign, issuerToSign, toSubmit, value, caller, beneficiarySignee, issuerSignee, revertMessage)
                                    })
                                  })
                                })
                              })
                            })
                          })
                          context('when the hardDeposits partly cover the cheque', function() {
                            describe(describePreCondition + 'shouldIncreaseHardDeposit', function() {
                              shouldIncreaseHardDeposit(beneficiary, firstCumulativePayout.div(new BN(2)), issuer)
                              describe(describeTest + 'shouldCashCheque', function() {
                                shouldCashCheque(beneficiary, recipient, firstCumulativePayout, callerPayout, caller, beneficiarySignee, issuerSignee)
                              })
                            })
                          })
                        })
                      })
                    })
                    context('when the money partly covers the cheque', function() {
                      const depositAmount = firstCumulativePayout.div(new BN(2))
                      describe(describePreCondition + 'shouldDeposit', function () {
                        shouldDeposit(depositAmount, issuer)
                        describe(describeTest + 'shouldCashCheque', function() {
                          shouldCashCheque(beneficiary, recipient, firstCumulativePayout, callerPayout, caller, beneficiarySignee, issuerSignee)
                        })
                      })
                    })                  
                  })
                  context('when no money is deposited', function () {
                    const revertMessage = 'SimpleSwap: cannot pay caller'
                    const beneficiaryToSign = {
                      cumulativePayout: firstCumulativePayout,
                      recipient,
                      callerPayout
                    }
                    const issuerToSign = {
                      beneficiary,
                      cumulativePayout: firstCumulativePayout,
                    }
                    const toSubmit = Object.assign({}, beneficiaryToSign, issuerToSign)
                    describe(describeTest + 'shouldNotCashCheque', function() {
                      shouldNotCashCheque(beneficiaryToSign, issuerToSign, toSubmit, value, caller, beneficiarySignee, issuerSignee, revertMessage)
                    })
                  })
                })
                context('when the callerPayout is zero', function () {
                  const callerPayout = new BN(0)
                  describe(describeTest + 'shouldCashCheque', function() {
                    shouldCashCheque(beneficiary, recipient, firstCumulativePayout, callerPayout, caller, beneficiarySignee, issuerSignee)
                  })
                })
              })
              context('when the issuer does not provide the issuerSig', function () {
                const issuerSignee = alice
                const callerPayout = defaults.firstCumulativePayout.div(new BN(100))
                const revertMessage = 'SimpleSwap: invalid issuerSig'
                const beneficiaryToSign = {
                  cumulativePayout: firstCumulativePayout,
                  recipient,
                  callerPayout
                }
                const issuerToSign = {
                  beneficiary,
                  cumulativePayout: firstCumulativePayout,
                }
                const toSubmit = Object.assign({}, beneficiaryToSign, issuerToSign)
                describe(describeTest + 'shouldNotCashCheque', function() {
                  shouldNotCashCheque(beneficiaryToSign, issuerToSign, toSubmit, value, caller, beneficiarySignee, issuerSignee, revertMessage)
                })
              })
  
            })
            context('when the beneficiary does not provide the beneficiarySig', function () {
              const beneficiarySignee = alice
              const issuerSignee = issuer
              const callerPayout = defaults.firstCumulativePayout.div(new BN(100))
              const revertMessage = 'SimpleSwap: invalid beneficiarySig'
              const beneficiaryToSign = {
                cumulativePayout: firstCumulativePayout,
                recipient,
                callerPayout
              }
              const issuerToSign = {
                beneficiary,
                cumulativePayout: firstCumulativePayout,
              }
              const toSubmit = Object.assign({}, beneficiaryToSign, issuerToSign)
              describe(describeTest + 'shouldNotCashCheque', function() {
                shouldNotCashCheque(beneficiaryToSign, issuerToSign, toSubmit, value, caller, beneficiarySignee, issuerSignee, revertMessage)
              })
            })
          })
          context('when we send value along', function () {
            const value = new BN(50)
            const beneficiarySignee = alice
            const issuerSignee = issuer
            const callerPayout = defaults.firstCumulativePayout.div(new BN(100))
            const revertMessage = 'revert'
            const beneficiaryToSign = {
              cumulativePayout: firstCumulativePayout,
              recipient,
              callerPayout
            }
            const issuerToSign = {
              beneficiary,
              cumulativePayout: firstCumulativePayout,
            }
            const toSubmit = Object.assign({}, beneficiaryToSign, issuerToSign)
            describe(describeTest + 'shouldNotCashCheque', function() {
              shouldNotCashCheque(beneficiaryToSign, issuerToSign, toSubmit, value, caller, beneficiarySignee, issuerSignee, revertMessage)
            })
          })
        })
        context('when the sender is the issuer', function() {
          const caller = issuer
          const callerPayout = new BN(0)
          const beneficiarySignee = beneficiary
          const issuerSignee = beneficiary // on purpose not the correct signee, as it is not needed
          describe(describeTest + 'shouldCashCheque', function() {
            shouldCashCheque(beneficiary, recipient, firstCumulativePayout, callerPayout, caller, beneficiarySignee, issuerSignee)
          })
        })
        
      }
    })

    describe(describeFunction + 'cashChequeBeneficiary', function () {
      if (enabledTests.cashChequeBeneficiary) {
        const beneficiary = defaults.beneficiary
        context("when we don't send value along", function () {
          const value = new BN(0)
          context('when the issuer is a signee', function () {
            const sender = beneficiary
            const signee = issuer
            context('when the signee signs the correct fields', function () {
              context('when the recipient is not the beneficiary', function () {
                const recipient = defaults.recipient
                context('when we have not cashed a cheque before', function () {
                  const requestPayout = defaults.firstCumulativePayout
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
                                  shouldCashChequeBeneficiary(recipient, requestPayout, signee, sender)
                                })
                              })
                            })
                            context('when the hardDeposit equals the requestPayout', function () {
                              const hardDeposit = requestPayout
                              describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                                shouldIncreaseHardDeposit(hardDepositReceiver, hardDeposit, issuer)
                                describe(describeTest + 'shouldCashChequeBeneficiary', function () {
                                  shouldCashChequeBeneficiary(recipient, requestPayout, signee, sender)
                                })
                              })
                            })
                            context('when the hardDeposit is less than the requestPayout', function () {
                              const hardDeposit = requestPayout.sub(new BN(1))
                              describe(describePreCondition + 'shouldIncreaseHardDeposit', function () {
                                shouldIncreaseHardDeposit(hardDepositReceiver, hardDeposit, issuer)
                                describe(describeTest + 'shouldCashChequeBeneficiary', function () {
                                  shouldCashChequeBeneficiary(recipient, requestPayout, signee, sender)
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
                                shouldCashChequeBeneficiary(recipient, requestPayout, signee, sender)
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
                          shouldCashChequeBeneficiary(recipient, requestPayout, signee, sender)
                        })
                      })
                    })
                  })
                  context('when there is no balance', function () {
                    describe(describeTest + 'shouldCashChequeBeneficiary', function () {
                      shouldCashChequeBeneficiary(recipient, requestPayout, signee, sender)
                    })
                  })
                })

                context('when we have cashed a cheque before', function () {
                  describe(describePreCondition + 'shouldDeposit', function () {
                    shouldDeposit(defaults.deposit, issuer)
                    describe(describePreCondition + 'shouldCashChequeBeneficiary', function () {
                      shouldCashChequeBeneficiary(recipient, defaults.firstCumulativePayout, signee, sender)
                      describe(describeTest + 'shouldCashChequeBeneficiary', function () {
                        shouldCashChequeBeneficiary(recipient, defaults.firstCumulativePayout, signee, sender)
                      })
                    })
                  })
                })
              })
              context('when the recipient is the beneficiary', function () {
                const recipient = defaults.beneficiary
                const requestPayout = defaults.firstCumulativePayout
                describe(describeTest + 'shouldCashChequeBeneficiary', function () {
                  shouldCashChequeBeneficiary(recipient, requestPayout, signee, sender)
                })
              })
            })
            context('when the signee does not sign the correct fields', function () {
              const revertMessage = "SimpleSwap: invalid issuerSig"
              const recipient = defaults.recipient
              const toSubmitCumulativePayment = defaults.firstCumulativePayout
              const toSignCumulativePayment = new BN(1)
              const sender = beneficiary
              shouldNotCashChequeBeneficiary(recipient, toSubmitCumulativePayment, toSignCumulativePayment, signee, sender, value, revertMessage)
            })
          })
          context('when the issuer is not a signee', function () {
            const revertMessage = "SimpleSwap: invalid issuerSig"
            const signee = alice
            const recipient = defaults.recipient
            const toSubmitCumulativePayment = defaults.firstCumulativePayout
            const toSignCumulativePayment = toSubmitCumulativePayment
            const sender = beneficiary
            shouldNotCashChequeBeneficiary(recipient, toSubmitCumulativePayment, toSignCumulativePayment, signee, sender, value, revertMessage)
          })
        })
        context('when we send value along', function () {
          const value = new BN(1)
          const revertMessage = "revert"
          const signee = alice
          const recipient = defaults.recipient
          const toSubmitCumulativePayment = defaults.firstCumulativePayout
          const toSignCumulativePayment = toSubmitCumulativePayment
          const sender = beneficiary
          describe(describeTest + 'shouldNotCashChequeBeneficiary', function () {
            shouldNotCashChequeBeneficiary(recipient, toSubmitCumulativePayment, toSignCumulativePayment, signee, sender, value, revertMessage)
          })
        })
      }
    })


    describe(describeFunction + 'prepareDecreaseHardDeposit', function () {
      if (enabledTests.prepareDecreaseHardDeposit) {
        const beneficiary = defaults.beneficiary
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
                    context('when we have set a custom timeout', function () {
                      describe(describePreCondition + 'shouldSetCustomHardDepositTimeout', function () {
                        const customTimeout = new BN(10)
                        shouldSetCustomHardDepositTimeout(beneficiary, customTimeout, issuer)
                        context('when we have not set a custom timeout', function () {
                          describe(describeTest + 'prepareDecreaseHardDeposit', function () {
                            shouldPrepareDecreaseHardDeposit(beneficiary, decreaseAmount, sender)
                          })
                        })
                      })
                    })
                    context('when we have not set a custom timeout', function () {
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
        const beneficiary = defaults.beneficiary
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
                    context('when we have waited more than timeout time', function () {
                      beforeEach(async function () {
                        await time.increase(await this.ERC20SimpleSwap.defaultHardDepositTimeout())
                      })
                      describe(describeTest + 'shouldDecreaseHardDeposit', function () {
                        shouldDecreaseHardDeposit(beneficiary, sender)
                      })
                    })
                    context('when we have not waited more than defaultHardDepositTimeout time', function () {
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
        const beneficiary = defaults.beneficiary
        context("when we don't send value along", function () {
          const value = new BN(0)
          context('when the sender is the issuer', function () {
            const sender = issuer
            context('when there is more liquidBalance than the requested hardDepositIncrease', function () {
              const deposit = hardDepositIncrease.mul(new BN(2))
              describe(describePreCondition + 'shouldDeposit', function () {
                shouldDeposit(deposit, issuer)
                context('when we have set a customHardDepositTimeout', function () {
                  const customHardDepositTimeout = new BN(60)
                  describe(describePreCondition + 'shouldSetCustomHardDepositTimeout', function () {
                    shouldSetCustomHardDepositTimeout(beneficiary, customHardDepositTimeout, issuer)
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
          const beneficiary = defaults.beneficiary
          const sender = issuer
          describe(describeTest + 'shouldNotIncreaseHardDeposit', function () {
            const revertMessage = "revert"
            shouldNotIncreaseHardDeposit(beneficiary, hardDepositIncrease, sender, value, revertMessage)
          })
        })
      }
    })

    describe(describeFunction + 'setCustomHardDepositTimeout', function () {
      if (enabledTests.setCustomHardDepositTimeout) {
        const beneficiary = defaults.beneficiary
        const timeout = new BN(60)
        context("when we don't send value along", function () {
          const value = new BN(0)
          context('when the sender is the issuer', function () {
            const sender = issuer
            context('when the beneficiary is a signee', function () {
              const signee = beneficiary
              context('when the beneficiary signs the correct fields', function () {
                describe(describeTest + 'shouldSetCustomHardDepositTimeout', function () {
                  shouldSetCustomHardDepositTimeout(beneficiary, timeout, sender)
                })
              })
              context('when the beneficiary does not sign the correct fields', function () {
                describe(describeTest + 'shouldNotSetCustomHardDepositTimeout', function () {
                  const toSubmit = { beneficiary, timeout }
                  const toSign = { beneficiary, timeout: timeout.sub(new BN(1)) }
                  const revertMessage = "SimpleSwap: invalid beneficiarySig"
                  shouldNotSetCustomHardDepositTimeout(toSubmit, toSign, signee, sender, value, revertMessage)
                })
              })
            })
            context('when the beneficiary is not a signee', function () {
              const signee = alice
              describe(describeTest + 'shouldNotSetCustomHardDepositTimeout', function () {
                const toSubmit = { beneficiary, timeout }
                const toSign = toSubmit
                const revertMessage = "SimpleSwap: invalid beneficiarySig"
                shouldNotSetCustomHardDepositTimeout(toSubmit, toSign, signee, sender, value, revertMessage)
              })
            })
          })
          context('when the sender is not the issuer', function () {
            const sender = alice
            describe(describeTest + 'shouldNotSetCustomHardDepositTimeout', function () {
              const toSubmit = { beneficiary, timeout }
              const toSign = toSubmit
              const signee = beneficiary
              const revertMessage = "SimpleSwap: not issuer"
              shouldNotSetCustomHardDepositTimeout(toSubmit, toSign, signee, sender, value, revertMessage)
            })
          })
        })
        context('when we send value along', function () {
          const value = new BN(1)
          const sender = issuer
          const signee = beneficiary
          describe(describeTest + 'shouldNotSetCustomHardDepositTimeout', function () {
            const toSubmit = { beneficiary, timeout }
            const toSign = toSubmit
            const revertMessage = "revert"
            shouldNotSetCustomHardDepositTimeout(toSubmit, toSign, signee, sender, value, revertMessage)
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
        context('when the depositAmount is not zero', function () {
          const depositAmount = new BN(1)
          describe(describeTest + 'shouldDeposit', function () {
            shouldDeposit(depositAmount, issuer)
          })
        })
      }
    })
  })
}

module.exports = {
  shouldBehaveLikeERC20SimpleSwap
};