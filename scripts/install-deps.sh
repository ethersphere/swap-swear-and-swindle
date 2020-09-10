#!/bin/sh
#
# Download and install dependencies for abigen.
set -e

# Print a log-formatted message to STDOUT.
# Arguments:
#   Message to print
log() {
  echo "[$(date +'%Y-%m-%dT%H:%M:%S%z')]: $*"
}

# Print a log-formatted error to STDERR, then quit.
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
gethVersion="${GETH_VERSION:=1.9.20-979fc968}"
solidityVersion="${SOLIDITY_VERSION:=0.6.12}"

# Check all system-level dependencies.
checkDep wget
checkDep jq
checkDep envsubst
checkDep mktemp

log "installing solc v$solidityVersion"

wget -q "https://github.com/ethereum/solidity/releases/download/v$solidityVersion/solc-static-linux" ||
  fatal "failed to download solc $solidityVersion"
chmod +x solc-static-linux > /dev/null
mv solc-static-linux /usr/local/bin/solc


log "installing abigen from go-ethereum v$gethVersion"

wget -q "https://gethstore.blob.core.windows.net/builds/geth-alltools-linux-amd64-$gethVersion.tar.gz" ||
  fatal "failed to download geth $gethVersion"
tar xvzf "geth-alltools-linux-amd64-$gethVersion.tar.gz" > /dev/null
mv "geth-alltools-linux-amd64-$gethVersion/abigen" /usr/local/bin/abigen
chmod +x /usr/local/bin/abigen > /dev/null
rm -rf "geth-alltools-linux-amd64-$gethVersion"
rm "geth-alltools-linux-amd64-$gethVersion.tar.gz"

log "done"
