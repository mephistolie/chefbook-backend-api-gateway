# Public API contract snapshot

`chefbook.yaml` is vendored from the separate `chefbook-contracts` repository.
Its source commit and SHA-256 are recorded in `chefbook.lock.json`.
A local draft additionally records `workingTree: true`; the commit then identifies
the base revision rather than claiming the uncommitted YAML is published.
Edit the source repository, not this copy.

From the gateway root:

```sh
python3 scripts/sync_contract.py --source ../../chefbook-contracts --ref <commit-or-tag>
python3 scripts/sync_contract.py --check
```

By default an update reads committed content only. To test an explicitly selected
local draft before publishing it, use `--source ../../chefbook-contracts --working-tree`
instead of `--ref`. Normal builds never consult the source checkout or network
for the OpenAPI snapshot. For a standalone checkout,
clone the contracts repository separately and pass its local path explicitly.

The gateway registers its public routes with the generated Gin interface in
`internal/transport/http/contract/api.gen.go`. Explicit adapters in
`router/v1/contract_adapter.go` implement that interface. Auth handlers use the typed
AuthenticationService gRPC API; the other domains retain their service handlers. Adding/removing/changing an operation produces a compile-time
adapter obligation. The generated security metadata drives JWT enforcement;
account deletion and session management permit pending-deletion accounts.
Public signIn/signUp processes use an opaque Flow-Token after creation; sensitive
purposes additionally require the original access bearer. The auth service enforces
this conditional AND binding. Sensitive actions accept Reauthentication-Token.
Logout accepts an access bearer for own sessions or a refresh bearer for that exact
path session. Every protected request checks the live session and account with auth;
a revoked session cannot continue using a still-unexpired JWT. Legacy JWTs without
sid must sign in again.

```sh
python3 scripts/generate_contract.py
python3 scripts/generate_contract.py --check
go test ./...
```

Generation is pinned to oapi-codegen 2.4.1. Generated Go source is committed, so
ordinary builds do not need the generator. The legacy `generate_doc.sh` command
now forwards to contract generation; Swagger annotations and generated Swagger
2.0 files were removed. `/openapi.yaml` serves the exact embedded vendored file,
and `/docs` renders it in Scalar in non-production mode.
The Scalar bundle is pinned and served by the gateway; browser requests do not use
an external API proxy. See the gateway README for asset verification and updates.

Existing Go DTOs keep their gRPC conversion functions. The test in
`contracts/schema_test.go` compares all provider DTO JSON fields with the shared
schema, catching accidental drift during ordinary `go test` runs. Router tests
verify protection of every authenticated operation and profile-deletion policy.
See [MIGRATION.md](MIGRATION.md) for the discrepancies reconciled during adoption.

The auth gRPC API is updated in the sibling `../services/auth/api` module. The
checked-in `go.work` selects that local API for `go test ./...` and local builds.
Before a standalone gateway/container release, publish and pin that API revision,
or use a backend workspace build context containing the sibling API module.
