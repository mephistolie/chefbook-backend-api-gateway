# ChefBook Backend API Gateway

The API gateway is the public HTTP entrypoint for ChefBook clients. It owns REST routing, request authentication, rate limiting, OpenAPI contract integration, and translation from HTTP DTOs to backend gRPC calls.

## Responsibilities

- Expose the public API under `/v1`.
- Validate protected requests with JWT auth middleware.
- Fetch and cache the auth service public key for JWT validation.
- Convert HTTP request and response bodies to service gRPC messages.
- Apply recovery, request logging, and rate limiting middleware.
- Generate public routes from the pinned OpenAPI contract; serve it at `/openapi.yaml` and render it with Scalar at `/docs` in non-release mode.
- Expose health at `/healthz`.

## HTTP Route Groups

- `/authentications` - server-owned multi-step sign-in, registration and reauthentication.

- `/sessions`, `/account`, `/oauth`, `/usernames` - sessions, account security, factor management and provider linking.
- `/subscriptions` - subscription reads and Google subscription confirmation.
- `/profile` - current user profile and avatar management.
- `/profiles/:profileId` - another user's profile.
- `/recipes` - recipe CRUD, book, favourites, pictures, rating, translations, and recipe-to-collection binding.
- `/recipes/tags` - tag lookup inside recipe flows.
- `/collections` - collection CRUD and save/remove from recipe book.
- `/encryption/vault` - encrypted vault lifecycle.
- `/encryption/recipes/:recipeId` - recipe key ownership and sharing.
- `/shopping-lists` - personal/shared shopping lists, users, and invite links.

Authentication starts with a purpose only. Typed `/steps` collect registration
email, verify it, and then set up the first credential. The flat `oneOf` body of
`/steps/{stepId}/completion` is validated against the pinned OpenAPI before an
RPC is made. Username selection happens after registration. State includes
`steps` and `next.options`; the auth service owns prerequisites and finalization.
OAuth setup URLs/nonces are returned only when creating their step, flow tokens
only when starting the process, and confirmation grants only by mutations.

## Downstream Services

- `auth` for account, session, OAuth, password, username, public key, and auth-info RPCs.
- `user` for social profile fields and avatar lifecycle RPCs.
- `profile` for aggregated profile read models.
- `tag` for tag and tag-group lookup.
- `recipe` for recipe, collection, translation, picture, rating, and recipe policy RPCs.
- `encryption` for vault and recipe key RPCs.
- `shopping-list` for shopping list and membership RPCs.
- `subscription` for subscription reads and Google subscription confirmation.

## Change Guidance

- Change this module when public HTTP shape, middleware, request DTOs, response DTOs, or REST-to-gRPC mapping changes.
- Change the owning service when business rules, persistence, or gRPC contract behavior changes.
- If a gRPC contract changes, update this gateway and the provider service together.

## Public contract

See [contracts/README.md](contracts/README.md) for pinning, generation, and checks,
and [contracts/MIGRATION.md](contracts/MIGRATION.md) for the reconciled client/server mismatches.

## API reference UI

Scalar 1.69.0 is embedded in the gateway as a compressed standalone bundle. Its
JavaScript, initialization and OpenAPI document are served from this gateway;
API requests go directly to the selected API origin, without a Scalar proxy.
Remote fonts, auth persistence and telemetry are disabled. The versioned script
can be cached; HTML and initialization are revalidated. Documentation remains
disabled in production mode.

The npm archive integrity and JavaScript checksums are pinned under
`internal/transport/http/router/docs/assets/scalar.lock.json`, with the MIT license.
Verify offline with `python3 scripts/vendor_scalar.py --check`. To reproduce the
bundle, run `python3 scripts/vendor_scalar.py` (or supply `--archive <npm-tarball>`).
Ordinary Go builds embed the vendored asset and do not download frontend packages.

Auth operation IDs follow the agreed table in the contracts repository's
`AUTH_OPERATION_NAMES.md`. Renaming these generated methods does not change the
HTTP paths, bodies, status codes or internal service RPC names.
