import { BN, balance, expectEvent, expectRevert } from '@openzeppelin/test-helpers';
import { time } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from 'chai';
import { ethers, deployments, getNamedAccounts, getUnnamedAccounts } from 'hardhat';
import { BigNumber, Contract, ContractTransaction } from 'ethers';

import { signCheque, signCashOut, signCustomDecreaseTimeout } from './swutils';

function shouldDeploy(issuer, defaultHardDepositTimeout, from, value) {
  const salt = '0x000000000000000000000000000000000000000000000000000000000000abcd';

  beforeEach(async function () {
    this.TestToken = await ethers.getContract('TestToken');
    await this.TestToken.mint(issuer, 1000000000, { from: issuer });

    const SimpleSwapFactory = await ethers.getContractFactory('SimpleSwapFactory');
    const simpleSwapFactory = await SimpleSwapFactory.deploy(this.TestToken.address);
    await simpleSwapFactory.deployed();
    
    const deployTx = await simpleSwapFactory.deploySimpleSwap(issuer, defaultHardDepositTimeout, salt);
    // Wait for the transaction to be mined to access the events and their arguments
    const receipt = await deployTx.wait();
    const ERC20SimpleSwapAddress = receipt.events?.filter((x) => x.event === 'SimpleSwapDeployed')[0].args.contractAddress;

    // // Assuming 'YourEventName' is the event name that logs the contract address, replace it accordingly

    const ERC20SimpleSwap = await ethers.getContractFactory('ERC20SimpleSwap');
    this.ERC20SimpleSwap = await ERC20SimpleSwap.attach(ERC20SimpleSwapAddress);
  
    // If value is not equal to 0, transfer tokens from issuer to the ERC20SimpleSwap contract
    
    if (value != 0) {
      await this.TestToken.transfer(this.ERC20SimpleSwap.address, value);
      //await this.TestToken.connect(issuer).transfer(this.ERC20SimpleSwap.address, value);
    }
  
    // Assuming 'issuer' is a Signer or you have access to the Signer to perform the transfer

    // Read postconditions from the deployed ERC20SimpleSwap contract
    const postconditions = {
      issuer: await this.ERC20SimpleSwap.issuer(),
      defaultHardDepositTimeout: await this.ERC20SimpleSwap.defaultHardDepositTimeout(),
    };

  });

  it('should not allow a second init', async function () {
    await expectRevert(this.ERC20SimpleSwap.init(issuer, this.TestToken.address, 0), 'already initialized');
  });

  it('should set the issuer', function () {
    expect(this.postconditions.issuer).to.be.equal(issuer);
  });
  it('should set the defaultHardDepositTimeout', function () {
    expect(this.postconditions.defaultHardDepositTimeout).to.equal(defaultHardDepositTimeout);
  });
}
function shouldReturnDefaultHardDepositTimeout(expected) {
  it('should return the expected defaultHardDepositTimeout', async function () {
    expect(await this.ERC20SimpleSwap.defaultHardDepositTimeout()).to.equal(expected.toString());
  });
}

function shouldReturnPaidOut(beneficiary, expectedAmount) {
  beforeEach(async function () {
    this.paidOut = await this.ERC20SimpleSwap.paidOut(beneficiary);
  });
  it('should return the expected amount', function () {
    expect(expectedAmount.toString()).to.equal(this.paidOut);
  });
}

function shouldReturnTotalPaidOut(expectedAmount) {
  beforeEach(async function () {
    this.totalPaidOut = await this.ERC20SimpleSwap.totalPaidOut();
  });
  it('should return the expected amount', function () {
    expect(expectedAmount.toString()).to.equal(this.totalPaidOut);
  });
}

