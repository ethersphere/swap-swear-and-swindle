#!/bin/sh
#
# Uninstall dependencies for abigen.
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

# Script must be run as root to uninstall dependencies.
[ "$(id -u)" -ne 0 ] && fatal "must run as root"

log "Removing solc"
rm /usr/local/bin/solc >/dev/null 2>&1 || log "solc not installed, skipping"

log "Removing abigen"
rm /usr/local/bin/abigen >/dev/null 2>&1 || log "abigen not installed, skipping"

log "done"
