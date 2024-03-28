#!/usr/bin/env sh
yarn compile

TARGET="$1"
VERSION="$(echo $2 | tr . _)"

PRICE_ORACLE_ABI="$(cat artifacts/src/PriceOracle.sol/PriceOracle.json | jq .abi)"
PRICE_ORACLE_BYTECODE="$(cat artifacts/src/PriceOracle.sol/PriceOracle.json | jq .bytecode)"
PRICE_ORACLE_DEPLOYED_BYTECODE="$(cat artifacts/src/PriceOracle.sol/PriceOracle.json | jq .deployedBytecode)"

cat <<EOF > "$TARGET"
// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package priceoracleabi

const PriceOracleABI$VERSION = \`$PRICE_ORACLE_ABI\`
const PriceOracleBin$VERSION = $PRICE_ORACLE_BYTECODE
const PriceOracleDeployedBin$VERSION = $PRICE_ORACLE_DEPLOYED_BYTECODE
EOF