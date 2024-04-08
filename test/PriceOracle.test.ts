import { expect } from './util/chai';
import { ethers, deployments, getNamedAccounts, getUnnamedAccounts } from 'hardhat';
import { Contract } from 'ethers';

// Named accounts used by tests.
let deployer: string;
let others: string[];

// Before the tests, set named accounts and read deployments.
before(async function () {
  const namedAccounts = await getNamedAccounts();
  deployer = namedAccounts.deployer;
  others = await getUnnamedAccounts();
});

describe('PriceOracle', function () {
  let priceOracle: Contract;
  describe('when deploying contract', function () {
    beforeEach(async function () {
      await deployments.fixture();
    });

    it('should deploy PriceOracle', async function () {
      priceOracle = await ethers.getContract('PriceOracle');
      expect(priceOracle.address).to.be.properAddress;
    });

    it('should set the initial price', async function () {
      priceOracle = await ethers.getContract('PriceOracle');
      expect(await priceOracle.price()).to.equal(100);
    });

    it('should set the initial chequeValueDeduction', async function () {
      priceOracle = await ethers.getContract('PriceOracle');
      expect(await priceOracle.chequeValueDeduction()).to.equal(200);
    });
  });

  describe('with deployed contract', function () {
    beforeEach(async function () {
      await deployments.fixture();
      priceOracle = await ethers.getContract('PriceOracle', deployer);
    });

    describe('when updating the price', function () {
      describe('when called by owner', function () {
        it('should emit the PriceUpdate event', async function () {
          const newPrice = 100;
          await expect(priceOracle.updatePrice(newPrice)).to.emit(priceOracle, 'PriceUpdate').withArgs(newPrice);
        });

        it('should update the price', async function () {
          const newPrice = 100;
          await priceOracle.updatePrice(newPrice);
          expect(await priceOracle.price()).to.equal(newPrice);
        });
      });

      describe('when not called by owner', function () {
        it('should revert', async function () {
          const priceOracle = await ethers.getContract('PriceOracle', others[0]);
          const newPrice = 100;
          await expect(priceOracle.updatePrice(newPrice)).to.be.revertedWith('Ownable: caller is not the owner');
        });
      });
    });

    describe('when updating the chequeValueDeduction', function () {
      describe('when called by owner', function () {
        it('should emit the ChequeValueDeductionUpdate event', async function () {
          const newValue = 100;
          await expect(priceOracle.updateChequeValueDeduction(newValue))
            .to.emit(priceOracle, 'ChequeValueDeductionUpdate')
            .withArgs(newValue);
        });

        it('should update the chequeValueDeduction', async function () {
          const newValue = 100;
          await priceOracle.updateChequeValueDeduction(newValue);
          expect(await priceOracle.chequeValueDeduction()).to.equal(newValue);
        });
      });

      describe('when not called by owner', function () {
        it('should revert', async function () {
          const priceOracle = await ethers.getContract('PriceOracle', others[0]);
          await expect(priceOracle.updateChequeValueDeduction(100)).to.be.revertedWith(
            'Ownable: caller is not the owner'
          );
        });
      });
    });

    describe('when getting the price', function () {
      beforeEach(async function () {
        await deployments.fixture();
        priceOracle = await ethers.getContract('PriceOracle', deployer);
      });

      it('should get both values', async function () {
        const price = 100;
        const chequeValueDeduction = 200;
        await priceOracle.updatePrice(price);
        await priceOracle.updateChequeValueDeduction(chequeValueDeduction);

        const values = await priceOracle.getPrice();
        expect(values[0]).to.equal(price);
        expect(values[1]).to.equal(chequeValueDeduction);
      });
    });
  });
});
