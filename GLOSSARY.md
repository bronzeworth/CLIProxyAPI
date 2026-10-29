# CLIProxyAPI

Multi-provider API proxy that load-balances requests across OAuth and API-key credentials.

## Language

**Auth Directory**:
The directory that holds credential files, configured as `auth-dir`.
_Avoid_: auth folder, token dir

**Home Shorthand**:
A leading `~` in an Auth Directory value, meaning the current user's home directory. Only `~`, `~/…` and `~\…` are Home Shorthand.
_Avoid_: tilde expansion, `~user` form
