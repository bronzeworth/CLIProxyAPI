Status: ready-for-agent

## Problem Statement

An operator who writes `auth-dir: "~bob/auths"` expects either bob's directory or a clear error. Today the server silently
treats it as `<current user's home>/bob/auths`, so credentials are read from and written to a directory the operator did not
intend. The function that decides this has no tests, and 12+ call sites depend on it.

## Solution

The Auth Directory resolver accepts only the Home Shorthand forms (`~`, `~/…`, `~\…`). Any other value starting with `~`
(for example `~bob/auths`) is rejected with a clear error instead of being mapped under the current user's home. Existing
behaviour for all other inputs is pinned by tests.

## User Stories

1. As an operator, I want `~bob/auths` rejected with a clear error, so that credentials never land in an unintended directory.
2. As an operator, I want `~/auths` to keep resolving under my home, so that my existing configuration keeps working.
3. As an operator, I want `~` alone to resolve to my home directory, so that the documented shorthand still works.
4. As a Windows operator, I want `~\auths` to keep resolving under my home, so that backslash configs keep working.
5. As an operator, I want an empty `auth-dir` to keep defaulting to the standard directory, so that zero-config setups work.
6. As an operator, I want absolute and relative paths returned cleaned and otherwise untouched, so that explicit paths behave predictably.
7. As a maintainer, I want the resolver covered by behaviour tests, so that its 12+ callers are protected from regressions.
8. As an operator, I want the config example to state exactly which `~` forms are supported, so that the contract is documented.

## Implementation Decisions

- The Auth Directory resolver returns a wrapped error for any `~`-prefixed value that is not Home Shorthand. The error names the offending value and keeps the existing `resolve auth dir:` prefix.
- All other inputs keep their current results (empty, `~`, `~/x`, `~\x`, absolute, relative).
- No caller changes: every caller already handles the resolver's error return.
- The config example comment is clarified to list the supported forms.

## Testing Decisions

- Test external behaviour only, through the resolver's public function; no mocking.
- One seam: the resolver. Home directory is controlled with `t.Setenv` (HOME and USERPROFILE).
- Prior art: table-driven tests in `internal/util` (e.g. `header_helpers_test.go`).

## Out of Scope

- Resolving another user's home directory.
- Environment-variable expansion (`$HOME`, `%USERPROFILE%`).
- Any change to storage backends or management endpoints.

## Further Notes

Behaviour change is intentionally breaking for the (undocumented) `~name` form only.
