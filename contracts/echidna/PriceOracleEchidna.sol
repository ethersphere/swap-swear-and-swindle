// SPDX-License-Identifier: BSD-3-Clause
pragma solidity ^0.8.4;

import "../PriceOracle.sol";

contract PriceOracleActor {
    function deployPriceOracle(
        uint256 initialPrice,
        uint256 initialChequeValueDeduction
    ) external returns (PriceOracle) {
        return new PriceOracle(initialPrice, initialChequeValueDeduction);
    }

    function updatePrice(
        PriceOracle oracle,
        uint256 newPrice
    ) external returns (bool) {
        try oracle.updatePrice(newPrice) {
            return true;
        } catch {
            return false;
        }
    }

    function updateChequeValueDeduction(
        PriceOracle oracle,
        uint256 newChequeValueDeduction
    ) external returns (bool) {
        try oracle.updateChequeValueDeduction(newChequeValueDeduction) {
            return true;
        } catch {
            return false;
        }
    }
}

contract PriceOracleEchidna {
    uint256 private constant ACTOR_COUNT = 3;

    PriceOracle private oracle;
    PriceOracleActor[3] private actors;
    address[3] private actorAddresses;
    bool private invariantFailed;

    constructor() {
        for (uint256 i = 0; i < ACTOR_COUNT; i++) {
            actors[i] = new PriceOracleActor();
            actorAddresses[i] = address(actors[i]);
        }

        oracle = actors[0].deployPriceOracle(100, 200);
    }

    function ownerUpdatePrice(uint256 newPrice) public {
        if (!actors[0].updatePrice(oracle, newPrice)) {
            invariantFailed = true;
            return;
        }

        if (oracle.price() != newPrice) {
            invariantFailed = true;
        }
    }

    function ownerUpdateChequeValueDeduction(
        uint256 newChequeValueDeduction
    ) public {
        if (
            !actors[0].updateChequeValueDeduction(
                oracle,
                newChequeValueDeduction
            )
        ) {
            invariantFailed = true;
            return;
        }

        if (oracle.chequeValueDeduction() != newChequeValueDeduction) {
            invariantFailed = true;
        }
    }

    function nonOwnerUpdatePrice(
        uint256 actorSeed,
        uint256 newPrice
    ) public {
        uint256 actorIndex = 1 + (actorSeed % (ACTOR_COUNT - 1));
        uint256 priceBefore = oracle.price();
        uint256 deductionBefore = oracle.chequeValueDeduction();

        if (actors[actorIndex].updatePrice(oracle, newPrice)) {
            invariantFailed = true;
        }

        if (oracle.price() != priceBefore) {
            invariantFailed = true;
        }
        if (oracle.chequeValueDeduction() != deductionBefore) {
            invariantFailed = true;
        }
    }

    function nonOwnerUpdateChequeValueDeduction(
        uint256 actorSeed,
        uint256 newChequeValueDeduction
    ) public {
        uint256 actorIndex = 1 + (actorSeed % (ACTOR_COUNT - 1));
        uint256 priceBefore = oracle.price();
        uint256 deductionBefore = oracle.chequeValueDeduction();

        if (
            actors[actorIndex].updateChequeValueDeduction(
                oracle,
                newChequeValueDeduction
            )
        ) {
            invariantFailed = true;
        }

        if (oracle.price() != priceBefore) {
            invariantFailed = true;
        }
        if (oracle.chequeValueDeduction() != deductionBefore) {
            invariantFailed = true;
        }
    }

    function echidna_owner_is_stable() public view returns (bool) {
        return oracle.owner() == actorAddresses[0];
    }

    function echidna_get_price_matches_storage() public view returns (bool) {
        (uint256 price, uint256 chequeValueDeduction) = oracle.getPrice();
        return
            price == oracle.price() &&
            chequeValueDeduction == oracle.chequeValueDeduction();
    }

    function echidna_no_postcondition_failures() public view returns (bool) {
        return !invariantFailed;
    }
}
