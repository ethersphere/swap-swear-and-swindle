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
  
  Swindle public swindle;
  SwapWitness public swapWitness;
  AckWitness public ackWitness;
  UpdateWitness public updateWitness;
  MerkleWitness public merkleWitness;
  
  constructor(Swindle _swindle, UpdateWitness _updateWitness, MerkleWitness _merkleWitness, SwapWitness _swapWitness, AckWitness _ackWitness) public {
    swindle = _swindle;
    swapWitness = _swapWitness;
    ackWitness = _ackWitness;
    updateWitness = _updateWitness;
    merkleWitness = _merkleWitness;
  }

  function getSwearNoteRemark(  
    uint payment
  ) public view returns (bytes32) {
    return keccak256(encodeNoteRemarkData(address(this), payment));
  }

  function getInitialStatus(bytes memory payload)
  public pure returns (uint8, bytes memory) {
    (uint payment) = abi.decode(payload, (uint));
    return (TRIAL_STATUS_SWEAR_NOTE,
      abi.encode(TrialData({
        payment: payment,
        serviceNoteHash: bytes32(0),
        dataHash: bytes32(0),
        dataTime: 0,
        updateRoot: bytes32(0)
      }))
    );
  }
  
  function nextStatus(AbstractWitness.TestimonyStatus witnessStatus, uint8 trialStatus)
  public pure returns (uint8 status) {
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

  function updateData(AbstractWitness.TestimonyStatus witnessStatus, uint8 status, bytes memory trialData, bytes memory roles)
  public pure returns (bytes memory) {
    TrialData memory data = abi.decode(trialData, (TrialData));

    require(witnessStatus == AbstractWitness.TestimonyStatus.VALID);
    if(status == TRIAL_STATUS_SWEAR_NOTE) {
      bytes[] memory swapRoles = abi.decode(roles, (bytes[]));
      data.serviceNoteHash = abi.decode(swapRoles[9], (bytes32));
    } else if(status == TRIAL_STATUS_PAYMENT_NOTE) {      
    } else if(status == TRIAL_STATUS_ACK) {
      (uint dataTime, bytes32 dataHash) = abi.decode(roles, (uint, bytes32));
      data.dataHash = dataHash;
      data.dataTime = dataTime;
    } else if(status == TRIAL_STATUS_UPDATE) {
      data.updateRoot = abi.decode(roles, (bytes32));
    } else if(status == TRIAL_STATUS_INCLUSION) {      
    } else revert("setRoles");

    return abi.encode(data);
  }


  /// @notice return witness for a given status
  function getWitnessPayload(uint8 trialStatus, address payable provider, address payable plaintiff, bytes memory trialData)
  public view returns (bytes memory specification) {
    TrialData memory data = abi.decode(trialData, (TrialData));

    if(trialStatus == TRIAL_STATUS_SWEAR_NOTE) {
      bytes32 expectedRemark = getSwearNoteRemark(data.payment);
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
          amount: data.payment,
          beneficiary: address(0),
          witness: address(0),
          validFrom: 0,
          validUntil: 0,
          remark: data.serviceNoteHash,
          timeout: 0
        }),
        plaintiff
      );
    }

    if(trialStatus == TRIAL_STATUS_ACK) {
      return abi.encode(provider, plaintiff);
    }

    if(trialStatus == TRIAL_STATUS_UPDATE) {
      return abi.encode(provider, plaintiff, data.dataTime);
    }

    if(trialStatus == TRIAL_STATUS_INCLUSION) {
      return abi.encode(data.updateRoot, data.dataHash);
    }

    revert("no payload");
  }

  function encodeInitialPayload(uint payment)
  public view returns (bytes memory) {
    return abi.encode(payment);
  }

  function encodePayloadSwearNote(bytes memory encodedNote, bytes memory sig) 
  public view returns (bytes memory) {
    return abi.encode(encodedNote, sig);
  }

  function encodeNoteRemarkData(address trial, uint payment) 
  public pure returns (bytes memory) {
    return abi.encode(trial, abi.encode(payment));
  }
}
