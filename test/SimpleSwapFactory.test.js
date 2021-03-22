const {
  BN,
  balance,
  constants,
  expectEvent
} = require("@openzeppelin/test-helpers");

const { expect } = require('chai');
const SimpleSwapFactory = artifacts.require('./SimpleSwapFactory')
const ERC20SimpleSwap = artifacts.require('./ERC20SimpleSwap')
const ERC20PresetMinterPauser = artifacts.require("ERC20PresetMinterPauser")

contract('SimpleSwapFactory', function([issuer, other]) {

  const salt = "0x000000000000000000000000000000000000000000000000000000000000abcd"

  function shouldDeployERC20SimpleSwap(issuer, DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT, value) {
    beforeEach(async function() {
      this.ERC20PresetMinterPauser = await ERC20PresetMinterPauser.new("TestToken", "TEST", {from: issuer})
      this.simpleSwapFactory = await SimpleSwapFactory.new(this.ERC20PresetMinterPauser.address)
      let { logs } = await this.simpleSwapFactory.deploySimpleSwap(issuer, DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT, salt)
      this.ERC20SimpleSwapAddress = logs[0].args.contractAddress
      this.ERC20SimpleSwap = await ERC20SimpleSwap.at(this.ERC20SimpleSwapAddress)
      if(value != 0) {
        await this.ERC20PresetMinterPauser.mint(issuer, value) // mint tokens
        await this.ERC20PresetMinterPauser.transfer(this.ERC20SimpleSwap.address, value, {from: issuer}); // deposit those tokens in chequebook
      }
    })

    it('should allow other addresses to deploy with same salt', async function() {
      await this.simpleSwapFactory.deploySimpleSwap(issuer, DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT, salt, { from: other })
    })

    it('should deploy with the right issuer', async function() {
      expect(await this.ERC20SimpleSwap.issuer()).to.be.equal(issuer)
    })

    it('should deploy with the right DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT', async function() {
      expect(await this.ERC20SimpleSwap.defaultHardDepositTimeout()).to.be.bignumber.equal(DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT)
    })

    if(value.gtn(0)) {
      it('should forward the deposit to SimpleSwap', async function() {
        expect(await this.ERC20SimpleSwap.balance()).to.bignumber.equal(value)
      })
    }

    it('should record the deployed address', async function() {
      expect(await this.simpleSwapFactory.deployedContracts(this.ERC20SimpleSwapAddress)).to.be.true
    })

    it('should have set the ERC20 address correctly', async function() {
      expect(await this.ERC20SimpleSwap.token()).to.be.equal(this.ERC20PresetMinterPauser.address)
    })
  }
  
  describe('when we deploy ERC20 SimpleSwap', function() {
    describe("when we don't deposit while deploying SimpleSwap", function() {
      shouldDeployERC20SimpleSwap(issuer, new BN(86400), new BN(0))
    })
  
    describe("when we deposit while deploying SimpleSwap", function() {
      shouldDeployERC20SimpleSwap(issuer, new BN(86400), new BN(10))
    })
  })
})