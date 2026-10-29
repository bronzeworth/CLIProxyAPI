# 01: Pin current Auth Directory resolution with behaviour tests

**What to build:** Behaviour tests that lock in how the resolver treats empty, `~`, `~/x`, `~\x`, absolute and relative values, so the follow-up change is made against a safety net. No production change.

**Blocked by:** None (can start immediately)

**Status:** ready-for-agent

- [ ] Tests cover: empty -> default dir under home, `~`, `~/x`, `~\x`, absolute path, relative path, path needing cleaning
- [ ] Tests run hermetically (home dir set via env, no real home access)
- [ ] `go test ./internal/util/...` passes with no production code changed
