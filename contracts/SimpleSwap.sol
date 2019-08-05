pragma solidity ^0.5.10;
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/math/Math.sol";
import "openzeppelin-solidity/contracts/cryptography/ECDSA.sol";

/// @title Chequebook wihout waivers
contract SimpleSwap {
  using SafeMath for uint;

  event Deposit(address depositor, uint amount);
  event ChequeCashed(
    address indexed beneficiary,
    address indexed recipient,
    address indexed callee,
    uint totalPayout,
    uint cumulativePayout,
    uint calleePayout
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

  /* associates every beneficiary with their paidOutCheques */
  mapping (address => uint) public paidOutCheques;
  /* associates every beneficiary with their HardDeposit */
  mapping (address => HardDeposit) public hardDeposits;
  /* sum of all hard deposits */
  uint public totalHardDeposit;

  /* issuer of the contract, set at construction */
  address payable public issuer;

  /// @notice constructor, allows setting the issuer (needed for "setup wallet as payment")
  constructor(address payable _issuer, uint defaultHardDepositTimeoutDuration) public payable {
    DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT = defaultHardDepositTimeoutDuration;
    issuer = _issuer;
    if(msg.value > 0) {
      emit Deposit(msg.sender, msg.value);
    }
  }

  /// @return the part of the balance that is not covered by hard deposits
  function liquidBalance() public view returns(uint) {
    return address(this).balance.sub(totalHardDeposit);
  }

  /// @return the part of the balance usable for a specific beneficiary
  function balanceFor(address beneficiary) public view returns(uint) {
    return liquidBalance().add(hardDeposits[beneficiary].amount);
  }

  function _cashChequeInternal(address beneficiary, address payable recipient, uint cumulativePayout, uint calleePayout, bytes memory issuerSig) public {
    /* The issuer must have given explicit approval to the cumulativePayout, either by being the callee or by signature*/
    if(msg.sender != issuer) {
      require(issuer == recover(chequeHash(address(this), beneficiary, cumulativePayout), issuerSig),
      "SimpleSwap: invalid issuerSig");
    }
    /* the requestPayout is the amount requested for payment processing */
    uint requestPayout = cumulativePayout.sub(paidOutCheques[beneficiary]);
    /* calculates acutal payout */
    uint totalPayout = Math.min(requestPayout, balanceFor(beneficiary));
    /* calculates hard-deposit usage */
    uint hardDepositUsage = Math.min(totalPayout, hardDeposits[beneficiary].amount);
    require(totalPayout >= calleePayout, "SimpleSwap: cannot pay callee");
    /* if there are some of the hard deposit used, update hardDeposits*/
    if(hardDepositUsage != 0) {
      hardDeposits[beneficiary].amount = hardDeposits[beneficiary].amount.sub(hardDepositUsage);
      totalHardDeposit = totalHardDeposit.sub(hardDepositUsage);
    }
    /* increase the stored paidOut amount to avoid double payout */
    paidOutCheques[beneficiary] = paidOutCheques[beneficiary].add(totalPayout);
    /* do the actual payments */
    recipient.transfer(totalPayout.sub(calleePayout));
    /* do a transfer to the callee if specified*/
    if(calleePayout != 0) {
      msg.sender.transfer(calleePayout);
    }
    emit ChequeCashed(beneficiary, recipient, msg.sender, totalPayout, cumulativePayout, calleePayout);
    /* let the world know that the issuer has over-promised on outstanding cheques */
    if(requestPayout != totalPayout) {
      emit ChequeBounced();
    }
  }

  function cashCheque(
    address beneficiary,
    address payable recipient,
    uint cumulativePayout,
    bytes memory beneficiarySig,
    uint256 calleePayout,
    bytes memory issuerSig
  ) public {
    /* the beneficiary must have given explicit approval for cashing out the cumulativePayout and the calleePayout payment */
    require(beneficiary == recover(cashOutHash(
      address(this),
      msg.sender,
      cumulativePayout,
      recipient,
      calleePayout
      ), beneficiarySig), "SimpleSwap: invalid beneficiarySig");
    _cashChequeInternal(beneficiary, recipient, cumulativePayout, calleePayout, issuerSig);
  }
  /// @notice Cash a cheque
  /// @param recipient agent (of the beneficiary) who receives the payment (i.e. other chequebook contract or the beneficiary)
  /// @param cumulativePayout amount requested to pay out
  function cashChequeBeneficiary(address payable recipient, uint cumulativePayout, bytes memory issuerSig) public {
    _cashChequeInternal(msg.sender, recipient, cumulativePayout, 0, issuerSig);
  }

  /// @notice prepare to decrease the hard deposit
  /// @param beneficiary beneficiary whose hard deposit should be decreased
  /// @param decreaseAmount amount that the deposit is supposed to be decreased by
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

  /// @notice actually decrease the hard deposit
  /// @param beneficiary beneficiary whose hard deposit should be decreased
  function decreaseHardDeposit(address beneficiary) public {
    HardDeposit storage hardDeposit = hardDeposits[beneficiary];
    require(now >= hardDeposit.canBeDecreasedAt && hardDeposit.canBeDecreasedAt != 0, "SimpleSwap: deposit not yet timed out");
    /* this throws if decreaseAmount > amount */
    hardDeposit.amount = hardDeposit.amount.sub(hardDeposit.decreaseAmount);
    /* reset the canBeDecreasedAt to avoid a double decrease */
    hardDeposit.canBeDecreasedAt = 0;
    /* keep totalDeposit in sync */
    totalHardDeposit = totalHardDeposit.sub(hardDeposit.decreaseAmount);
    emit HardDepositAmountChanged(beneficiary, hardDeposit.amount);
  }

  /// @notice increase the hard deposit
  /// @param beneficiary beneficiary whose hard deposit should be decreased
  /// @param amount the new hard deposit
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


  function setCustomHardDepositDecreaseTimeout(
    address beneficiary,
    uint decreaseTimeout,
    bytes memory beneficiarySig
  ) public {
    require(msg.sender == issuer, "SimpleSwap: not issuer");
    require(beneficiary == recover(customDecreaseTimeoutHash(address(this), beneficiary, decreaseTimeout), beneficiarySig), "SimpleSwap: invalid beneficiarySig");
    hardDeposits[beneficiary].decreaseTimeout = decreaseTimeout;
    emit HardDepositDecreaseTimeoutChanged(beneficiary, hardDeposits[beneficiary].decreaseTimeout);
  }

  /// @notice withdraw ether
  /// @param amount amount to withdraw
  function withdraw(uint amount) public {
    /* only issuer can do this */
    require(msg.sender == issuer, "SimpleSwap: not issuer");
    /* ensure we don't take anything from the hard deposit */
    require(amount <= liquidBalance(), "SimpleSwap: liquidBalance not sufficient");
    issuer.transfer(amount);
    emit Withdraw(amount);
  }

  /// @notice deposit ether
  function() payable external {
    if(msg.value > 0) {
      emit Deposit(msg.sender, msg.value);
    }
  }

  function recover(bytes32 hash, bytes memory sig) internal pure returns (address) {
    return ECDSA.recover(ECDSA.toEthSignedMessageHash(hash), sig);
  }

  function chequeHash(address swap, address beneficiary, uint cumulativePayout)
  public pure returns (bytes32) {
    return keccak256(abi.encodePacked(swap, beneficiary, cumulativePayout));
  }

  function cashOutHash(address swap, address sender, uint requestPayout, address recipient, uint calleePayout)
  public pure returns (bytes32) {
    return keccak256(abi.encodePacked(swap, sender, requestPayout, recipient, calleePayout));
  }

  function customDecreaseTimeoutHash(address swap, address beneficiary, uint decreaseTimeout) public pure returns (bytes32) {
    return keccak256(abi.encode(swap, beneficiary, decreaseTimeout));
  }

}


