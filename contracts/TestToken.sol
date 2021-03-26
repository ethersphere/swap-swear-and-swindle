// SPDX-License-Identifier: BSD-3-Clause
pragma solidity ^0.6.12;

import "@openzeppelin/contracts/presets/ERC20PresetMinterPauser.sol";

contract TestToken is ERC20PresetMinterPauser {

  constructor() ERC20PresetMinterPauser("Test", "TST") public {    
  }

}
