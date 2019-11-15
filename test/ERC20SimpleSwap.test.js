const {
  BN
} = require("@openzeppelin/test-helpers");

const { shouldBehaveLikeERC20SimpleSwap } = require('./ERC20SimpleSwap.behavior')
const { shouldDeploy } = require('./ERC20SimpleSwap.should')

contract('ERC20SimpleSwap', function([issuer, alice, bob, agent]) {
  describe("when we don't deposit while deploying", async function() {
    const sender = issuer
    const DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT = new BN(86400)
    const value = new BN(0)
    shouldDeploy(issuer, DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT, sender, value)
    shouldBehaveLikeERC20SimpleSwap([issuer, alice, bob, agent], new BN(86400))
  })
  describe('when we deposit while deploying', function() {
    const sender = issuer
    const DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT = new BN(86400)
    const value = new BN(50)
    shouldDeploy(issuer, DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT, sender, value)
  })
})