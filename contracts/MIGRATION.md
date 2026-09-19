# OpenAPI 0.3 draft consumer migration

Gateway and mobile now vendor the same SHA-256 snapshot of the explicit local
OpenAPI draft. This work changes source code; it does not update the running dev
service. The older adoption notes below describe the preceding migration.

- Auth uses `/account`, `/sessions`, `/session/tokens`, `/session`, and `/oauth/*`.
  Password login/recovery accept `login`; registration is an opaque `202` response.
- Email verification and recovery accept opaque tokens. Email links now open
  `/auth/email/confirmation?token=...` and
  `/auth/password/recovery/confirmation?token=...` on the configured frontend.
  That frontend must submit the corresponding POST and render the result before
  the browser email flow can be deployed. Mobile can submit a pasted token/link.
- Login returns `201` with Location; refresh returns `200`. Both include `userId`,
  stable `sessionId` and `restrictions`. Ordinary operations check current account
  state, while pending accounts may cancel/update deletion and manage sessions.
- Logout uses a refresh token in Authorization and returns `204`. Session deletion
  rejects legacy bodies instead of silently revoking every session.
- OAuth init sets a Secure, HttpOnly, SameSite=Lax cookie for the provider. Code/state
  completion must preserve that cookie, including when the frontend calls the API.
  Google native ID-token login does not require the browser-flow cookie.
- Identity responses use string Google IDs and numeric VK IDs. Linking returns
  `201` with Location when created and `204` for the same verified identity.
- The shared DTO renames also apply to recipe, collection, shopping-list, profile,
  and tag IDs. Former response arrays use their named contract envelopes.
- Auth body validation uses the pinned OpenAPI schemas to enforce `oneOf`, `not`,
  required fields and PATCH's prohibition on extra properties. Token/OAuth responses
  set no-store/no-cache before rate limiting, including errors. Bearer 401s include
  WWW-Authenticate.

Deploy auth's `000004_account_workflows` SQL migration and matching service, gateway,
and client together. The new HTTP shapes are incompatible with the previous `/v1`
client; no legacy route aliases are installed. Existing dev was not changed here.
Firebase bulk migration and frontend pages are separate work. Google sensitive
reauthentication requires a fresh `auth_time` and a one-use proof; native providers
that omit this claim cannot use that reauthentication variant yet.

## Previous adoption notes


The old Swagger document, live Go handlers, and Kotlin SDK did differ. OpenAPI
syntax validation alone did not expose those incompatibilities.

| Area | Previous discrepancy | Result |
| --- | --- | --- |
| Current profile | Gateway forwarded an empty lookup target | Explicit authenticated user ID |
| Public profile | Mobile used `/profile/{id}`, provider used `/profiles/{id}` | Generated `/profiles/{profile_id}` call and separate handler |
| Recipe book | Mobile used `/save`, provider used `/book` | Generated book operations |
| Favourites | Mobile used singular `/favourite`, provider used `/favourites` | Generated plural routes |
| Recipe collections | Mobile used `/categories`, provider used `/collections` | Generated collections call |
| Recipe search | Client sent `tag`, provider expected `tags`; Swagger described a GET body | Repeated `tags` query parameters and explicit search query fields |
| Joining shopping lists | Client accepted a key but never transmitted it | Generated required JSON body containing `key` |
| Vault deletion | Client used DELETE `/vault/delete`, provider used DELETE `/vault` | Correct route and generated `deleteCode` body |
| Session termination | Swagger omitted the DELETE array body | Generated JSON array of session IDs |
| Session response | Client required `isMobile`, server emitted `mobile` | Client serializer reads `mobile` |
| Recipe owner | Client required nested `owner`, server emitted `ownerId` plus `profilesInfo` | Owner resolved from the response envelope |
| Swagger response shapes | Profile, avatar, collections, shopping-list creation, encryption keys were incorrectly described | Contract follows concrete provider DTOs |
| Missing request shapes | Collection save, recipe rating/collections and shopping-list naming were absent or wrong | Explicit request schemas |
| Translation deletion | Swagger omitted the language path segment | Generated `/translations/{language_code}` route |
| Tag details | Embedded Go tag fields need flattening in the schema | Full tag response fields preserved |
| Optional response fields | SDK required `owned`, `encrypted`, Google OAuth ID, deletion status fields, and recipe item text | Defaults/nullability follow actual server responses |
| Upload form fields | Go maps could encode `null` | Upload handlers normalize absent form data to `{}` |
| Empty shopping-list users | Go could encode a nil list as `null` | Empty response is an array |

The unused mobile API for deleting an owner's recipe key had no HTTP endpoint.
The service explicitly forbids owner-key deletion through its user-access removal
operation. The unused remote capability and its unused SDK interface method were
removed instead of inventing a destructive endpoint. Existing key cleanup belongs
to the service lifecycle.

All 83 public operations now come from the shared contract. Request/response
mapping to domain models still belongs to each consumer; generated models do not
replace local storage models. Optional legacy client fields such as `broccoins`
and collection `emoji` remain readable locally but are not HTTP guarantees.

Checks include Go DTO/schema parity, all protected routes, current/public profile
targets, Kotlin host tests for transport serialization and token refresh, and
37 recursive SDK/generated response comparisons across 11 modules, and
compilation of all affected SDK modules. No deployment or live integration test
is performed by this migration.
