pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;
import "./abstracts/AbstractWitness.sol";
import "./Swindle.sol";

contract Swear is AbstractWitness, AbstractConstants {

  Swindle public swindle;

  /// @notice constructor, allows setting the swindle
  constructor(address _swindle) public {
    swindle = Swindle(_swindle);
  }

  function testimonyFor(bytes memory specification, bytes memory payload)
  public view returns (TestimonyStatus, bytes memory) {
    bytes32 remark = abi.decode(specification, (bytes32));    
    (address expectedRules, bytes memory expectedInput, bytes32 caseId) = abi.decode(payload, (address,bytes, bytes32));
    require(remark == keccak256(abi.encode(expectedRules, expectedInput)), "input does not match remark");

    (uint8 status, address rules, bytes32 inputHash) = swindle.getTrialInfo(caseId);

    require(status == TRIAL_STATUS_GUILTY, "not guilty");
    require(rules == expectedRules, "wrong rules");
    require(inputHash == keccak256(expectedInput), "wrong input data");

    return (TestimonyStatus.VALID, new bytes(0));
  }

  function encodePayload(address expectedRules, bytes memory expectedInput, bytes32 caseId)
  public pure returns (bytes memory) {
    return abi.encode(expectedRules, expectedInput, caseId);
  }

}