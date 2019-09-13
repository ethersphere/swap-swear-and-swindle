const {
  BN,
  balance,
  expectEvent
} = require("openzeppelin-test-helpers");

const { expect } = require('chai');
const SimpleSwapFactory = artifacts.require('./SimpleSwapFactory')
const SimpleSwap = artifacts.require('./SimpleSwap')

contract('SimpleSwapFactory', function([issuer]) {

  function shouldDeploySimpleSwaps(issuer, DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT, value) {            
    beforeEach(async function() {
      this.simpleSwapFactory = await SimpleSwapFactory.new()
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

  describe("when we don't deposit while deploying SimpleSwap", function() {    
    shouldDeploySimpleSwaps(issuer, new BN(86400), new BN(0))
  })

  describe("when we deposit while deploying SimpleSwap", function() {    
    shouldDeploySimpleSwaps(issuer, new BN(86400), new BN(10))
  })
})