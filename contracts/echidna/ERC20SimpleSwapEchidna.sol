// SPDX-License-Identifier: BSD-3-Clause
pragma solidity =0.7.6;
pragma abicoder v2;

import "../ERC20SimpleSwap.sol";
import "../TestToken.sol";

interface IERC20SimpleSwapEchidnaTarget {
    function increaseHardDeposit(address beneficiary, uint256 amount) external;

    function prepareDecreaseHardDeposit(
        address beneficiary,
        uint256 decreaseAmount
    ) external;

    function decreaseHardDeposit(address beneficiary) external;

    function withdraw(uint256 amount) external;

    function harnessSetCustomHardDepositTimeout(
        address beneficiary,
        uint256 timeout
    ) external;

    function harnessCashCheque(
        address beneficiary,
        address recipient,
        uint256 cumulativePayout,
        uint256 callerPayout
    ) external;

    function harnessCashChequeBeneficiary(
        address recipient,
        uint256 cumulativePayout
    ) external;
}

contract SwapActor {
    function depositToken(
        TestToken token,
        address recipient,
        uint256 amount
    ) external returns (bool) {
        try token.transfer(recipient, amount) returns (bool ok) {
            return ok;
        } catch {
            return false;
        }
    }

    function increaseHardDeposit(
        IERC20SimpleSwapEchidnaTarget target,
        address beneficiary,
        uint256 amount
    ) external returns (bool) {
        try target.increaseHardDeposit(beneficiary, amount) {
            return true;
        } catch {
            return false;
        }
    }

    function prepareDecreaseHardDeposit(
        IERC20SimpleSwapEchidnaTarget target,
        address beneficiary,
        uint256 decreaseAmount
    ) external returns (bool) {
        try target.prepareDecreaseHardDeposit(beneficiary, decreaseAmount) {
            return true;
        } catch {
            return false;
        }
    }

    function decreaseHardDeposit(
        IERC20SimpleSwapEchidnaTarget target,
        address beneficiary
    ) external returns (bool) {
        try target.decreaseHardDeposit(beneficiary) {
            return true;
        } catch {
            return false;
        }
    }

    function withdraw(
        IERC20SimpleSwapEchidnaTarget target,
        uint256 amount
    ) external returns (bool) {
        try target.withdraw(amount) {
            return true;
        } catch {
            return false;
        }
    }

    function setCustomHardDepositTimeout(
        IERC20SimpleSwapEchidnaTarget target,
        address beneficiary,
        uint256 timeout
    ) external returns (bool) {
        try target.harnessSetCustomHardDepositTimeout(beneficiary, timeout) {
            return true;
        } catch {
            return false;
        }
    }

    function cashCheque(
        IERC20SimpleSwapEchidnaTarget target,
        address beneficiary,
        address recipient,
        uint256 cumulativePayout,
        uint256 callerPayout
    ) external returns (bool) {
        try
            target.harnessCashCheque(
                beneficiary,
                recipient,
                cumulativePayout,
                callerPayout
            )
        {
            return true;
        } catch {
            return false;
        }
    }

    function cashChequeBeneficiary(
        IERC20SimpleSwapEchidnaTarget target,
        address recipient,
        uint256 cumulativePayout
    ) external returns (bool) {
        try target.harnessCashChequeBeneficiary(recipient, cumulativePayout) {
            return true;
        } catch {
            return false;
        }
    }
}

contract ERC20SimpleSwapEchidnaTarget is ERC20SimpleSwap {
    using SafeMath for uint256;

    address private immutable harness;
    mapping(address => bool) private fuzzActors;

    constructor(
        address issuerActor,
        address tokenAddress,
        uint256 defaultHardDepositTimeout
    ) {
        harness = msg.sender;
        init(issuerActor, tokenAddress, defaultHardDepositTimeout);
    }

    modifier onlyHarness() {
        require(msg.sender == harness, "only harness");
        _;
    }

    modifier onlyFuzzActor() {
        require(fuzzActors[msg.sender], "only actor");
        _;
    }

    function authorizeFuzzActor(address actor) external onlyHarness {
        fuzzActors[actor] = true;
    }

    function harnessSetCustomHardDepositTimeout(
        address beneficiary,
        uint256 hardDepositTimeout
    ) external {
        require(msg.sender == issuer, "not issuer");
        hardDeposits[beneficiary].timeout = hardDepositTimeout;
        emit HardDepositTimeoutChanged(beneficiary, hardDepositTimeout);
    }

    function harnessCashCheque(
        address beneficiary,
        address recipient,
        uint256 cumulativePayout,
        uint256 callerPayout
    ) external onlyFuzzActor {
        _cashChequeWithoutSignatures(
            beneficiary,
            recipient,
            cumulativePayout,
            callerPayout
        );
    }

    function harnessCashChequeBeneficiary(
        address recipient,
        uint256 cumulativePayout
    ) external onlyFuzzActor {
        _cashChequeWithoutSignatures(msg.sender, recipient, cumulativePayout, 0);
    }

    function _cashChequeWithoutSignatures(
        address beneficiary,
        address recipient,
        uint256 cumulativePayout,
        uint256 callerPayout
    ) internal {
        uint256 requestPayout = cumulativePayout.sub(paidOut[beneficiary]);
        uint256 totalPayout = Math.min(requestPayout, liquidBalanceFor(beneficiary));
        uint256 hardDepositUsage = Math.min(
            totalPayout,
            hardDeposits[beneficiary].amount
        );

        require(totalPayout >= callerPayout, "SimpleSwap: cannot pay caller");

        if (hardDepositUsage != 0) {
            hardDeposits[beneficiary].amount = hardDeposits[beneficiary]
                .amount
                .sub(hardDepositUsage);
            totalHardDeposit = totalHardDeposit.sub(hardDepositUsage);
        }

        paidOut[beneficiary] = paidOut[beneficiary].add(totalPayout);
        totalPaidOut = totalPaidOut.add(totalPayout);

        if (requestPayout != totalPayout) {
            bounced = true;
            emit ChequeBounced();
        }

        if (callerPayout != 0) {
            require(token.transfer(msg.sender, callerPayout), "transfer failed");
            require(
                token.transfer(recipient, totalPayout.sub(callerPayout)),
                "transfer failed"
            );
        } else {
            require(token.transfer(recipient, totalPayout), "transfer failed");
        }

        emit ChequeCashed(
            beneficiary,
            recipient,
            msg.sender,
            totalPayout,
            cumulativePayout,
            callerPayout
        );
    }
}

