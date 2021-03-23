#!/usr/bin/env sh
yarn compile

TARGET="$1"
VERSION="$(echo $2 | tr . _)"

SIMPLESWAP_ABI="$(cat artifacts/contracts/ERC20SimpleSwap.sol/ERC20SimpleSwap.json | jq .abi)"
SIMPLESWAP_FACTORY_ABI="$(cat artifacts/contracts/SimpleSwapFactory.sol/SimpleSwapFactory.json | jq .abi)"
SIMPLESWAP_FACTORY_BYTECODE="$(cat artifacts/contracts/SimpleSwapFactory.sol/SimpleSwapFactory.json | jq .bytecode)"
ERC20_ABI="$(cat artifacts/@openzeppelin/contracts/token/ERC20/ERC20.sol/ERC20.json | jq .abi)"

cat <<EOF > "$TARGET"
// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sw3abi

const ERC20SimpleSwapABI$VERSION = \`$SIMPLESWAP_ABI\`

const SimpleSwapFactoryABI$VERSION = \`$SIMPLESWAP_FACTORY_ABI\`

const ERC20ABI$VERSION = \`$ERC20_ABI\`

const SimpleSwapFactoryBin$VERSION = $SIMPLESWAP_FACTORY_BYTECODE
EOF