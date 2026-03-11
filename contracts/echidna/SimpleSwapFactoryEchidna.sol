// SPDX-License-Identifier: BSD-3-Clause
pragma solidity =0.7.6;
pragma abicoder v2;

import "../SimpleSwapFactory.sol";
import "../TestToken.sol";

contract FactoryActor {
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
}

contract SimpleSwapFactoryEchidna {
    uint256 private constant ACTOR_COUNT = 3;
    uint256 private constant TRACKED_DEPLOYMENTS = 16;

    TestToken private token;
    SimpleSwapFactory private factory;
    FactoryActor[3] private actors;
    address[3] private actorAddresses;
    address[16] private deployedSwaps;
    address[16] private swapIssuers;
    uint256[16] private swapTimeouts;
    uint256 private deployedSwapCount;
    bool private invariantFailed;

    constructor() {
        token = new TestToken();
        factory = new SimpleSwapFactory(address(token));

        for (uint256 i = 0; i < ACTOR_COUNT; i++) {
            actors[i] = new FactoryActor();
            actorAddresses[i] = address(actors[i]);
        }
    }

    function deploySimpleSwap(
        uint256 callerSeed,
        uint256 issuerSeed,
        uint256 defaultHardDepositTimeout,
        bytes32 salt
    ) public {
        uint256 callerIndex = callerSeed % ACTOR_COUNT;
        address issuer = actorAddresses[issuerSeed % ACTOR_COUNT];

        (bool ok, address deployedAddress) = actors[callerIndex].deploySimpleSwap(
            factory,
            issuer,
            defaultHardDepositTimeout,
            salt
        );

        if (!ok) {
            return;
        }

        if (!factory.deployedContracts(deployedAddress)) {
            invariantFailed = true;
        }

        ERC20SimpleSwap swap = ERC20SimpleSwap(deployedAddress);

        if (swap.issuer() != issuer) {
            invariantFailed = true;
        }
        if (address(swap.token()) != address(token)) {
            invariantFailed = true;
        }
        if (swap.defaultHardDepositTimeout() != defaultHardDepositTimeout) {
            invariantFailed = true;
        }

        if (_seen(deployedAddress)) {
            invariantFailed = true;
            return;
        }

        if (deployedSwapCount < TRACKED_DEPLOYMENTS) {
            deployedSwaps[deployedSwapCount] = deployedAddress;
            swapIssuers[deployedSwapCount] = issuer;
            swapTimeouts[deployedSwapCount] = defaultHardDepositTimeout;
            deployedSwapCount++;
        }
    }

    function echidna_factory_token_is_constant() public view returns (bool) {
        return factory.ERC20Address() == address(token);
    }

    function echidna_master_is_initialized() public view returns (bool) {
        ERC20SimpleSwap master = ERC20SimpleSwap(factory.master());
        return
            master.issuer() == address(1) &&
            address(master.token()) == address(0) &&
            master.defaultHardDepositTimeout() == 0;
    }

    function echidna_tracked_deployments_are_registered()
        public
        view
        returns (bool)
    {
        for (uint256 i = 0; i < deployedSwapCount; i++) {
            if (!factory.deployedContracts(deployedSwaps[i])) {
                return false;
            }
        }

        return true;
    }

    function echidna_tracked_deployments_keep_configuration()
        public
        view
        returns (bool)
    {
        for (uint256 i = 0; i < deployedSwapCount; i++) {
            ERC20SimpleSwap swap = ERC20SimpleSwap(deployedSwaps[i]);

            if (swap.issuer() != swapIssuers[i]) {
                return false;
            }
            if (address(swap.token()) != address(token)) {
                return false;
            }
            if (swap.defaultHardDepositTimeout() != swapTimeouts[i]) {
                return false;
            }
        }

        return true;
    }

    function echidna_no_postcondition_failures() public view returns (bool) {
        return !invariantFailed;
    }

    function _seen(address deployedAddress) internal view returns (bool) {
        for (uint256 i = 0; i < deployedSwapCount; i++) {
            if (deployedSwaps[i] == deployedAddress) {
                return true;
            }
        }

        return false;
    }
}
