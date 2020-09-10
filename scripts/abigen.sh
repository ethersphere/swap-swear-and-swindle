#!/bin/sh
#
# Generate the go-ethereum bind ABI code for all contracts.
set -e

# Print a log-formatted message to STDOUT.
# Arguments:
#   Message to print
log() {
  echo "[$(date +'%Y-%m-%dT%H:%M:%S%z')]: $*"
}

# Print a log-formatted error to the STDERR, then quit.
# Arguments:
#   Error to print.
fatal() {
  echo "[$(date +'%Y-%m-%dT%H:%M:%S%z')]: $*" >&2
  exit 1
}

# Check if a dependency is missing form PATH.
# Arguments:
#   Name of dependency
checkDep() {
  command -v "$1" >/dev/null 2>&1 || fatal "dependency $1 not found"
}

# Cleanup temp dir on exit.
trap 'rm -r $tempDir' EXIT

# Check all dependencies.
checkDep mktemp
checkDep solc
checkDep abigen
checkDep jq

# Get script directory.
scriptDir="$( cd "$( dirname "$0" )" >/dev/null 2>&1 && pwd )"
rootDir=$( dirname "$scriptDir" )

# Ensure the build directory exists.
mkdir -p "$rootDir/build"

# Create a temp compilation directory.
tempDir="$( mktemp -d )" || fatal "Failed to make temp dir"

# Create the outputDir.
outDir="$rootDir/build/gosw3"

# Locate the code emplate file.
templateFile="$rootDir/templates/code-template.txt"
[ -f "$templateFile" ] || fatal "$templateFile not found"

# Parse all contracts.
for contractName in "$@"; do
  contractPath="contracts/$contractName.sol"

  log "processing $contractPath"

  # Ensure contract exists.
  [ -f "$contractPath" ] || fatal "$contractPath" not found

  # Define the package name for the contract (all lower case).
  package=$(echo "$contractName" | tr '[:upper:]' '[:lower:]')

  # Compile the contract allowing imports from openzeppelin-solidity.
  solc \
    @openzeppelin="$rootDir/node_modules/@openzeppelin" \
    --allow-paths node_modules/@openzeppelin/contracts/ \
    --combined-json=bin,abi,userdoc,devdoc,metadata,bin-runtime \
    --optimize --optimize-runs 200 --evm-version istanbul \
    "$contractPath" > "$tempDir/$contractName.json"

  # Generate the abigen bind code.
  mkdir -p "$outDir/$package"
  abigen -pkg "$package" \
    -out "$outDir/$package/$package.go" \
    --combined-json "$tempDir/$contractName.json"

  # Extract the binRuntime value from the compiled contract json.
  jqSelector=".contracts | .\"contracts/$contractName.sol:$contractName\" | .\"bin-runtime\""
  binRuntime=$( jq -r "$jqSelector" < "$tempDir/$contractName.json" )

  # Substitute the templates using envsubst.
  packageName="$package" contractName="$contractName" binRuntime="$binRuntime" \
    envsubst < "$templateFile" \
    > "$outDir/$package/code.go"

done
