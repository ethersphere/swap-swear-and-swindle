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
        uint256[3] actorBalances;
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

        if (
            ERC20SimpleSwap(clones[cloneIndex]).balance() !=
            snapshot.balances[cloneIndex].add(amount)
        ) {
            invariantFailed = true;
        }
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

        (uint256 hardAmountAfter, , , ) = swap.hardDeposits(
            actorAddresses[beneficiaryIndex]
        );

        if (swap.totalHardDeposit() != snapshot.totalHardDeposits[cloneIndex].add(amount)) {
            invariantFailed = true;
        }
        if (hardAmountAfter != hardAmountBefore.add(amount)) {
            invariantFailed = true;
        }
        if (swap.balance() != snapshot.balances[cloneIndex]) {
            invariantFailed = true;
        }

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

        (
            uint256 hardAmountAfter,
            uint256 decreaseAmountAfter,
            uint256 timeoutAfter,
            uint256 canDecreaseAtAfter
        ) = swap.hardDeposits(actorAddresses[beneficiaryIndex]);

        if (hardAmountAfter != hardAmountBefore) {
            invariantFailed = true;
        }
        if (decreaseAmountAfter != amount) {
            invariantFailed = true;
        }
        if (timeoutAfter != timeoutBefore) {
            invariantFailed = true;
        }
        if (canDecreaseAtAfter == 0) {
            invariantFailed = true;
        }

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
            if (
                swap.balance() != snapshot.balances[cloneIndex] ||
                swap.totalHardDeposit() != snapshot.totalHardDeposits[cloneIndex] ||
                swap.totalPaidOut() != snapshot.totalPaidOut[cloneIndex]
            ) {
                invariantFailed = true;
            }
            _assertOtherClonesUnchanged(cloneIndex, snapshot);
            _assertFactoryImmutable();
            return;
        }

        (
            uint256 hardAmountAfter,
            uint256 decreaseAmountAfter,
            uint256 timeoutAfter,
            uint256 canDecreaseAtAfter
        ) = swap.hardDeposits(actorAddresses[beneficiaryIndex]);

        if (
            hardAmountAfter != hardAmountBefore.sub(decreaseAmountBefore)
        ) {
            invariantFailed = true;
        }
        if (decreaseAmountAfter != decreaseAmountBefore) {
            invariantFailed = true;
        }
        if (timeoutAfter != timeoutBefore) {
            invariantFailed = true;
        }
        if (canDecreaseAtAfter != 0) {
            invariantFailed = true;
        }
        if (
            swap.totalHardDeposit() !=
            snapshot.totalHardDeposits[cloneIndex].sub(decreaseAmountBefore)
        ) {
            invariantFailed = true;
        }
        if (swap.balance() != snapshot.balances[cloneIndex]) {
            invariantFailed = true;
        }

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

        if (
            swap.balance() != snapshot.balances[cloneIndex].sub(amount)
        ) {
            invariantFailed = true;
        }
        if (
            token.balanceOf(actorAddresses[issuerIndex]) !=
            snapshot.actorBalances[issuerIndex].add(amount)
        ) {
            invariantFailed = true;
        }
        if (
            swap.totalHardDeposit() != snapshot.totalHardDeposits[cloneIndex]
        ) {
            invariantFailed = true;
        }

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
}
