const {
    BN,
    balance,
    time,
    expectRevert,
    constants,
    expectEvent
} = require("openzeppelin-test-helpers");

const { expect } = require('chai');

const { signCheque } = require("./swutils");



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
  describe('as a simple swap', function() {
    it('should have a correct owner', async function() {
      expect(await this.simpleSwap.owner()).to.equal(owner)          
    })
    describe('deposit', function() {
      shouldDeposit(new BN(1), owner)
    })
    describe('submitCheque', function() {
      describe('when the sender is the owner', function() {
        submitChequeBySender(owner)
      })
      describe('when the sender is the beneficiary', function() {
        submitChequeBySender(defaultCheque.beneficiary)  
      })
      describe('when the sender is a third party', function() {
        submitChequeBySender(alice)   
      })
      function submitChequeBySender(sender) {
        describe('when the first serial is higher than 0', function() {
          expect(defaultCheque.serial).bignumber.to.be.above(new BN(0), "Serial of defaultCheque not above 0")
          describe('when the first serial is below MAX_UINT256', function() {
            expect(defaultCheque.serial).bignumber.to.be.below(constants.MAX_UINT256, "Serial of defaultCheque not above 0")
            describe('when the beneficiary and owners are both a signee', function() {
              let unsignedCheque = Object.assign({}, defaultCheque, {signee: [defaultCheque.beneficiary, owner]})
              expect(unsignedCheque.signee).to.be.include(unsignedCheque.beneficiary, "Signee of unsignedCheque is not beneficiary")
              expect(unsignedCheque.signee).to.be.include(owner, "Signee of unsignedCheque is not owner")
              describe('when the signees signs the correct fields', function() {
                describe('when we send one cheque', function() {
                  describe('when there is a liquidBalance to cover the cheque', function() {
                    shouldDeposit(unsignedCheque.amount + new BN(1), owner)
                    shouldSubmitCheque(unsignedCheque, sender)
                  })
                  describe('when there is no liquidBalance to cover the cheque', function() {
                    shouldSubmitCheque(unsignedCheque, sender)  
                  })
                })
                describe('when we send more than one cheque', async function() {
                  shouldSubmitCheque(unsignedCheque, sender)
                  describe('when the serial number is increasing', function() {
                    let secondSerial = new BN(parseInt(unsignedCheque.serial) + 1)
                    let increasing_serial_unsignedCheque = Object.assign({}, defaultCheque, {serial: secondSerial, signee: [defaultCheque.beneficiary, owner]})
                    shouldSubmitCheque(increasing_serial_unsignedCheque, sender)
                  })
                  describe('when the serial number stays the same', function() {
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
                  describe('when the serial number is decreasing', function() {
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
              describe('when the signee does not sign the correct fields', function() {
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
            describe('when the owner is not a signee', function() {
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
            describe('when the beneficiary is not a signee', function() {
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
          describe('when the first serial is at MAX_UINT256', function() {
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
        describe('when the serial is 0', function() {
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
        describe('uses _submitChequeInternal', function() {
          _shouldSubmitChequeInternal() 
        })
      }
    })
    describe('submitChequeBeneficiary', function() {
      describe('when the sender is the beneficiary', function() {
        let sender = defaultCheque.beneficiary
        describe('when the first serial is higher than 0', function() {
          expect(defaultCheque.serial).bignumber.to.be.above(new BN(0), "Serial of defaultCheque not above 0")
          describe('when the first serial is below MAX_UINT256', function() {
            expect(defaultCheque.serial).bignumber.to.be.below(constants.MAX_UINT256, "Serial of defaultCheque not above 0")
            describe('when the owner is a signee', function() {
              expect(defaultCheque.signee).to.be.equal(owner, "Signee of defaultCheque is not owner")
              describe('when the signee signs the correct fields', function() {
                let unsignedCheque = Object.assign({}, defaultCheque)
                describe('when we send one cheque', function() {
                  describe('when there is a liquidBalance to cover the cheque', function() {
                    shouldDeposit(unsignedCheque.amount + new BN(1), owner)
                    shouldSubmitChequeBeneficiary(unsignedCheque, sender)
                  })
                  describe('when there is no liquidBalance to cover the cheque', function() {
                    shouldSubmitChequeBeneficiary(unsignedCheque, sender)  
                  })
                })
                describe('when we send more than one cheque', async function() {
                  shouldSubmitChequeBeneficiary(unsignedCheque, sender)
                  describe('when the serial number is increasing', function() {
                    let secondSerial = new BN(parseInt(unsignedCheque.serial) + 1)
                    let increasing_serial_unsignedCheque = Object.assign({}, defaultCheque, {serial: secondSerial})
                    shouldSubmitChequeBeneficiary(increasing_serial_unsignedCheque, sender)
                  })
                  describe('when the serial number stays the same', function() {
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
                  describe('when the serial number is decreasing', function() {
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
              describe('when the signee does not sign the correct fields', function() {
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
            describe('when the owner is not the signee', function() {
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
          describe('when the first serial is at MAX_UINT256', function() {
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
        describe('when the serial is 0', function() {
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
        describe('uses _submitChequeInternal', function() {
          _shouldSubmitChequeInternal() 
        })
      }
    })
    describe('submitChequeOwner', function() {
      describe('when the sender is the owner', function() {
        let sender = owner
        describe('when the first serial is higher than 0', function() {
          expect(defaultCheque.serial).bignumber.to.be.above(new BN(0), "Serial of defaultCheque not above 0")
          describe('when the first serial is below MAX_UINT256', function() {
            expect(defaultCheque.serial).bignumber.to.be.below(constants.MAX_UINT256, "Serial of defaultCheque not above 0")
            describe('when the beneficiary is a signee', function() {
              let unsignedCheque = Object.assign({}, defaultCheque, {signee: defaultCheque.beneficiary})
              expect(unsignedCheque.signee).to.be.equal(unsignedCheque.beneficiary, "Signee of unsignedCheque is not beneficiary")
              describe('when the signee signs the correct fields', function() {
                describe('when we send one cheque', function() {
                  describe('when there is a liquidBalance to cover the cheque', function() {
                    shouldDeposit(unsignedCheque.amount + new BN(1), owner)
                    shouldSubmitChequeOwner(unsignedCheque, sender)
                  })
                  describe('when there is no liquidBalance to cover the cheque', function() {
                    shouldSubmitChequeOwner(unsignedCheque, sender)  
                  })
                })
                describe('when we send more than one cheque', async function() {
                  shouldSubmitChequeOwner(unsignedCheque, sender)
                  describe('when the serial number is increasing', function() {
                    let secondSerial = new BN(parseInt(unsignedCheque.serial) + 1)
                    let increasing_serial_unsignedCheque = Object.assign({}, defaultCheque, {serial: secondSerial, signee: defaultCheque.beneficiary})
                    shouldSubmitChequeOwner(increasing_serial_unsignedCheque, sender)
                  })
                  describe('when the serial number stays the same', function() {
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
                  describe('when the serial number is decreasing', function() {
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
              describe('when the signee does not sign the correct fields', function() {
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
            describe('when the beneficiary is not the signee', function() {
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
          describe('when the first serial is at MAX_UINT256', function() {
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
        describe('when the serial is 0', function() {
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
        describe('uses _submitChequeInternal', function() {
          _shouldSubmitChequeInternal() 
        })
      }
    })
  })
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
      expect(parseInt(this.currentCheque.timeout)).is.equal(parseInt(await time.latest()) + parseInt(this.signedCheque.timeout))
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
}

module.exports = {
  shouldBehaveLikeSimpleSwap
};