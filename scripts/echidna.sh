#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
IMAGE="${ECHIDNA_IMAGE:-ghcr.io/crytic/echidna/echidna:latest}"
PLATFORM="${ECHIDNA_DOCKER_PLATFORM:-}"
CONFIG_FILE="${ECHIDNA_CONFIG:-echidna/echidna.yaml}"
TARGET_CONTRACT="${ECHIDNA_CONTRACT:-}"
TARGET_FILTER="${ECHIDNA_TARGET_FILTER:-}"
LOG_DIR="${ECHIDNA_LOG_DIR:-echidna/logs}"
TMP_CONFIG_DIR="$ROOT_DIR/echidna/.tmp"

if [[ -z "$PLATFORM" ]]; then
  case "$(uname -m)" in
    arm64|aarch64)
      PLATFORM="linux/amd64"
      ;;
  esac
fi

if ! command -v docker >/dev/null 2>&1; then
  echo "docker is required to run Echidna."
  exit 1
fi

if [[ ! -f "$ROOT_DIR/$CONFIG_FILE" ]]; then
  echo "Echidna config not found: $CONFIG_FILE"
  exit 1
fi

if ! docker image inspect "$IMAGE" >/dev/null 2>&1; then
  docker pull "$IMAGE"
fi

rm -rf "$ROOT_DIR/artifacts/build-info" "$ROOT_DIR/crytic-export"
yarn -s hardhat compile --force

mkdir -p "$TMP_CONFIG_DIR" "$ROOT_DIR/$LOG_DIR"
trap 'rm -rf "$TMP_CONFIG_DIR"' EXIT

shopt -s nullglob
HARNESS_FILES=( "$ROOT_DIR"/contracts/echidna/*Echidna.sol )
shopt -u nullglob

if [[ ${#HARNESS_FILES[@]} -eq 0 ]]; then
  echo "No Echidna harnesses found under contracts/echidna."
  exit 1
fi

if [[ -n "$TARGET_CONTRACT" ]]; then
  MATCHED=false
  for harness_file in "${HARNESS_FILES[@]}"; do
    if [[ "$(basename "$harness_file" .sol)" == "$TARGET_CONTRACT" ]]; then
      MATCHED=true
      break
    fi
  done

  if [[ "$MATCHED" == false ]]; then
    echo "Harness not found: $TARGET_CONTRACT"
    exit 1
  fi
fi

DOCKER_ARGS=(
  run
  --rm
  -v "$ROOT_DIR:/src"
  -w /src
)

if [[ -n "$PLATFORM" ]]; then
  DOCKER_ARGS+=(--platform "$PLATFORM")
fi

strip_config_keys() {
  # Always rewrite corpusDir. Only drop other yaml keys when an env override
  # will replace them, otherwise CI/local runs would lose testLimit/seqLen.
  local pattern="^corpusDir:"
  [[ -n "${ECHIDNA_TEST_LIMIT:-}" ]] && pattern+="|^testLimit:"
  [[ -n "${ECHIDNA_SEQ_LEN:-}" ]] && pattern+="|^seqLen:"
  [[ -n "${ECHIDNA_WORKERS:-}" ]] && pattern+="|^workers:"
  [[ -n "${ECHIDNA_MAX_TIME_DELAY:-}" ]] && pattern+="|^maxTimeDelay:"
  [[ -n "${ECHIDNA_MAX_BLOCK_DELAY:-}" ]] && pattern+="|^maxBlockDelay:"
  [[ -n "${ECHIDNA_FORMAT:-}" ]] && pattern+="|^format:"
  awk -v pat="$pattern" '$0 !~ pat'
}

append_optional_yaml() {
  local key="$1"
  local value="$2"

  if [[ -n "$value" ]]; then
    printf '%s: %s\n' "$key" "$value"
  fi
}

run_contract() {
  local contract_name="$1"
  local corpus_dir="echidna/corpus/by-contract/${contract_name}"
  local temp_config="$TMP_CONFIG_DIR/${contract_name}.yaml"
  local log_file="$ROOT_DIR/$LOG_DIR/${contract_name}.log"
  local extra_cli=""
  local config_path="${temp_config#"$ROOT_DIR/"}"
  local ec

  if [[ -n "${ECHIDNA_SEED:-}" ]]; then
    corpus_dir="${corpus_dir}/seed-${ECHIDNA_SEED}"
    extra_cli+=" --seed ${ECHIDNA_SEED}"
  fi
  if [[ -n "${ECHIDNA_TIMEOUT:-}" ]]; then
    extra_cli+=" --timeout ${ECHIDNA_TIMEOUT}"
  fi
  if [[ -n "${ECHIDNA_ARGS:-}" ]]; then
    extra_cli+=" ${ECHIDNA_ARGS}"
  fi

  mkdir -p "$ROOT_DIR/$corpus_dir"
  mkdir -p "$ROOT_DIR/echidna/out"

  strip_config_keys < "$ROOT_DIR/$CONFIG_FILE" > "$temp_config"
  {
    printf '\ncorpusDir: "%s"\n' "$corpus_dir"
    append_optional_yaml "testLimit" "${ECHIDNA_TEST_LIMIT:-}"
    append_optional_yaml "seqLen" "${ECHIDNA_SEQ_LEN:-}"
    append_optional_yaml "workers" "${ECHIDNA_WORKERS:-}"
    append_optional_yaml "maxTimeDelay" "${ECHIDNA_MAX_TIME_DELAY:-}"
    append_optional_yaml "maxBlockDelay" "${ECHIDNA_MAX_BLOCK_DELAY:-}"
    append_optional_yaml "format" "${ECHIDNA_FORMAT:-}"
  } >> "$temp_config"

  echo "==> Running ${contract_name}"
  # Capture the fuzzer exit code while still teeing logs for CI artifacts.
  set +e
  docker "${DOCKER_ARGS[@]}" "$IMAGE" /bin/bash -lc \
    "rm -rf /src/crytic-export && echidna . --contract \"$contract_name\" --config \"$config_path\"${extra_cli}" \
    2>&1 | tee "$log_file"
  ec="${PIPESTATUS[0]}"
  set -e

  if [[ "$ec" -ne 0 ]]; then
    echo "==> echidna: ${contract_name} failed (exit ${ec}); log: ${log_file}" >&2
    exit "$ec"
  fi
}

for harness_file in "${HARNESS_FILES[@]}"; do
  contract_name="$(basename "$harness_file" .sol)"

  if [[ -n "$TARGET_CONTRACT" && "$contract_name" != "$TARGET_CONTRACT" ]]; then
    continue
  fi

  if [[ -z "$TARGET_CONTRACT" && -n "$TARGET_FILTER" && "$contract_name" != *"$TARGET_FILTER"* ]]; then
    continue
  fi

  run_contract "$contract_name"
done
