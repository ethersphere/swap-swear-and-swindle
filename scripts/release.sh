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

# Read env variables.
gethVersion="${GETH_VERSION:=1.9.20-979fc968}"

# Get script directory.
scriptDir="$( cd "$( dirname "$0" )" >/dev/null 2>&1 && pwd )"
rootDir=$( dirname "$scriptDir" )

# Ensure that dependencies are installed.
checkDep git
checkDep go

# Format gethVersion semver string for go get.
# We assume the version MAY contain a build number after "-".
gethVersion="v$( echo $gethVersion | cut -d "-" -f1 )"


# Print header information.
log "creating new release"
log "$(go version)"
log "geth version: $gethVersion"
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

# Determine major package version from tag.
tag=$( git describe --tags --always | cut -d "." -f 1 )

# By convention, if there is no git tag, version is 0.
majorVersion=0

# If we've found a git version tag (starting with "v"), parse the major version.
[ "$( echo "$tag" | cut -c1 )" = "v" ] \
  && majorVersion=$( echo "$tag" | cut -d "v" -f 2 )

# Increment the major version.
majorVersion=$((majorVersion + 1))

log "incrementing bindings version to v$majorVersion"

# Verstion 1 is a special case.
if [ "$majorVersion" -eq 1 ]; then
  # Copy the code into the root of the repo.
  cp -R "$bindingsDir"/* "$repoDir"

  log "fetching 'go-ethereum@$gethVersion' on base module"
  GO111MODULE=on go get "github.com/ethereum/go-ethereum@$gethVersion"
else
  # Create a new ../v<MAJOR_VERSION> directory as a copy of the abigen bindings.
  newVersionDir="$repoDir/v$majorVersion"
  cp -R "$bindingsDir" "$newVersionDir"

  # Run go mod init in the new module.
  cd "$newVersionDir"

  log "init go module in $newVersionDir"
  go mod init "github.com/ethersphere/sw3-bindings/v$majorVersion"

  log "fetching 'go-ethereum@$gethVersion' on base $newVersionDir"
  GO111MODULE=on go get "github.com/ethereum/go-ethereum@$gethVersion"

  go mod tidy
fi

# Run go mod tidy at the top leel
cd "$repoDir"
go mod tidy

git add . 
git commit --message="Release v$majorVersion.0.0"
git tag "v$majorVersion.0.0"
git push --tags origin master

# Clean up the git directory.
rm -rf "$repoDir"

log "Sucessfully pushed v$majorVersion"
