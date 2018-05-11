pragma solidity ^0.4.0;

/// @title AbstractConstants
contract AbstractConstants {
  /* constants for the special trial states */
  uint8 constant public TRIAL_STATUS_UNCHALLENGED = 0;
  uint8 constant public TRIAL_STATUS_GUILTY = 1;
  uint8 constant public TRIAL_STATUS_NOT_GUILTY = 2;
}
