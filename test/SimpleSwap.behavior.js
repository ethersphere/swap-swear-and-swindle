const {
    BN,
    balance,
    time,
    expectRevert,
    constants,
    expectEvent
} = require("openzeppelin-test-helpers");

const { expect } = require('chai');

const { signCheque, sign } = require("./swutils");
const { computeCost } = require("./testutils");



// @param balance total ether deposited in checkbook
// @param liquidBalance totalDeposit - harddeposits
// @param owner the owner of the checkbook
// @param alice a counterparty of the checkbook 
// @param bob a counterparty of the checkbook
function shouldBehaveLikeSimpleSwap([owner, alice, bob]) {
  const defaultCheque = {
    beneficiary: bob,
    serial: new BN(3),
    amount: new BN(Math.floor(Math.random() * 100000)),
    timeout: new BN(86400),
    signee: owner,
    signature: ""
  }
  context('as a simple swap', function() {
    it('should have a correct owner', async function() {
      expect(await this.simpleSwap.owner()).to.equal(owner)          
    })
    describe('deposit', function() {
      shouldDeposit(new BN(1), owner)
    })
    describe('submitCheque', function() {
      context('when the sender is the owner', function() {
        submitChequeBySender(owner)
      })
      context('when the sender is the beneficiary', function() {
        submitChequeBySender(defaultCheque.beneficiary)  
      })
      context('when the sender is a third party', function() {
        submitChequeBySender(alice)   
      })
      function submitChequeBySender(sender) {
        context('when the first serial is higher than 0', function() {
          expect(defaultCheque.serial).bignumber.to.be.above(new BN(0), "Serial of defaultCheque not above 0")
          context('when the first serial is below MAX_UINT256', function() {
            expect(defaultCheque.serial).bignumber.to.be.below(constants.MAX_UINT256, "Serial of defaultCheque not above 0")
            context('when the beneficiary and owners are both a signee', function() {
              let unsignedCheque = Object.assign({}, defaultCheque, {signee: [defaultCheque.beneficiary, owner]})
              expect(unsignedCheque.signee).to.be.include(unsignedCheque.beneficiary, "Signee of unsignedCheque is not beneficiary")
              expect(unsignedCheque.signee).to.be.include(owner, "Signee of unsignedCheque is not owner")
              context('when the signees signs the correct fields', function() {
                context('when we send one cheque', function() {
                  context('when there is a liquidBalance to cover the cheque', function() {
                    shouldDeposit(unsignedCheque.amount + new BN(1), owner)
                    shouldSubmitCheque(unsignedCheque, sender)
                  })
                  context('when there is no liquidBalance to cover the cheque', function() {
                    shouldSubmitCheque(unsignedCheque, sender)  
                  })
                })
                context('when we send more than one cheque', async function() {
                  shouldSubmitCheque(unsignedCheque, sender)
                  context('when the serial number is increasing', function() {
                    let secondSerial = new BN(parseInt(unsignedCheque.serial) + 1)
                    let increasing_serial_unsignedCheque = Object.assign({}, defaultCheque, {serial: secondSerial, signee: [defaultCheque.beneficiary, owner]})
                    shouldSubmitCheque(increasing_serial_unsignedCheque, sender)
                  })
                  context('when the serial number stays the same', function() {
                    let secondSerial = new BN(parseInt(unsignedCheque.serial))
                    let same_serial_unsignedCheque = Object.assign({}, defaultCheque, {serial: secondSerial, signee: [defaultCheque.beneficiary, owner]})
                    it('reverts', async function() {
                      this.signedCheque = await signCheque(this.simpleSwap, same_serial_unsignedCheque)
                      await expectRevert(this.simpleSwap.submitCheque(
                        this.signedCheque.beneficiary,
                        this.signedCheque.serial, 
                        this.signedCheque.amount, 
                        this.signedCheque.timeout,
                        this.signedCheque.signature[1],
                        this.signedCheque.signature[0], 
                        {from: sender}), "SimpleSwap: invalid serial")
                    })
                  })
                  context('when the serial number is decreasing', function() {
                    let secondSerial = new BN(parseInt(unsignedCheque.serial) + -1)
                    let decreasing_serial_unsignedCheque = Object.assign({}, defaultCheque, {serial: secondSerial, signee: [defaultCheque.beneficiary, owner]})
                    it('reverts', async function() {
                      this.signedCheque = await signCheque(this.simpleSwap, decreasing_serial_unsignedCheque)
                      await expectRevert(this.simpleSwap.submitCheque(
                        this.signedCheque.beneficiary,
                        this.signedCheque.serial, 
                        this.signedCheque.amount, 
                        this.signedCheque.timeout,
                        this.signedCheque.signature[1],
                        this.signedCheque.signature[0], 
                        {from: sender}), "SimpleSwap: invalid serial")
                    })
                  })
                })
              })
              context('when the signee does not sign the correct fields', function() {
                let wrongBeneficiary = constants.ZERO_ADDRESS
                let wrong_beneficiary_unsignedCheque = Object.assign({}, defaultCheque, {beneficiary: wrongBeneficiary, signee: [defaultCheque.beneficiary, owner]})
                it('reverts', async function() {
                  this.signedCheque = await signCheque(this.simpleSwap, wrong_beneficiary_unsignedCheque)
                  await expectRevert(this.simpleSwap.submitCheque(
                    this.signedCheque.beneficiary,
                    this.signedCheque.serial, 
                    this.signedCheque.amount, 
                    this.signedCheque.timeout,
                    this.signedCheque.signature[1],
                    this.signedCheque.signature[0], 
                    {from: sender}), "SimpleSwap: invalid beneficiarySig")
                })
              })
            })
            context('when the owner is not a signee', function() {
              const wrong_signee_unsignedCheque = Object.assign({}, defaultCheque, {signee: [defaultCheque.beneficiary, defaultCheque.beneficiary]})
              it('reverts', async function() {
                this.signedCheque = await signCheque(this.simpleSwap, wrong_signee_unsignedCheque)
                await expectRevert(this.simpleSwap.submitCheque(
                  this.signedCheque.beneficiary,
                  this.signedCheque.serial, 
                  this.signedCheque.amount, 
                  this.signedCheque.timeout,
                  this.signedCheque.signature[1],
                  this.signedCheque.signature[0], 
                  {from: sender}), "SimpleSwap: invalid ownerSig")
              })
            })
            context('when the beneficiary is not a signee', function() {
              const wrong_signee_unsignedCheque = Object.assign({}, defaultCheque, {signee: [owner, owner]})
              it('reverts', async function() {
                this.signedCheque = await signCheque(this.simpleSwap, wrong_signee_unsignedCheque)
                await expectRevert(this.simpleSwap.submitCheque(
                  this.signedCheque.beneficiary,
                  this.signedCheque.serial, 
                  this.signedCheque.amount, 
                  this.signedCheque.timeout,
                  this.signedCheque.signature[1],
                  this.signedCheque.signature[0], 
                  {from: sender}), "SimpleSwap: invalid beneficiarySig")
              })
            })
          })
          context('when the first serial is at MAX_UINT256', function() {
            const MAX_UINT256_unsignedCheque = Object.assign({}, defaultCheque, {serial: constants.MAX_UINT256, signee: [defaultCheque.beneficiary, owner]})
            shouldSubmitCheque(MAX_UINT256_unsignedCheque, sender)
            // Solidity wraps integers
            const MAX_UINT256_wrap_unsignedCheque = Object.assign({}, defaultCheque, {serial: MAX_UINT256_unsignedCheque.serial + new BN(1), signee: [defaultCheque.beneficiary, owner]})
            it('should not be possible to submit a cheque afterwards', async function() {
              this.signedCheque = await signCheque(this.simpleSwap, MAX_UINT256_wrap_unsignedCheque)
              await expectRevert(this.simpleSwap.submitCheque(
                this.signedCheque.beneficiary,
                this.signedCheque.serial, 
                this.signedCheque.amount, 
                this.signedCheque.timeout,
                this.signedCheque.signature[1],
                this.signedCheque.signature[0], 
                {from: sender}), "SimpleSwap: invalid serial")
            })
          })
        })
        context('when the serial is 0', function() {
          let serial = new BN(0)
          const zero_serial_unsignedCheque = Object.assign({}, defaultCheque, {serial: serial, signee: [defaultCheque.beneficiary, owner]})
          it('reverts', async function() {
            this.signedCheque = await signCheque(this.simpleSwap, zero_serial_unsignedCheque)
            await expectRevert(this.simpleSwap.submitCheque(
              this.signedCheque.beneficiary,
              this.signedCheque.serial, 
              this.signedCheque.amount, 
              this.signedCheque.timeout,
              this.signedCheque.signature[1],
              this.signedCheque.signature[0], 
              {from: sender}), "SimpleSwap: invalid serial")
          })
        })         
      }
      
      function shouldSubmitCheque(unsignedCheque, sender) {
        beforeEach(async function() {
          let lastCheque = await this.simpleSwap.cheques(unsignedCheque.beneficiary)
          expect(unsignedCheque.serial).bignumber.is.above(new BN(0), "serial is not positive")
          expect(unsignedCheque.amount).bignumber.to.be.above(new BN(0), "amount is not positive")
          expect(unsignedCheque.serial).bignumber.is.above(lastCheque.serial, "serial is not above the serial of the last submitted cheque")   
          expect(unsignedCheque.signee.length).to.equal(2, "no two signers present")
          this.signedCheque = await signCheque(this.simpleSwap, unsignedCheque)
          const { logs } = await this.simpleSwap.submitCheque(
            this.signedCheque.beneficiary, 
            this.signedCheque.serial, 
            this.signedCheque.amount, 
            this.signedCheque.timeout, 
            this.signedCheque.signature[1], 
            this.signedCheque.signature[0], 
            {from: sender}
          )
          this.logs = logs
        })
        context('uses _submitChequeInternal', function() {
          _shouldSubmitChequeInternal() 
        })
      }
    })
    describe('submitChequeBeneficiary', function() {
      context('when the sender is the beneficiary', function() {
        let sender = defaultCheque.beneficiary
        context('when the first serial is higher than 0', function() {
          expect(defaultCheque.serial).bignumber.to.be.above(new BN(0), "Serial of defaultCheque not above 0")
          context('when the first serial is below MAX_UINT256', function() {
            expect(defaultCheque.serial).bignumber.to.be.below(constants.MAX_UINT256, "Serial of defaultCheque not above 0")
            context('when the owner is a signee', function() {
              expect(defaultCheque.signee).to.be.equal(owner, "Signee of defaultCheque is not owner")
              context('when the signee signs the correct fields', function() {
                let unsignedCheque = Object.assign({}, defaultCheque)
                context('when we send one cheque', function() {
                  context('when there is a liquidBalance to cover the cheque', function() {
                    shouldDeposit(unsignedCheque.amount + new BN(1), owner)
                    shouldSubmitChequeBeneficiary(unsignedCheque, sender)
                  })
                  context('when there is no liquidBalance to cover the cheque', function() {
                    shouldSubmitChequeBeneficiary(unsignedCheque, sender)  
                  })
                })
                context('when we send more than one cheque', async function() {
                  shouldSubmitChequeBeneficiary(unsignedCheque, sender)
                  context('when the serial number is increasing', function() {
                    let secondSerial = new BN(parseInt(unsignedCheque.serial) + 1)
                    let increasing_serial_unsignedCheque = Object.assign({}, defaultCheque, {serial: secondSerial})
                    shouldSubmitChequeBeneficiary(increasing_serial_unsignedCheque, sender)
                  })
                  context('when the serial number stays the same', function() {
                    let secondSerial = new BN(parseInt(unsignedCheque.serial))
                    let same_serial_unsignedCheque = Object.assign({}, defaultCheque, {serial: secondSerial})
                    it('reverts', async function() {
                      this.signedCheque = await signCheque(this.simpleSwap, same_serial_unsignedCheque)
                      await expectRevert(this.simpleSwap.submitChequeBeneficiary(
                        this.signedCheque.serial, 
                        this.signedCheque.amount, 
                        this.signedCheque.timeout,
                        this.signedCheque.signature, {from: sender}), "SimpleSwap: invalid serial")
                    })
                  })
                  context('when the serial number is decreasing', function() {
                    let secondSerial = new BN(parseInt(unsignedCheque.serial) + -1)
                    let decreasing_serial_unsignedCheque = Object.assign({}, defaultCheque, {serial: secondSerial})
                    it('reverts', async function() {
                      this.signedCheque = await signCheque(this.simpleSwap, decreasing_serial_unsignedCheque)
                      await expectRevert(this.simpleSwap.submitChequeBeneficiary(
                        this.signedCheque.serial, 
                        this.signedCheque.amount, 
                        this.signedCheque.timeout,
                        this.signedCheque.signature, {from: sender}), "SimpleSwap: invalid serial")
                    })
                  })
                })
              })
              context('when the signee does not sign the correct fields', function() {
                let wrongBeneficiary = constants.ZERO_ADDRESS
                let wrong_beneficiary_unsignedCheque = Object.assign({}, defaultCheque, {beneficiary: wrongBeneficiary})
                it('reverts', async function() {
                  this.signedCheque = await signCheque(this.simpleSwap, wrong_beneficiary_unsignedCheque)
                  await expectRevert(this.simpleSwap.submitChequeBeneficiary(
                    this.signedCheque.serial, 
                    this.signedCheque.amount, 
                    this.signedCheque.timeout,
                    this.signedCheque.signature, {from: sender}), "SimpleSwap: invalid ownerSig")
                })
              })
            })
            context('when the owner is not the signee', function() {
              let signee = alice
              const wrong_signee_unsignedCheque = Object.assign({}, defaultCheque, {signee: signee})
              it('reverts', async function() {
                this.signedCheque = await signCheque(this.simpleSwap, wrong_signee_unsignedCheque)
                await expectRevert(this.simpleSwap.submitChequeBeneficiary(
                  this.signedCheque.serial, 
                  this.signedCheque.amount, 
                  this.signedCheque.timeout,
                  this.signedCheque.signature, {from: sender}), "SimpleSwap: invalid ownerSig")
              })
            })
          })
          context('when the first serial is at MAX_UINT256', function() {
            const MAX_UINT256_unsignedCheque = Object.assign({}, defaultCheque, {serial: constants.MAX_UINT256})
            shouldSubmitChequeBeneficiary(MAX_UINT256_unsignedCheque, defaultCheque.beneficiary)
            // Solidity wraps integers
            const MAX_UINT256_wrap_unsignedCheque = Object.assign({}, defaultCheque, {serial: MAX_UINT256_unsignedCheque.serial + new BN(1)})
            it('should not be possible to submit a cheque afterwards', async function() {
              this.signedCheque = await signCheque(this.simpleSwap, MAX_UINT256_wrap_unsignedCheque)
              await expectRevert(this.simpleSwap.submitChequeBeneficiary(
                this.signedCheque.serial, 
                this.signedCheque.amount, 
                this.signedCheque.timeout,
                this.signedCheque.signature, {from: sender}), "SimpleSwap: invalid serial")
            })
          })
        })
        context('when the serial is 0', function() {
          let serial = new BN(0)
          const zero_serial_unsignedCheque = Object.assign({}, defaultCheque, {serial: serial})
          it('reverts', async function() {
            this.signedCheque = await signCheque(this.simpleSwap, zero_serial_unsignedCheque)
            await expectRevert(this.simpleSwap.submitChequeBeneficiary(
              this.signedCheque.serial, 
              this.signedCheque.amount, 
              this.signedCheque.timeout,
              this.signedCheque.signature, {from: sender}), "SimpleSwap: invalid serial")
          })
        })         
      })
    })
    describe('submitChequeOwner', function() {
      context('when the sender is the owner', function() {
        let sender = owner
        context('when the first serial is higher than 0', function() {
          expect(defaultCheque.serial).bignumber.to.be.above(new BN(0), "Serial of defaultCheque not above 0")
          context('when the first serial is below MAX_UINT256', function() {
            expect(defaultCheque.serial).bignumber.to.be.below(constants.MAX_UINT256, "Serial of defaultCheque not above 0")
            context('when the beneficiary is a signee', function() {
              let unsignedCheque = Object.assign({}, defaultCheque, {signee: defaultCheque.beneficiary})
              expect(unsignedCheque.signee).to.be.equal(unsignedCheque.beneficiary, "Signee of unsignedCheque is not beneficiary")
              context('when the signee signs the correct fields', function() {
                context('when we send one cheque', function() {
                  context('when there is a liquidBalance to cover the cheque', function() {
                    shouldDeposit(unsignedCheque.amount + new BN(1), owner)
                    shouldSubmitChequeOwner(unsignedCheque, sender)
                  })
                  context('when there is no liquidBalance to cover the cheque', function() {
                    shouldSubmitChequeOwner(unsignedCheque, sender)  
                  })
                })
                context('when we send more than one cheque', async function() {
                  shouldSubmitChequeOwner(unsignedCheque, sender)
                  context('when the serial number is increasing', function() {
                    let secondSerial = new BN(parseInt(unsignedCheque.serial) + 1)
                    let increasing_serial_unsignedCheque = Object.assign({}, defaultCheque, {serial: secondSerial, signee: defaultCheque.beneficiary})
                    shouldSubmitChequeOwner(increasing_serial_unsignedCheque, sender)
                  })
                  context('when the serial number stays the same', function() {
                    let secondSerial = new BN(parseInt(unsignedCheque.serial))
                    let same_serial_unsignedCheque = Object.assign({}, defaultCheque, {serial: secondSerial, signee: defaultCheque.beneficiary})
                    it('reverts', async function() {
                      this.signedCheque = await signCheque(this.simpleSwap, same_serial_unsignedCheque)
                      await expectRevert(this.simpleSwap.submitChequeOwner(
                        this.signedCheque.beneficiary,
                        this.signedCheque.serial, 
                        this.signedCheque.amount, 
                        this.signedCheque.timeout,
                        this.signedCheque.signature, {from: sender}), "SimpleSwap: invalid serial")
                    })
                  })
                  context('when the serial number is decreasing', function() {
                    let secondSerial = new BN(parseInt(unsignedCheque.serial) + -1)
                    let decreasing_serial_unsignedCheque = Object.assign({}, defaultCheque, {serial: secondSerial, signee: defaultCheque.beneficiary})
                    it('reverts', async function() {
                      this.signedCheque = await signCheque(this.simpleSwap, decreasing_serial_unsignedCheque)
                      await expectRevert(this.simpleSwap.submitChequeOwner(
                        this.signedCheque.beneficiary,
                        this.signedCheque.serial, 
                        this.signedCheque.amount, 
                        this.signedCheque.timeout,
                        this.signedCheque.signature, {from: sender}), "SimpleSwap: invalid serial")
                    })
                  })
                })
              })
              context('when the signee does not sign the correct fields', function() {
                let wrongBeneficiary = constants.ZERO_ADDRESS
                let wrong_beneficiary_unsignedCheque = Object.assign({}, defaultCheque, {beneficiary: wrongBeneficiary, signee: defaultCheque.beneficiary})
                it('reverts', async function() {
                  this.signedCheque = await signCheque(this.simpleSwap, wrong_beneficiary_unsignedCheque)
                  await expectRevert(this.simpleSwap.submitChequeOwner(
                    this.signedCheque.beneficiary,
                    this.signedCheque.serial, 
                    this.signedCheque.amount, 
                    this.signedCheque.timeout,
                    this.signedCheque.signature, {from: sender}), "SimpleSwap: invalid beneficiarySig")
                })
              })
            })
            context('when the beneficiary is not the signee', function() {
              let signee = alice
              const wrong_signee_unsignedCheque = Object.assign({}, defaultCheque, {signee: signee})
              it('reverts', async function() {
                this.signedCheque = await signCheque(this.simpleSwap, wrong_signee_unsignedCheque)
                await expectRevert(this.simpleSwap.submitChequeOwner(
                  this.signedCheque.beneficiary,
                  this.signedCheque.serial, 
                  this.signedCheque.amount, 
                  this.signedCheque.timeout,
                  this.signedCheque.signature, {from: sender}), "SimpleSwap: invalid beneficiarySig")
              })
            })
          })
          context('when the first serial is at MAX_UINT256', function() {
            const MAX_UINT256_unsignedCheque = Object.assign({}, defaultCheque, {serial: constants.MAX_UINT256, signee: defaultCheque.beneficiary})
            shouldSubmitChequeOwner(MAX_UINT256_unsignedCheque, owner)
            // Solidity wraps integers
            const MAX_UINT256_wrap_unsignedCheque = Object.assign({}, defaultCheque, {serial: MAX_UINT256_unsignedCheque.serial + new BN(1), signee: defaultCheque.beneficiary})
            it('should not be possible to submit a cheque afterwards', async function() {
              this.signedCheque = await signCheque(this.simpleSwap, MAX_UINT256_wrap_unsignedCheque)
              await expectRevert(this.simpleSwap.submitChequeOwner(
                this.signedCheque.beneficiary,
                this.signedCheque.serial, 
                this.signedCheque.amount, 
                this.signedCheque.timeout,
                this.signedCheque.signature, {from: sender}), "SimpleSwap: invalid serial")
            })
          })
        })
        context('when the serial is 0', function() {
          let serial = new BN(0)
          const zero_serial_unsignedCheque = Object.assign({}, defaultCheque, {serial: serial, signee: defaultCheque.beneficiary})
          it('reverts', async function() {
            this.signedCheque = await signCheque(this.simpleSwap, zero_serial_unsignedCheque)
            await expectRevert(this.simpleSwap.submitChequeOwner(
              this.signedCheque.beneficiary,
              this.signedCheque.serial, 
              this.signedCheque.amount, 
              this.signedCheque.timeout,
              this.signedCheque.signature, {from: sender}), "SimpleSwap: invalid serial")
          })
        })         
      })
      function shouldSubmitChequeOwner(unsignedCheque, sender) {
        beforeEach(async function() {
          let lastCheque = await this.simpleSwap.cheques(unsignedCheque.beneficiary)
          expect(unsignedCheque.serial).bignumber.is.above(new BN(0), "serial is not positive")
          expect(unsignedCheque.amount).bignumber.to.be.above(new BN(0), "amount is not positive")
          expect(owner).to.equal(sender, "owner is not the sender")
          expect(unsignedCheque.serial).bignumber.is.above(lastCheque.serial, "serial is not above the serial of the last submitted cheque")   
          this.signedCheque = await signCheque(this.simpleSwap, unsignedCheque)
          const { logs } = await this.simpleSwap.submitChequeOwner(this.signedCheque.beneficiary, this.signedCheque.serial, this.signedCheque.amount, this.signedCheque.timeout, this.signedCheque.signature, {from: sender})
          this.logs = logs
        })
        context('uses _submitChequeInternal', function() {
          _shouldSubmitChequeInternal() 
        })
      }
    })
    describe('increaseHardDeposit', function() {
      let amount = new BN(50)
      let beneficiary = bob
      context('when the sender is the owner', function() {
        let sender = owner
        context('when the totalHardDeposit is below the swap balance', function() {
          shouldDeposit(amount.muln(2), owner)
          describe('when there is no prior deposit', function() {
            shouldIncreaseHardDeposit(sender, amount)
          })
          context('when there is a prior deposit', function() {
            shouldIncreaseHardDeposit(sender, amount)
            describe('when the totalHardDeposit is below the swap balance', function() {
              shouldIncreaseHardDeposit(sender, amount)
            })
          })
        })
        context('when the totalHardDeposit exceeds the swap balance', function() {
          it('reverts', async function() {
            await expectRevert(this.simpleSwap.increaseHardDeposit(
              bob,
              new BN(amount),
              { from: sender }), "SimpleSwap: hard deposit cannot be more than balance ")
          })
        })
        function shouldIncreaseHardDeposit(sender, amount) {
          beforeEach(async function() {
            this.previousTotalHardDeposit = await this.simpleSwap.totalHardDeposit()
            this.previousHardDeposit = (await this.simpleSwap.hardDeposits(beneficiary))[0]
            let { logs } = await this.simpleSwap.increaseHardDeposit(
              beneficiary,
              amount,
              { from: sender }
            )
            this.logs = logs
          })

          it('should fire the HardDepositAmountChanged event', async function() {
            expectEvent.inLogs(this.logs, 'HardDepositAmountChanged', {
              beneficiary,
              amount: this.previousHardDeposit.add(amount)
            })
          })
          it('increases the totalHardDeposit', async function() {
            expect(await this.simpleSwap.totalHardDeposit()).bignumber.is.equal(this.previousTotalHardDeposit.add(amount))
          })
          it('increases the hardDeposit amount', async function() {
            expect((await this.simpleSwap.hardDeposits(beneficiary))[0]).bignumber.is.equal(this.previousHardDeposit.add(amount))
          })
          it('reset the canBeDecreasedAt  value', async function() {
            expect((await this.simpleSwap.hardDeposits(beneficiary))[3]).bignumber.is.equal(new BN(0))
          })
        }
      })
      context('when the sender is not the owner', function() {
        let sender = bob
        it('reverts', async function() {
          await expectRevert(this.simpleSwap.increaseHardDeposit(
            bob,
            new BN(amount),
            { from: sender }), "SimpleSwap: not owner")
        })
      })
    })

    describe('prepareDecreaseHardDeposit', function() {
      let amount = new BN(50)
      let beneficiary = bob
      context('when the sender is the owner', function() {
        context('when the hard deposit is high enough', function() {
          context('when no custom decreaseTimeout is set', function() {
            decreaseHardDeposit()
          })
          context('when a custom decreaseTimeout is set', function() {
            let decreaseTimeout = new BN(100)
            beforeEach(async function() {
              const data = web3.utils.keccak256(web3.eth.abi.encodeParameters(['address', 'address', 'uint256'], [this.simpleSwap.address, beneficiary, decreaseTimeout.toString()]))              
              await this.simpleSwap.setCustomHardDepositDecreaseTimeout(
                beneficiary,
                decreaseTimeout, 
                await sign(data, owner),
                await sign(data, beneficiary), {
                  from: owner
                })
            })
            decreaseHardDeposit()
          })
          
          function decreaseHardDeposit() {
            beforeEach(async function() {
              await this.simpleSwap.send(amount)
              await this.simpleSwap.increaseHardDeposit(beneficiary, amount)

              let { logs } = await this.simpleSwap.prepareDecreaseHardDeposit(
                beneficiary,
                amount, {
                  from: owner
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
        })
        context('when the hard deposit is not high enough', function() {
          beforeEach(async function() {
            await this.simpleSwap.send(amount)
            await this.simpleSwap.increaseHardDeposit(beneficiary, amount.divn(2))
          })
          it('reverts', async function() {
            await expectRevert(this.simpleSwap.prepareDecreaseHardDeposit(
              beneficiary,
              amount, {
                from: owner
            }), "SimpleSwap: hard deposit not sufficient")
          })

        })
      })
      context('when the sender is not the owner', function() {
        let sender = bob
        it('reverts', async function() {
          await expectRevert(this.simpleSwap.prepareDecreaseHardDeposit(
            beneficiary,
            amount,
            { from: sender }), "SimpleSwap: not owner")
        })
      })
    })

    describe('setCustomHardDepositDecreaseTimeout', function() {
      let beneficiary = bob
      let decreaseTimeout = new BN(10)
      beforeEach(function() {
        this.data = web3.utils.keccak256(web3.eth.abi.encodeParameters(['address', 'address', 'uint256'], [this.simpleSwap.address, beneficiary, decreaseTimeout.toString()]))
      })
      describe('when both signature are valid', function() {
        beforeEach(async function() {
          let { logs } = await this.simpleSwap.setCustomHardDepositDecreaseTimeout(
            beneficiary,
            decreaseTimeout,
            await sign(this.data, owner),
            await sign(this.data, beneficiary)
          )

          this.logs = logs
        })

        it('should set the decreaseTimeout', async function() {
          expect((await this.simpleSwap.hardDeposits(beneficiary))[2]).bignumber.is.equal(decreaseTimeout)
        })

        it('should fire the HardDepositDecreaseTimeoutChanged', async function() {
          expectEvent.inLogs(this.logs, 'HardDepositDecreaseTimeoutChanged', {
            beneficiary,
            decreaseTimeout
          })
        })
      })
      context('when ownerSig invalid', function() {
        it('reverts', async function() {
          await expectRevert.unspecified(this.simpleSwap.setCustomHardDepositDecreaseTimeout(
            beneficiary,
            decreaseTimeout,
            '0x',
            await sign(this.data, beneficiary)
          ))
        })
      })
      context('when beneficiarySig invalid', function() {
        it('reverts', async function() {
          await expectRevert.unspecified(this.simpleSwap.setCustomHardDepositDecreaseTimeout(
            beneficiary,
            decreaseTimeout,
            await sign(this.data, owner),
            '0x'
          ))
        })
      })
    })

    describe('decreaseHardDeposit', function() {
      let beneficiary = bob
      let amount = new BN(500)
      let decrease = new BN(400)
      beforeEach(async function() {
        await this.simpleSwap.send(amount)
        await this.simpleSwap.increaseHardDeposit(beneficiary, amount)
      })
      context('when decrease is ready', function() {
        context('when there is enough hard deposit left', function() {
          beforeEach(async function() {
            await this.simpleSwap.prepareDecreaseHardDeposit(beneficiary, decrease)
            await time.increase(await this.simpleSwap.DEFAULT_HARDDEPPOSIT_DECREASE_TIMEOUT())            
            let { logs } = await this.simpleSwap.decreaseHardDeposit(beneficiary)
            this.logs = logs
          })

          it('should fire the HardDepositAmountChanged event', async function() {
            expectEvent.inLogs(this.logs, 'HardDepositAmountChanged', {
              beneficiary,
              amount: amount.sub(decrease)
            })
          })

          it('should set the new amount', async function() {
            expect((await this.simpleSwap.hardDeposits(beneficiary))[0]).bignumber.is.equal(amount.sub(decrease))
          })
        })
        // TODO: when there is not enough left
      })

      context('when timeout not yet expired', function() {
        beforeEach(async function() {
          await this.simpleSwap.prepareDecreaseHardDeposit(beneficiary, amount)
        })
        it('reverts', async function() {
          await expectRevert(
            this.simpleSwap.decreaseHardDeposit(beneficiary, { from: owner }),
            "SimpleSwap: deposit not yet timed out"
          )
        })
      })

      context('when no decrease prepared', async function() {
        it('reverts', async function() {
          await expectRevert(
            this.simpleSwap.decreaseHardDeposit(beneficiary, { from: owner }),
            "SimpleSwap: deposit not yet timed out"
          )
        })
      })
    })

    describe('withdraw', function() {
      let amount = new BN(100)

      beforeEach(async function() {
        await this.simpleSwap.send(amount)
      })

      context('when the sender is the owner', function() {
        let sender = owner
        context('when the liquid balance is high enough', function() {
          beforeEach(async function() {
            let ownerBalancePrior = await balance.current(owner)
            let { logs, receipt } = await this.simpleSwap.withdraw(
              amount,
              { from: sender }
            )

            this.logs = logs
            this.expectedBalance = ownerBalancePrior.add(amount).sub(await computeCost(receipt))
          })

          it('should change the owner balance correctly', async function() {
            expect(await balance.current(owner)).bignumber.is.equal(this.expectedBalance)
          })
        })
        context('when the liquid balance is too low', function() {
          beforeEach(async function() {
            await this.simpleSwap.increaseHardDeposit(bob, new BN(1), { from: sender })
          })

          it('reverts', async function() {
            await expectRevert(this.simpleSwap.withdraw(amount, {
              from: sender
            }), "SimpleSwap: liquidBalance not sufficient")
          })
        })
      })

      context('when the sender is not the owner', function() {
        let sender = bob
        it('reverts', async function() {
          await expectRevert(this.simpleSwap.withdraw(amount, {
            from: sender
          }), 'SimpleSwap: not owner')
        })
      })
    })

    describe('cashcheque', function() {

    })
    describe('cashChequeBeneficiary', function() {
      let cheque = defaultCheque
      shouldDeposit(defaultCheque.amount, owner)
      shouldSubmitChequeBeneficiary(cheque, unsignedCheque.beneficiary)
      time.increase(new BN(86400))
      _shouldCashChequeInternal(unsignedCheque.beneficiary, unsignedCheque.beneficiary, defaultCheque.amount, 0)
    })
  })

  function _shouldCashChequeInternal() {
    // it should test all updates to state variables and tests

  }
  function _shouldSubmitChequeInternal() {    
    beforeEach(async function() {
      this.currentCheque = await this.simpleSwap.cheques(this.signedCheque.beneficiary)
    })
    it('should update the cheque serial number', async function() {
      expect(this.currentCheque.serial).bignumber.is.equal(this.signedCheque.serial, "serial was not updated")
    })
    it('should update the cheque amount', async function() {
      expect(this.currentCheque.amount).bignumber.is.equal(this.signedCheque.amount, "amount was not updated")
    })
    it('should update the cheque timeout', async function() {
      expect(parseInt(this.currentCheque.cashTimeout)).is.equal(parseInt(await time.latest()) + parseInt(this.signedCheque.timeout))
    })
    it('should emit a chequeSubmitted event', async function() {
      expectEvent.inLogs(this.logs, "ChequeSubmitted", {
        amount: this.signedCheque.amount,
        beneficiary: this.signedCheque.beneficiary,
        serial: this.signedCheque.serial
      })
    })
  }
  function shouldDeposit(amount, sender) {
    beforeEach(async function() {
      this.currentBalance = await balance.current(this.simpleSwap.address)
      this.currentTotalHardDeposit = await this.simpleSwap.totalHardDeposit()
      this.liquidBalance = this.currentBalance - this.currentTotalHardDeposit
      const { logs } = await this.simpleSwap.send(amount, {from: sender})
      this.depositLogs = logs
    })
    it('should update the liquidBalance of the checkbook', async function() {
      expect(await balance.current(this.simpleSwap.address)).bignumber.to.equal(this.currentBalance + amount)
    })
    it('should update the balance of the checkbook', async function() {
      expect(await balance.current(this.simpleSwap.address)).bignumber.to.equal(this.currentBalance + amount)
    })
    it('should not afect the totalHardDeposit', async function() {
      expect(await this.simpleSwap.totalHardDeposit()).bignumber.to.equal(this.currentTotalHardDeposit)
    })
    it('should emit a deposit event', async function() {
      expectEvent.inLogs(this.depositLogs, "Deposit", {
        depositor: sender,
        amount: amount
      })
    })
  }

  function shouldSubmitChequeBeneficiary(unsignedCheque, sender) {
    beforeEach(async function() {
      let lastCheque = await this.simpleSwap.cheques(unsignedCheque.beneficiary)
      expect(unsignedCheque.serial).bignumber.is.above(new BN(0), "serial is not positive")
      expect(unsignedCheque.amount).bignumber.to.be.above(new BN(0), "amount is not positive")
      expect(unsignedCheque.beneficiary).to.equal(sender, "beneficiary is not the sender")
      expect(unsignedCheque.serial).bignumber.is.above(lastCheque.serial, "serial is not above the serial of the last submitted cheque")   
      this.signedCheque = await signCheque(this.simpleSwap, unsignedCheque)
      const { logs } = await this.simpleSwap.submitChequeBeneficiary(this.signedCheque.serial, this.signedCheque.amount, this.signedCheque.timeout, this.signedCheque.signature, {from: sender})
      this.logs = logs
    })
    context('uses _submitChequeInternal', function() {
      _shouldSubmitChequeInternal() 
    })
  }
}

module.exports = {
  shouldBehaveLikeSimpleSwap
};