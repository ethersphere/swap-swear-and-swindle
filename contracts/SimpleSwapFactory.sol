pragma solidity ^0.5.11;
import "./ERC20SimpleSwap.sol";
import "@openzeppelin/contracts/ownership/Ownable.sol";

/**
@title Factory contract for SimpleSwap
@author The Swarm Authors
@notice This contract deploys SimpleSwap contracts
*/
contract SimpleSwapFactory is Ownable {

  /* a tax that is paid (in PPM) on any withdrawal of chequebook profits */
  uint256 public tax;

  /* base value used in PPM calculations (calculation tax due) */
  uint256 public constant PPMBase = 1000000;

  /* the address to which tax is paid */
  address public taxCollector;


  /* event fired on every new SimpleSwap deployment */
  event SimpleSwapDeployed(address contractAddress);

  /* mapping to keep track of which contracts were deployed by this factory */
  mapping (address => bool) private deployedChequebooks;

  /* address of the ERC20-token, to be used by the to-be-deployed chequebooks */
  address public ERC20Address;

  constructor(address _ERC20Address, uint256 initialTax, address _taxCollector) public {
    ERC20Address = _ERC20Address;
    tax = initialTax;
    taxCollector = _taxCollector;
  }

  /**
  @notice deploys a new SimpleSwap contract
  @param issuer the issuer of cheques for the new chequebook
  @param defaultHardDepositTimeoutDuration duration in seconds which by default will be used to reduce hardDeposit allocations
  */
  function deploySimpleSwap(address issuer, uint defaultHardDepositTimeoutDuration)
  public returns (address) {
    address contractAddress = address(new ERC20SimpleSwap(issuer, ERC20Address, defaultHardDepositTimeoutDuration));
    deployedChequebooks[contractAddress] = true;
    emit SimpleSwapDeployed(contractAddress);
    return contractAddress;
  }

  /**
  @notice sets a new tax
  @param newTax the new value of tax (in PPM)
  */
  function setTax(uint256 newTax) public onlyOwner {
    tax = newTax;
  }

  /**
  @notice sets the taxCollector
  @param newTaxCollector the beneficiary of all tax payments, paid by chequebooks deployed by this factory
   */
  function setNewTaxCollector(address newTaxCollector) public onlyOwner {
    taxCollector = newTaxCollector;
  }

  function isDeployedChequebook(address chequebook) public view returns(bool) {
    return deployedChequebooks[chequebook];
  }
}