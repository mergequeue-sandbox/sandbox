# mergequeue sandbox

Throwaway repo for prototyping mergequeue against.
Nothing here matters except the CI behaviour:

| File | Effect |
| --- | --- |
| `ci/status` | `tests` job outcome: `PASS`, `FAIL`, or `FLAKY` (fails on first attempt only) |
| `ci/sleep` | seconds the `tests` job takes (default 60) |
| `ci/lint` | `FAIL` makes the `lint` job fail |
| `shared.txt` | edit the same line in two PRs to get a conflict |

The `label-check` job (PRs only) fails unless the PR has the `ready` label.

`develop` is protected like projectsapigo's `develop`: required checks `lint`,
`tests`, `label-check`, one approval, branches must be up to date.