function shouldReturnHardDeposits(
  beneficiary,
  expectedAmount,
  expectedDecreaseAmount,
  expectedDecreaseTimeout,
  expectedCanBeDecreasedAt
) {

  beforeEach(async function () {
    // If we expect this not to be the default value, we have to set the value here, as it depends on the most current time
    if (!expectedCanBeDecreasedAt.eq(new BN(0))) {
      this.expectedCanBeDecreasedAt = (BigNumber.from(await time.latest())).add(await this.ERC20SimpleSwap.defaultHardDepositTimeout());
    } else {
      this.expectedCanBeDecreasedAt = expectedCanBeDecreasedAt;
    }
    
    this.exptectedCanBeDecreasedAt = (BigNumber.from(await time.latest())).add(await this.ERC20SimpleSwap.defaultHardDepositTimeout());  
    this.hardDeposits = await this.ERC20SimpleSwap.hardDeposits(beneficiary);;
  });
  it('should return the expected amount', function () {
    expect(expectedAmount.toString()).to.equal(this.hardDeposits.amount.toString());
  });
  it('should return the expected decreaseAmount', function () {
    expect(expectedDecreaseAmount.toString()).to.equal(this.hardDeposits.decreaseAmount.toString());
  });
  it('should return the expected timeout', function () {
    expect(expectedDecreaseTimeout.toString()).to.equal(this.hardDeposits.timeout.toString());
  });
  it('should return the exptected canBeDecreasedAt', function () {    
    expect(this.expectedCanBeDecreasedAt.toNumber()).to.be.closeTo(this.hardDeposits.canBeDecreasedAt.toNumber(), 5);
  });
}

function shouldReturnTotalHardDeposit(expectedTotalHardDeposit) {
  beforeEach(async function () {
    this.totalHardDeposit = await this.ERC20SimpleSwap.totalHardDeposit();
  });

  it('should return the expectedTotalHardDeposit', function () {
    expect(BigNumber.from(expectedTotalHardDeposit.toString())).to.equal(this.totalHardDeposit);
  });
}

function shouldReturnIssuer(expectedIssuer) {
  it('should return the expected issuer', async function () {
    expect(await this.ERC20SimpleSwap.issuer()).to.be.equal(expectedIssuer);
  });
}

function shouldReturnLiquidBalance(expectedLiquidBalance) {
  it('should return the expected liquidBalance', async function () {
    expect(await this.ERC20SimpleSwap.liquidBalance()).to.equal(BigNumber.from(expectedLiquidBalance.toString()));
  });
}

function shouldReturnLiquidBalanceFor(beneficiary, expectedLiquidBalanceFor) {
  it('should return the expected liquidBalance', async function () {
    expect(await this.ERC20SimpleSwap.liquidBalanceFor(beneficiary)).to.equal(BigNumber.from(expectedLiquidBalanceFor.toString()));
  });
}

