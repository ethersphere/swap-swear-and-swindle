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
  const defaultCheck = {
    beneficiary: bob,
    serial: new BN(3),
    amount: new BN(Math.floor(Math.random() * 100000)),
    timeout: new BN(Math.floor(Math.random() * 100000)),
    signee: owner,
    signature: ""
  }
  describe('as a simple swap', function() {
    beforeEach(async function() {
      it('should have a correct owner', async function() {
        expect(await this.simpleSwap.owner()).to.equal(owner)          
      })
    })
    describe('submitChequeBeneficiary', function() {
        //TODO: explicit check on allowing overdraft
      describe('when the sender is the beneficiary', function() {
        let sender = defaultCheck.beneficiary
        describe('when the first serial is higher than 0', function() {
          expect(defaultCheck.serial).bignumber.to.be.above(new BN(0), "Serial of defaultCheck not above 0")
          describe('when the owner is a signee', function() {
            expect(defaultCheck.signee).to.be.equal(owner, "Signee of defaultCheck is not owner")
            describe('when the signee signs the correct fields', function() {
              let unsignedCheque = Object.assign({}, defaultCheck)
              describe('when we send one cheque', function() {
                shouldSubmitChequeBeneficiary(unsignedCheque, sender)
              })
              describe('when we send more than one cheque', async function() {
                shouldSubmitChequeBeneficiary(unsignedCheque, sender)
                describe('when the serial number is increasing', function() {
                  let secondSerial = new BN(parseInt(unsignedCheque.serial) + 1)
                  let increasing_serial_unsignedCheque = Object.assign({}, defaultCheck, {serial: secondSerial})
                  shouldSubmitChequeBeneficiary(increasing_serial_unsignedCheque, sender)
                })
                describe('when the serial number stays the same', function() {
                  let secondSerial = new BN(parseInt(unsignedCheque.serial))
                  let same_serial_unsignedCheque = Object.assign({}, defaultCheck, {serial: secondSerial})
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
                  let decreasing_serial_unsignedCheque = Object.assign({}, defaultCheck, {serial: secondSerial})
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
              let wrong_beneficiary_unsignedCheque = Object.assign({}, defaultCheck, {beneficiary: wrongBeneficiary})
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
            const wrong_signee_unsignedCheque = Object.assign({}, defaultCheck, {signee: signee})
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
        describe('when the serial is 0', function() {
          let serial = new BN(0)
          const zero_serial_unsignedCheque = Object.assign({}, defaultCheck, {serial: serial})
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
    function shouldSubmitChequeBeneficiary(unsignedCheque, sender) {
      beforeEach(async function() {
        let lastCheque = await this.simpleSwap.cheques(unsignedCheque.beneficiary)
        expect(unsignedCheque.serial).bignumber.is.above(new BN(0), "serial is not a positive")
        expect(unsignedCheque.amount).bignumber.to.be.above(new BN(0), "amount is not positive")
        expect(unsignedCheque.beneficiary).to.equal(sender, "beneficiary is not the sender")
        expect(unsignedCheque.serial).bignumber.is.above(lastCheque.serial, "serial is not above the last submitted cheque")   
        this.signedCheque = await signCheque(this.simpleSwap, unsignedCheque)
        const { logs } = await this.simpleSwap.submitChequeBeneficiary(this.signedCheque.serial, this.signedCheque.amount, this.signedCheque.timeout, this.signedCheque.signature, {from: sender})
        this.logs = logs
      })
      describe('uses _submitChequeInternal', function() {
        _submitChequeInternal() 
      })
    }
  })
  function _submitChequeInternal() {    
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
}


module.exports = {
  shouldBehaveLikeSimpleSwap
};