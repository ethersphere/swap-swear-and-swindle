// SPDX-License-Identifier: BSD-3-Clause
pragma solidity =0.7.6;
pragma abicoder v2;

import "../SimpleSwapFactory.sol";
import "../TestToken.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";

interface IFactorySystemSwapTarget {
    function init(
        address issuer,
        address token,
        uint256 defaultHardDepositTimeout
    ) external;

    function increaseHardDeposit(address beneficiary, uint256 amount) external;

    function prepareDecreaseHardDeposit(
        address beneficiary,
        uint256 decreaseAmount
    ) external;

    function decreaseHardDeposit(address beneficiary) external;

    function withdraw(uint256 amount) external;

    function cashChequeBeneficiary(
        address recipient,
        uint256 cumulativePayout,
        bytes calldata issuerSig
    ) external;
}

contract FactorySystemActor {
    function deploySimpleSwap(
        SimpleSwapFactory factory,
        address issuer,
        uint256 defaultHardDepositTimeout,
        bytes32 salt
    ) external returns (bool, address) {
        try
            factory.deploySimpleSwap(issuer, defaultHardDepositTimeout, salt)
        returns (address deployedAddress) {
            return (true, deployedAddress);
        } catch {
            return (false, address(0));
        }
    }

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
        IFactorySystemSwapTarget target,
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
        IFactorySystemSwapTarget target,
        address beneficiary,
        uint256 amount
    ) external returns (bool) {
        try target.prepareDecreaseHardDeposit(beneficiary, amount) {
            return true;
        } catch {
            return false;
        }
    }

    function decreaseHardDeposit(
        IFactorySystemSwapTarget target,
        address beneficiary
    ) external returns (bool) {
        try target.decreaseHardDeposit(beneficiary) {
            return true;
        } catch {
            return false;
        }
    }

    function withdraw(
        IFactorySystemSwapTarget target,
        uint256 amount
    ) external returns (bool) {
        try target.withdraw(amount) {
            return true;
        } catch {
            return false;
        }
    }

    function cashChequeBeneficiary(
        IFactorySystemSwapTarget target,
        address recipient,
        uint256 cumulativePayout
    ) external returns (bool) {
        try target.cashChequeBeneficiary(recipient, cumulativePayout, "") {
            return true;
        } catch {
            return false;
        }
    }

    function reinitSwap(
        IFactorySystemSwapTarget target,
        address issuer,
        address token,
        uint256 timeout
    ) external returns (bool) {
        try target.init(issuer, token, timeout) {
            return true;
        } catch {
            return false;
        }
    }
}

