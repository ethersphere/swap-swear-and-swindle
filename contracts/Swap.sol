pragma solidity ^0.4.23;
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/math/Math.sol";
import "./SW3Utils.sol";
import "./abstracts/AbstractWitness.sol";
import "./SimpleSwap.sol";

/// @title Swap Channel Contract
contract Swap is SimpleSwap {
  /* structure to keep track of a note */
  /* most of this probably does not need to be stored, could be resubmitted on payout to save gas */
  struct NoteInfo {
    uint paidOut; /* total amount paid out */
    uint timeout; /* timeout after which payout can happen */
  }

  /* associates every noteId with a NoteInfo */
  mapping (bytes32 => NoteInfo) public notes;

  constructor(address _owner) SimpleSwap(_owner) public { }

  /// @dev verify the conditions of a note
  function verifyNote(Note memory note) internal view {
    /* if there is validFrom make sure it's in the past */
    if(note.validFrom != 0) require(now >= note.validFrom);
    /* if there is validUntil make sure it's in the future */
    if(note.validUntil != 0) require(now <= note.validUntil);

    /* if there is a witness check the escrow condition */
    if(note.witness != address(0x0)) {
      /* TODO: should be STATIC_CALL, will be so automatically in Solidity 0.5 */
      require(AbstractWitness(note.witness).testimonyFor(owner, note.beneficiary, note.id) == AbstractWitness.TestimonyStatus.VALID);
    }
  }

  /// @notice submit a note
  /// @param sig signature of the note
  function submitNote(bytes encoded, bytes sig) public {
    Note memory note = decodeNote(encoded);

    /* verify the signature of the owner */
    require(owner == recoverSignature(note.id, sig));
    /* make sure the note has not been submitted before */
    require(notes[note.id].timeout == 0);

    notes[note.id] = NoteInfo({
      paidOut: 0,
      timeout: now + timeout
    });

    /* verify that the note conditions hold, else revert everything */
    verifyNote(note);
  }

  /// @notice cash a note
  /// @param amount amount to be paid out
  function cashNote(bytes encoded, uint amount) public {
    Note memory note = decodeNote(encoded);
    NoteInfo storage noteInfo = notes[note.id];

    /* check the note has been submitted */
    require(noteInfo.timeout != 0);
    /* check that the security delay is over */
    require(now >= noteInfo.timeout);
    /* only the beneficiary of the note may call this */
    require(msg.sender == note.beneficiary);
    /* verify that the note conditions hold, WARNING: re-entrance possible until Solidity 0.5 */
    verifyNote(note);

    /* if there is a limit make sure we don't exceed it */
    if(note.amount != 0) {
      require(noteInfo.paidOut.add(amount) <= note.amount);
    }

    /* actual payout */
    (uint payout,) = _payout(note.beneficiary, amount);
    /* increase the stored paidOut amount to avoid double payout */
    noteInfo.paidOut += payout;
  }

  /// @notice demonstrate that an invoice was paid
  /// @param swapBalance swapBalance in the invoice
  /// @param serial serial in the invoice
  /// @param invoiceSig beneficiary signature of the invoice
  /// @param amount of the cheque / note
  /// @param chequeSig owner signature of the cheque
  function submitPaidInvoice(bytes encoded, uint swapBalance, uint serial, bytes invoiceSig, uint amount, bytes chequeSig) public {
    /* only the owner may do this */
    require(msg.sender == owner);
    Note memory note = decodeNote(encoded);
    bytes32 invoiceId = invoiceHash(note.id, swapBalance, serial);

    NoteInfo storage noteInfo = notes[note.id];
    /* ensure the note has been submitted */
    require(noteInfo.timeout != 0);
    /* ensure the security delay is not yet over */
    require(noteInfo.timeout > now);

    /* the expected amount in the cheque is the old swapBalance plus the note amount  */
    uint cumulativeTotal = swapBalance.add(amount);

    /* TODO: this breaks with note.beneficiary = 0 */
    /* check signature of the invoice */
    require(note.beneficiary == recoverSignature(invoiceId, invoiceSig));
    /* check signature of the cheque */
    require(owner == recoverSignature(chequeHash(address(this), note.beneficiary, serial + 1, cumulativeTotal), chequeSig));

    /* cheque needs to be an exact match */
    require(note.amount == amount);
    /* TODO: this breaks with note.amount = 0 */
    /* set paidOut to amount to prevent further payout */
    noteInfo.paidOut = amount;

    /* process the cheque if it is newer than the previous one */
    if(serial > cheques[note.beneficiary].serial)
      _submitChequeInternal(note.beneficiary, serial + 1, cumulativeTotal);
  }

}
