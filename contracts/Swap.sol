pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/math/Math.sol";
import "./SW3Utils.sol";
import "./abstracts/AbstractWitness.sol";
import "./SimpleSwap.sol";

/// @title Swap Channel Contract
contract Swap is SimpleSwap {
  event NoteSubmitted(bytes32 indexed noteId);
  event NoteCashed(bytes32 indexed noteId, uint amount);
  event NoteBounced(bytes32 indexed noteId, uint paid, uint bounced);

  /* structure to keep track of a note */
  /* most of this probably does not need to be stored, could be resubmitted on payout to save gas */
  struct NoteInfo {
    uint paidOut; /* total amount paid out */
    uint timeout; /* timeout after which payout can happen */
  }

  /* associates every noteId with a NoteInfo */
  mapping (bytes32 => NoteInfo) public notes;

  constructor(address payable _owner) SimpleSwap(_owner) public { }

  /// @dev verify the conditions of a note
  function verifyNote(bytes32 id, Note memory note) internal view {
    /* if there is validFrom make sure it's in the past */
    if(note.validFrom != 0) require(now >= note.validFrom);
    /* if there is validUntil make sure it's in the future */
    if(note.validUntil != 0) require(now <= note.validUntil);

    /* if there is a witness check the escrow condition */
    if(note.witness != address(0x0)) {
      /* static call */
      require(AbstractWitness(note.witness).testimonyFor(owner, note.beneficiary, id) == AbstractWitness.TestimonyStatus.VALID);
    }
  }

  /// @notice submit a note
  /// @param sig signature of the note
  function submitNote(bytes memory encoded, bytes memory sig) public {
    Note memory note = abi.decode(encoded, (Note));
    bytes32 id = keccak256(encoded);

    /* verify the signature of the owner */
    require(owner == recover(id, sig));
    /* make sure the note has not been submitted before */
    require(notes[id].timeout == 0);

    notes[id] = NoteInfo({
      paidOut: 0,
      timeout: now + timeout
    });

    /* verify that the note conditions hold, else revert everything */
    verifyNote(id, note);

    emit NoteSubmitted(id);
  }

  /// @notice cash a note
  /// @param amount amount to be paid out
  function cashNote(bytes memory encoded, uint amount) public {
    Note memory note = abi.decode(encoded, (Note));
    bytes32 id = keccak256(encoded);
    NoteInfo storage noteInfo = notes[id];

    /* check the note has been submitted */
    require(noteInfo.timeout != 0);
    /* check that the security delay is over */
    require(now >= noteInfo.timeout);
    /* only the beneficiary of the note may call this */
    require(msg.sender == note.beneficiary); // necessary because of blank cheques
    /* verify that the note conditions hold, static call */
    verifyNote(id, note);

    /* if there is a limit make sure we don't exceed it */
    if(note.amount != 0) {
      require(noteInfo.paidOut.add(amount) <= note.amount);
    }

    /* compute the actual payout */
    (uint payout, uint bounced) = _computePayout(note.beneficiary, amount);

    /* TODO: event */

    /* increase the stored paidOut amount to avoid double payout */
    noteInfo.paidOut += payout;

    if(bounced != 0) emit NoteBounced(id, payout, bounced);
    else emit NoteCashed(id, payout);

    /* do the payout */
    note.beneficiary.transfer(payout); // TODO: test
  }

  /// @notice demonstrate that an invoice was paid
  /// @param swapBalance swapBalance in the invoice
  /// @param serial serial in the invoice
  /// @param invoiceSig beneficiary signature of the invoice
  /// @param amount of the cheque / note
  /// @param chequeSig owner signature of the cheque
  function submitPaidInvoice(bytes memory encoded, uint swapBalance, uint serial, bytes memory invoiceSig, uint amount, uint timeout, bytes memory chequeSig) public {
    /* only the owner may do this */
    require(msg.sender == owner);
    Note memory note = abi.decode(encoded, (Note));
    bytes32 id = keccak256(encoded);
    bytes32 invoiceId = invoiceHash(id, swapBalance, serial);

    NoteInfo storage noteInfo = notes[id];
    /* ensure the note has been submitted */
    require(noteInfo.timeout != 0);
    /* ensure the security delay is not yet over */
    require(noteInfo.timeout > now);

    /* the expected amount in the cheque is the old swapBalance plus the note amount  */
    uint cumulativeTotal = swapBalance.add(amount);

    /* TODO: this breaks with note.beneficiary = 0 */
    /* check signature of the invoice */
    require(note.beneficiary == recover(invoiceId, invoiceSig));
    /* check signature of the cheque */
    require(owner == recover(chequeHash(address(this), note.beneficiary, serial + 1, cumulativeTotal, timeout), chequeSig));

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
