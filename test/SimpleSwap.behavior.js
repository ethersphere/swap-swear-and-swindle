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


enabledTests = {
  issuer: true,
  deposit: true,
  submitCheque: true,
  submitChequeBeneficiary: true,
  submitChequeissuer: true,
  increaseHardDeposit: true,
  cashCheque: true,
  cashChequeBeneficiary: true

}

// @param balance total ether deposited in checkbook
// @param liquidBalance totalDeposit - harddeposits
// @param issuer the issuer of the checkbook
// @param alice a counterparty of the checkbook 
// @param bob a counterparty of the checkbook
function shouldBehaveLikeSimpleSwap([issuer, alice, bob]) {
  const defaultCheque = {
    beneficiary: bob,
    serial: new BN(3),
    amount: new BN(Math.floor(Math.random() * 100000)),
    timeout: new BN(86400),
    signee: issuer,
    signature: ""
  }
  context('as a simple swap', function() {
    if(enabledTests.issuer) {
      it('should have a correct issuer', async function() {
        expect(await this.simpleSwap.issuer()).to.equal(issuer)          
      })
    }
    describe('deposit', function() {
      if(enabledTests.deposit) {  
        shouldDeposit(new BN(1), issuer)
      }
    })
    describe('submitCheque', function() {
      if(enabledTests.submitCheque) {
        context('when the sender is the issuer', function() {
          submitChequeBySender(issuer)
        })
        context('when the sender is the beneficiary', function() {
          submitChequeBySender(defaultCheque.beneficiary)  
        })
        context('when the sender is a third party', function() {
          submitChequeBySender(alice)   
        })
      }
      function submitChequeBySender(sender) {
        context('when the first serial is higher than 0', function() {
          expect(defaultCheque.serial).bignumber.to.be.above(new BN(0), "Serial of defaultCheque not above 0")
          context('when the first serial is below MAX_UINT256', function() {
            expect(defaultCheque.serial).bignumber.to.be.below(constants.MAX_UINT256, "Serial of defaultCheque not above 0")
            context('when the beneficiary and issuers are both a signee', function() {
              let unsignedCheque = Object.assign({}, defaultCheque, {signee: [defaultCheque.beneficiary, issuer]})
              expect(unsignedCheque.signee).to.be.include(unsignedCheque.beneficiary, "Signee of unsignedCheque is not beneficiary")
              expect(unsignedCheque.signee).to.be.include(issuer, "Signee of unsignedCheque is not issuer")
              context('when the signees signs the correct fields', function() {
                context('when we send one cheque', function() {
                  context('when there is a liquidBalance to cover the cheque', function() {
                    shouldDeposit(unsignedCheque.amount + new BN(1), issuer)
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
                    let increasing_serial_unsignedCheque = Object.assign({}, defaultCheque, {serial: secondSerial, signee: [defaultCheque.beneficiary, issuer]})
                    shouldSubmitCheque(increasing_serial_unsignedCheque, sender)
                  })
                  context('when the serial number stays the same', function() {
                    let secondSerial = new BN(parseInt(unsignedCheque.serial))
                    let same_serial_unsignedCheque = Object.assign({}, defaultCheque, {serial: secondSerial, signee: [defaultCheque.beneficiary, issuer]})
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
                    let decreasing_serial_unsignedCheque = Object.assign({}, defaultCheque, {serial: secondSerial, signee: [defaultCheque.beneficiary, issuer]})
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
                let wrong_beneficiary_unsignedCheque = Object.assign({}, defaultCheque, {beneficiary: wrongBeneficiary, signee: [defaultCheque.beneficiary, issuer]})
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
            context('when the issuer is not a signee', function() {
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
                  {from: sender}), "SimpleSwap: invalid issuerSig")
              })
            })
            context('when the beneficiary is not a signee', function() {
              const wrong_signee_unsignedCheque = Object.assign({}, defaultCheque, {signee: [issuer, issuer]})
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
            const MAX_UINT256_unsignedCheque = Object.assign({}, defaultCheque, {serial: constants.MAX_UINT256, signee: [defaultCheque.beneficiary, issuer]})
            shouldSubmitCheque(MAX_UINT256_unsignedCheque, sender)
            // Solidity wraps integers
            const MAX_UINT256_wrap_unsignedCheque = Object.assign({}, defaultCheque, {serial: MAX_UINT256_unsignedCheque.serial + new BN(1), signee: [defaultCheque.beneficiary, issuer]})
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
          const zero_serial_unsignedCheque = Object.assign({}, defaultCheque, {serial: serial, signee: [defaultCheque.beneficiary, issuer]})
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
      if(enabledTests.submitChequeBeneficiary) {
        context('when the sender is the beneficiary', function() {
          let sender = defaultCheque.beneficiary
          context('when the first serial is higher than 0', function() {
            expect(defaultCheque.serial).bignumber.to.be.above(new BN(0), "Serial of defaultCheque not above 0")
            context('when the first serial is below MAX_UINT256', function() {
              expect(defaultCheque.serial).bignumber.to.be.below(constants.MAX_UINT256, "Serial of defaultCheque not above 0")
              context('when the issuer is a signee', function() {
                expect(defaultCheque.signee).to.be.equal(issuer, "Signee of defaultCheque is not issuer")
                context('when the signee signs the correct fields', function() {
                  let unsignedCheque = Object.assign({}, defaultCheque)
                  context('when we send one cheque', function() {
                    context('when there is a liquidBalance to cover the cheque', function() {
                      shouldDeposit(unsignedCheque.amount + new BN(1), issuer)
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
                      this.signedCheque.signature, {from: sender}), "SimpleSwap: invalid issuerSig")
                  })
                })
              })
              context('when the issuer is not the signee', function() {
                let signee = alice
                const wrong_signee_unsignedCheque = Object.assign({}, defaultCheque, {signee: signee})
                it('reverts', async function() {
                  this.signedCheque = await signCheque(this.simpleSwap, wrong_signee_unsignedCheque)
                  await expectRevert(this.simpleSwap.submitChequeBeneficiary(
                    this.signedCheque.serial, 
                    this.signedCheque.amount, 
                    this.signedCheque.timeout,
                    this.signedCheque.signature, {from: sender}), "SimpleSwap: invalid issuerSig")
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
        context('when the sender is not the beneficiary', function() {
          //TODO: reverts
        })
      }
    })
    describe('submitChequeissuer', function() {
      if(enabledTests.submitChequeissuer) {
        context('when the sender is the issuer', function() {
          let sender = issuer
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
                      shouldDeposit(unsignedCheque.amount + new BN(1), issuer)
                      shouldSubmitChequeissuer(unsignedCheque, sender)
                    })
                    context('when there is no liquidBalance to cover the cheque', function() {
                      shouldSubmitChequeissuer(unsignedCheque, sender)  
                    })
                  })
                  context('when we send more than one cheque', async function() {
                    shouldSubmitChequeissuer(unsignedCheque, sender)
                    context('when the serial number is increasing', function() {
                      let secondSerial = new BN(parseInt(unsignedCheque.serial) + 1)
                      let increasing_serial_unsignedCheque = Object.assign({}, defaultCheque, {serial: secondSerial, signee: defaultCheque.beneficiary})
                      shouldSubmitChequeissuer(increasing_serial_unsignedCheque, sender)
                    })
                    context('when the serial number stays the same', function() {
                      let secondSerial = new BN(parseInt(unsignedCheque.serial))
                      let same_serial_unsignedCheque = Object.assign({}, defaultCheque, {serial: secondSerial, signee: defaultCheque.beneficiary})
                      it('reverts', async function() {
                        this.signedCheque = await signCheque(this.simpleSwap, same_serial_unsignedCheque)
                        await expectRevert(this.simpleSwap.submitChequeissuer(
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
                        await expectRevert(this.simpleSwap.submitChequeissuer(
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
                    await expectRevert(this.simpleSwap.submitChequeissuer(
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
                  await expectRevert(this.simpleSwap.submitChequeissuer(
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
              shouldSubmitChequeissuer(MAX_UINT256_unsignedCheque, issuer)
              // Solidity wraps integers
              const MAX_UINT256_wrap_unsignedCheque = Object.assign({}, defaultCheque, {serial: MAX_UINT256_unsignedCheque.serial + new BN(1), signee: defaultCheque.beneficiary})
              it('should not be possible to submit a cheque afterwards', async function() {
                this.signedCheque = await signCheque(this.simpleSwap, MAX_UINT256_wrap_unsignedCheque)
                await expectRevert(this.simpleSwap.submitChequeissuer(
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
              await expectRevert(this.simpleSwap.submitChequeissuer(
                this.signedCheque.beneficiary,
                this.signedCheque.serial, 
                this.signedCheque.amount, 
                this.signedCheque.timeout,
                this.signedCheque.signature, {from: sender}), "SimpleSwap: invalid serial")
            })
          })         
        })
        context('when the sender is not the issuer', function() {
          //TODO: reverts
        })
      }
      function shouldSubmitChequeissuer(unsignedCheque, sender) {
        beforeEach(async function() {
          let lastCheque = await this.simpleSwap.cheques(unsignedCheque.beneficiary)
          expect(unsignedCheque.serial).bignumber.is.above(new BN(0), "serial is not positive")
          expect(unsignedCheque.amount).bignumber.to.be.above(new BN(0), "amount is not positive")
          expect(issuer).to.equal(sender, "issuer is not the sender")
          expect(unsignedCheque.serial).bignumber.is.above(lastCheque.serial, "serial is not above the serial of the last submitted cheque")   
          this.signedCheque = await signCheque(this.simpleSwap, unsignedCheque)
          const { logs } = await this.simpleSwap.submitChequeissuer(this.signedCheque.beneficiary, this.signedCheque.serial, this.signedCheque.amount, this.signedCheque.timeout, this.signedCheque.signature, {from: sender})
          this.logs = logs
        })
        context('uses _submitChequeInternal', function() {
          _shouldSubmitChequeInternal() 
        })
      }
    })
    describe('increaseHardDeposit', function() {
      if(enabledTests.increaseHardDeposit) {
        let amount = new BN(50)
        let beneficiary = bob
        context('when the sender is the issuer', function() {
          let sender = issuer
          context('when the totalHardDeposit is below the swap balance', function() {
            shouldDeposit(amount.muln(2), issuer)
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
        context('when the sender is not the issuer', function() {
          let sender = bob
          it('reverts', async function() {
            await expectRevert(this.simpleSwap.increaseHardDeposit(
              bob,
              new BN(amount),
              { from: sender }), "SimpleSwap: not issuer")
          })
        })
      }
    })
    describe('cashCheque', function() {
      let cheque = defaultCheque
      context('when there is sufficient balance in the chequebook', function() {
        shouldDeposit(defaultCheque.amount, issuer)
        context('when the beneficiary has submitted a cheque', function() {
          shouldSubmitChequeBeneficiary(cheque, cheque.beneficiary)
          context('when sufficient time has passed', function() {
            before(function() {time.increase(new BN(86400))})
            const sender = alice
            it('is here', function() {

            })
          })
        })
      })
    })
    describe('liquidBalanceFor', function() {
      let beneficiary = bob
      let amount = new BN(50)
      beforeEach(async function() {
        await this.simpleSwap.send(amount)
      })

      context('when no hard deposits', function() {
        it('should equal the liquid balance', async function() {
          expect(await this.simpleSwap.liquidBalanceFor(beneficiary)).bignumber.equal(await this.simpleSwap.liquidBalance())
        })
      })

      context('when hard deposits', function() {
        let hardDeposit = new BN(10)
        beforeEach(async function() {
          await this.simpleSwap.increaseHardDeposit(beneficiary, hardDeposit)
          await this.simpleSwap.increaseHardDeposit(alice, hardDeposit)
        })
        it('should be higher than the liquid balance', async function() {
          expect(await this.simpleSwap.liquidBalanceFor(beneficiary)).bignumber.equal((await this.simpleSwap.liquidBalance()).add(hardDeposit))
        })
      })
    })
    describe('prepareDecreaseHardDeposit', function() {
      let amount = new BN(50)
      let beneficiary = bob
      context('when the sender is the issuer', function() {
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
                await sign(data, issuer),
                await sign(data, beneficiary), {
                  from: issuer
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
                from: issuer
            }), "SimpleSwap: hard deposit not sufficient")
          })

        })
      })
      context('when the sender is not the issuer', function() {
        let sender = bob
        it('reverts', async function() {
          await expectRevert(this.simpleSwap.prepareDecreaseHardDeposit(
            beneficiary,
            amount,
            { from: sender }), "SimpleSwap: not issuer")
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
            await sign(this.data, issuer),
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
      context('when issuerSig invalid', function() {
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
            await sign(this.data, issuer),
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
            this.simpleSwap.decreaseHardDeposit(beneficiary, { from: issuer }),
            "SimpleSwap: deposit not yet timed out"
          )
        })
      })

      context('when no decrease prepared', async function() {
        it('reverts', async function() {
          await expectRevert(
            this.simpleSwap.decreaseHardDeposit(beneficiary, { from: issuer }),
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

      context('when the sender is the issuer', function() {
        let sender = issuer
        context('when the liquid balance is high enough', function() {
          beforeEach(async function() {
            let issuerBalancePrior = await balance.current(issuer)
            let { logs, receipt } = await this.simpleSwap.withdraw(
              amount,
              { from: sender }
            )

            this.logs = logs
            this.expectedBalance = issuerBalancePrior.add(amount).sub(await computeCost(receipt))
          })

          it('should change the issuer balance correctly', async function() {
            expect(await balance.current(issuer)).bignumber.is.equal(this.expectedBalance)
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

      context('when the sender is not the issuer', function() {
        let sender = bob
        it('reverts', async function() {
          await expectRevert(this.simpleSwap.withdraw(amount, {
            from: sender
          }), 'SimpleSwap: not issuer')
        })
      })
    })

    describe('cashcheque', function() {

    })
    describe('cashChequeBeneficiary', function() {
      if(enabledTests.cashChequeBeneficiary) {
        let cheque = defaultCheque
        context('when there is sufficient balance in the chequebook', function() {
          shouldDeposit(defaultCheque.amount, issuer)
          context('when the beneficiary has submitted a cheque', function() {
            shouldSubmitChequeBeneficiary(cheque, cheque.beneficiary)
            context('when sufficient time has passed', function() {
              beforeEach(async function() {
                time.increase(new BN(86400))
                this.currentBeneficiaryBalance = await balance.current(cheque.beneficiary)
                this.currentCheque = await this.simpleSwap.cheques(cheque.beneficiary)
                const { logs, receipt } = await this.simpleSwap.cashChequeBeneficiary(cheque.beneficiary, cheque.amount, {from: cheque.beneficiary})
                this.logs = logs
                this.receipt = receipt
              })
              _shouldCashChequeInternal(cheque.beneficiary, cheque.beneficiary, defaultCheque.amount, new BN(0))
  
            })
          })
        }) 
      }
    })
  })
  function _shouldCashChequeInternal(benefciaryPrincipal, beneficiaryAgent, amount, calleepayout) {
    it('should update the harddeposit usage', function() {
      //TODO => minor importance
    })
    it('should update paidOut', async function() {
      expect((await this.simpleSwap.cheques(benefciaryPrincipal)).paidOut).bignumber.to.be.equal(this.currentCheque.paidOut.add(amount).add(calleepayout), "Did not update paidOut")
    })
    it('should transfer the correct amount to the beneficiaryAgent', async function() {
      expect(await balance.current(beneficiaryAgent)).bignumber.to.be.equal(this.currentBeneficiaryBalance.add(amount).sub(calleepayout).sub(await computeCost(this.receipt)))
    })
    it('should emit a ChequeCashed event', function() {
      expectEvent.inLogs(this.logs, "ChequeCashed", {
        beneficiaryPrincipal: benefciaryPrincipal,
        beneficiaryAgent: beneficiaryAgent,
        callee: benefciaryPrincipal,
        serial: this.currentCheque.serial,
        totalPayout: this.currentCheque.amount,
        requestPayout: this.currentCheque.amount,
        calleePayout: new BN(0)
      })
    })
    it('should emit a ChequeBounced event', function() {
      //TODO => less important => only if there is no balance
    })

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