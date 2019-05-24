pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;
import "../../abstracts/AbstractTrialRules.sol";
import "../witnesses/SwapWitness.sol";
import "../witnesses/AckWitness.sol";
import "../witnesses/UpdateWitness.sol";
import "../witnesses/MerkleWitness.sol";
import "../../Swindle.sol";

contract MailboxTrial is AbstractTrialRules, SW3Utils {
  uint8 constant TRIAL_STATUS_SWEAR_NOTE = 3;
  uint8 constant TRIAL_STATUS_PAYMENT_NOTE = 4;
  uint8 constant TRIAL_STATUS_ACK = 5;
  uint8 constant TRIAL_STATUS_UPDATE = 6;
  uint8 constant TRIAL_STATUS_INCLUSION = 7;

  struct TrialData {
    uint payment;
    bytes32 serviceNoteHash;
    bytes32 dataHash;
    uint dataTime;
    bytes32 updateRoot;
  }

  mapping (bytes32 => TrialData) public data;  

  Swindle public swindle;
  SwapWitness public swapWitness;
  AckWitness public ackWitness;
  UpdateWitness public updateWitness;
  MerkleWitness public merkleWitness;

  modifier only_swindle {
    require(msg.sender == address(swindle), "only_swindle");
    _;
  }

  constructor(address _swindle, UpdateWitness _updateWitness, MerkleWitness _merkleWitness) public {
    swindle = Swindle(_swindle);
    swapWitness = new SwapWitness();
    ackWitness = new AckWitness();
    updateWitness = _updateWitness;
    merkleWitness = _merkleWitness;
  }

  function encodeNoteRemarkData(address trial, uint payment) 
  public pure returns (bytes memory) {
    return abi.encode(trial, abi.encode(payment));
  }

  function getSwearNoteRemark(  
    uint payment
  ) public view returns (bytes32) {
    return keccak256(encodeNoteRemarkData(address(this), payment));
  }

  function getInitialStatus()
  public view returns (uint8 status) {
    return TRIAL_STATUS_SWEAR_NOTE;
  }
  
  function nextStatus(AbstractWitness.TestimonyStatus witnessStatus, uint8 trialStatus)
  public view returns (uint8 status) {
    if(witnessStatus == AbstractWitness.TestimonyStatus.PENDING) {      
      if(trialStatus == TRIAL_STATUS_INCLUSION)
        return TRIAL_STATUS_GUILTY;
      else 
        return TRIAL_STATUS_NOT_GUILTY;
    }

    require(witnessStatus == AbstractWitness.TestimonyStatus.VALID, "...");

    if(trialStatus == TRIAL_STATUS_SWEAR_NOTE) {
      return TRIAL_STATUS_PAYMENT_NOTE;
    } else if(trialStatus == TRIAL_STATUS_PAYMENT_NOTE) {
      return TRIAL_STATUS_ACK;
    } else if(trialStatus == TRIAL_STATUS_ACK) {
      return TRIAL_STATUS_UPDATE;
    } else if(trialStatus == TRIAL_STATUS_UPDATE) {
      return TRIAL_STATUS_INCLUSION;
    } else if(trialStatus == TRIAL_STATUS_INCLUSION) {
      return TRIAL_STATUS_NOT_GUILTY;
    } else revert("no status");
  }

  /// @notice return witness for a given status
  function getWitness(uint8 trialStatus) 
  public view returns (address witness, uint expiry) {    
    if(trialStatus == TRIAL_STATUS_SWEAR_NOTE || trialStatus == TRIAL_STATUS_PAYMENT_NOTE) {
      return (address(swapWitness), 2 days);
    } else if(trialStatus == TRIAL_STATUS_ACK) {
      return (address(ackWitness), 2 days);
    } else if(trialStatus == TRIAL_STATUS_UPDATE) {
      return (address(updateWitness), 2 days);
    } else if(trialStatus == TRIAL_STATUS_INCLUSION) {
      return (address(merkleWitness), 2 days);
    }

    revert("no witness");
  }

  function setRoles(bytes32 caseId, AbstractWitness.TestimonyStatus witnessStatus, uint8 status, bytes memory roles)
  public only_swindle {
    require(witnessStatus == AbstractWitness.TestimonyStatus.VALID);
    if(status == TRIAL_STATUS_SWEAR_NOTE) {
      bytes[] memory swapRoles = abi.decode(roles, (bytes[]));
      data[caseId].serviceNoteHash = abi.decode(swapRoles[9], (bytes32));
    } else if(status == TRIAL_STATUS_PAYMENT_NOTE) {      
    } else if(status == TRIAL_STATUS_ACK) {
      bytes[] memory ackRoles = abi.decode(roles, (bytes[]));
      data[caseId].dataHash = abi.decode(ackRoles[1], (bytes32));
      data[caseId].dataTime = abi.decode(ackRoles[0], (uint));
    } else if(status == TRIAL_STATUS_UPDATE) {
      bytes[] memory decodedRoles = abi.decode(roles, (bytes[]));
      data[caseId].updateRoot = abi.decode(decodedRoles[0], (bytes32));
    } else if(status == TRIAL_STATUS_INCLUSION) {      
    } else revert("setRoles");
  }


  /// @notice return witness for a given status
  function getWitnessPayload(uint8 trialStatus, bytes32 caseId, address payable provider, address payable plaintiff)
  public view returns (bytes memory specification) {    
    if(trialStatus == TRIAL_STATUS_SWEAR_NOTE) {
      uint payment = data[caseId].payment;
      bytes32 expectedRemark = getSwearNoteRemark(payment);
      return abi.encode(
        Note({
          swap: address(0),
          index: 0,
          amount: 0,
          beneficiary: plaintiff,
          witness: address(0),
          validFrom: 0,
          validUntil: 0,
          remark: expectedRemark,
          timeout: 0
        }),
        provider
      );
    }
    
    if(trialStatus == TRIAL_STATUS_PAYMENT_NOTE) {
      return abi.encode(
        Note({
          swap: address(0),
          index: 0,
          amount: data[caseId].payment,
          beneficiary: address(0),
          witness: address(0),
          validFrom: 0,
          validUntil: 0,
          remark: data[caseId].serviceNoteHash,
          timeout: 0
        }),
        plaintiff
      );
    }

    if(trialStatus == TRIAL_STATUS_ACK) {
      return abi.encode(provider, plaintiff);
    }

    if(trialStatus == TRIAL_STATUS_UPDATE) {
      return abi.encode(provider, plaintiff, data[caseId].dataTime);
    }

    if(trialStatus == TRIAL_STATUS_INCLUSION) {
      return abi.encode(data[caseId].updateRoot, data[caseId].dataHash);
    }

    revert("no payload");
  }

  function initialize(bytes32 caseId, bytes memory payload) public only_swindle {
    (uint payment) = abi.decode(payload, (uint));
    data[caseId].payment = payment;
  }

  function encodeInitialPayload(uint payment)
  public view returns (bytes memory) {
    return abi.encode(payment);
  }

  function encodePayloadSwearNote(bytes memory encodedNote, bytes memory sig) 
  public view returns (bytes memory) {
    return abi.encode(encodedNote, sig);
  }
}
