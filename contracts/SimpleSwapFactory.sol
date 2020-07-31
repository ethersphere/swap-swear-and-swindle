pragma solidity =0.6.12;
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

  /* address of the ERC20-token, to be used by the to-be-deployed chequebooks */
  address public ERC20Address;

  constructor(address _ERC20Address) public {
    ERC20Address = _ERC20Address;
  }
  /**
  @notice deployes a new SimpleSwap contract
  @param issuer the issuer of cheques for the new chequebook
  @param defaultHardDepositTimeoutDuration duration in seconds which by default will be used to reduce hardDeposit allocations
  */
  function deploySimpleSwap(address issuer, uint defaultHardDepositTimeoutDuration)
  public returns (address) {
    address contractAddress = address(new ERC20SimpleSwap(issuer, ERC20Address, defaultHardDepositTimeoutDuration));
    deployedContracts[contractAddress] = true;
    emit SimpleSwapDeployed(contractAddress);
    return contractAddress;
  }
}