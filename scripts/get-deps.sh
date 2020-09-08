#!/bin/sh
#
# Download and install dependencies for abigen.
set -e

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

# Script must be run as root to install dependencies.
[ "$(id -u)" -ne 0 ] && fatal "must run as root"

# Try and get the versions from the environment.
GETH_VERSION="${GETH_VERSION:=1.9.20-979fc968}"
SOLIDITY_VERSION="${SOLIDITY_VERSION:=0.6.12}"

# Check all dependencies.
checkDep wget

# Get solc.
wget -q "https://github.com/ethereum/solidity/releases/download/v$SOLIDITY_VERSION/solc-static-linux" ||
  fatal "failed to download solc $SOLIDITY_VERSION"
chmod +x solc-static-linux
mv solc-static-linux /usr/local/bin/solc

# Get geth.
wget -q "https://gethstore.blob.core.windows.net/builds/geth-alltools-linux-amd64-$GETH_VERSION.tar.gz" ||
  fatal "failed to download geth $GETH_VERSION"
tar xvzf "geth-alltools-linux-amd64-$GETH_VERSION.tar.gz" > /dev/null
mv "geth-alltools-linux-amd64-$GETH_VERSION/abigen" /usr/local/bin/abigen
rm -rf "geth-alltools-linux-amd64-$GETH_VERSION"
rm "geth-alltools-linux-amd64-$GETH_VERSION.tar.gz"
