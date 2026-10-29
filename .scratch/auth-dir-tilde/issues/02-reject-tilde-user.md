# 02: Reject `~name` Auth Directory values

**What to build:** An Auth Directory value starting with `~` that is not Home Shorthand (e.g. `~bob/auths`) returns a clear error instead of resolving under the current user's home. The config example comment states the supported forms.

**Blocked by:** 01 (Pin current Auth Directory resolution with behaviour tests)

**Status:** ready-for-agent

- [ ] Test first (red): `~bob/auths` and `~bob` return an error naming the value
- [ ] Tests from 01 still pass unchanged
- [ ] `config.example.yaml` auth-dir comment lists the supported `~` forms
- [ ] `go build ./cmd/server` succeeds; `gofmt -l` clean
