const {
  balance,
  constants,
  expectEvent,
  expectRevert
} = require("@openzeppelin/test-helpers");
const { expect } = require('chai');
const { ethers, deployments, getNamedAccounts } = require("hardhat");
const { BigNumber } = ethers;

let issuer, other, TestToken, SimpleSwapFactory, ERC20SimpleSwap;
    let testToken, simpleSwapFactory, erc20SimpleSwap;

before(async function () {
  const namedAccounts = await getNamedAccounts();
  issuer = namedAccounts.deployer;
  console.log(issuer);
  TestToken = await ethers.getContractFactory("TestToken");
  SimpleSwapFactory = await ethers.getContractFactory("SimpleSwapFactory");
  ERC20SimpleSwap = await ethers.getContractFactory('ERC20SimpleSwap');
});

describe('SimpleSwapFactory', function () {

  const salt = "0x000000000000000000000000000000000000000000000000000000000000abcd";

  async function shouldDeployERC20SimpleSwap(deployer, DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT, value) {

    beforeEach(async function () {
      testToken = await TestToken.deploy();
      await testToken.deployed();

      simpleSwapFactory = await SimpleSwapFactory.deploy(testToken.address);
      await simpleSwapFactory.deployed();

      const deployTx = await simpleSwapFactory.deploySimpleSwap(deployer, DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT, salt);
      const receipt = await deployTx.wait();
      const event = receipt.events.find(e => e.event === "SimpleSwapDeployed");
      const erc20SimpleSwapAddress = event.args.contractAddress;

      erc20SimpleSwap = ERC20SimpleSwap.attach(erc20SimpleSwapAddress);

      if (value.gt(0)) {
        await testToken.mint(deployer.address, value);
        await testToken.connect(deployer).transfer(erc20SimpleSwap.address, value);
      }
    });

    it('should allow other addresses to deploy with the same salt', async function () {
      await expect(simpleSwapFactory.connect(other).deploySimpleSwap(other.address, DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT, salt))
        .to.emit(simpleSwapFactory, 'SimpleSwapDeployed');
    });

    it('should deploy with the right issuer', async function () {
      expect(await erc20SimpleSwap.issuer()).to.equal(deployer.address);
    });

    it('should deploy with the right DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT', async function () {
      expect(await erc20SimpleSwap.defaultHardDepositTimeout()).to.equal(DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT);
    });

    if (value.gt(0)) {
      it('should forward the deposit to SimpleSwap', async function () {
        expect(await balance.current(erc20SimpleSwap.address)).to.equal(value);
      });
    }

    it('should record the deployed address', async function () {
      expect(await simpleSwapFactory.deployedContracts(erc20SimpleSwapAddress)).to.be.true;
    });

    it('should have set the ERC20 address correctly', async function () {
      expect(await erc20SimpleSwap.token()).to.equal(testToken.address);
    });
  }

  describe('when we deploy ERC20 SimpleSwap', function () {
    describe("without depositing during deployment", function () {
      shouldDeployERC20SimpleSwap(issuer, 86400, BigNumber.from(0));
    });

    describe("with deposit during deployment", function () {
  console.log(issuer);
      shouldDeployERC20SimpleSwap(issuer, 86400, ethers.utils.parseUnits("10", 18));
    });

    describe("with issuer being zero address", function () {
      it('should fail', async function () {
        testToken = await TestToken.deploy();
        await testToken.deployed();

        simpleSwapFactory = await SimpleSwapFactory.deploy(testToken.address);
        await simpleSwapFactory.deployed();

        await expect(simpleSwapFactory.deploySimpleSwap(constants.ZERO_ADDRESS, 0, salt))
          .to.be.revertedWith('invalid issuer');
      });
    });
  });
});
