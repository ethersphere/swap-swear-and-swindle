pragma solidity =0.6.12;
import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/math/Math.sol";
import "@openzeppelin/contracts/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/**
@title Chequebook contract without waivers
@author The Swarm Authors
@notice The chequebook contract allows the issuer of the chequebook to send cheques to an unlimited amount of counterparties.
Furthermore, solvency can be guaranteed via hardDeposits
@dev as an issuer, no cheques should be send if the cumulative worth of a cheques send is above the cumulative worth of all deposits
as a beneficiary, we should always take into account the possibility that a cheque bounces (when no hardDeposits are assigned)
*/
contract ERC20SimpleSwap {
  using SafeMath for uint;

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
  event HardDepositTimeoutChanged(address indexed beneficiary, uint timeout);
  event Withdraw(uint amount);

  uint public defaultHardDepositTimeout;
  /* structure to keep track of the hard deposits (on-chain guarantee of solvency) per beneficiary*/
  struct HardDeposit {
    uint amount; /* hard deposit amount allocated */
    uint decreaseAmount; /* decreaseAmount substranced from amount when decrease is requested */
    uint timeout; /* issuer has to wait timeout seconds to decrease hardDeposit, 0 implies applying defaultHardDepositTimeout */
    uint canBeDecreasedAt; /* point in time after which harddeposit can be decreased*/
  }

  struct EIP712Domain {
    string name;
    string version;
    uint256 chainId;
  }

  bytes32 public constant EIP712DOMAIN_TYPEHASH = keccak256(
    "EIP712Domain(string name,string version,uint256 chainId)"
  );
  bytes32 public constant CHEQUE_TYPEHASH = keccak256(
    "Cheque(address chequebook,address beneficiary,uint256 cumulativePayout)"
  );
  bytes32 public constant CASHOUT_TYPEHASH = keccak256(
    "Cashout(address chequebook,address sender,uint256 requestPayout,address recipient,uint256 callerPayout)"
  );
  bytes32 public constant CUSTOMDECREASETIMEOUT_TYPEHASH = keccak256(
    "CustomDecreaseTimeout(address chequebook,address beneficiary,uint256 decreaseTimeout)"
  );

  // the EIP712 domain this contract uses
  function domain() internal pure returns (EIP712Domain memory) {
    uint256 chainId;
    assembly {
      chainId := chainid()
    }
    return EIP712Domain({
      name: "Chequebook",
      version: "1.0",
      chainId: chainId
    });
  }

  // compute the EIP712 domain separator. this cannot be constant because it depends on chainId
  function domainSeparator(EIP712Domain memory eip712Domain) internal pure returns (bytes32) {
    return keccak256(abi.encode(
        EIP712DOMAIN_TYPEHASH,
        keccak256(bytes(eip712Domain.name)),
        keccak256(bytes(eip712Domain.version)),
        eip712Domain.chainId
    ));
  }

  // recover a signature with the EIP712 signing scheme
  function recoverEIP712(bytes32 hash, bytes memory sig) internal pure returns (address) {
    bytes32 digest = keccak256(abi.encodePacked(
        "\x19\x01",
        domainSeparator(domain()),
        hash
    ));
    return ECDSA.recover(digest, sig);
  }

  /* The token against which this chequebook writes cheques */
  ERC20 public token;
  /* associates every beneficiary with how much has been paid out to them */
  mapping (address => uint) public paidOut;
  /* total amount paid out */
  uint public totalPaidOut;
  /* associates every beneficiary with their HardDeposit */
  mapping (address => HardDeposit) public hardDeposits;
  /* sum of all hard deposits */
  uint public totalHardDeposit;
  /* issuer of the contract, set at construction */
  address public issuer;
  /* indicates wether a cheque bounced in the past */
  bool public bounced;

  /**
  @notice sets the issuer, defaultHardDepositTimeout and receives an initial deposit
  @param _issuer the issuer of cheques from this chequebook (needed as an argument for "Setting up a chequebook as a payment").
  _issuer must be an Externally Owned Account, or it must support calling the function cashCheque
  @param _defaultHardDepositTimeout duration in seconds which by default will be used to reduce hardDeposit allocations
  */
  constructor(address _issuer, address _token, uint _defaultHardDepositTimeout) public {
    issuer = _issuer;
    token = ERC20(_token);
    defaultHardDepositTimeout = _defaultHardDepositTimeout;
  }

  /// @return the balance of the chequebook
  function balance() public view returns(uint) {
    return token.balanceOf(address(this));
  }
  /// @return the part of the balance that is not covered by hard deposits
  function liquidBalance() public view returns(uint) {
    return balance().sub(totalHardDeposit);
  }

  /// @return the part of the balance available for a specific beneficiary
  function liquidBalanceFor(address beneficiary) public view returns(uint) {
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
    address recipient,
    uint cumulativePayout,
    uint callerPayout,
    bytes memory issuerSig
  ) internal {
    /* The issuer must have given explicit approval to the cumulativePayout, either by being the caller or by signature*/
    if (msg.sender != issuer) {
      require(issuer == recoverEIP712(chequeHash(address(this), beneficiary, cumulativePayout), issuerSig),
      "SimpleSwap: invalid issuer signature");
    }
    /* the requestPayout is the amount requested for payment processing */
    uint requestPayout = cumulativePayout.sub(paidOut[beneficiary]);
    /* calculates acutal payout */
    uint totalPayout = Math.min(requestPayout, liquidBalanceFor(beneficiary));
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
    totalPaidOut = totalPaidOut.add(totalPayout);
    /* do the actual payments */

    require(token.transfer(recipient, totalPayout.sub(callerPayout)), "SimpleSwap: SimpleSwap: transfer failed");
    /* do a transfer to the caller if specified*/
    if (callerPayout != 0) {
      require(token.transfer(msg.sender, callerPayout), "SimpleSwap: SimpleSwap: transfer failed");
    }
    emit ChequeCashed(beneficiary, recipient, msg.sender, totalPayout, cumulativePayout, callerPayout);
    /* let the world know that the issuer has over-promised on outstanding cheques */
    if (requestPayout != totalPayout) {
      bounced = true;
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
    address recipient,
    uint cumulativePayout,
    bytes memory beneficiarySig,
    uint256 callerPayout,
    bytes memory issuerSig
  ) public {
    require(
      beneficiary == recoverEIP712(
        cashOutHash(
          address(this),
          msg.sender,
          cumulativePayout,
          recipient,
          callerPayout
        ), beneficiarySig
      ), "SimpleSwap: invalid beneficiary signature");
    _cashChequeInternal(beneficiary, recipient, cumulativePayout, callerPayout, issuerSig);
  }

  /**
  @notice cash a cheque as beneficiary
  @param recipient receives the differences between cumulativePayment and what was already paid-out to the beneficiary minus callerPayout
  @param cumulativePayout amount requested to pay out
  @param issuerSig issuer must have given explicit approval on the cumulativePayout to the beneficiary
  */
  function cashChequeBeneficiary(address recipient, uint cumulativePayout, bytes memory issuerSig) public {
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
    // if hardDeposit.timeout was never set, apply defaultHardDepositTimeout
    uint timeout = hardDeposit.timeout == 0 ? defaultHardDepositTimeout : hardDeposit.timeout;
    hardDeposit.canBeDecreasedAt = now + timeout;
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
    require(totalHardDeposit.add(amount) <= balance(), "SimpleSwap: hard deposit cannot be more than balance");

    HardDeposit storage hardDeposit = hardDeposits[beneficiary];
    hardDeposit.amount = hardDeposit.amount.add(amount);
    // we don't explicitely set hardDepositTimout, as zero means using defaultHardDepositTimeout
    totalHardDeposit = totalHardDeposit.add(amount);
    /* disable any pending decrease */
    hardDeposit.canBeDecreasedAt = 0;
    emit HardDepositAmountChanged(beneficiary, hardDeposit.amount);
  }

  /**
  @notice allows for setting a custom hardDepositDecreaseTimeout per beneficiary
  @dev this is required when solvency must be guaranteed for a period longer than the defaultHardDepositDecreaseTimeout
  @param beneficiary beneficiary whose hard deposit decreaseTimeout must be changed
  @param hardDepositTimeout new hardDeposit.timeout for beneficiary
  @param beneficiarySig beneficiary must give explicit approval by giving his signature on the new decreaseTimeout
  */
  function setCustomHardDepositTimeout(
    address beneficiary,
    uint hardDepositTimeout,
    bytes memory beneficiarySig
  ) public {
    require(msg.sender == issuer, "SimpleSwap: not issuer");
    require(
      beneficiary == recoverEIP712(customDecreaseTimeoutHash(address(this), beneficiary, hardDepositTimeout), beneficiarySig),
      "SimpleSwap: invalid beneficiary signature"
    );
    hardDeposits[beneficiary].timeout = hardDepositTimeout;
    emit HardDepositTimeoutChanged(beneficiary, hardDepositTimeout);
  }

  /// @notice withdraw ether
  /// @param amount amount to withdraw
  // solhint-disable-next-line no-simple-event-func-name
  function withdraw(uint amount) public {
    /* only issuer can do this */
    require(msg.sender == issuer, "SimpleSwap: not issuer");
    /* ensure we don't take anything from the hard deposit */
    require(amount <= liquidBalance(), "SimpleSwap: liquidBalance not sufficient");
    require(token.transfer(issuer, amount), "SimpleSwap: SimpleSwap: transfer failed");
  }

  function chequeHash(address chequebook, address beneficiary, uint cumulativePayout)
  internal pure returns (bytes32) {
    return keccak256(abi.encode(
      CHEQUE_TYPEHASH,
      chequebook,
      beneficiary,
      cumulativePayout
    ));
  }  

  function cashOutHash(address chequebook, address sender, uint requestPayout, address recipient, uint callerPayout)
  internal pure returns (bytes32) {
    return keccak256(abi.encode(
      CASHOUT_TYPEHASH,
      chequebook,
      sender,
      requestPayout,
      recipient,
      callerPayout
    ));
  }

  function customDecreaseTimeoutHash(address chequebook, address beneficiary, uint decreaseTimeout)
  internal pure returns (bytes32) {
    return keccak256(abi.encode(
      CUSTOMDECREASETIMEOUT_TYPEHASH,
      chequebook,
      beneficiary,
      decreaseTimeout
    ));
  }
}


