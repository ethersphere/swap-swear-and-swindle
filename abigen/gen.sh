#!/usr/bin/env sh
set -e

COMPILED_JSON=compiled.json
OUTPUT=bindings/SimpleSwap

solc \
  openzeppelin-solidity=$(pwd)/node_modules/openzeppelin-solidity\
  --allow-paths node_modules/openzeppelin-solidity/contracts\
  --combined-json=bin,abi,userdoc,devdoc,metadata,bin-runtime\
  contracts/SimpleSwap.sol > "$COMPILED_JSON"

mkdir -p "$OUTPUT"
abigen -pkg contract -out "$OUTPUT/simpleswap.go" --combined-json "$COMPILED_JSON"
node abigen/code.go.js contract "$COMPILED_JSON" SimpleSwap > "$OUTPUT/code.go"

rm "$COMPILED_JSON"