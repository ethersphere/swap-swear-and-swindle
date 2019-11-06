pragma solidity ^0.5.11;
import "./SimpleSwap.sol";
import "./ERC20SimpleSwap.sol";

/**
@title Factory contract for SimpleSwap
@author The Swarm Authors
@notice This contract deploys SimpleSwap contracts
*/
contract SimpleSwapFactory {

  /* event fired on every new SimpleSwap deployment */
  event SimpleSwapDeployed(address contractAddress);

  /* mapping to keep track of which contracts were deployed by this factory */
  mapping (address => bool) public deployedContracts;

  /* address of the ERC20-token, to be used by the to-be-deployed chequebooks.
  If 0, deploy SimpleSwap, otherwise ERC20SimpleSwap */
  address public ERC20Address;
  constructor(address _ERC20Address) public {
    ERC20Address = _ERC20Address;
  }
  /**
  @notice deployes a new SimpleSwap contract
  @param issuer the issuer of cheques for the new chequebook
  @param defaultHardDepositTimeoutDuration duration in seconds which by default will be used to reduce hardDeposit allocations
  */
  function deploySimpleSwap(address payable issuer, uint defaultHardDepositTimeoutDuration)
  public payable returns (address) {
    if(ERC20Address != address(0)) {
      return _deployERC20SimpleSwap(issuer, defaultHardDepositTimeoutDuration);
    } else {
      return _deploySimpleSwap(issuer, defaultHardDepositTimeoutDuration);
    }
  }

  function _deployERC20SimpleSwap(address payable issuer, uint defaultHardDepositTimeoutDuration)
  internal returns (address) {
    address contractAddress = address((new ERC20SimpleSwap).value(msg.value)(issuer, ERC20Address, defaultHardDepositTimeoutDuration));
    deployedContracts[contractAddress] = true;
    emit SimpleSwapDeployed(contractAddress);
    return contractAddress;
  }

  function _deploySimpleSwap(address payable issuer, uint defaultHardDepositTimeoutDuration)
    internal returns (address) {
    address contractAddress = address((new SimpleSwap).value(msg.value)(issuer, defaultHardDepositTimeoutDuration));
    deployedContracts[contractAddress] = true;
    emit SimpleSwapDeployed(contractAddress);
    return contractAddress;
  }
}