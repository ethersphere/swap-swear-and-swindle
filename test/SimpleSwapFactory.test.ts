import { expect } from 'chai';
import { ethers, getNamedAccounts, getUnnamedAccounts } from 'hardhat';
import { BigNumber, Contract, ContractTransaction } from 'ethers';
import { constants } from '@openzeppelin/test-helpers';

describe('SimpleSwapFactory', function () {
  const salt = '0x000000000000000000000000000000000000000000000000000000000000abcd';
  let simpleSwapFactory: Contract, testToken: Contract;

  let deployer: string;
  let admin: string;
  let others: string[];

  beforeEach(async function () {
    const namedAccounts = await getNamedAccounts();
    deployer = namedAccounts.deployer;
    admin = namedAccounts.admin;
    others = await getUnnamedAccounts();
  });

  function shouldDeployERC20SimpleSwap(DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT: number, value: BigNumber) {
    beforeEach(async function () {
      const namedAccounts = await getNamedAccounts();
      deployer = namedAccounts.deployer;
      admin = namedAccounts.admin;

      const TestToken = await ethers.getContractFactory('TestToken');
      testToken = await TestToken.deploy();
      await testToken.deployed();

      const SimpleSwapFactory = await ethers.getContractFactory('SimpleSwapFactory');
      simpleSwapFactory = await SimpleSwapFactory.deploy(testToken.address);
      await simpleSwapFactory.deployed();

      const tx = await simpleSwapFactory.deploySimpleSwap(deployer, DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT, salt);
      const receipt = await tx.wait();
      const event = receipt.events.find((e) => e.event === 'SimpleSwapDeployed');
      this.ERC20SimpleSwapAddress = event.args.contractAddress;

      const ERC20SimpleSwap = await ethers.getContractFactory('ERC20SimpleSwap');
      this.ERC20SimpleSwap = await ERC20SimpleSwap.attach(this.ERC20SimpleSwapAddress);

      if (value.gt(0)) {
        await testToken.mint(deployer, value);
        await testToken.transfer(this.ERC20SimpleSwap.address, value);
      }
    });

    it('should allow other addresses to deploy with same salt', async function () {
      // We must get other as signer object, switch context and then try to use smart contract
      const [, account1] = await ethers.getSigners();
      await simpleSwapFactory
        .connect(account1)
        .deploySimpleSwap(account1.address, DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT, salt);
    });

    it('should deploy with the right issuer', async function () {
      expect(await this.ERC20SimpleSwap.issuer()).to.equal(deployer);
    });

    it('should deploy with the right DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT', async function () {
      expect(await this.ERC20SimpleSwap.defaultHardDepositTimeout()).to.equal(DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT);
    });

    if (value.gt(0)) {
      it('should forward the deposit to SimpleSwap', async function () {
        expect(await this.ERC20SimpleSwap.balance()).to.equal(value);
      });
    }

    it('should record the deployed address', async function () {
      expect(await simpleSwapFactory.deployedContracts(this.ERC20SimpleSwapAddress)).to.be.true;
    });

    it('should have set the ERC20 address correctly', async function () {
      expect(await this.ERC20SimpleSwap.token()).to.equal(testToken.address);
    });
  }

  describe('when we deploy ERC20 SimpleSwap', function () {
    describe("when we don't deposit while deploying SimpleSwap", function () {
      shouldDeployERC20SimpleSwap(86400, ethers.constants.Zero);
    });

    describe('when we deposit while deploying SimpleSwap', function () {
      shouldDeployERC20SimpleSwap(86400, ethers.utils.parseUnits('10', 18));
    });

    describe('when we deposit while issuer 0', function () {
      it('should fail', async function () {
        const TestToken = await ethers.getContractFactory('TestToken');
        testToken = await TestToken.deploy();
        await testToken.deployed();

        const SimpleSwapFactory = await ethers.getContractFactory('SimpleSwapFactory');
        simpleSwapFactory = await SimpleSwapFactory.deploy(testToken.address);
        await simpleSwapFactory.deployed();

        console.log(constants.ZERO_ADDRESS);
        await expect(simpleSwapFactory.deploySimpleSwap(constants.ZERO_ADDRESS, 86400, salt)).to.be.revertedWith(
          'invalid issuer'
        );
      });
    });
  });
});