function cashChequeInternal(beneficiary, recipient, cumulativePayout, callerPayout, from) {
  beforeEach(async function () {
    let requestPayout = BigNumber.from(cumulativePayout.toString()).sub(this.preconditions.paidOut);
    //if the requested payout is less than the liquidBalance available for beneficiary
    if (requestPayout.lt(this.preconditions.liquidBalanceFor)) {
      // full amount requested can be paid out
      this.totalPayout = requestPayout;
    } else {
      // partial amount requested can be paid out (the liquid balance available to the node)
      this.totalPayout = this.preconditions.liquidBalanceFor;
    }
    
    this.totalPaidOut = this.preconditions.totalPaidOut + this.totalPayout;
  });

  it('should update the totalHardDeposit and hardDepositFor ', function () {
    let expectedDecreaseHardDeposit;
    // if the hardDeposits can cover the totalPayout
    if (this.totalPayout.lt(this.preconditions.hardDepositFor.amount)) {
      // hardDeposit decreases by totalPayout
      expectedDecreaseHardDeposit = this.totalPayout;
    } else {
      // hardDeposit decreases by the full amount (and rest is from global liquid balance)
      expectedDecreaseHardDeposit = this.preconditions.hardDepositFor.amount;
    }
    // totalHarddeposit
    expect(this.postconditions.totalHardDeposit).to.equal(
      this.preconditions.totalHardDeposit.sub(expectedDecreaseHardDeposit)
    );
    // hardDepositFor
    expect(this.postconditions.hardDepositFor.amount).to.equal(
      this.preconditions.hardDepositFor.amount.sub(expectedDecreaseHardDeposit)
    );
  });

  it('should update paidOut', async function () {
    expect(this.postconditions.paidOut).to.equal(this.preconditions.paidOut.add(this.totalPayout));
  });

  it('should update totalPaidOut', async function () {
    expect(this.postconditions.totalPaidOut).to.equal(this.preconditions.paidOut.add(this.totalPayout));
  });

  it('should transfer the correct amount to the recipient', async function () {
    expect(this.postconditions.recipientBalance).to.equal(
      this.preconditions.recipientBalance.add(this.totalPayout.sub(BigNumber.from(callerPayout.toString())))
    );
  });
  it('should transfer the correct amount to the caller', async function () {
    let expectedAmountCaller;
    if (recipient == from) {
      expectedAmountCaller = this.totalPayout;
    } else {
      expectedAmountCaller = callerPayout;
    }
    expect(this.postconditions.callerBalance).to.equal(this.preconditions.callerBalance.add(BigNumber.from(expectedAmountCaller.toString())));
  });

  it('should emit a ChequeCashed event', function () {
     expectEvent.inLogs(this.receipt.events, 'ChequeCashed', {
      beneficiary,
      recipient: recipient,
      caller: from,
      totalPayout: this.totalPayout,
      cumulativePayout: BigNumber.from(cumulativePayout.toString()),
      callerPayout: BigNumber.from(callerPayout.toString())
    });
  });
  it('should only emit a chequeBounced event when insufficient funds', function () {
    if (this.totalPayout.lt(BigNumber.from(cumulativePayout.toString()).sub(this.preconditions.paidOut))) {
      expectEvent.inLogs(this.receipt.events, 'ChequeBounced', {});
    } else {
      const events = this.receipt.events.filter((e) => e.event === 'ChequeBounced');
      expect(events.length > 0).to.equal(false, `There is a ChequeBounced event`);
    }
  });

  it('should only set the bounced field when insufficient funds', function () {
    if (this.totalPayout.lt(BigNumber.from(cumulativePayout.toString()).sub(this.preconditions.paidOut))) {
      expect(this.postconditions.bounced).to.be.true;
    } else {
      expect(this.postconditions.bounced).to.be.false;
    }
  });
}

