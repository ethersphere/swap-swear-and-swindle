// SPDX-License-Identifier: BSD-3-Clause
pragma solidity =0.7.6;
import "@openzeppelin/contracts/presets/ERC20PresetMinterPauser.sol";

// Used only for testing, token is deployed with storage incentives repo for testnets
// and for mainet its part of special bonding curve mechanisam
contract TestToken is ERC20PresetMinterPauser {
    constructor() ERC20PresetMinterPauser("Test", "TST") {}
}
