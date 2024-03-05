const {
  balance,
  constants,
} = require("@openzeppelin/test-helpers");

const { expect } = require("chai");
const { waffle } = require("ethereum-waffle");
const { ethers, deployments, getNamedAccounts } = require("hardhat");
const { BigNumber } = ethers;

let issuer, alice, bob, carol, other;
let TestToken, SimpleSwapFactory, ERC20SimpleSwap;
const salt = "0x000000000000000000000000000000000000000000000000000000000000abcd";

before(async function () {
  // // Get signer accounts from Hardhat's ethers provider
  // [issuer, alice, bob, carol, other] = await ethers.getSigners();

  TestToken = await ethers.getContractFactory("TestToken");
  SimpleSwapFactory = await ethers.getContractFactory("SimpleSwapFactory");
  ERC20SimpleSwap = await ethers.getContractFactory('ERC20SimpleSwap');
});



describe('SimpleSwapFactory', function () {
  describe('when we deploy ERC20 SimpleSwap', async function () {
    describe("without depositing during deployment", async function () {

      DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT = 86400;
      value = BigNumber.from(0);
      let testToken, simpleSwapFactory, erc20SimpleSwap;

      beforeEach(async function () {
        [deployer, other] = await ethers.getSigners();
        testToken = await TestToken.deploy();
        await testToken.deployed();

        simpleSwapFactory = await SimpleSwapFactory.deploy(testToken.address);
        await simpleSwapFactory.deployed();

        const deployTx = await simpleSwapFactory.deploySimpleSwap(deployer.address, DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT, salt);
        const receipt = await deployTx.wait();
        const event = receipt.events.find(e => e.event === "SimpleSwapDeployed");
        const erc20SimpleSwapAddress = event.args.contractAddress;

        erc20SimpleSwap = ERC20SimpleSwap.attach(erc20SimpleSwapAddress);

        if (value.gt(0)) {
          await testToken.mint(deployer.address, value);
          await testToken.connect(deployer.address).transfer(erc20SimpleSwap.address, value);
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
    });

    describe("with deposit during deployment", function () {
      //   shouldDeployERC20SimpleSwap( 86400, ethers.utils.parseUnits("10", 18));
    });

    describe("with issuer being zero address", function () {
      it('should fail', async function () {
        testToken = await TestToken.deploy();
        await testToken.deployed();

        simpleSwapFactory = await SimpleSwapFactory.deploy(testToken.address);
        await simpleSwapFactory.deployed();



        // await expect(simpleSwapFactory.deploySimpleSwap(constants.ZERO_ADDRESS, 0, salt))
        //   .to.be.revertedWith('invalid issuer');
      });
    });
  });
});