function shouldCashChequeBeneficiary(recipient, cumulativePayout, signee, from) {
  beforeEach(async function () {
    this.preconditions = {
      callerBalance: await this.TestToken.balanceOf(from),
      recipientBalance: await this.TestToken.balanceOf(recipient),
      totalHardDeposit: await this.ERC20SimpleSwap.totalHardDeposit(),
      hardDepositFor: await this.ERC20SimpleSwap.hardDeposits(from),
      liquidBalance: await this.ERC20SimpleSwap.liquidBalance(),
      liquidBalanceFor: await this.ERC20SimpleSwap.liquidBalanceFor(from),
      chequebookBalance: await this.ERC20SimpleSwap.balance(),
      paidOut: await this.ERC20SimpleSwap.paidOut(from),
      totalPaidOut: await this.ERC20SimpleSwap.totalPaidOut(),
    };

    const issuerSig = await signCheque(this.ERC20SimpleSwap, from, cumulativePayout, signee);

    const fromAddressSigner = await ethers.getSigner(from)
    const tx = await this.ERC20SimpleSwap.connect(fromAddressSigner).cashChequeBeneficiary(recipient, cumulativePayout.toString(), issuerSig);
    
    const receipt = await tx.wait();
    this.logs = receipt.logs;
    this.receipt = receipt;

    this.postconditions = {
      callerBalance: await this.TestToken.balanceOf(from),
      recipientBalance: await this.TestToken.balanceOf(recipient),
      totalHardDeposit: await this.ERC20SimpleSwap.totalHardDeposit(),
      hardDepositFor: await this.ERC20SimpleSwap.hardDeposits(from),
      liquidBalance: await this.ERC20SimpleSwap.liquidBalance(),
      liquidBalanceFor: await this.ERC20SimpleSwap.liquidBalanceFor(from),
      chequebookBalance: await this.ERC20SimpleSwap.balance(),
      paidOut: await this.ERC20SimpleSwap.paidOut(from),
      totalPaidOut: await this.ERC20SimpleSwap.totalPaidOut(),
      bounced: await this.ERC20SimpleSwap.bounced(),
    };

  });
  
  cashChequeInternal(from, recipient, cumulativePayout, new BN(0), from);
}
function shouldNotCashChequeBeneficiary(
  recipient,
  toSubmitCumulativePayout,
  toSignCumulativePayout,
  signee,
  from,
  value,
  revertMessage
) {
  beforeEach(async function () {
    this.issuerSig = await signCheque(this.ERC20SimpleSwap, from, toSignCumulativePayout, signee);
  });
  it('reverts', async function () {
    const fromAddressSigner = await ethers.getSigner(from);
    try {
      await expect(this.ERC20SimpleSwap.connect(fromAddressSigner).cashChequeBeneficiary(recipient, toSubmitCumulativePayout.toString(), this.issuerSig, { value: value.toString() }))
      .to.be.revertedWith(revertMessage);
    } catch(error) {
      await expect(error.toString()).to.include('non-payable method cannot override value')
    }
  });
}
function shouldCashCheque(
  beneficiary,
  recipient,
  cumulativePayout,
  callerPayout,
  from,
  beneficiarySignee,
  issuerSignee
) {
  beforeEach(async function () {
    const beneficiarySig = await signCashOut(
      this.ERC20SimpleSwap,
      from,
      cumulativePayout,
      recipient,
      callerPayout,
      beneficiarySignee
    );
    const issuerSig = await signCheque(this.ERC20SimpleSwap, beneficiary, cumulativePayout, issuerSignee);
    this.preconditions = {
      callerBalance: await this.TestToken.balanceOf(from),
      recipientBalance: await this.TestToken.balanceOf(recipient),
      totalHardDeposit: await this.ERC20SimpleSwap.totalHardDeposit(),
      hardDepositFor: await this.ERC20SimpleSwap.hardDeposits(beneficiary),
      liquidBalance: await this.ERC20SimpleSwap.liquidBalance(),
      liquidBalanceFor: await this.ERC20SimpleSwap.liquidBalanceFor(beneficiary),
      chequebookBalance: await this.ERC20SimpleSwap.balance(),
      paidOut: await this.ERC20SimpleSwap.paidOut(beneficiary),
      totalPaidOut: await this.ERC20SimpleSwap.totalPaidOut(),
    };

    const fromAddressSigner = await ethers.getSigner(from);
    const tax = await this.ERC20SimpleSwap.connect(fromAddressSigner).cashCheque(
      beneficiary,
      recipient,
      cumulativePayout.toString(),
      beneficiarySig,
      callerPayout.toString(),
      issuerSig
    );
    const receipt = await tax.wait();
    this.logs = receipt.logs;
    this.receipt = receipt;

    this.postconditions = {
      callerBalance: await this.TestToken.balanceOf(from),
      recipientBalance: await this.TestToken.balanceOf(recipient),
      totalHardDeposit: await this.ERC20SimpleSwap.totalHardDeposit(),
      hardDepositFor: await this.ERC20SimpleSwap.hardDeposits(beneficiary),
      liquidBalance: await this.ERC20SimpleSwap.liquidBalance(),
      liquidBalanceFor: await this.ERC20SimpleSwap.liquidBalanceFor(beneficiary),
      chequebookBalance: await this.ERC20SimpleSwap.balance(),
      paidOut: await this.ERC20SimpleSwap.paidOut(beneficiary),
      totalPaidOut: await this.ERC20SimpleSwap.totalPaidOut(),
      bounced: await this.ERC20SimpleSwap.bounced(),
    };
  });

  cashChequeInternal(beneficiary, recipient, cumulativePayout, callerPayout, from);
}
function shouldNotCashCheque(
  beneficiaryToSign,
  issuerToSign,
  toSubmitFields,
  value,
  from,
  beneficiarySignee,
  issuerSignee,
  revertMessage
) { 
  beforeEach(async function () {
    this.beneficiarySig = await signCashOut(
      this.ERC20SimpleSwap,
      from,
      beneficiaryToSign.cumulativePayout,
      beneficiaryToSign.recipient,
      beneficiaryToSign.callerPayout,
      beneficiarySignee
    );
    this.issuerSig = await signCheque(
      this.ERC20SimpleSwap,
      issuerToSign.beneficiary,
      issuerToSign.cumulativePayout,
      issuerSignee
    );
  });
  it('reverts', async function () {
    const fromAddressSigner = await ethers.getSigner(from);
    try {
      await expect(
        this.ERC20SimpleSwap.connect(fromAddressSigner).cashCheque(
          toSubmitFields.beneficiary,
          toSubmitFields.recipient,
          toSubmitFields.cumulativePayout.toString(),
          this.beneficiarySig,
          toSubmitFields.callerPayout.toString(),
          this.issuerSig,
          { value: value.toString() }
        )
      ).to.be.revertedWith(revertMessage);
    } catch(error) {
      await expect(error.toString()).to.include('non-payable method cannot override value')
    }
  });
}
function shouldPrepareDecreaseHardDeposit(beneficiary, decreaseAmount, from) {
  beforeEach(async function () {
    this.preconditions = {
      hardDepositFor: await this.ERC20SimpleSwap.hardDeposits(beneficiary),
    };

    const tx = await this.ERC20SimpleSwap.prepareDecreaseHardDeposit(beneficiary, decreaseAmount.toString(), { from: from });
    const receipt = await tx.wait();
    this.logs = receipt.logs;
    this.receipt = receipt;

    this.postconditions = {
      hardDepositFor: await this.ERC20SimpleSwap.hardDeposits(beneficiary),
    };
  });

  it('should update the canBeDecreasedAt', async function () {
    let expectedCanBeDecreasedAt;
    let personalDecreaseTimeout = (await this.ERC20SimpleSwap.hardDeposits(beneficiary)).timeout;
    // if personalDecreaseTimeout is zero
    if (personalDecreaseTimeout.eq(BigNumber.from(0))) {
      // use the contract's default
      expectedCanBeDecreasedAt = await this.ERC20SimpleSwap.defaultHardDepositTimeout();
    } else {
      // use the value that was set
      expectedCanBeDecreasedAt = personalDecreaseTimeout;
    }

    expect(this.postconditions.hardDepositFor.canBeDecreasedAt.toNumber()).to.be.closeTo(
      (BigNumber.from(await time.latest())).add(expectedCanBeDecreasedAt).toNumber(),
      5
    );

  });

  it('should update the decreaseAmount', function () {
    expect(this.postconditions.hardDepositFor.decreaseAmount).to.equal(BigNumber.from(decreaseAmount.toString()));
  });

  it('should emit a HardDepositDecreasePrepared event', function () {
    expectEvent.inLogs(this.receipt.events, 'HardDepositDecreasePrepared', {
      beneficiary,
      decreaseAmount: decreaseAmount.toString(),
    });
  });
}
function shouldNotPrepareDecreaseHardDeposit(beneficiary, decreaseAmount, from, value, revertMessage) {  
  it('reverts', async function () {
    const fromAddressSigner = await ethers.getSigner(from);
    try {
      await expect(this.ERC20SimpleSwap.connect(fromAddressSigner).prepareDecreaseHardDeposit(beneficiary, decreaseAmount.toString(), { from: from, value: value.toString() }))
      .to.be.revertedWith(revertMessage);
    } catch(error) {
      await expect(error.toString()).to.include('non-payable method cannot override value')
    }
  });
}
function shouldDecreaseHardDeposit(beneficiary, from) {
  beforeEach(async function () {
    this.preconditions = {
      hardDeposit: await this.ERC20SimpleSwap.totalHardDeposit(),
      hardDepositFor: await this.ERC20SimpleSwap.hardDeposits(beneficiary),
    };

    const tx = await this.ERC20SimpleSwap.decreaseHardDeposit(beneficiary, { from: from });
    const receipt = await tx.wait();
    this.logs = receipt.logs;
    this.receipt = receipt;

    this.postconditions = {
      hardDeposit: await this.ERC20SimpleSwap.totalHardDeposit(),
      hardDepositFor: await this.ERC20SimpleSwap.hardDeposits(beneficiary),
    };
  });

  it('decreases the hardDeposit amount for the beneficiary', function () {
    expect(this.postconditions.hardDepositFor.amount).to.equal(
      this.preconditions.hardDepositFor.amount.sub(this.preconditions.hardDepositFor.decreaseAmount)
    );
  });

  it('decreases the total hardDeposits', function () {
    expect(this.postconditions.hardDeposit).to.equal(
      this.preconditions.hardDeposit.sub(this.preconditions.hardDepositFor.decreaseAmount)
    );
  });

  it('resets the canBeDecreased at', function () {
    expect(this.postconditions.hardDepositFor.canBeDecreasedAt.toString()).to.equal((new BN(0)).toString());
  });

  it('emits a hardDepositAmountChanged event', function () {
    expectEvent.inLogs(this.receipt.events, 'HardDepositAmountChanged', {
      beneficiary,
      amount: this.postconditions.hardDepositFor.amount,
    });
  });
}
function shouldNotDecreaseHardDeposit(beneficiary, from, value, revertMessage) {  
  it('reverts', async function () {
    const fromAddressSigner = await ethers.getSigner(from);
    try {
      await expect(this.ERC20SimpleSwap.decreaseHardDeposit(beneficiary, { value: value.toString() }))
      .to.be.revertedWith(revertMessage);
    } catch(error) {
      await expect(error.toString()).to.include('non-payable method cannot override value')
    }
  });
}
function shouldIncreaseHardDeposit(beneficiary, amount, from) {
  beforeEach(async function () {
    this.preconditions = {
      balance: await this.ERC20SimpleSwap.balance(),
      liquidBalance: await this.ERC20SimpleSwap.liquidBalance(),
      liquidBalanceFor: await this.ERC20SimpleSwap.liquidBalanceFor(beneficiary),
      totalHardDeposit: await this.ERC20SimpleSwap.totalHardDeposit(),
      hardDepositFor: await this.ERC20SimpleSwap.hardDeposits(beneficiary),
    };

    const tx = await this.ERC20SimpleSwap.increaseHardDeposit(beneficiary, amount.toString(), { from: from });
    const receipt = await tx.wait();
    this.logs = receipt.logs;
    this.receipt = receipt;

    this.postconditions = {
      balance: await this.ERC20SimpleSwap.balance(),
      liquidBalance: await this.ERC20SimpleSwap.liquidBalance(),
      liquidBalanceFor: await this.ERC20SimpleSwap.liquidBalanceFor(beneficiary),
      totalHardDeposit: await this.ERC20SimpleSwap.totalHardDeposit(),
      hardDepositFor: await this.ERC20SimpleSwap.hardDeposits(beneficiary),
    };

  });

  it('should decrease the liquidBalance', function () {
    expect(this.postconditions.liquidBalance).to.equal(this.preconditions.liquidBalance.sub(BigNumber.from(amount.toString())));
  });

  it('should not affect the liquidBalanceFor', function () {
    expect(this.postconditions.liquidBalanceFor).to.equal(this.preconditions.liquidBalanceFor);
  });

  it('should not affect the balance', function () {
    expect(this.postconditions.balance).to.equal(this.preconditions.balance);
  });

  it('should increase the totalHardDeposit', function () {
    expect(this.postconditions.totalHardDeposit).to.equal(this.preconditions.totalHardDeposit.add(BigNumber.from(amount.toString())));
  });

  it('should increase the hardDepositFor', function () {
    expect(this.postconditions.hardDepositFor.amount).to.equal(this.preconditions.hardDepositFor.amount.add(BigNumber.from(amount.toString())));
  });

  it('should not influence the timeout', function () {
    expect(this.postconditions.hardDepositFor.timeout).to.equal(this.preconditions.hardDepositFor.timeout);
  });

  it('should set canBeDecreasedAt to zero', function () {
    expect(this.postconditions.hardDepositFor.canBeDecreasedAt).to.equal(BigNumber.from(0));
  });

  it('emits a hardDepositAmountChanged event', function () {
    expectEvent.inLogs(this.receipt.events, 'HardDepositAmountChanged', {
      beneficiary,
      amount: BigNumber.from(amount.toString()),
    });
  });
}
function shouldNotIncreaseHardDeposit(beneficiary, amount, from, value, revertMessage) {  
  it('reverts', async function () {
    const fromAddressSigner = await ethers.getSigner(from);
    try {
      await expect(this.ERC20SimpleSwap.connect(fromAddressSigner).increaseHardDeposit(beneficiary, amount.toString(), { value: value.toString()}))
      .to.be.revertedWith(revertMessage);
    } catch (error) {
      await expect(error.toString()).to.include('non-payable method cannot override value')
    }
  });
}
function shouldSetCustomHardDepositTimeout(beneficiary, timeout, from) {
  beforeEach(async function () {
    const beneficiarySig = await signCustomDecreaseTimeout(this.ERC20SimpleSwap, beneficiary, timeout.toString(), beneficiary);

    const fromAdressSigner = await ethers.getSigner(from);
    const tax = await this.ERC20SimpleSwap.connect(fromAdressSigner).setCustomHardDepositTimeout(beneficiary, timeout.toString(), beneficiarySig);
    const receipt = await tax.wait();

    this.receipt = receipt;
    this.logs = receipt.logs;

    this.postconditions = {
      hardDepositFor: await this.ERC20SimpleSwap.hardDeposits(beneficiary),
    };
  });

  it('should have set the timeout', async function () {
    expect(this.postconditions.hardDepositFor.timeout).to.equal(BigNumber.from(timeout.toString()));
  });

  it('emits a HardDepositTimeoutChanged event', function () {
    expectEvent.inLogs(this.receipt.events, 'HardDepositTimeoutChanged', {
      beneficiary,
      timeout: timeout.toString(),
    });
  });
}
function shouldNotSetCustomHardDepositTimeout(toSubmit, toSign, signee, from, value, revertMessage) { 
  beforeEach(async function () {
    this.beneficiarySig = await signCustomDecreaseTimeout(
      this.ERC20SimpleSwap,
      toSign.beneficiary,
      toSign.timeout,
      signee
    );
  });

  it('reverts', async function () {
    const fromAddressSigner = await ethers.getSigner(from);
    try {
      await expect(this.ERC20SimpleSwap.connect(fromAddressSigner).setCustomHardDepositTimeout(toSubmit.beneficiary, toSubmit.timeout.toString(), this.beneficiarySig, {value: value.toString()}))
      .to.be.revertedWith(revertMessage);
    } catch(error) {
      await expect(error.toString()).to.include('non-payable method cannot override value')
    }
  });
}

