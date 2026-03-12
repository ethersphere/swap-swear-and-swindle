// SPDX-License-Identifier: BSD-3-Clause
pragma solidity =0.7.6;
pragma abicoder v2;

import "../ERC20SimpleSwap.sol";
import "../TestToken.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";

contract ERC20SimpleSwapSignatureEchidna {
    using SafeMath for uint256;

    uint256 private constant VALID_CUMULATIVE_PAYOUT = 100;
    uint256 private constant VALID_CALLER_PAYOUT = 7;
    address private constant EXPECTED_HARNESS =
        0x00a329c0648769A73afAc7F9381E08FB43dBEA72;
    address private constant EXPECTED_SWAP =
        0x62d69f6867A0A084C6d313943dC22023Bc263691;
    address private constant ISSUER =
        0x7E71bA1aB8AF3454a01CFafe358BEbb7691d02f8;
    address private constant BENEFICIARY =
        0xFCA295bC36F47A3Eb53F657b88f3f324374656C6;
    address private constant RECIPIENT =
        0xB5963cAcF590909407433024cD3BA0319542E99D;

    struct SwapSnapshot {
        uint256 balanceBefore;
        uint256 totalHardBefore;
        uint256 totalPaidBefore;
        uint256 paidOutBefore;
        uint256 callerBalanceBefore;
        uint256 recipientBalanceBefore;
        uint256 issuerBalanceBefore;
        uint256 beneficiaryBalanceBefore;
        bool bouncedBefore;
    }

    TestToken private token;
    ERC20SimpleSwap private swap;
    bool private invariantFailed;

    constructor() {
        token = new TestToken();
        swap = new ERC20SimpleSwap();
        swap.init(ISSUER, address(token), 1);
    }

    function cashChequeWithValidSignatures() public {
        if (!_signatureFixtureReady()) {
            invariantFailed = true;
            return;
        }

        if (swap.paidOut(BENEFICIARY) != 0) {
            return;
        }

        _ensureSwapBalance(VALID_CUMULATIVE_PAYOUT);
        SwapSnapshot memory snapshot = _snapshot();

        try
            swap.cashCheque(
                BENEFICIARY,
                RECIPIENT,
                VALID_CUMULATIVE_PAYOUT,
                _beneficiarySig(),
                VALID_CALLER_PAYOUT,
                _issuerSig()
            )
        {
            _assertSuccessfulCash(snapshot);
        } catch {
            invariantFailed = true;
        }
    }

    function cashChequeWithInvalidIssuerSignature() public {
        _expectCashChequeRevert(
            BENEFICIARY,
            RECIPIENT,
            VALID_CUMULATIVE_PAYOUT,
            _beneficiarySig(),
            VALID_CALLER_PAYOUT,
            _mutatedSignature(_issuerSig())
        );
    }

    function cashChequeWithInvalidBeneficiarySignature() public {
        _expectCashChequeRevert(
            BENEFICIARY,
            RECIPIENT,
            VALID_CUMULATIVE_PAYOUT,
            _mutatedSignature(_beneficiarySig()),
            VALID_CALLER_PAYOUT,
            _issuerSig()
        );
    }

    function cashChequeWithMismatchedRecipient() public {
        _expectCashChequeRevert(
            BENEFICIARY,
            ISSUER,
            VALID_CUMULATIVE_PAYOUT,
            _beneficiarySig(),
            VALID_CALLER_PAYOUT,
            _issuerSig()
        );
    }

    function cashChequeWithMismatchedCallerPayout() public {
        _expectCashChequeRevert(
            BENEFICIARY,
            RECIPIENT,
            VALID_CUMULATIVE_PAYOUT,
            _beneficiarySig(),
            VALID_CALLER_PAYOUT.add(1),
            _issuerSig()
        );
    }

    function echidna_signature_fixture_ready() public view returns (bool) {
        return _signatureFixtureReady();
    }

    function echidna_swap_is_initialized() public view returns (bool) {
        return
            swap.issuer() == ISSUER &&
            address(swap.token()) == address(token) &&
            swap.defaultHardDepositTimeout() == 1;
    }

    function echidna_no_postcondition_failures() public view returns (bool) {
        return !invariantFailed;
    }

    function _expectCashChequeRevert(
        address beneficiary,
        address recipient,
        uint256 cumulativePayout,
        bytes memory beneficiarySig,
        uint256 callerPayout,
        bytes memory issuerSig
    ) internal {
        if (!_signatureFixtureReady()) {
            invariantFailed = true;
            return;
        }

        _ensureSwapBalance(VALID_CUMULATIVE_PAYOUT);
        SwapSnapshot memory snapshot = _snapshot();

        try
            swap.cashCheque(
                beneficiary,
                recipient,
                cumulativePayout,
                beneficiarySig,
                callerPayout,
                issuerSig
            )
        {
            invariantFailed = true;
        } catch {
            _assertStateUnchanged(snapshot);
        }
    }

    function _snapshot() internal view returns (SwapSnapshot memory snapshot) {
        snapshot.balanceBefore = swap.balance();
        snapshot.totalHardBefore = swap.totalHardDeposit();
        snapshot.totalPaidBefore = swap.totalPaidOut();
        snapshot.paidOutBefore = swap.paidOut(BENEFICIARY);
        snapshot.callerBalanceBefore = token.balanceOf(address(this));
        snapshot.recipientBalanceBefore = token.balanceOf(RECIPIENT);
        snapshot.issuerBalanceBefore = token.balanceOf(ISSUER);
        snapshot.beneficiaryBalanceBefore = token.balanceOf(BENEFICIARY);
        snapshot.bouncedBefore = swap.bounced();
    }

    function _assertSuccessfulCash(SwapSnapshot memory snapshot) internal {
        if (
            swap.paidOut(BENEFICIARY) !=
            snapshot.paidOutBefore.add(VALID_CUMULATIVE_PAYOUT)
        ) {
            invariantFailed = true;
        }
        if (
            swap.totalPaidOut() !=
            snapshot.totalPaidBefore.add(VALID_CUMULATIVE_PAYOUT)
        ) {
            invariantFailed = true;
        }
        if (
            swap.balance() !=
            snapshot.balanceBefore.sub(VALID_CUMULATIVE_PAYOUT)
        ) {
            invariantFailed = true;
        }
        if (swap.totalHardDeposit() != snapshot.totalHardBefore) {
            invariantFailed = true;
        }
        if (
            token.balanceOf(address(this)) !=
            snapshot.callerBalanceBefore.add(VALID_CALLER_PAYOUT)
        ) {
            invariantFailed = true;
        }
        if (
            token.balanceOf(RECIPIENT) !=
            snapshot.recipientBalanceBefore.add(
                VALID_CUMULATIVE_PAYOUT.sub(VALID_CALLER_PAYOUT)
            )
        ) {
            invariantFailed = true;
        }
        if (token.balanceOf(ISSUER) != snapshot.issuerBalanceBefore) {
            invariantFailed = true;
        }
        if (
            token.balanceOf(BENEFICIARY) != snapshot.beneficiaryBalanceBefore
        ) {
            invariantFailed = true;
        }
        if (swap.bounced() != snapshot.bouncedBefore) {
            invariantFailed = true;
        }
    }

    function _assertStateUnchanged(SwapSnapshot memory snapshot) internal {
        if (swap.balance() != snapshot.balanceBefore) {
            invariantFailed = true;
        }
        if (swap.totalHardDeposit() != snapshot.totalHardBefore) {
            invariantFailed = true;
        }
        if (swap.totalPaidOut() != snapshot.totalPaidBefore) {
            invariantFailed = true;
        }
        if (swap.paidOut(BENEFICIARY) != snapshot.paidOutBefore) {
            invariantFailed = true;
        }
        if (token.balanceOf(address(this)) != snapshot.callerBalanceBefore) {
            invariantFailed = true;
        }
        if (token.balanceOf(RECIPIENT) != snapshot.recipientBalanceBefore) {
            invariantFailed = true;
        }
        if (token.balanceOf(ISSUER) != snapshot.issuerBalanceBefore) {
            invariantFailed = true;
        }
        if (
            token.balanceOf(BENEFICIARY) != snapshot.beneficiaryBalanceBefore
        ) {
            invariantFailed = true;
        }
        if (swap.bounced() != snapshot.bouncedBefore) {
            invariantFailed = true;
        }
    }

    function _ensureSwapBalance(uint256 amount) internal {
        uint256 currentBalance = swap.balance();
        if (currentBalance < amount) {
            token.mint(address(swap), amount.sub(currentBalance));
        }
    }

    function _signatureFixtureReady() internal view returns (bool) {
        return
            address(this) == EXPECTED_HARNESS &&
            address(swap) == EXPECTED_SWAP &&
            _supportedChainId(_chainId());
    }

    function _supportedChainId(uint256 chainId) internal pure returns (bool) {
        return chainId == 1 || chainId == 12345 || chainId == 31337;
    }

    function _chainId() internal pure returns (uint256 id) {
        assembly {
            id := chainid()
        }
    }

    function _issuerSig() internal pure returns (bytes memory) {
        uint256 chainId = _chainId();

        if (chainId == 1) {
            return
                hex"b304812ead72cc612d551dc8986d82c8b49c1ed7ca487d5c2464b8a5371f0f5508436138bfeb19db19294d367215838a8c2a8210ec22f39f19ebf545b63d0c061b";
        }
        if (chainId == 12345) {
            return
                hex"c429ba5cb9b5d52f58b704553f06083ed5ee72d459535a4ec947d414d6fb55137284b146aecef4198b1211c9dd4e648597a93e1dc05cf68486d63441caacabb71b";
        }
        if (chainId == 31337) {
            return
                hex"b0c2ccc2b8c02738f19a5533a534a02f65b86a26349fe42cb76e9f8f19f6f24d729193634b9d7dea386308a3a9aaabac2032faefd4a430fb13d5215a7a97c25d1c";
        }

        revert("unsupported chain id");
    }

    function _beneficiarySig() internal pure returns (bytes memory) {
        uint256 chainId = _chainId();

        if (chainId == 1) {
            return
                hex"4ca1edf6fa62e8d154a1db1bd6e675f1288eec8f946ac9e6eb60002ac39ffc4430973c06732e5714ddb09781d51436e1c93c1961593986c7de9a5595fcf8d4401c";
        }
        if (chainId == 12345) {
            return
                hex"18ee93c890bd3dcf77d61379965ecaad42525c88d1c309c41a9682e13b39bc3807ff14847fe58d6e54eeb4e31a514506c85db43609430051fd0489dc1a92c9b21b";
        }
        if (chainId == 31337) {
            return
                hex"256082480d67e83ab614c15b7aeae47b7a05e6692133b1c966c9e552a24dbb2c677c707c935364be0d53d6c100e781a0d17808b5a67e0306e1a20175f83c92891b";
        }

        revert("unsupported chain id");
    }

    function _mutatedSignature(bytes memory signature)
        internal
        pure
        returns (bytes memory mutated)
    {
        mutated = abi.encodePacked(signature);
        mutated[mutated.length - 1] = bytes1(
            uint8(mutated[mutated.length - 1]) ^ uint8(1)
        );
    }
}
