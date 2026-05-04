# ChefBook Backend API Gateway

The API gateway is the public HTTP entrypoint for ChefBook clients. It owns REST routing, request authentication, rate limiting, Swagger exposure, and translation from HTTP DTOs to backend gRPC calls.

## Responsibilities

- Expose the public API under `/v1`.
- Validate protected requests with JWT auth middleware.
- Fetch and cache the auth service public key for JWT validation.
- Convert HTTP request and response bodies to service gRPC messages.
- Apply recovery, request logging, and rate limiting middleware.
- Serve Swagger documentation in non-release mode at `/doc/index.html`.
- Expose health at `/healthz`.

## HTTP Route Groups

- `/auth` - sign-up, activation, sign-in, refresh, sign-out, OAuth, sessions, password flows, nickname flows.
- `/subscriptions` - subscription reads and Google subscription confirmation.
- `/profile` - current user profile and avatar management.
- `/profiles/:profileId` - another user's profile.
- `/recipes` - recipe CRUD, book, favourites, pictures, rating, translations, and recipe-to-collection binding.
- `/recipes/tags` - tag lookup inside recipe flows.
- `/collections` - collection CRUD and save/remove from recipe book.
- `/encryption/vault` - encrypted vault lifecycle.
- `/encryption/recipes/:recipeId` - recipe key ownership and sharing.
- `/shopping-lists` - personal/shared shopping lists, users, and invite links.

## Downstream Services

- `auth` for account, session, OAuth, password, nickname, public key, and auth-info RPCs.
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
