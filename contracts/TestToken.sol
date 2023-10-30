// SPDX-License-Identifier: BSD-3-Clause
pragma solidity =0.8.19;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol";

contract TestToken is ERC20PresetMinterPauser {

  constructor() ERC20PresetMinterPauser("Test", "TST") {
  }

}