function shouldWithdraw(amount, from) {
  beforeEach(async function () {
    this.preconditions = {
      callerBalance: await this.TestToken.balanceOf(from),
      liquidBalance: await this.ERC20SimpleSwap.liquidBalance(),
    };

    await this.ERC20SimpleSwap.withdraw(amount.toString(), { from: from });

    this.postconditions = {
      callerBalance: await this.TestToken.balanceOf(from),
      liquidBalance: await this.ERC20SimpleSwap.liquidBalance(),
    };
  });

  it('should have updated the liquidBalance', function () {
    expect(this.postconditions.liquidBalance).to.equal(this.preconditions.liquidBalance.sub(BigNumber.from(amount.toString())));
  });

  it('should have updated the callerBalance', function () {
    expect(this.postconditions.callerBalance).to.equal(this.preconditions.callerBalance.add(BigNumber.from(amount.toString())));
  });
}
function shouldNotWithdraw(amount, from, value, revertMessage) { 
  it('reverts', async function () {  
    const fromAddressSigner = await ethers.getSigner(from);
    try {
      await expect(this.ERC20SimpleSwap.connect(fromAddressSigner).withdraw(amount.toString(), {value: value.toString()}))
      .to.be.revertedWith(revertMessage);
    } catch(error) {
      await expect(error.toString()).to.include('non-payable method cannot override value')
    }
  });
}

