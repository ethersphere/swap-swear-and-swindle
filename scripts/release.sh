#!/bin/sh
#
# Create a new release on the sw3-bindings repo.
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

# Get script directory.
scriptDir="$( cd "$( dirname "$0" )" >/dev/null 2>&1 && pwd )"
rootDir=$( dirname "$scriptDir" )

# Ensure that dependencies are installed.
checkDep git
checkDep go

# Print header information.
log "creating new release"
log "$(go version)"
log "git user: $(git config user.name) <$(git config user.email)>"

buildDir="$rootDir/build"
bindingsDir="$buildDir/bindings"
repoDir="$buildDir/sw3-bindings"

mkdir -p "$buildDir"
cd "$buildDir"

# Ensure bindings dir exists.
[ -d "$bindingsDir" ] || fatal "$bindingsDir not found"

# Ensure we start with a clean repo.
rm -rf "$repoDir"

# Clone the repo.
git clone "git@github.com:ethersphere/sw3-bindings.git"
cd "$repoDir"

# Count the number of "v.." dirs to get the major version.
majorVersion=$( ls -p | grep "v[0-9]*/" | wc -l )
majorVersion=$((majorVersion + 1))

# Create a new ../v<MAJOR_VERSION> directory as a copy of the abigen bindings.
newVersionDir="$repoDir/v$majorVersion"
cp -r "$bindingsDir" "$newVersionDir"

# Run go mod init in the new module.
cd "$newVersionDir"
go mod init "github.com/ethersphere/sw3-bindings/v$majorVersion"

# Run go mod tidy at the top leel
cd "$repoDir"
go mod tidy

# Commit the changes (new version and version file).
git add "$newVersionDir"
git commit --message="Release v$majorVersion.0.0"

# Tag the new release and push to remote.
git tag "v$majorVersion.0.0"
git push --tags origin master

# Clean up the git directory.
rm -rf "$repoDir"

log "Sucessfully pushed v$majorVersion"