contract SimpleSwapFactorySystemEchidna {
    using SafeMath for uint256;

    uint256 private constant ACTOR_COUNT = 3;
    uint256 private constant TRACKED_CLONES = 8;
    uint256 private constant INITIAL_TOKEN_BALANCE = 1_000_000;

    struct SystemSnapshot {
        uint256 cloneCount;
        uint256[8] balances;
        uint256[8] totalHardDeposits;
        uint256[8] totalPaidOut;
        bool[8] bounced;
        uint256[3][8] hardAmounts;
        uint256[3][8] decreaseAmounts;
        uint256[3][8] timeouts;
        uint256[3][8] canDecreaseAt;
        uint256[3][8] paidOut;
        uint256[3] actorBalances;
    }

    struct CloneCashSnapshot {
        uint256 paidOutBefore;
        uint256 balanceBefore;
        uint256 totalHardBefore;
        uint256 totalPaidBefore;
        uint256 issuerHardDepositBefore;
        uint256 liquidForIssuerBefore;
        uint256 requestPayout;
        uint256 totalPayout;
        uint256 hardDepositUsage;
        bool bouncedBefore;
    }

    TestToken private token;
    SimpleSwapFactory private factory;
    FactorySystemActor[3] private actors;
    address[3] private actorAddresses;
    address private factoryMaster;
    address[8] private clones;
    address[8] private cloneCallers;
    address[8] private cloneIssuers;
    bytes32[8] private cloneSalts;
    uint256[8] private cloneTimeouts;
    uint256 private cloneCount;
    bool private invariantFailed;

    constructor() {
        token = new TestToken();
        factory = new SimpleSwapFactory(address(token));
        factoryMaster = factory.master();

        for (uint256 i = 0; i < ACTOR_COUNT; i++) {
            actors[i] = new FactorySystemActor();
            actorAddresses[i] = address(actors[i]);
            token.mint(actorAddresses[i], INITIAL_TOKEN_BALANCE);
        }
    }

    function deploySimpleSwap(
        uint256 callerSeed,
        uint256 issuerSeed,
        uint256 defaultHardDepositTimeout,
        bytes32 salt
    ) public {
        uint256 callerIndex = callerSeed % ACTOR_COUNT;
        address caller = actorAddresses[callerIndex];
        address issuer = actorAddresses[issuerSeed % ACTOR_COUNT];
        bool alreadyTracked = _trackedIndex(caller, salt) != TRACKED_CLONES;

        (bool ok, address deployedAddress) = actors[callerIndex].deploySimpleSwap(
            factory,
            issuer,
            defaultHardDepositTimeout,
            salt
        );

        if (!ok) {
            if (!alreadyTracked && issuer != address(0)) {
                invariantFailed = true;
            }
            return;
        }

        if (issuer == address(0) || alreadyTracked) {
            invariantFailed = true;
            return;
        }

        if (cloneCount < TRACKED_CLONES) {
            clones[cloneCount] = deployedAddress;
            cloneCallers[cloneCount] = caller;
            cloneIssuers[cloneCount] = issuer;
            cloneSalts[cloneCount] = salt;
            cloneTimeouts[cloneCount] = defaultHardDepositTimeout;
            cloneCount++;
        }

        _checkCloneMetadata(_cloneIndex(deployedAddress));
    }

    function depositToTrackedClone(
        uint256 trackedSeed,
        uint256 actorSeed,
        uint256 rawAmount
    ) public {
        if (cloneCount == 0) {
            return;
        }

        uint256 cloneIndex = trackedSeed % cloneCount;
        uint256 actorIndex = actorSeed % ACTOR_COUNT;
        uint256 amount = _bounded(
            rawAmount,
            token.balanceOf(actorAddresses[actorIndex])
        );
        SystemSnapshot memory snapshot = _snapshotSystem();

        if (amount == 0) {
            return;
        }

        if (!actors[actorIndex].depositToken(token, clones[cloneIndex], amount)) {
            invariantFailed = true;
            return;
        }

        _assertCloneScalarState(
            cloneIndex,
            snapshot.balances[cloneIndex].add(amount),
            snapshot.totalHardDeposits[cloneIndex],
            snapshot.totalPaidOut[cloneIndex],
            snapshot.bounced[cloneIndex]
        );
        _assertCloneActorStateUnchanged(cloneIndex, snapshot);
        if (
            token.balanceOf(actorAddresses[actorIndex]) !=
            snapshot.actorBalances[actorIndex].sub(amount)
        ) {
            invariantFailed = true;
        }

        _assertOtherClonesUnchanged(cloneIndex, snapshot);
        _assertFactoryImmutable();
    }

    function increaseHardDepositOnTrackedClone(
        uint256 trackedSeed,
        uint256 beneficiarySeed,
        uint256 rawAmount
    ) public {
        if (cloneCount == 0) {
            return;
        }

        uint256 cloneIndex = trackedSeed % cloneCount;
        ERC20SimpleSwap swap = ERC20SimpleSwap(clones[cloneIndex]);
        uint256 amount = _bounded(rawAmount, swap.liquidBalance());
        uint256 issuerIndex = _actorIndexForAddress(cloneIssuers[cloneIndex]);
        uint256 beneficiaryIndex = beneficiarySeed % ACTOR_COUNT;
        SystemSnapshot memory snapshot = _snapshotSystem();
        (uint256 hardAmountBefore, , , ) = swap.hardDeposits(
            actorAddresses[beneficiaryIndex]
        );

        if (amount == 0) {
            return;
        }

        if (
            !actors[issuerIndex].increaseHardDeposit(
                IFactorySystemSwapTarget(clones[cloneIndex]),
                actorAddresses[beneficiaryIndex],
                amount
            )
        ) {
            invariantFailed = true;
            return;
        }

        _assertIncreaseHardDepositPostconditions(
            cloneIndex,
            beneficiaryIndex,
            hardAmountBefore,
            amount,
            snapshot
        );

        _assertOtherClonesUnchanged(cloneIndex, snapshot);
        _assertFactoryImmutable();
    }

    function prepareDecreaseHardDepositOnTrackedClone(
        uint256 trackedSeed,
        uint256 beneficiarySeed,
        uint256 rawAmount
    ) public {
        if (cloneCount == 0) {
            return;
        }

        uint256 cloneIndex = trackedSeed % cloneCount;
        ERC20SimpleSwap swap = ERC20SimpleSwap(clones[cloneIndex]);
        uint256 issuerIndex = _actorIndexForAddress(cloneIssuers[cloneIndex]);
        uint256 beneficiaryIndex = beneficiarySeed % ACTOR_COUNT;
        (uint256 hardAmountBefore, , uint256 timeoutBefore, ) = swap.hardDeposits(
            actorAddresses[beneficiaryIndex]
        );
        uint256 amount = _bounded(rawAmount, hardAmountBefore);
        SystemSnapshot memory snapshot = _snapshotSystem();

        if (amount == 0) {
            return;
        }

        if (
            !actors[issuerIndex].prepareDecreaseHardDeposit(
                IFactorySystemSwapTarget(clones[cloneIndex]),
                actorAddresses[beneficiaryIndex],
                amount
            )
        ) {
            invariantFailed = true;
            return;
        }

        _assertPrepareDecreaseHardDepositPostconditions(
            cloneIndex,
            beneficiaryIndex,
            hardAmountBefore,
            timeoutBefore,
            amount,
            snapshot
        );

        _assertOtherClonesUnchanged(cloneIndex, snapshot);
        _assertFactoryImmutable();
    }

    function decreaseHardDepositOnTrackedClone(
        uint256 trackedSeed,
        uint256 beneficiarySeed
    ) public {
        if (cloneCount == 0) {
            return;
        }

        uint256 cloneIndex = trackedSeed % cloneCount;
        uint256 beneficiaryIndex = beneficiarySeed % ACTOR_COUNT;
        ERC20SimpleSwap swap = ERC20SimpleSwap(clones[cloneIndex]);
        SystemSnapshot memory snapshot = _snapshotSystem();
        (
            uint256 hardAmountBefore,
            uint256 decreaseAmountBefore,
            uint256 timeoutBefore,
            uint256 canDecreaseAtBefore
        ) = swap.hardDeposits(actorAddresses[beneficiaryIndex]);
        bool ok = actors[0].decreaseHardDeposit(
            IFactorySystemSwapTarget(clones[cloneIndex]),
            actorAddresses[beneficiaryIndex]
        );

        if (!ok) {
            _assertCloneScalarState(
                cloneIndex,
                snapshot.balances[cloneIndex],
                snapshot.totalHardDeposits[cloneIndex],
                snapshot.totalPaidOut[cloneIndex],
                snapshot.bounced[cloneIndex]
            );
            _assertCloneActorStateUnchanged(cloneIndex, snapshot);
            _assertOtherClonesUnchanged(cloneIndex, snapshot);
            _assertFactoryImmutable();
            return;
        }

        _assertDecreaseHardDepositPostconditions(
            cloneIndex,
            beneficiaryIndex,
            hardAmountBefore,
            decreaseAmountBefore,
            timeoutBefore,
            snapshot
        );

        if (
            canDecreaseAtBefore != 0 && block.timestamp < canDecreaseAtBefore
        ) {
            invariantFailed = true;
        }

        _assertOtherClonesUnchanged(cloneIndex, snapshot);
        _assertFactoryImmutable();
    }

    function withdrawFromTrackedClone(
        uint256 trackedSeed,
        uint256 rawAmount
    ) public {
        if (cloneCount == 0) {
            return;
        }

        uint256 cloneIndex = trackedSeed % cloneCount;
        ERC20SimpleSwap swap = ERC20SimpleSwap(clones[cloneIndex]);
        uint256 issuerIndex = _actorIndexForAddress(cloneIssuers[cloneIndex]);
        uint256 amount = _bounded(rawAmount, swap.liquidBalance());
        SystemSnapshot memory snapshot = _snapshotSystem();

        if (amount == 0) {
            return;
        }

        if (
            !actors[issuerIndex].withdraw(
                IFactorySystemSwapTarget(clones[cloneIndex]),
                amount
            )
        ) {
            invariantFailed = true;
            return;
        }

        _assertCloneScalarState(
            cloneIndex,
            snapshot.balances[cloneIndex].sub(amount),
            snapshot.totalHardDeposits[cloneIndex],
            snapshot.totalPaidOut[cloneIndex],
            snapshot.bounced[cloneIndex]
        );
        _assertCloneActorStateUnchanged(cloneIndex, snapshot);
        if (
            token.balanceOf(actorAddresses[issuerIndex]) !=
            snapshot.actorBalances[issuerIndex].add(amount)
        ) {
            invariantFailed = true;
        }

        _assertOtherClonesUnchanged(cloneIndex, snapshot);
        _assertFactoryImmutable();
    }

    function cashChequeBeneficiaryOnTrackedClone(
        uint256 trackedSeed,
        uint256 recipientSeed,
        uint256 cumulativeDelta
    ) public {
        if (cloneCount == 0) {
            return;
        }

        uint256 cloneIndex = trackedSeed % cloneCount;
        uint256 issuerIndex = _actorIndexForAddress(cloneIssuers[cloneIndex]);
        uint256 recipientIndex = recipientSeed % ACTOR_COUNT;
        ERC20SimpleSwap swap = ERC20SimpleSwap(clones[cloneIndex]);
        SystemSnapshot memory snapshot = _snapshotSystem();
        CloneCashSnapshot memory cashSnapshot;
        cashSnapshot.paidOutBefore = swap.paidOut(cloneIssuers[cloneIndex]);
        cashSnapshot.balanceBefore = swap.balance();
        cashSnapshot.totalHardBefore = swap.totalHardDeposit();
        cashSnapshot.totalPaidBefore = swap.totalPaidOut();
        cashSnapshot.issuerHardDepositBefore = _hardDepositAmount(
            swap,
            cloneIssuers[cloneIndex]
        );
        cashSnapshot.liquidForIssuerBefore = swap.liquidBalanceFor(
            cloneIssuers[cloneIndex]
        );
        cashSnapshot.bouncedBefore = swap.bounced();
        cashSnapshot.requestPayout = _bounded(cumulativeDelta, token.totalSupply());
        uint256 cumulativePayout = cashSnapshot.paidOutBefore.add(
            cashSnapshot.requestPayout
        );
        cashSnapshot.totalPayout = _min(
            cashSnapshot.requestPayout,
            cashSnapshot.liquidForIssuerBefore
        );
        cashSnapshot.hardDepositUsage = _min(
            cashSnapshot.totalPayout,
            cashSnapshot.issuerHardDepositBefore
        );

        if (
            !actors[issuerIndex].cashChequeBeneficiary(
                IFactorySystemSwapTarget(clones[cloneIndex]),
                actorAddresses[recipientIndex],
                cumulativePayout
            )
        ) {
            invariantFailed = true;
            return;
        }

        _checkCloneCashPostconditions(
            cloneIndex,
            recipientIndex,
            snapshot,
            cashSnapshot,
            swap
        );

        _assertOtherClonesUnchanged(cloneIndex, snapshot);
        _assertFactoryImmutable();
    }

    function reinitTrackedClone(
        uint256 trackedSeed,
        uint256 issuerSeed,
        uint256 timeoutSeed
    ) public {
        if (cloneCount == 0) {
            return;
        }

        uint256 cloneIndex = trackedSeed % cloneCount;
        SystemSnapshot memory snapshot = _snapshotSystem();
        bool ok = actors[0].reinitSwap(
            IFactorySystemSwapTarget(clones[cloneIndex]),
            actorAddresses[issuerSeed % ACTOR_COUNT],
            address(token),
            timeoutSeed
        );

        if (ok) {
            invariantFailed = true;
        }

        _assertEntireSystemUnchanged(snapshot);
        _assertFactoryImmutable();
    }

    function redeployTrackedCloneWithDifferentParams(
        uint256 trackedSeed,
        uint256 issuerSeed,
        uint256 timeoutSeed
    ) public {
        if (cloneCount == 0) {
            return;
        }

        uint256 cloneIndex = trackedSeed % cloneCount;
        uint256 callerIndex = _actorIndexForAddress(cloneCallers[cloneIndex]);
        address issuer = actorAddresses[issuerSeed % ACTOR_COUNT];
        uint256 timeout = cloneTimeouts[cloneIndex].add(1).add(timeoutSeed % 3);
        SystemSnapshot memory snapshot = _snapshotSystem();
        (bool ok, address deployedAddress) = actors[callerIndex].deploySimpleSwap(
            factory,
            issuer,
            timeout,
            cloneSalts[cloneIndex]
        );

        if (ok || deployedAddress != address(0)) {
            invariantFailed = true;
        }

        _assertEntireSystemUnchanged(snapshot);
        _assertFactoryImmutable();
    }

    function echidna_factory_state_is_immutable() public view returns (bool) {
        return
            factory.master() == factoryMaster &&
            factory.ERC20Address() == address(token);
    }

    function echidna_master_remains_sentinel_initialized()
        public
        view
        returns (bool)
    {
        ERC20SimpleSwap master = ERC20SimpleSwap(factory.master());
        return
            master.issuer() == address(1) &&
            address(master.token()) == address(0) &&
            master.defaultHardDepositTimeout() == 0;
    }

    function echidna_tracked_clones_remain_initialized()
        public
        view
        returns (bool)
    {
        for (uint256 i = 0; i < cloneCount; i++) {
            ERC20SimpleSwap swap = ERC20SimpleSwap(clones[i]);

            if (swap.issuer() != cloneIssuers[i]) {
                return false;
            }
            if (swap.issuer() == address(0)) {
                return false;
            }
            if (address(swap.token()) != address(token)) {
                return false;
            }
            if (swap.defaultHardDepositTimeout() != cloneTimeouts[i]) {
                return false;
            }
        }

        return true;
    }

    function echidna_tracked_clones_are_registered()
        public
        view
        returns (bool)
    {
        for (uint256 i = 0; i < cloneCount; i++) {
            if (!factory.deployedContracts(clones[i])) {
                return false;
            }
        }

        return true;
    }

    function echidna_tracked_clones_are_unique() public view returns (bool) {
        for (uint256 i = 0; i < cloneCount; i++) {
            for (uint256 j = i + 1; j < cloneCount; j++) {
                if (clones[i] == clones[j]) {
                    return false;
                }
            }
        }

        return true;
    }

    function echidna_no_postcondition_failures() public view returns (bool) {
        return !invariantFailed;
    }

    function _snapshotSystem()
        internal
        view
        returns (SystemSnapshot memory snapshot)
    {
        snapshot.cloneCount = cloneCount;

        for (uint256 i = 0; i < cloneCount; i++) {
            ERC20SimpleSwap swap = ERC20SimpleSwap(clones[i]);
            snapshot.balances[i] = swap.balance();
            snapshot.totalHardDeposits[i] = swap.totalHardDeposit();
            snapshot.totalPaidOut[i] = swap.totalPaidOut();
            snapshot.bounced[i] = swap.bounced();

            for (uint256 j = 0; j < ACTOR_COUNT; j++) {
                (
                    snapshot.hardAmounts[i][j],
                    snapshot.decreaseAmounts[i][j],
                    snapshot.timeouts[i][j],
                    snapshot.canDecreaseAt[i][j]
                ) = swap.hardDeposits(actorAddresses[j]);
                snapshot.paidOut[i][j] = swap.paidOut(actorAddresses[j]);
            }
        }

        for (uint256 i = 0; i < ACTOR_COUNT; i++) {
            snapshot.actorBalances[i] = token.balanceOf(actorAddresses[i]);
        }
    }

    function _assertOtherClonesUnchanged(
        uint256 targetCloneIndex,
        SystemSnapshot memory snapshot
    ) internal {
        for (uint256 i = 0; i < snapshot.cloneCount; i++) {
            if (i == targetCloneIndex) {
                continue;
            }

            ERC20SimpleSwap swap = ERC20SimpleSwap(clones[i]);

            if (swap.balance() != snapshot.balances[i]) {
                invariantFailed = true;
            }
            if (swap.totalHardDeposit() != snapshot.totalHardDeposits[i]) {
                invariantFailed = true;
            }
            if (swap.totalPaidOut() != snapshot.totalPaidOut[i]) {
                invariantFailed = true;
            }
            if (swap.bounced() != snapshot.bounced[i]) {
                invariantFailed = true;
            }
            _assertCloneActorStateUnchanged(i, snapshot);
            _checkCloneMetadata(i);
        }
    }

    function _assertEntireSystemUnchanged(SystemSnapshot memory snapshot)
        internal
    {
        if (cloneCount != snapshot.cloneCount) {
            invariantFailed = true;
        }

        for (uint256 i = 0; i < snapshot.cloneCount; i++) {
            ERC20SimpleSwap swap = ERC20SimpleSwap(clones[i]);

            if (swap.balance() != snapshot.balances[i]) {
                invariantFailed = true;
            }
            if (swap.totalHardDeposit() != snapshot.totalHardDeposits[i]) {
                invariantFailed = true;
            }
            if (swap.totalPaidOut() != snapshot.totalPaidOut[i]) {
                invariantFailed = true;
            }
            if (swap.bounced() != snapshot.bounced[i]) {
                invariantFailed = true;
            }
            _assertCloneActorStateUnchanged(i, snapshot);
            _checkCloneMetadata(i);
        }

        for (uint256 i = 0; i < ACTOR_COUNT; i++) {
            if (token.balanceOf(actorAddresses[i]) != snapshot.actorBalances[i]) {
                invariantFailed = true;
            }
        }
    }

    function _checkCloneMetadata(uint256 cloneIndex) internal {
        ERC20SimpleSwap swap = ERC20SimpleSwap(clones[cloneIndex]);

        if (swap.issuer() != cloneIssuers[cloneIndex]) {
            invariantFailed = true;
        }
        if (address(swap.token()) != address(token)) {
            invariantFailed = true;
        }
        if (
            swap.defaultHardDepositTimeout() != cloneTimeouts[cloneIndex]
        ) {
            invariantFailed = true;
        }
    }

    function _assertCloneActorStateUnchanged(
        uint256 cloneIndex,
        SystemSnapshot memory snapshot
    ) internal {
        _assertCloneActorStateMatches(
            cloneIndex,
            snapshot,
            ACTOR_COUNT,
            0,
            0,
            0,
            0,
            0
        );
    }

    function _assertCloneScalarState(
        uint256 cloneIndex,
        uint256 expectedBalance,
        uint256 expectedTotalHardDeposit,
        uint256 expectedTotalPaidOut,
        bool expectedBounced
    ) internal {
        ERC20SimpleSwap swap = ERC20SimpleSwap(clones[cloneIndex]);

        if (swap.balance() != expectedBalance) {
            invariantFailed = true;
        }
        if (swap.totalHardDeposit() != expectedTotalHardDeposit) {
            invariantFailed = true;
        }
        if (swap.totalPaidOut() != expectedTotalPaidOut) {
            invariantFailed = true;
        }
        if (swap.bounced() != expectedBounced) {
            invariantFailed = true;
        }
        _checkCloneMetadata(cloneIndex);
    }

    function _assertIncreaseHardDepositPostconditions(
        uint256 cloneIndex,
        uint256 beneficiaryIndex,
        uint256 hardAmountBefore,
        uint256 amount,
        SystemSnapshot memory snapshot
    ) internal {
        _assertCloneScalarState(
            cloneIndex,
            snapshot.balances[cloneIndex],
            snapshot.totalHardDeposits[cloneIndex].add(amount),
            snapshot.totalPaidOut[cloneIndex],
            snapshot.bounced[cloneIndex]
        );
        _assertCloneActorStateMatches(
            cloneIndex,
            snapshot,
            beneficiaryIndex,
            hardAmountBefore.add(amount),
            snapshot.decreaseAmounts[cloneIndex][beneficiaryIndex],
            snapshot.timeouts[cloneIndex][beneficiaryIndex],
            0,
            snapshot.paidOut[cloneIndex][beneficiaryIndex]
        );
    }

    function _assertPrepareDecreaseHardDepositPostconditions(
        uint256 cloneIndex,
        uint256 beneficiaryIndex,
        uint256 hardAmountBefore,
        uint256 timeoutBefore,
        uint256 amount,
        SystemSnapshot memory snapshot
    ) internal {
        uint256 effectiveTimeout = timeoutBefore == 0
            ? cloneTimeouts[cloneIndex]
            : timeoutBefore;

        _assertCloneScalarState(
            cloneIndex,
            snapshot.balances[cloneIndex],
            snapshot.totalHardDeposits[cloneIndex],
            snapshot.totalPaidOut[cloneIndex],
            snapshot.bounced[cloneIndex]
        );
        _assertCloneActorStateMatches(
            cloneIndex,
            snapshot,
            beneficiaryIndex,
            hardAmountBefore,
            amount,
            timeoutBefore,
            block.timestamp.add(effectiveTimeout),
            snapshot.paidOut[cloneIndex][beneficiaryIndex]
        );
    }

    function _assertDecreaseHardDepositPostconditions(
        uint256 cloneIndex,
        uint256 beneficiaryIndex,
        uint256 hardAmountBefore,
        uint256 decreaseAmountBefore,
        uint256 timeoutBefore,
        SystemSnapshot memory snapshot
    ) internal {
        _assertCloneScalarState(
            cloneIndex,
            snapshot.balances[cloneIndex],
            snapshot.totalHardDeposits[cloneIndex].sub(decreaseAmountBefore),
            snapshot.totalPaidOut[cloneIndex],
            snapshot.bounced[cloneIndex]
        );
        _assertCloneActorStateMatches(
            cloneIndex,
            snapshot,
            beneficiaryIndex,
            hardAmountBefore.sub(decreaseAmountBefore),
            decreaseAmountBefore,
            timeoutBefore,
            0,
            snapshot.paidOut[cloneIndex][beneficiaryIndex]
        );
    }

    function _assertCloneActorStateMatches(
        uint256 cloneIndex,
        SystemSnapshot memory snapshot,
        uint256 changedActorIndex,
        uint256 expectedHardAmount,
        uint256 expectedDecreaseAmount,
        uint256 expectedTimeout,
        uint256 expectedCanDecreaseAt,
        uint256 expectedPaidOut
    ) internal {
        ERC20SimpleSwap swap = ERC20SimpleSwap(clones[cloneIndex]);

        for (uint256 i = 0; i < ACTOR_COUNT; i++) {
            (
                uint256 hardAmount,
                uint256 decreaseAmount,
                uint256 timeout,
                uint256 canDecreaseAtValue
            ) = swap.hardDeposits(actorAddresses[i]);
            uint256 paidOutAmount = swap.paidOut(actorAddresses[i]);

            if (i == changedActorIndex) {
                if (hardAmount != expectedHardAmount) {
                    invariantFailed = true;
                }
                if (decreaseAmount != expectedDecreaseAmount) {
                    invariantFailed = true;
                }
                if (timeout != expectedTimeout) {
                    invariantFailed = true;
                }
                if (canDecreaseAtValue != expectedCanDecreaseAt) {
                    invariantFailed = true;
                }
                if (paidOutAmount != expectedPaidOut) {
                    invariantFailed = true;
                }
                continue;
            }

            if (hardAmount != snapshot.hardAmounts[cloneIndex][i]) {
                invariantFailed = true;
            }
            if (decreaseAmount != snapshot.decreaseAmounts[cloneIndex][i]) {
                invariantFailed = true;
            }
            if (timeout != snapshot.timeouts[cloneIndex][i]) {
                invariantFailed = true;
            }
            if (canDecreaseAtValue != snapshot.canDecreaseAt[cloneIndex][i]) {
                invariantFailed = true;
            }
            if (paidOutAmount != snapshot.paidOut[cloneIndex][i]) {
                invariantFailed = true;
            }
        }
    }

    function _assertFactoryImmutable() internal {
        if (factory.master() != factoryMaster) {
            invariantFailed = true;
        }
        if (factory.ERC20Address() != address(token)) {
            invariantFailed = true;
        }

        ERC20SimpleSwap master = ERC20SimpleSwap(factory.master());

        if (master.issuer() != address(1)) {
            invariantFailed = true;
        }
        if (address(master.token()) != address(0)) {
            invariantFailed = true;
        }
        if (master.defaultHardDepositTimeout() != 0) {
            invariantFailed = true;
        }
    }

    function _actorIndexForAddress(address actor) internal view returns (uint256) {
        for (uint256 i = 0; i < ACTOR_COUNT; i++) {
            if (actorAddresses[i] == actor) {
                return i;
            }
        }

        revert("unknown actor");
    }

    function _trackedIndex(address caller, bytes32 salt)
        internal
        view
        returns (uint256)
    {
        for (uint256 i = 0; i < cloneCount; i++) {
            if (cloneCallers[i] == caller && cloneSalts[i] == salt) {
                return i;
            }
        }

        return TRACKED_CLONES;
    }

    function _hardDepositAmount(ERC20SimpleSwap swap, address beneficiary)
        internal
        view
        returns (uint256 amount)
    {
        (amount, , , ) = swap.hardDeposits(beneficiary);
    }

    function _cloneIndex(address cloneAddress) internal view returns (uint256) {
        for (uint256 i = 0; i < cloneCount; i++) {
            if (clones[i] == cloneAddress) {
                return i;
            }
        }

        revert("unknown clone");
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

    function _checkCloneCashPostconditions(
        uint256 cloneIndex,
        uint256 recipientIndex,
        SystemSnapshot memory snapshot,
        CloneCashSnapshot memory cashSnapshot,
        ERC20SimpleSwap swap
    ) internal {
        uint256 issuerIndex = _actorIndexForAddress(cloneIssuers[cloneIndex]);
        bool expectedBounced = cashSnapshot.requestPayout > cashSnapshot.totalPayout
            ? true
            : cashSnapshot.bouncedBefore;

        _assertCloneScalarState(
            cloneIndex,
            cashSnapshot.balanceBefore.sub(cashSnapshot.totalPayout),
            cashSnapshot.totalHardBefore.sub(cashSnapshot.hardDepositUsage),
            cashSnapshot.totalPaidBefore.add(cashSnapshot.totalPayout),
            expectedBounced
        );
        _assertCloneActorStateMatches(
            cloneIndex,
            snapshot,
            issuerIndex,
            cashSnapshot.issuerHardDepositBefore.sub(
                cashSnapshot.hardDepositUsage
            ),
            snapshot.decreaseAmounts[cloneIndex][issuerIndex],
            snapshot.timeouts[cloneIndex][issuerIndex],
            snapshot.canDecreaseAt[cloneIndex][issuerIndex],
            cashSnapshot.paidOutBefore.add(cashSnapshot.totalPayout)
        );
        if (
            token.balanceOf(actorAddresses[recipientIndex]) !=
            snapshot.actorBalances[recipientIndex].add(cashSnapshot.totalPayout)
        ) {
            invariantFailed = true;
        }

        for (uint256 i = 0; i < ACTOR_COUNT; i++) {
            if (i == recipientIndex) {
                continue;
            }

            if (token.balanceOf(actorAddresses[i]) != snapshot.actorBalances[i]) {
                invariantFailed = true;
            }
        }

        if (
            cashSnapshot.requestPayout > cashSnapshot.totalPayout &&
            !swap.bounced()
        ) {
            invariantFailed = true;
        }
        if (
            cashSnapshot.requestPayout == cashSnapshot.totalPayout &&
            swap.bounced() != cashSnapshot.bouncedBefore
        ) {
            invariantFailed = true;
        }
    }
}
