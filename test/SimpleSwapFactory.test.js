const {
  BN,
  balance,
  constants,
  expectEvent
} = require("openzeppelin-test-helpers");

const { expect } = require('chai');
const SimpleSwapFactory = artifacts.require('./SimpleSwapFactory')
const SimpleSwap = artifacts.require('./SimpleSwap')
const ERC20SimpleSwap = artifacts.require('./ERC20SimpleSwap')

contract('SimpleSwapFactory', function([issuer]) {

  function shouldDeploySimpleSwap(issuer, DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT, value) {            
    beforeEach(async function() {
      this.simpleSwapFactory = await SimpleSwapFactory.new(constants.ZERO_ADDRESS)
      let { logs } = await this.simpleSwapFactory.deploySimpleSwap(issuer, DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT, {
        value
      })

      expectEvent.inLogs(logs, 'SimpleSwapDeployed')
      this.simpleSwapAddress = logs[0].args.contractAddress
      this.simpleSwap = await SimpleSwap.at(this.simpleSwapAddress)
    })
    
    it('should deploy the correct bytecode', async function() {
      expect(await web3.eth.getCode(this.simpleSwapAddress)).to.be.equal(SimpleSwap.deployedBytecode)
    })

    it('should deploy with the right issuer', async function() {
      expect(await this.simpleSwap.issuer()).to.be.equal(issuer)
    })

    it('should deploy with the right DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT', async function() {
      expect(await this.simpleSwap.DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT()).to.be.bignumber.equal(DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT)
    })

    if(value.gtn(0)) {
      it('should forward the deposit to SimpleSwap', async function() {
        expect(await balance.current(this.simpleSwapAddress)).to.bignumber.equal(value)
      })
    }

    it('should record the deployed address', async function() {
      expect(await this.simpleSwapFactory.deployedContracts(this.simpleSwapAddress)).to.be.true
    })
  }

  function shouldDeployERC20SimpleSwap(issuer, ERC20Address, DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT, value) {
    beforeEach(async function() {
      this.simpleSwapFactory = await SimpleSwapFactory.new(ERC20Address)
      let { logs } = await this.simpleSwapFactory.deploySimpleSwap(issuer, DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT, {
        value
      })
      expectEvent.inLogs(logs, 'SimpleSwapDeployed')
      this.ERC20SimpleSwapAddress = logs[0].args.contractAddress
      this.ERC20SimpleSwap = await ERC20SimpleSwap.at(this.ERC20SimpleSwapAddress)
    })

    it('should deploy the correct bytecode', async function() {
      expect(await web3.eth.getCode(this.ERC20SimpleSwapAddress)).to.be.equal(ERC20SimpleSwap.deployedBytecode)
    })

    it('should deploy with the right issuer', async function() {
      expect(await this.ERC20SimpleSwap.issuer()).to.be.equal(issuer)
    })

    it('should deploy with the right DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT', async function() {
      expect(await this.ERC20SimpleSwap.DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT()).to.be.bignumber.equal(DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT)
    })

    if(value.gtn(0)) {
      it('should forward the deposit to SimpleSwap', async function() {
        expect(await balance.current(this.ERC20SimpleSwapAddress)).to.bignumber.equal(value)
      })
    }

    it('should record the deployed address', async function() {
      expect(await this.simpleSwapFactory.deployedContracts(this.ERC20SimpleSwapAddress)).to.be.true
    })

    it('should have set the ERC20 address correctly', async function() {
      expect(await this.ERC20SimpleSwap.token()).to.be.equal(ERC20Address)
    })

  }

  describe('when we deploy native SimpleSwap', function() {
    describe("when we don't deposit while deploying SimpleSwap", function() {
      shouldDeploySimpleSwap(issuer, new BN(86400), new BN(0))
    })
    describe("when we deposit while deploying SimpleSwap", function() {
      shouldDeploySimpleSwap(issuer, new BN(86400), new BN(10))
    })
  })
  
  describe('when we deploy ERC20 SimpleSwap', function() {
    const mockERC20Address = issuer;
    describe("when we don't deposit while deploying SimpleSwap", function() {
      shouldDeployERC20SimpleSwap(issuer, mockERC20Address, new BN(86400), new BN(0))
    })
  
    describe("when we deposit while deploying SimpleSwap", function() {
      shouldDeployERC20SimpleSwap(issuer, mockERC20Address, new BN(86400), new BN(10))
    })
  })
})