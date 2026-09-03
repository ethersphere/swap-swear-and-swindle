## Echidna Harnesses

This repo uses stateful Echidna harnesses under `contracts/echidna/` to fuzz the swap, factory, and oracle flows with multiple actors.

### Harness layout

- `ERC20SimpleSwapEchidna.sol`: multi-actor swap accounting, hard deposits, cheque cashing, auth failures, and post-conditions.
- `ERC20SimpleSwapSignatureEchidna.sol`: deterministic signature fixture checks for valid and invalid cheque signatures.
- `SimpleSwapFactoryEchidna.sol`: deterministic clone deployment, uniqueness, and configuration invariants.
- `SimpleSwapFactorySystemEchidna.sol`: cross-contract factory plus tracked clone lifecycle with deposits, hard deposits, cashing, and re-init resistance.
- `PriceOracleEchidna.sol`: owner-only oracle updates and storage consistency.

### Harness design notes

- Actor wrappers isolate roles so Echidna can explore issuer, beneficiary, caller, and non-owner behavior through distinct `msg.sender` values.
- Action functions are written to absorb expected reverts through `try/catch`, while `echidna_*` properties and explicit post-condition checks stay strict.
- Timeout-heavy flows are intentionally bounded to a small domain so stateful campaigns can reach `prepareDecreaseHardDeposit` and `decreaseHardDeposit` paths in a single run.
- The system harness tracks clone metadata and token balances so cross-contract invariants catch configuration drift and accounting leaks.

### Runner behavior

Run the full suite with:

```sh
yarn echidna
```

Run one harness:

```sh
ECHIDNA_CONTRACT=ERC20SimpleSwapEchidna yarn echidna
```

Fast smoke campaign:

```sh
ECHIDNA_TEST_LIMIT=5000 ECHIDNA_SEQ_LEN=120 yarn echidna
```

Reproducible / time-boxed run:

```sh
ECHIDNA_SEED=1 ECHIDNA_TIMEOUT=600 yarn echidna
```

Per-harness corpora are stored under `echidna/corpus/by-contract/<HarnessName>/`. Logs go to `echidna/logs/` (gitignored).

| Setting | Default | Override env |
|---------|---------|----------------|
| `testLimit` | 60000 | `ECHIDNA_TEST_LIMIT` |
| `seqLen` | 320 | `ECHIDNA_SEQ_LEN` |
| workers | yaml | `ECHIDNA_WORKERS` |
| seed | random | `ECHIDNA_SEED` |
| timeout (seconds) | none | `ECHIDNA_TIMEOUT` |

## CI

Workflow: [`.github/workflows/echidna.yml`](../.github/workflows/echidna.yml). Runs on every pull request (and manual `workflow_dispatch`).

One matrix job per harness, using the same campaign as local `yarn echidna` (`echidna/echidna.yaml`: `testLimit` 60000, `seqLen` 320). Each job has a 180-minute cap and Echidna `--timeout` 10200s so the fuzzer stops cleanly. On failure the workflow uploads `echidna/logs/`, corpus reproducers under `echidna/corpus/by-contract/`, and `crytic-export/` as artifacts.

Reproduce a CI counterexample locally:

```sh
ECHIDNA_CONTRACT=ERC20SimpleSwapEchidna \
yarn echidna
```