contract ERC20SimpleSwapEchidna {
    using SafeMath for uint256;

    uint256 private constant ACTOR_COUNT = 4;
    uint256 private constant INITIAL_TOKEN_BALANCE = 1_000_000;
    uint256 private constant MAX_TIMEOUT = 7;

    struct CashSnapshot {
        uint256 totalHardBefore;
        uint256 totalPaidBefore;
        uint256 swapBalanceBefore;
        uint256 callerBalanceBefore;
        uint256 recipientBalanceBefore;
        uint256 cumulativePayout;
        uint256 requestPayout;
        uint256 totalPayout;
        uint256 hardDepositUsage;
        uint256 callerPayout;
        bool bouncedBefore;
    }

    struct FullStateSnapshot {
        uint256 swapBalance;
        uint256 totalHardDeposit;
        uint256 totalPaidOut;
        bool bounced;
        uint256[4] hardAmounts;
        uint256[4] decreaseAmounts;
        uint256[4] timeouts;
        uint256[4] canDecreaseAt;
        uint256[4] paidOut;
        uint256[4] actorBalances;
    }

    TestToken private token;
    ERC20SimpleSwapEchidnaTarget private simpleSwap;
    SwapActor[4] private actors;
    address[4] private actorAddresses;
    bool private invariantFailed;
    uint256[4] private lastPaidOut;
    uint256 private lastTotalPaidOut;
    bool private bounceSeen;

    constructor() {
        token = new TestToken();

        for (uint256 i = 0; i < ACTOR_COUNT; i++) {
            actors[i] = new SwapActor();
            actorAddresses[i] = address(actors[i]);
            token.mint(actorAddresses[i], INITIAL_TOKEN_BALANCE);
        }

        simpleSwap = new ERC20SimpleSwapEchidnaTarget(
            actorAddresses[0],
            address(token),
            0
        );

        for (uint256 i = 0; i < ACTOR_COUNT; i++) {
            simpleSwap.authorizeFuzzActor(actorAddresses[i]);
        }

        _syncTemporalState();
    }

    function deposit(uint256 actorSeed, uint256 rawAmount) public {
        uint256 actorIndex = _actorIndex(actorSeed);
        uint256 amount = _bounded(
            rawAmount,
            token.balanceOf(actorAddresses[actorIndex])
        );
        uint256[4] memory amountsBefore;
        uint256[4] memory decreaseAmountsBefore;
        uint256[4] memory timeoutsBefore;
        uint256[4] memory canDecreaseBefore;

        _snapshotBeneficiaries(
            amountsBefore,
            decreaseAmountsBefore,
            timeoutsBefore,
            canDecreaseBefore
        );

        if (amount == 0) {
            return;
        }

        if (
            !actors[actorIndex].depositToken(token, address(simpleSwap), amount)
        ) {
            return;
        }

        _assertUntouched(
            ACTOR_COUNT,
            amountsBefore,
            decreaseAmountsBefore,
            timeoutsBefore,
            canDecreaseBefore
        );

        _checkTemporalInvariants();
    }

    function increaseHardDeposit(uint256 beneficiarySeed, uint256 rawAmount) public {
        uint256 beneficiaryIndex = _actorIndex(beneficiarySeed);
        uint256[4] memory amountsBefore;
        uint256[4] memory decreaseAmountsBefore;
        uint256[4] memory timeoutsBefore;
        uint256[4] memory canDecreaseBefore;

        _snapshotBeneficiaries(
            amountsBefore,
            decreaseAmountsBefore,
            timeoutsBefore,
            canDecreaseBefore
        );

        uint256 balanceBefore = simpleSwap.balance();
        uint256 liquidBefore = simpleSwap.liquidBalance();
        uint256 totalHardBefore = simpleSwap.totalHardDeposit();
        uint256 amount = _bounded(rawAmount, liquidBefore);

        if (amount == 0) {
            return;
        }

        if (
            !actors[0].increaseHardDeposit(
                IERC20SimpleSwapEchidnaTarget(address(simpleSwap)),
                actorAddresses[beneficiaryIndex],
                amount
            )
        ) {
            return;
        }

        _checkIncreaseHardDepositPostconditions(
            beneficiaryIndex,
            amount,
            balanceBefore,
            liquidBefore,
            totalHardBefore,
            amountsBefore,
            decreaseAmountsBefore,
            timeoutsBefore
        );

        _assertUntouched(
            beneficiaryIndex,
            amountsBefore,
            decreaseAmountsBefore,
            timeoutsBefore,
            canDecreaseBefore
        );

        _checkTemporalInvariants();
    }

    function prepareDecreaseHardDeposit(
        uint256 beneficiarySeed,
        uint256 rawAmount
    ) public {
        uint256 beneficiaryIndex = _actorIndex(beneficiarySeed);
        uint256[4] memory amountsBefore;
        uint256[4] memory decreaseAmountsBefore;
        uint256[4] memory timeoutsBefore;
        uint256[4] memory canDecreaseBefore;

        _snapshotBeneficiaries(
            amountsBefore,
            decreaseAmountsBefore,
            timeoutsBefore,
            canDecreaseBefore
        );

        uint256 amount = _bounded(rawAmount, amountsBefore[beneficiaryIndex]);

        if (amount == 0) {
            return;
        }

        if (
            !actors[0].prepareDecreaseHardDeposit(
                IERC20SimpleSwapEchidnaTarget(address(simpleSwap)),
                actorAddresses[beneficiaryIndex],
                amount
            )
        ) {
            return;
        }

        (
            uint256 targetAmountAfter,
            uint256 targetDecreaseAfter,
            uint256 targetTimeoutAfter,
            uint256 targetCanDecreaseAfter
        ) = simpleSwap.hardDeposits(actorAddresses[beneficiaryIndex]);

        if (targetAmountAfter != amountsBefore[beneficiaryIndex]) {
            invariantFailed = true;
        }
        if (targetDecreaseAfter != amount) {
            invariantFailed = true;
        }
        if (targetTimeoutAfter != timeoutsBefore[beneficiaryIndex]) {
            invariantFailed = true;
        }
        if (targetCanDecreaseAfter == 0) {
            invariantFailed = true;
        }
        if (
            targetCanDecreaseAfter !=
            block.timestamp.add(_effectiveTimeout(timeoutsBefore[beneficiaryIndex]))
        ) {
            invariantFailed = true;
        }

        _assertUntouched(
            beneficiaryIndex,
            amountsBefore,
            decreaseAmountsBefore,
            timeoutsBefore,
            canDecreaseBefore
        );

        _checkTemporalInvariants();
    }

    function decreaseHardDeposit(
        uint256 callerSeed,
        uint256 beneficiarySeed
    ) public {
        uint256 callerIndex = _actorIndex(callerSeed);
        uint256 beneficiaryIndex = _actorIndex(beneficiarySeed);
        uint256[4] memory amountsBefore;
        uint256[4] memory decreaseAmountsBefore;
        uint256[4] memory timeoutsBefore;
        uint256[4] memory canDecreaseBefore;

        _snapshotBeneficiaries(
            amountsBefore,
            decreaseAmountsBefore,
            timeoutsBefore,
            canDecreaseBefore
        );

        uint256 totalHardBefore = simpleSwap.totalHardDeposit();

        if (
            !actors[callerIndex].decreaseHardDeposit(
                IERC20SimpleSwapEchidnaTarget(address(simpleSwap)),
                actorAddresses[beneficiaryIndex]
            )
        ) {
            return;
        }

        (
            uint256 targetAmountAfter,
            uint256 targetDecreaseAfter,
            uint256 targetTimeoutAfter,
            uint256 targetCanDecreaseAfter
        ) = simpleSwap.hardDeposits(actorAddresses[beneficiaryIndex]);

        if (
            targetAmountAfter !=
            amountsBefore[beneficiaryIndex].sub(
                decreaseAmountsBefore[beneficiaryIndex]
            )
        ) {
            invariantFailed = true;
        }
        if (targetDecreaseAfter != decreaseAmountsBefore[beneficiaryIndex]) {
            invariantFailed = true;
        }
        if (targetTimeoutAfter != timeoutsBefore[beneficiaryIndex]) {
            invariantFailed = true;
        }
        if (targetCanDecreaseAfter != 0) {
            invariantFailed = true;
        }
        if (
            simpleSwap.totalHardDeposit() !=
            totalHardBefore.sub(decreaseAmountsBefore[beneficiaryIndex])
        ) {
            invariantFailed = true;
        }

        _assertUntouched(
            beneficiaryIndex,
            amountsBefore,
            decreaseAmountsBefore,
            timeoutsBefore,
            canDecreaseBefore
        );

        _checkTemporalInvariants();
    }

    function withdraw(uint256 rawAmount) public {
        uint256[4] memory amountsBefore;
        uint256[4] memory decreaseAmountsBefore;
        uint256[4] memory timeoutsBefore;
        uint256[4] memory canDecreaseBefore;

        _snapshotBeneficiaries(
            amountsBefore,
            decreaseAmountsBefore,
            timeoutsBefore,
            canDecreaseBefore
        );

        uint256 amount = _bounded(rawAmount, simpleSwap.liquidBalance());
        uint256 issuerBalanceBefore = token.balanceOf(actorAddresses[0]);
        uint256 liquidBefore = simpleSwap.liquidBalance();

        if (amount == 0) {
            return;
        }

        if (
            !actors[0].withdraw(
                IERC20SimpleSwapEchidnaTarget(address(simpleSwap)),
                amount
            )
        ) {
            return;
        }

        if (simpleSwap.liquidBalance() != liquidBefore.sub(amount)) {
            invariantFailed = true;
        }
        if (
            token.balanceOf(actorAddresses[0]) != issuerBalanceBefore.add(amount)
        ) {
            invariantFailed = true;
        }

        _assertUntouched(
            ACTOR_COUNT,
            amountsBefore,
            decreaseAmountsBefore,
            timeoutsBefore,
            canDecreaseBefore
        );

        _checkTemporalInvariants();
    }

    function setCustomHardDepositTimeout(
        uint256 beneficiarySeed,
        uint256 rawTimeout
    ) public {
        uint256 beneficiaryIndex = _actorIndex(beneficiarySeed);
        uint256[4] memory amountsBefore;
        uint256[4] memory decreaseAmountsBefore;
        uint256[4] memory timeoutsBefore;
        uint256[4] memory canDecreaseBefore;

        _snapshotBeneficiaries(
            amountsBefore,
            decreaseAmountsBefore,
            timeoutsBefore,
            canDecreaseBefore
        );

        uint256 timeout = rawTimeout % (MAX_TIMEOUT + 1);

        if (
            !actors[0].setCustomHardDepositTimeout(
                IERC20SimpleSwapEchidnaTarget(address(simpleSwap)),
                actorAddresses[beneficiaryIndex],
                timeout
            )
        ) {
            return;
        }

        (
            uint256 targetAmountAfter,
            uint256 targetDecreaseAfter,
            uint256 targetTimeoutAfter,
            uint256 targetCanDecreaseAfter
        ) = simpleSwap.hardDeposits(actorAddresses[beneficiaryIndex]);

        if (targetAmountAfter != amountsBefore[beneficiaryIndex]) {
            invariantFailed = true;
        }
        if (targetDecreaseAfter != decreaseAmountsBefore[beneficiaryIndex]) {
            invariantFailed = true;
        }
        if (targetTimeoutAfter != timeout) {
            invariantFailed = true;
        }
        if (targetCanDecreaseAfter != canDecreaseBefore[beneficiaryIndex]) {
            invariantFailed = true;
        }

        _assertUntouched(
            beneficiaryIndex,
            amountsBefore,
            decreaseAmountsBefore,
            timeoutsBefore,
            canDecreaseBefore
        );

        _checkTemporalInvariants();
    }

    function unauthorizedIncreaseHardDeposit(
        uint256 callerSeed,
        uint256 beneficiarySeed,
        uint256 rawAmount
    ) public {
        uint256 callerIndex = _nonIssuerIndex(callerSeed);
        uint256 beneficiaryIndex = _actorIndex(beneficiarySeed);
        FullStateSnapshot memory snapshot = _snapshotFullState();
        uint256 amount = _bounded(rawAmount, simpleSwap.balance().add(1));

        if (
            actors[callerIndex].increaseHardDeposit(
                IERC20SimpleSwapEchidnaTarget(address(simpleSwap)),
                actorAddresses[beneficiaryIndex],
                amount
            )
        ) {
            invariantFailed = true;
        }

        _assertFullStateUnchanged(snapshot);
        _checkTemporalInvariants();
    }

    function unauthorizedWithdraw(uint256 callerSeed, uint256 rawAmount) public {
        uint256 callerIndex = _nonIssuerIndex(callerSeed);
        FullStateSnapshot memory snapshot = _snapshotFullState();
        uint256 amount = _bounded(rawAmount, simpleSwap.balance().add(1));

        if (
            actors[callerIndex].withdraw(
                IERC20SimpleSwapEchidnaTarget(address(simpleSwap)),
                amount
            )
        ) {
            invariantFailed = true;
        }

        _assertFullStateUnchanged(snapshot);
        _checkTemporalInvariants();
    }

    function prepareDecreaseHardDepositTooLarge(uint256 beneficiarySeed) public {
        uint256 beneficiaryIndex = _actorIndex(beneficiarySeed);
        FullStateSnapshot memory snapshot = _snapshotFullState();
        uint256 amount = snapshot.hardAmounts[beneficiaryIndex].add(1);

        if (
            actors[0].prepareDecreaseHardDeposit(
                IERC20SimpleSwapEchidnaTarget(address(simpleSwap)),
                actorAddresses[beneficiaryIndex],
                amount
            )
        ) {
            invariantFailed = true;
        }

        _assertFullStateUnchanged(snapshot);
        _checkTemporalInvariants();
    }

    function prematureDecreaseHardDeposit(uint256 beneficiarySeed) public {
        uint256 beneficiaryIndex = _actorIndex(beneficiarySeed);
        FullStateSnapshot memory snapshot = _snapshotFullState();

        if (
            snapshot.canDecreaseAt[beneficiaryIndex] != 0 &&
            block.timestamp >= snapshot.canDecreaseAt[beneficiaryIndex]
        ) {
            return;
        }

        if (
            actors[1].decreaseHardDeposit(
                IERC20SimpleSwapEchidnaTarget(address(simpleSwap)),
                actorAddresses[beneficiaryIndex]
            )
        ) {
            invariantFailed = true;
        }

        _assertFullStateUnchanged(snapshot);
        _checkTemporalInvariants();
    }

    function cashChequeCannotPayCaller(
        uint256 callerSeed,
        uint256 beneficiarySeed,
        uint256 recipientSeed,
        uint256 cumulativeDelta
    ) public {
        uint256 callerIndex = _actorIndex(callerSeed);
        uint256 beneficiaryIndex = _actorIndex(beneficiarySeed);
        uint256 recipientIndex = _actorIndex(recipientSeed);
        FullStateSnapshot memory snapshot = _snapshotFullState();
        uint256 cumulativePayout = snapshot.paidOut[beneficiaryIndex].add(
            _bounded(cumulativeDelta, token.totalSupply())
        );
        (, uint256 totalPayout, ) = _expectedCashEffects(
            actorAddresses[beneficiaryIndex],
            cumulativePayout
        );
        uint256 callerPayout = totalPayout.add(1);

        if (
            actors[callerIndex].cashCheque(
                IERC20SimpleSwapEchidnaTarget(address(simpleSwap)),
                actorAddresses[beneficiaryIndex],
                actorAddresses[recipientIndex],
                cumulativePayout,
                callerPayout
            )
        ) {
            invariantFailed = true;
        }

        _assertFullStateUnchanged(snapshot);
        _checkTemporalInvariants();
    }

    function cashChequeBeneficiary(
        uint256 beneficiarySeed,
        uint256 recipientSeed,
        uint256 cumulativeDelta
    ) public {
        uint256 beneficiaryIndex = _actorIndex(beneficiarySeed);
        uint256 recipientIndex = _actorIndex(recipientSeed);
        uint256[4] memory amountsBefore;
        uint256[4] memory decreaseAmountsBefore;
        uint256[4] memory timeoutsBefore;
        uint256[4] memory canDecreaseBefore;
        uint256[4] memory paidOutBefore;

        _snapshotAccounting(
            amountsBefore,
            decreaseAmountsBefore,
            timeoutsBefore,
            canDecreaseBefore,
            paidOutBefore
        );

        CashSnapshot memory snapshot;
        snapshot.totalHardBefore = simpleSwap.totalHardDeposit();
        snapshot.totalPaidBefore = simpleSwap.totalPaidOut();
        snapshot.swapBalanceBefore = simpleSwap.balance();
        snapshot.recipientBalanceBefore = token.balanceOf(
            actorAddresses[recipientIndex]
        );
        snapshot.bouncedBefore = simpleSwap.bounced();
        snapshot.cumulativePayout = paidOutBefore[beneficiaryIndex].add(
            _bounded(cumulativeDelta, token.totalSupply())
        );
        (
            snapshot.requestPayout,
            snapshot.totalPayout,
            snapshot.hardDepositUsage
        ) = _expectedCashEffects(
            actorAddresses[beneficiaryIndex],
            snapshot.cumulativePayout
        );

        if (
            !actors[beneficiaryIndex].cashChequeBeneficiary(
                IERC20SimpleSwapEchidnaTarget(address(simpleSwap)),
                actorAddresses[recipientIndex],
                snapshot.cumulativePayout
            )
        ) {
            return;
        }

        (
            uint256 targetAmountAfter,
            ,
            ,
            uint256 targetCanDecreaseAfter
        ) = simpleSwap.hardDeposits(actorAddresses[beneficiaryIndex]);

        if (
            simpleSwap.paidOut(actorAddresses[beneficiaryIndex]) !=
            paidOutBefore[beneficiaryIndex].add(snapshot.totalPayout)
        ) {
            invariantFailed = true;
        }
        if (
            simpleSwap.totalPaidOut() !=
            snapshot.totalPaidBefore.add(snapshot.totalPayout)
        ) {
            invariantFailed = true;
        }
        if (
            simpleSwap.totalHardDeposit() !=
            snapshot.totalHardBefore.sub(snapshot.hardDepositUsage)
        ) {
            invariantFailed = true;
        }
        if (
            targetAmountAfter !=
            amountsBefore[beneficiaryIndex].sub(snapshot.hardDepositUsage)
        ) {
            invariantFailed = true;
        }
        if (
            simpleSwap.balance() !=
            snapshot.swapBalanceBefore.sub(snapshot.totalPayout)
        ) {
            invariantFailed = true;
        }
        if (
            token.balanceOf(actorAddresses[recipientIndex]) !=
            snapshot.recipientBalanceBefore.add(snapshot.totalPayout)
        ) {
            invariantFailed = true;
        }
        if (
            snapshot.requestPayout > snapshot.totalPayout && !simpleSwap.bounced()
        ) {
            invariantFailed = true;
        }
        if (
            snapshot.requestPayout == snapshot.totalPayout &&
            snapshot.bouncedBefore != simpleSwap.bounced()
        ) {
            invariantFailed = true;
        }
        if (
            targetCanDecreaseAfter != canDecreaseBefore[beneficiaryIndex]
        ) {
            invariantFailed = true;
        }

        _assertAccountingUntouched(
            beneficiaryIndex,
            amountsBefore,
            decreaseAmountsBefore,
            timeoutsBefore,
            canDecreaseBefore,
            paidOutBefore
        );

        _checkTemporalInvariants();
    }

    function cashCheque(
        uint256 callerSeed,
        uint256 beneficiarySeed,
        uint256 recipientSeed,
        uint256 cumulativeDelta,
        uint256 rawCallerPayout
    ) public {
        uint256 callerIndex = _actorIndex(callerSeed);
        uint256 beneficiaryIndex = _actorIndex(beneficiarySeed);
        uint256 recipientIndex = _actorIndex(recipientSeed);
        uint256[4] memory amountsBefore;
        uint256[4] memory decreaseAmountsBefore;
        uint256[4] memory timeoutsBefore;
        uint256[4] memory canDecreaseBefore;
        uint256[4] memory paidOutBefore;

        _snapshotAccounting(
            amountsBefore,
            decreaseAmountsBefore,
            timeoutsBefore,
            canDecreaseBefore,
            paidOutBefore
        );

        CashSnapshot memory snapshot;
        snapshot.totalHardBefore = simpleSwap.totalHardDeposit();
        snapshot.totalPaidBefore = simpleSwap.totalPaidOut();
        snapshot.swapBalanceBefore = simpleSwap.balance();
        snapshot.callerBalanceBefore = token.balanceOf(actorAddresses[callerIndex]);
        snapshot.recipientBalanceBefore = token.balanceOf(
            actorAddresses[recipientIndex]
        );
        snapshot.bouncedBefore = simpleSwap.bounced();
        snapshot.cumulativePayout = paidOutBefore[beneficiaryIndex].add(
            _bounded(cumulativeDelta, token.totalSupply())
        );
        (
            snapshot.requestPayout,
            snapshot.totalPayout,
            snapshot.hardDepositUsage
        ) = _expectedCashEffects(
            actorAddresses[beneficiaryIndex],
            snapshot.cumulativePayout
        );
        snapshot.callerPayout = _bounded(
            rawCallerPayout,
            snapshot.totalPayout
        );

        if (
            !actors[callerIndex].cashCheque(
                IERC20SimpleSwapEchidnaTarget(address(simpleSwap)),
                actorAddresses[beneficiaryIndex],
                actorAddresses[recipientIndex],
                snapshot.cumulativePayout,
                snapshot.callerPayout
            )
        ) {
            return;
        }

        _checkCashChequePostconditions(
            callerIndex,
            beneficiaryIndex,
            recipientIndex,
            amountsBefore,
            canDecreaseBefore,
            paidOutBefore,
            snapshot
        );

        _assertAccountingUntouched(
            beneficiaryIndex,
            amountsBefore,
            decreaseAmountsBefore,
            timeoutsBefore,
            canDecreaseBefore,
            paidOutBefore
        );

        _checkTemporalInvariants();
    }

    function echidna_balance_conservation() public view returns (bool) {
        return
            _sumActorBalances().add(simpleSwap.balance()) == token.totalSupply();
    }

    function echidna_total_hard_deposit_bounded() public view returns (bool) {
        return simpleSwap.totalHardDeposit() <= simpleSwap.balance();
    }

    function echidna_liquid_balance_matches_balance() public view returns (bool) {
        return
            simpleSwap.liquidBalance().add(simpleSwap.totalHardDeposit()) ==
            simpleSwap.balance();
    }

    function echidna_tracked_hard_deposits_match_total()
        public
        view
        returns (bool)
    {
        uint256 totalTrackedHardDeposits = 0;

        for (uint256 i = 0; i < ACTOR_COUNT; i++) {
            (uint256 amount, , , ) = simpleSwap.hardDeposits(actorAddresses[i]);
            totalTrackedHardDeposits = totalTrackedHardDeposits.add(amount);
        }

        return totalTrackedHardDeposits == simpleSwap.totalHardDeposit();
    }

    function echidna_liquid_balance_for_matches_formula()
        public
        view
        returns (bool)
    {
        uint256 liquid = simpleSwap.liquidBalance();

        for (uint256 i = 0; i < ACTOR_COUNT; i++) {
            (uint256 amount, , , ) = simpleSwap.hardDeposits(actorAddresses[i]);
            if (
                simpleSwap.liquidBalanceFor(actorAddresses[i]) !=
                liquid.add(amount)
            ) {
                return false;
            }
        }

        return true;
    }

    function echidna_bounced_is_sticky() public view returns (bool) {
        return !bounceSeen || simpleSwap.bounced();
    }

    function echidna_no_postcondition_failures() public view returns (bool) {
        return !invariantFailed;
    }

    function _snapshotBeneficiaries(
        uint256[4] memory amounts,
        uint256[4] memory decreaseAmounts,
        uint256[4] memory timeouts,
        uint256[4] memory canDecreaseAt
    ) internal view {
        for (uint256 i = 0; i < ACTOR_COUNT; i++) {
            (
                amounts[i],
                decreaseAmounts[i],
                timeouts[i],
                canDecreaseAt[i]
            ) = simpleSwap.hardDeposits(actorAddresses[i]);
        }
    }

    function _snapshotAccounting(
        uint256[4] memory amounts,
        uint256[4] memory decreaseAmounts,
        uint256[4] memory timeouts,
        uint256[4] memory canDecreaseAt,
        uint256[4] memory paidOutAmounts
    ) internal view {
        _snapshotBeneficiaries(
            amounts,
            decreaseAmounts,
            timeouts,
            canDecreaseAt
        );

        for (uint256 i = 0; i < ACTOR_COUNT; i++) {
            paidOutAmounts[i] = simpleSwap.paidOut(actorAddresses[i]);
        }
    }

    function _assertUntouched(
        uint256 targetIndex,
        uint256[4] memory amountsBefore,
        uint256[4] memory decreaseAmountsBefore,
        uint256[4] memory timeoutsBefore,
        uint256[4] memory canDecreaseBefore
    ) internal {
        for (uint256 i = 0; i < ACTOR_COUNT; i++) {
            if (i == targetIndex) {
                continue;
            }

            (
                uint256 amountAfter,
                uint256 decreaseAmountAfter,
                uint256 timeoutAfter,
                uint256 canDecreaseAfter
            ) = simpleSwap.hardDeposits(actorAddresses[i]);

            if (amountAfter != amountsBefore[i]) {
                invariantFailed = true;
            }
            if (decreaseAmountAfter != decreaseAmountsBefore[i]) {
                invariantFailed = true;
            }
            if (timeoutAfter != timeoutsBefore[i]) {
                invariantFailed = true;
            }
            if (canDecreaseAfter != canDecreaseBefore[i]) {
                invariantFailed = true;
            }
        }
    }

    function _assertAccountingUntouched(
        uint256 targetIndex,
        uint256[4] memory amountsBefore,
        uint256[4] memory decreaseAmountsBefore,
        uint256[4] memory timeoutsBefore,
        uint256[4] memory canDecreaseBefore,
        uint256[4] memory paidOutBefore
    ) internal {
        _assertUntouched(
            targetIndex,
            amountsBefore,
            decreaseAmountsBefore,
            timeoutsBefore,
            canDecreaseBefore
        );

        for (uint256 i = 0; i < ACTOR_COUNT; i++) {
            if (i == targetIndex) {
                continue;
            }
            if (simpleSwap.paidOut(actorAddresses[i]) != paidOutBefore[i]) {
                invariantFailed = true;
            }
        }
    }

    function _snapshotFullState()
        internal
        view
        returns (FullStateSnapshot memory snapshot)
    {
        snapshot.swapBalance = simpleSwap.balance();
        snapshot.totalHardDeposit = simpleSwap.totalHardDeposit();
        snapshot.totalPaidOut = simpleSwap.totalPaidOut();
        snapshot.bounced = simpleSwap.bounced();

        for (uint256 i = 0; i < ACTOR_COUNT; i++) {
            (
                snapshot.hardAmounts[i],
                snapshot.decreaseAmounts[i],
                snapshot.timeouts[i],
                snapshot.canDecreaseAt[i]
            ) = simpleSwap.hardDeposits(actorAddresses[i]);
            snapshot.paidOut[i] = simpleSwap.paidOut(actorAddresses[i]);
            snapshot.actorBalances[i] = token.balanceOf(actorAddresses[i]);
        }
    }

    function _assertFullStateUnchanged(FullStateSnapshot memory snapshot)
        internal
    {
        if (simpleSwap.balance() != snapshot.swapBalance) {
            invariantFailed = true;
        }
        if (simpleSwap.totalHardDeposit() != snapshot.totalHardDeposit) {
            invariantFailed = true;
        }
        if (simpleSwap.totalPaidOut() != snapshot.totalPaidOut) {
            invariantFailed = true;
        }
        if (simpleSwap.bounced() != snapshot.bounced) {
            invariantFailed = true;
        }

        for (uint256 i = 0; i < ACTOR_COUNT; i++) {
            (
                uint256 hardAmount,
                uint256 decreaseAmount,
                uint256 timeout,
                uint256 canDecreaseAt
            ) = simpleSwap.hardDeposits(actorAddresses[i]);

            if (hardAmount != snapshot.hardAmounts[i]) {
                invariantFailed = true;
            }
            if (decreaseAmount != snapshot.decreaseAmounts[i]) {
                invariantFailed = true;
            }
            if (timeout != snapshot.timeouts[i]) {
                invariantFailed = true;
            }
            if (canDecreaseAt != snapshot.canDecreaseAt[i]) {
                invariantFailed = true;
            }
            if (simpleSwap.paidOut(actorAddresses[i]) != snapshot.paidOut[i]) {
                invariantFailed = true;
            }
            if (token.balanceOf(actorAddresses[i]) != snapshot.actorBalances[i]) {
                invariantFailed = true;
            }
        }
    }

    function _checkIncreaseHardDepositPostconditions(
        uint256 beneficiaryIndex,
        uint256 amount,
        uint256 balanceBefore,
        uint256 liquidBefore,
        uint256 totalHardBefore,
        uint256[4] memory amountsBefore,
        uint256[4] memory decreaseAmountsBefore,
        uint256[4] memory timeoutsBefore
    ) internal {
        (
            uint256 targetAmountAfter,
            uint256 targetDecreaseAfter,
            uint256 targetTimeoutAfter,
            uint256 targetCanDecreaseAfter
        ) = simpleSwap.hardDeposits(actorAddresses[beneficiaryIndex]);

        if (simpleSwap.balance() != balanceBefore) {
            invariantFailed = true;
        }
        if (simpleSwap.liquidBalance() != liquidBefore.sub(amount)) {
            invariantFailed = true;
        }
        if (simpleSwap.totalHardDeposit() != totalHardBefore.add(amount)) {
            invariantFailed = true;
        }
        if (targetAmountAfter != amountsBefore[beneficiaryIndex].add(amount)) {
            invariantFailed = true;
        }
        if (targetDecreaseAfter != decreaseAmountsBefore[beneficiaryIndex]) {
            invariantFailed = true;
        }
        if (targetTimeoutAfter != timeoutsBefore[beneficiaryIndex]) {
            invariantFailed = true;
        }
        if (targetCanDecreaseAfter != 0) {
            invariantFailed = true;
        }
    }

    function _checkCashChequePostconditions(
        uint256 callerIndex,
        uint256 beneficiaryIndex,
        uint256 recipientIndex,
        uint256[4] memory amountsBefore,
        uint256[4] memory canDecreaseBefore,
        uint256[4] memory paidOutBefore,
        CashSnapshot memory snapshot
    ) internal {
        (
            uint256 targetAmountAfter,
            ,
            ,
            uint256 targetCanDecreaseAfter
        ) = simpleSwap.hardDeposits(actorAddresses[beneficiaryIndex]);

        uint256 expectedRecipientGain = recipientIndex == callerIndex
            ? snapshot.totalPayout
            : snapshot.totalPayout.sub(snapshot.callerPayout);
        uint256 expectedCallerGain = recipientIndex == callerIndex
            ? snapshot.totalPayout
            : snapshot.callerPayout;

        if (
            simpleSwap.paidOut(actorAddresses[beneficiaryIndex]) !=
            paidOutBefore[beneficiaryIndex].add(snapshot.totalPayout)
        ) {
            invariantFailed = true;
        }
        if (
            simpleSwap.totalPaidOut() !=
            snapshot.totalPaidBefore.add(snapshot.totalPayout)
        ) {
            invariantFailed = true;
        }
        if (
            simpleSwap.totalHardDeposit() !=
            snapshot.totalHardBefore.sub(snapshot.hardDepositUsage)
        ) {
            invariantFailed = true;
        }
        if (
            targetAmountAfter !=
            amountsBefore[beneficiaryIndex].sub(snapshot.hardDepositUsage)
        ) {
            invariantFailed = true;
        }
        if (
            simpleSwap.balance() !=
            snapshot.swapBalanceBefore.sub(snapshot.totalPayout)
        ) {
            invariantFailed = true;
        }
        if (
            token.balanceOf(actorAddresses[recipientIndex]) !=
            snapshot.recipientBalanceBefore.add(expectedRecipientGain)
        ) {
            invariantFailed = true;
        }
        if (
            token.balanceOf(actorAddresses[callerIndex]) !=
            snapshot.callerBalanceBefore.add(expectedCallerGain)
        ) {
            invariantFailed = true;
        }
        if (
            snapshot.requestPayout > snapshot.totalPayout &&
            !simpleSwap.bounced()
        ) {
            invariantFailed = true;
        }
        if (
            snapshot.requestPayout == snapshot.totalPayout &&
            snapshot.bouncedBefore != simpleSwap.bounced()
        ) {
            invariantFailed = true;
        }
        if (targetCanDecreaseAfter != canDecreaseBefore[beneficiaryIndex]) {
            invariantFailed = true;
        }
    }

    function _checkTemporalInvariants() internal {
        uint256 currentTotalPaidOut = simpleSwap.totalPaidOut();

        if (currentTotalPaidOut < lastTotalPaidOut) {
            invariantFailed = true;
        }

        for (uint256 i = 0; i < ACTOR_COUNT; i++) {
            uint256 currentPaidOut = simpleSwap.paidOut(actorAddresses[i]);

            if (currentPaidOut < lastPaidOut[i]) {
                invariantFailed = true;
            }

            lastPaidOut[i] = currentPaidOut;
        }

        if (bounceSeen && !simpleSwap.bounced()) {
            invariantFailed = true;
        }

        if (simpleSwap.bounced()) {
            bounceSeen = true;
        }

        lastTotalPaidOut = currentTotalPaidOut;
    }

    function _syncTemporalState() internal {
        lastTotalPaidOut = simpleSwap.totalPaidOut();

        for (uint256 i = 0; i < ACTOR_COUNT; i++) {
            lastPaidOut[i] = simpleSwap.paidOut(actorAddresses[i]);
        }

        bounceSeen = simpleSwap.bounced();
    }

    function _expectedCashEffects(address beneficiary, uint256 cumulativePayout)
        internal
        view
        returns (
            uint256 requestPayout,
            uint256 totalPayout,
            uint256 hardDepositUsage
        )
    {
        uint256 paidOutBefore = simpleSwap.paidOut(beneficiary);
        requestPayout = cumulativePayout.sub(paidOutBefore);
        totalPayout = _min(requestPayout, simpleSwap.liquidBalanceFor(beneficiary));
        (uint256 hardDepositAmount, , , ) = simpleSwap.hardDeposits(beneficiary);
        hardDepositUsage = _min(totalPayout, hardDepositAmount);
    }

    function _sumActorBalances() internal view returns (uint256 totalBalance) {
        for (uint256 i = 0; i < ACTOR_COUNT; i++) {
            totalBalance = totalBalance.add(token.balanceOf(actorAddresses[i]));
        }
    }

    function _actorIndex(uint256 seed) internal pure returns (uint256) {
        return seed % ACTOR_COUNT;
    }

    function _nonIssuerIndex(uint256 seed) internal pure returns (uint256) {
        uint256 actorIndex = seed % ACTOR_COUNT;
        return actorIndex == 0 ? 1 : actorIndex;
    }

    function _bounded(uint256 value, uint256 max) internal pure returns (uint256) {
        if (max == 0) {
            return 0;
        }

        return value % (max + 1);
    }

    function _min(uint256 a, uint256 b) internal pure returns (uint256) {
        return a < b ? a : b;
    }

    function _effectiveTimeout(uint256 customTimeout)
        internal
        view
        returns (uint256)
    {
        return
            customTimeout == 0
                ? simpleSwap.defaultHardDepositTimeout()
                : customTimeout;
    }
}
