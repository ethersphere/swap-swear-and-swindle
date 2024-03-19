import { BN } from '@openzeppelin/test-helpers';
import { shouldBehaveLikeERC20SimpleSwap } from './ERC20SimpleSwap.behavior';
import { shouldDeploy } from './ERC20SimpleSwap.should';

import { expect } from 'chai';
import { ethers, getNamedAccounts, getUnnamedAccounts } from 'hardhat';
import { BigNumber, Contract, ContractTransaction } from 'ethers';

describe('ERC20SimpleSwap', async function () {
  let issuer: string, alice: string, bob: string, agent: string;
  let defaultHardDepositTimeout: BN;
  let value: BN;

  beforeEach(async function () {
    const namedAccounts = await getNamedAccounts();
    issuer = namedAccounts.deployer;
    alice = namedAccounts.user_1;
    bob = namedAccounts.user_2;
    agent = namedAccounts.user_3;
    defaultHardDepositTimeout = new BN(86400);
  });

  describe("when we don't deposit while deploying", async function () {
    const defaultHardDepositTimeout = new BN(86400);
    const value = new BN(0);

    it('should check', async function () {
      const sender = issuer;
      console.log(alice);
      // shouldDeploy(issuer, defaultHardDepositTimeout, sender, value);
      shouldBehaveLikeERC20SimpleSwap([issuer, alice, bob, agent], new BN(86400));
    });
  });
  describe('when we deposit while deploying', async function () {
    const sender = issuer;
    const defaultHardDepositTimeout = new BN(86400);
    const value = new BN(50);

    it('should check', async function () {
      //shouldDeploy(issuer, defaultHardDepositTimeout, sender, value);
    });
  });
});
