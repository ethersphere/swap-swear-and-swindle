const {
  BN
} = require("@openzeppelin/test-helpers");

const { shouldBehaveLikeERC20SimpleSwap } = require('./ERC20SimpleSwap.behavior');
const { shouldDeploy } = require('./ERC20SimpleSwap.should');

describe('ERC20SimpleSwap', function() {
  let issuer, alice, bob, agent;

  before(async function() {
    // Get signer accounts from Hardhat's ethers provider
    const accounts = await ethers.getSigners();
    
    issuer = accounts[0];
    alice = accounts[1];
    bob = accounts[2];
    agent = accounts[3];
  });


  describe("when we don't deposit while deploying", function() {
    const sender = issuer;
    const defaultHardDepositTimeout = new BN(86400);
    const value = new BN(0);
    shouldDeploy(issuer, defaultHardDepositTimeout, sender, value);
    shouldBehaveLikeERC20SimpleSwap([issuer, alice, bob, agent], new BN(86400));
  });

  describe('when we deposit while deploying', function() {
    const sender = issuer;
    const defaultHardDepositTimeout = new BN(86400);
    const value = new BN(50);
    shouldDeploy(issuer, defaultHardDepositTimeout, sender, value);
  });
});
