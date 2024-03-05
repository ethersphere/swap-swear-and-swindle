const { ethers } = require('hardhat');
const { BigNumber } = ethers;

const { shouldBehaveLikeERC20SimpleSwap } = require('./ERC20SimpleSwap.behavior');
const { shouldDeploy } = require('./ERC20SimpleSwap.should');

describe('ERC20SimpleSwap', function() {
  let issuer, alice, bob, carol;

  before(async function() {
    // Get signer accounts from Hardhat's ethers provider
    [issuer, alice, bob, carol] = await ethers.getSigners();
  });

  describe("when we don't deposit while deploying", function() {
    const defaultHardDepositTimeout = BigNumber.from(86400);
    const value = BigNumber.from(0);
    shouldDeploy(issuer, defaultHardDepositTimeout, issuer, value);
    shouldBehaveLikeERC20SimpleSwap([issuer, alice, bob, carol], defaultHardDepositTimeout);
  });

  describe('when we deposit while deploying', function() {
    const defaultHardDepositTimeout = BigNumber.from(86400);
    const value = BigNumber.from(50);
    shouldDeploy(issuer, defaultHardDepositTimeout, issuer, value);
  });
});
