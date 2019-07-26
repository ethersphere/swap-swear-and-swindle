const {
  BN,
  balance,
  time,
  shouldFail,
  constants,
  expectEvent
} = require("openzeppelin-test-helpers");

const { shouldBehaveLikeSimpleSwap } = require('./SimpleSwap.behavior')

const SimpleSwap = artifacts.require('SimpleSwap')

contract('SimpleSwap', function([issuer, alice, bob, agent]) {

  beforeEach(async function() {
    this.simpleSwap = await SimpleSwap.new(issuer, new BN(86400))
  })
  shouldBehaveLikeSimpleSwap([issuer, alice, bob, agent], new BN(86400))
})