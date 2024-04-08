import { BN } from '@openzeppelin/test-helpers';
import { shouldBehaveLikeERC20SimpleSwap } from './ERC20SimpleSwap.behavior';
import { shouldDeploy } from './ERC20SimpleSwap.should';

import { expect } from 'chai';
import { ethers, getNamedAccounts, getUnnamedAccounts, deployments } from 'hardhat';
import { BigNumber, Contract, ContractTransaction } from 'ethers';
import internal from 'stream';

describe('ERC20SimpleSwap', async function () {
  let issuer: string, alice: string, bob: string, agent: string;
  let defaultHardDepositTimeout: BN;
  let value: BN;

  beforeEach(async function () {
    const namedAccounts = await getNamedAccounts();
    issuer = namedAccounts.deployer;
    alice = namedAccounts.alice;
    bob = namedAccounts.bob;
    agent = namedAccounts.carol;
    defaultHardDepositTimeout = new BN(86400);
    await deployments.fixture();
  });

  describe('when we deposit while deploying', async function () {
    const defaultHardDepositTimeout = new BN(86400);
    const value = new BN(50);

    it('should check 1', async function () {
      const sender = issuer;
      shouldDeploy(issuer, defaultHardDepositTimeout.toString(), sender, value.toString());
    });
  });

  describe("when we don't deposit while deploying", async function () {
    const defaultHardDepositTimeout = new BN(86400);
    const value = new BN(0);

    it('should check 2', function () {
      const sender = issuer;
      shouldDeploy(issuer, defaultHardDepositTimeout.toString(), sender, value.toString());
      shouldBehaveLikeERC20SimpleSwap([issuer, alice, bob, agent], new BN(86400));
    });
  });
});
