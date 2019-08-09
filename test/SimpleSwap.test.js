const {
  BN
} = require("openzeppelin-test-helpers");

const { shouldBehaveLikeSimpleSwap } = require('./SimpleSwap.behavior')
const { shouldDeploy } = require('./SimpleSwap.should')

contract('SimpleSwap', function([issuer, alice, bob, agent]) {

  describe("when we don't deposit while deploying", function() {
    const sender = issuer
    const DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT = new BN(86400)
    const value = new BN(0)
    shouldDeploy(issuer, DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT, sender, value)
    shouldBehaveLikeSimpleSwap([issuer, alice, bob, agent], new BN(86400))
  })
  describe('when we deposit while deploying', function() {
    const sender = issuer
    const DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT = new BN(86400)
    const value = new BN(50)
    shouldDeploy(issuer, DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT, sender, value)
  })
 
})