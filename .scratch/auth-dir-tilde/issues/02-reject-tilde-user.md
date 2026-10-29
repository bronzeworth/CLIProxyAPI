# 02: Reject `~name` Auth Directory values

**What to build:** An Auth Directory value starting with `~` that is not Home Shorthand (e.g. `~bob/auths`) returns a clear error instead of resolving under the current user's home. The config example comment states the supported forms.

**Blocked by:** 01 (Pin current Auth Directory resolution with behaviour tests)

**Status:** ready-for-agent

- [x] Test first (red): `~bob/auths` and `~bob` return an error naming the value
- [x] Tests from 01 still pass unchanged
- [x] `config.example.yaml` auth-dir comment lists the supported `~` forms
- [x] `go build ./cmd/server` succeeds; `gofmt -l` clean