function shouldDeposit(amount, from) {
  beforeEach(async function () {
    this.preconditions = {
      balance: await this.ERC20SimpleSwap.balance(),
      totalHardDeposit: await this.ERC20SimpleSwap.totalHardDeposit(),
      liquidBalance: await this.ERC20SimpleSwap.liquidBalance(),
    };

    const tx = await this.TestToken.transfer(this.ERC20SimpleSwap.address, amount.toString(), { from: from });
    const receipt = await tx.wait()
    this.receipt = receipt;
    
  });
  it('should update the liquidBalance of the checkbook', async function () {
    expect(await this.ERC20SimpleSwap.liquidBalance()).to.equal(this.preconditions.liquidBalance.add(BigNumber.from(amount.toString())));
  });
  it('should update the balance of the checkbook', async function () {
    expect(await this.ERC20SimpleSwap.balance()).to.equal(this.preconditions.balance.add(BigNumber.from(amount.toString())));
  });
  it('should not afect the totalHardDeposit', async function () {
    expect(await this.ERC20SimpleSwap.totalHardDeposit()).to.equal(this.preconditions.totalHardDeposit);
  });
  it('should emit a transfer event', async function () {
    expectEvent.inLogs(this.receipt.events, 'Transfer', {
      from: from,
      to: this.ERC20SimpleSwap.address,
      value: amount.toString(),
    });
  });
}

export {
  shouldDeploy,
  shouldReturnDefaultHardDepositTimeout,
  shouldReturnPaidOut,
  shouldReturnTotalPaidOut,
  shouldReturnHardDeposits,
  shouldReturnTotalHardDeposit,
  shouldReturnIssuer,
  shouldReturnLiquidBalance,
  shouldReturnLiquidBalanceFor,
  shouldCashChequeBeneficiary,
  shouldNotCashChequeBeneficiary,
  shouldCashCheque,
  shouldNotCashCheque,
  shouldPrepareDecreaseHardDeposit,
  shouldNotPrepareDecreaseHardDeposit,
  shouldDecreaseHardDeposit,
  shouldNotDecreaseHardDeposit,
  shouldIncreaseHardDeposit,
  shouldNotIncreaseHardDeposit,
  shouldSetCustomHardDepositTimeout,
  shouldNotSetCustomHardDepositTimeout,
  shouldWithdraw,
  shouldNotWithdraw,
  shouldDeposit,
};
