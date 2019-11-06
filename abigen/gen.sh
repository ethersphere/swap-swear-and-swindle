#!/usr/bin/env sh
set -e

# check if all tools are available
for tool in node solc abigen
do
  if ! which $tool > /dev/null
  then
    echo "$tool not in PATH" >&2
    exit 1
  fi
done

for CONTRACT in "$@"
do
  # go package name
  PACKAGE=$(echo $CONTRACT | tr '[:upper:]' '[:lower:]')
  # temporary file for compiler output
  COMPILED_JSON=compiled.json
  # output directory
  OUTPUT=bindings/$PACKAGE

  # compile the contract allowing imports from openzeppelin-solidity
  solc \
    openzeppelin-solidity=$(pwd)/node_modules/openzeppelin-solidity\
    --allow-paths node_modules/openzeppelin-solidity/contracts\
    --combined-json=bin,abi,userdoc,devdoc,metadata,bin-runtime\
    contracts/$CONTRACT.sol > "$COMPILED_JSON"

  # generate the bindings
  mkdir -p "$OUTPUT"
  abigen -pkg $PACKAGE -out "$OUTPUT/$PACKAGE.go" --combined-json "$COMPILED_JSON"
  # this creates a separate file for the runtime binary which is not included by abigen
  node abigen/code.go.js $PACKAGE "$COMPILED_JSON" $CONTRACT > "$OUTPUT/code.go"

  # clean up temporary file for compiler output
  rm "$COMPILED_JSON"

  echo "generated go bindings for $CONTRACT in $OUTPUT"
done