# 03: Vertex import stops on an unresolvable Auth Directory

**What to build:** `DoVertexImport` no longer continues with an unresolved `~bob/...` Auth Directory (which would write credentials to a literal `./~bob/...` directory). It logs the error and returns, like startup does. Also pin the `~~` and `~C:\x` rejections in the resolver tests.

**Blocked by:** 02 (Reject `~name` Auth Directory values)

**Status:** ready-for-agent

- [x] With `cfg.AuthDir = "~bob/auths"`, importing a key writes nothing and the function returns after logging
- [x] Resolver tests also reject `~~` and `~C:\x`
- [x] Existing tests in `internal/util` and `internal/cmd` pass; `go build ./cmd/server` succeeds; `gofmt -l` clean

## Follow-up (not built, needs an owner decision)
Hot reload (`internal/watcher/config_reload.go`): a bad `auth-dir` edit is logged but the raw value is kept, so file clients load as zero until fixed. Same outcome as before this change (the old wrong dir also did not exist) but now loud. Option: keep the previous AuthDir on error.
