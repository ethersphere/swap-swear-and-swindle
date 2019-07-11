const {
  BN,
  balance,
  time,
  shouldFail,
  constants,
  expectEvent
} = require("openzeppelin-test-helpers");

const { expect } = require('chai');

const { shouldBehaveLikeSimpleSwap } = require('./SimpleSwap.behavior')

const SimpleSwap = artifacts.require('SimpleSwap')

contract('SimpleSwap', function([issuer, alice, bob]) {

  beforeEach(async function() {
    this.simpleSwap = await SimpleSwap.new(issuer, 86400)
  })

  shouldBehaveLikeSimpleSwap([issuer, alice, bob])
})