#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
IMAGE="${ECHIDNA_IMAGE:-ghcr.io/crytic/echidna/echidna:latest}"
PLATFORM="${ECHIDNA_DOCKER_PLATFORM:-}"
TARGET_FILTER="${ECHIDNA_TARGET_FILTER:-}"
CONFIG_FILE="${ECHIDNA_CONFIG:-echidna/echidna.yaml}"
TMP_CONFIG_DIR="$ROOT_DIR/echidna/.tmp"

if [[ -z "$PLATFORM" ]]; then
  case "$(uname -m)" in
    arm64|aarch64)
      PLATFORM="linux/amd64"
      ;;
  esac
fi

TARGETS=(
  "ERC20SimpleSwapEchidna"
  "SimpleSwapFactoryEchidna"
  "PriceOracleEchidna"
)

if ! command -v docker >/dev/null 2>&1; then
  echo "docker is required to run Echidna."
  exit 1
fi

if ! docker image inspect "$IMAGE" >/dev/null 2>&1; then
  docker pull "$IMAGE"
fi

yarn hardhat compile
mkdir -p "$TMP_CONFIG_DIR"
trap 'rm -rf "$TMP_CONFIG_DIR"' EXIT

DOCKER_ARGS=(
  run
  --rm
  -v "$ROOT_DIR:/src"
  -w /src
)

if [[ -n "$PLATFORM" ]]; then
  DOCKER_ARGS+=(--platform "$PLATFORM")
fi

for target in "${TARGETS[@]}"; do
  contract_name="$target"

  if [[ -n "$TARGET_FILTER" && "$contract_name" != *"$TARGET_FILTER"* ]]; then
    continue
  fi

  corpus_dir="echidna/corpus/${contract_name}"
  temp_config="$TMP_CONFIG_DIR/${contract_name}.yaml"

  awk '!/^corpusDir:/' "$CONFIG_FILE" > "$temp_config"
  printf '\ncorpusDir: %s\n' "$corpus_dir" >> "$temp_config"

  echo "==> Running ${contract_name}"
  docker "${DOCKER_ARGS[@]}" "$IMAGE" \
    echidna . --contract "$contract_name" --config "${temp_config#"$ROOT_DIR/"}"
done
