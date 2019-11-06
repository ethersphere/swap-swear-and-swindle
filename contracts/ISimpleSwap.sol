pragma solidity ^0.5.11;

interface ISimpleSwap {
    /// @return the part of the balance that is not covered by hard deposits
  function liquidBalance() external view returns(uint);

  /// @return the part of the balance available for a specific beneficiary
  function availableBalanceFor(address beneficiary) external view returns(uint);
  
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
    bytes calldata beneficiarySig,
    uint256 callerPayout,
    bytes calldata issuerSig
  ) external;

  /**
  @notice cash a cheque as beneficiary
  @param recipient receives the differences between cumulativePayment and what was already paid-out to the beneficiary minus callerPayout
  @param cumulativePayout amount requested to pay out
  @param issuerSig issuer must have given explicit approval on the cumulativePayout to the beneficiary
  */
  function cashChequeBeneficiary(address payable recipient, uint cumulativePayout, bytes calldata issuerSig) external;

  /**
  @notice prepare to decrease the hard deposit
  @dev decreasing hardDeposits must be done in two steps to allow beneficiaries to cash any uncashed cheques (and make use of the assgined hard-deposits)
  @param beneficiary beneficiary whose hard deposit should be decreased
  @param decreaseAmount amount that the deposit is supposed to be decreased by
  */
  function prepareDecreaseHardDeposit(address beneficiary, uint decreaseAmount) external;

  /**
  @notice decrease the hard deposit after waiting the necesary amount of time since prepareDecreaseHardDeposit was called
  @param beneficiary beneficiary whose hard deposit should be decreased
  */
  function decreaseHardDeposit(address beneficiary) external;

  /**
  @notice increase the hard deposit
  @param beneficiary beneficiary whose hard deposit should be decreased
  @param amount the new hard deposit
  */
  function increaseHardDeposit(address beneficiary, uint amount) external;

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
    bytes calldata beneficiarySig
  ) external;

  /// @notice withdraw ether
  /// @param amount amount to withdraw
  // solhint-disable-next-line no-simple-event-func-name
  function withdraw(uint amount) external;

}
