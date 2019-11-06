pragma solidity ^0.5.11;
import "./ISimpleSwap.sol";
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/math/Math.sol";
import "openzeppelin-solidity/contracts/cryptography/ECDSA.sol";

/**
@title Chequebook contract without waivers
@author The Swarm Authors
@notice The chequebook contract allows the issuer of the chequebook to send cheques to an unlimited amount of counterparties.
Furthermore, solvency can be guaranteed via hardDeposits
@dev as an issuer, no cheques should be send if the cumulative worth of a cheques send is above the cumulative worth of all deposits
as a beneficiary, we should always take into account the possibility that a cheque bounces (when no hardDeposits are assigned)
*/
contract SimpleSwap is ISimpleSwap {
  using SafeMath for uint;

  event Deposit(address depositor, uint amount);
  event ChequeCashed(
    address indexed beneficiary,
    address indexed recipient,
    address indexed caller,
    uint totalPayout,
    uint cumulativePayout,
    uint callerPayout
  );
  event ChequeBounced();
  event HardDepositAmountChanged(address indexed beneficiary, uint amount);
  event HardDepositDecreasePrepared(address indexed beneficiary, uint decreaseAmount);
  event HardDepositDecreaseTimeoutChanged(address indexed beneficiary, uint decreaseTimeout);
  event Withdraw(uint amount);

  uint public DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT;
  /* structure to keep track of the hard deposits (on-chain guarantee of solvency) per beneficiary*/
  struct HardDeposit {
    uint amount; /* hard deposit amount allocated */
    uint decreaseAmount; /* decreaseAmount substranced from amount when decrease is requested */
    uint decreaseTimeout; /* issuer has to wait decreaseTimeout seconds after decrease is requested to decrease hardDeposit */
    uint canBeDecreasedAt; /* point in time after which harddeposit can be decreased*/
  }

  /* associates every beneficiary with how much has been paid out to them */
  mapping (address => uint) public paidOut;
  /* associates every beneficiary with their HardDeposit */
  mapping (address => HardDeposit) public hardDeposits;
  /* sum of all hard deposits */
  uint public totalHardDeposit;

  /* issuer of the contract, set at construction */
  address payable public issuer;
  /**
  @notice sets the issuer, defaultHardDepositTimeoutDuration and receives an initial deposit
  @param _issuer the issuer of cheques from this chequebook (needed as an argument for "Setting up a chequebook as a payment").
  _issuer must be an Externally Owned Account, or it must support calling the function cashCheque
  @param defaultHardDepositTimeoutDuration duration in seconds which by default will be used to reduce hardDeposit allocations
  */
  constructor(address payable _issuer, uint defaultHardDepositTimeoutDuration) public payable {
    DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT = defaultHardDepositTimeoutDuration;
    issuer = _issuer;
    if (msg.value > 0) {
      emit Deposit(msg.sender, msg.value);
    }
  }

  /// @return the part of the balance that is not covered by hard deposits
  function liquidBalance() public view returns(uint) {
    return address(this).balance.sub(totalHardDeposit);
  }

  /// @return the part of the balance available for a specific beneficiary
  function availableBalanceFor(address beneficiary) public view returns(uint) {
    return liquidBalance().add(hardDeposits[beneficiary].amount);
  }
  /**
  @dev internal function responsible for checking the issuerSignature, updating hardDeposit balances and doing transfers.
  Called by cashCheque and cashChequeBeneficary
  @param beneficiary the beneficiary to which cheques were assigned. Beneficiary must be an Externally Owned Account
  @param recipient receives the differences between cumulativePayment and what was already paid-out to the beneficiary minus callerPayout
  @param cumulativePayout cumulative amount of cheques assigned to beneficiary
  @param issuerSig if issuer is not the sender, issuer must have given explicit approval on the cumulativePayout to the beneficiary
  */
  function _cashChequeInternal(
    address beneficiary,
    address payable recipient,
    uint cumulativePayout,
    uint callerPayout,
    bytes memory issuerSig
  ) internal {
    /* The issuer must have given explicit approval to the cumulativePayout, either by being the caller or by signature*/
    if (msg.sender != issuer) {
      require(
        issuer == recover(
          chequeHash(
            address(this),
            beneficiary,
            cumulativePayout
          ), issuerSig
        ), "SimpleSwap: invalid issuerSig"
      );
    }
    /* the requestPayout is the amount requested for payment processing */
    uint requestPayout = cumulativePayout.sub(paidOut[beneficiary]);
    /* calculates acutal payout */
    uint totalPayout = Math.min(requestPayout, availableBalanceFor(beneficiary));
    /* calculates hard-deposit usage */
    uint hardDepositUsage = Math.min(totalPayout, hardDeposits[beneficiary].amount);
    require(totalPayout >= callerPayout, "SimpleSwap: cannot pay caller");
    /* if there are some of the hard deposit used, update hardDeposits*/
    if (hardDepositUsage != 0) {
      hardDeposits[beneficiary].amount = hardDeposits[beneficiary].amount.sub(hardDepositUsage);

      totalHardDeposit = totalHardDeposit.sub(hardDepositUsage);
    }
    /* increase the stored paidOut amount to avoid double payout */
    paidOut[beneficiary] = paidOut[beneficiary].add(totalPayout);
    /* do the actual payments */

    recipient.transfer(totalPayout.sub(callerPayout));
    /* do a transfer to the caller if specified*/
    if (callerPayout != 0) {
      msg.sender.transfer(callerPayout);
    }
    emit ChequeCashed(beneficiary, recipient, msg.sender, totalPayout, cumulativePayout, callerPayout);
    /* let the world know that the issuer has over-promised on outstanding cheques */
    if (requestPayout != totalPayout) {
      emit ChequeBounced();
    }
  }
  /**
  @notice cash a cheque of the beneficiary by a non-beneficiary and reward the sender for doing so with callerPayout
  @dev a beneficiary must be able to generate signatures (be an Externally Owned Account) to make use of this feature
  @param beneficiary the beneficiary to which cheques were assigned. Beneficiary must be an Externally Owned Account
  @param recipient receives the differences between cumulativePayment and what was already paid-out to the beneficiary minus callerPayout
  @param cumulativePayout cumulative amount of cheques assigned to beneficiary
  @param beneficiarySig beneficiary must have given explicit approval for cashing out the cumulativePayout by the sender and sending the callerPayout
  @param issuerSig if issuer is not the sender, issuer must have given explicit approval on the cumulativePayout to the beneficiary
  @param callerPayout when beneficiary does not have ether yet, he can incentivize other people to cash cheques with help of callerPayout
  @param issuerSig if issuer is not the sender, issuer must have given explicit approval on the cumulativePayout to the beneficiary
  */
  function cashCheque(
    address beneficiary,
    address payable recipient,
    uint cumulativePayout,
    bytes memory beneficiarySig,
    uint256 callerPayout,
    bytes memory issuerSig
  ) public {
    require(
      beneficiary == recover(
        cashOutHash(
          address(this),
          msg.sender,
          cumulativePayout,
          recipient,
          callerPayout
        ), beneficiarySig
      ), "SimpleSwap: invalid beneficiarySig");
    _cashChequeInternal(beneficiary, recipient, cumulativePayout, callerPayout, issuerSig);
  }

  /**
  @notice cash a cheque as beneficiary
  @param recipient receives the differences between cumulativePayment and what was already paid-out to the beneficiary minus callerPayout
  @param cumulativePayout amount requested to pay out
  @param issuerSig issuer must have given explicit approval on the cumulativePayout to the beneficiary
  */
  function cashChequeBeneficiary(address payable recipient, uint cumulativePayout, bytes memory issuerSig) public {
    _cashChequeInternal(msg.sender, recipient, cumulativePayout, 0, issuerSig);
  }

  /**
  @notice prepare to decrease the hard deposit
  @dev decreasing hardDeposits must be done in two steps to allow beneficiaries to cash any uncashed cheques (and make use of the assgined hard-deposits)
  @param beneficiary beneficiary whose hard deposit should be decreased
  @param decreaseAmount amount that the deposit is supposed to be decreased by
  */
  function prepareDecreaseHardDeposit(address beneficiary, uint decreaseAmount) public {
    require(msg.sender == issuer, "SimpleSwap: not issuer");
    HardDeposit storage hardDeposit = hardDeposits[beneficiary];
    /* cannot decrease it by more than the deposit */
    require(decreaseAmount <= hardDeposit.amount, "SimpleSwap: hard deposit not sufficient");
    // if hardDeposit.decreaseTimeout was never set, we DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT. Otherwise we use the one which was set.
    uint decreaseTimeout = hardDeposit.decreaseTimeout == 0 ? DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT : hardDeposit.decreaseTimeout;
    hardDeposit.canBeDecreasedAt = now + decreaseTimeout;
    hardDeposit.decreaseAmount = decreaseAmount;
    emit HardDepositDecreasePrepared(beneficiary, decreaseAmount);
  }

  /**
  @notice decrease the hard deposit after waiting the necesary amount of time since prepareDecreaseHardDeposit was called
  @param beneficiary beneficiary whose hard deposit should be decreased
  */
  function decreaseHardDeposit(address beneficiary) public {
    HardDeposit storage hardDeposit = hardDeposits[beneficiary];
    require(now >= hardDeposit.canBeDecreasedAt && hardDeposit.canBeDecreasedAt != 0, "SimpleSwap: deposit not yet timed out");
    /* this throws if decreaseAmount > amount */
    //TODO: if there is a cash-out in between prepareDecreaseHardDeposit and decreaseHardDeposit, decreaseHardDeposit will throw and reducing hard-deposits is impossible.
    hardDeposit.amount = hardDeposit.amount.sub(hardDeposit.decreaseAmount);
    /* reset the canBeDecreasedAt to avoid a double decrease */
    hardDeposit.canBeDecreasedAt = 0;
    /* keep totalDeposit in sync */
    totalHardDeposit = totalHardDeposit.sub(hardDeposit.decreaseAmount);
    emit HardDepositAmountChanged(beneficiary, hardDeposit.amount);
  }

  /**
  @notice increase the hard deposit
  @param beneficiary beneficiary whose hard deposit should be decreased
  @param amount the new hard deposit
  */
  function increaseHardDeposit(address beneficiary, uint amount) public {
    require(msg.sender == issuer, "SimpleSwap: not issuer");
    /* ensure hard deposits don't exceed the global balance */
    require(totalHardDeposit.add(amount) <= address(this).balance, "SimpleSwap: hard deposit cannot be more than balance");

    HardDeposit storage hardDeposit = hardDeposits[beneficiary];
    hardDeposit.amount = hardDeposit.amount.add(amount);
    // we don't explicitely set decreaseTimeout, as zero means using the DEFAULT_HARDDEPOSIT_TIMEOUT_DURATION
    totalHardDeposit = totalHardDeposit.add(amount);
    /* disable any pending decrease */
    hardDeposit.canBeDecreasedAt = 0;
    emit HardDepositAmountChanged(beneficiary, hardDeposit.amount);
  }

  /**
  @notice allows for setting a custom hardDepositDecreaseTimeout per beneficiary
  @dev this is required when solvency must be guaranteed for a period longer than the defaultHardDepositDecreaseTimeout
  @param beneficiary beneficiary whose hard deposit decreaseTimeout must be changed
  @param decreaseTimeout new decreaseTimeout for beneficiary
  @param beneficiarySig beneficiary must give explicit approval by giving his signature on the new decreaseTimeout
  */
  function setCustomHardDepositDecreaseTimeout(
    address beneficiary,
    uint decreaseTimeout,
    bytes memory beneficiarySig
  ) public {
    require(msg.sender == issuer, "SimpleSwap: not issuer");
    require(
      beneficiary == recover(customDecreaseTimeoutHash(address(this), beneficiary, decreaseTimeout), beneficiarySig),
      "SimpleSwap: invalid beneficiarySig"
    );
    hardDeposits[beneficiary].decreaseTimeout = decreaseTimeout;
    emit HardDepositDecreaseTimeoutChanged(beneficiary, hardDeposits[beneficiary].decreaseTimeout);
  }

  /// @notice withdraw ether
  /// @param amount amount to withdraw
  // solhint-disable-next-line no-simple-event-func-name
  function withdraw(uint amount) public {
    /* only issuer can do this */
    require(msg.sender == issuer, "SimpleSwap: not issuer");
    /* ensure we don't take anything from the hard deposit */
    require(amount <= liquidBalance(), "SimpleSwap: liquidBalance not sufficient");
    issuer.transfer(amount);
    emit Withdraw(amount);
  }

  /// @notice deposit ether
  function() external payable {
    if (msg.value > 0) {
      emit Deposit(msg.sender, msg.value);
    }
  }

  function recover(bytes32 hash, bytes memory sig) internal pure returns (address) {
    return ECDSA.recover(ECDSA.toEthSignedMessageHash(hash), sig);
  }

  function chequeHash(address swap, address beneficiary, uint cumulativePayout)
  internal pure returns (bytes32) {
    return keccak256(abi.encodePacked(swap, beneficiary, cumulativePayout));
  }

  function cashOutHash(address swap, address sender, uint requestPayout, address recipient, uint callerPayout)
  internal pure returns (bytes32) {
    return keccak256(abi.encodePacked(swap, sender, requestPayout, recipient, callerPayout));
  }

  function customDecreaseTimeoutHash(address swap, address beneficiary, uint decreaseTimeout) 
  internal pure returns (bytes32) {
    return keccak256(abi.encodePacked(swap, beneficiary, decreaseTimeout));
  }

}


