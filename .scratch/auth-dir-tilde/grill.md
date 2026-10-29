# Grill record: auth-dir home shorthand (pilot, self-answered)

NOTE: answers below are the skill's recommended answers, accepted by the agent, because the operator asked not to
be interrupted for non-irreversible choices. Real use would have the human answer each round.

Facts gathered (no user needed):
- `util.ResolveAuthDir` has 12+ callers, zero tests. Callers: main.go (aborts on error), watcher (logs), management CRUD (falls back to raw value on error).
- Today `~bob/auths` is stripped to `bob/auths` and joined under the CURRENT user's home -> `$HOME/bob/auths` (silent wrong dir).
- config.example.yaml documents only: "supports ~ for home directory".

## Round 1
Q1 What should `~name/...` (tilde followed by a non-separator) do?
   Options: (a) keep literal relative path, (b) return an error, (c) resolve the other user's home via os/user.
   -> (b) error. Silent wrong-dir is the defect; (c) is non-portable and surprising for a credential dir; (a) hides typos.
Q2 Do `~`, `~/x`, `~\x` (Windows style) keep their current meaning?  -> Yes, unchanged.
Q3 Does empty string keep defaulting to DefaultAuthDir?  -> Yes, unchanged.

## Round 2 (unblocked by Q1)
Q4 Test seam?  -> ResolveAuthDir only (highest seam, ideal count = one); HOME/USERPROFILE set via t.Setenv.
Q5 Error shape?  -> wrapped, names the offending value, consistent with existing "resolve auth dir: %w" prefix.
Q6 Update config.example.yaml comment?  -> Yes, one-line clarification (user-visible contract).
Q7 ADR?  -> No: easy to reverse, not surprising -> fails the ADR bar.

Frontier empty. Shared understanding assumed (self-confirmed).
