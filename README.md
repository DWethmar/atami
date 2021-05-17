# Atami - Messaging platform

![Tests](https://github.com/DWethmar/atami/workflows/Tests/badge.svg)

## setup

- `go mod download`
- `make serve`

currently implementing: https://github.com/thecodenation/clean-architecture-go-v2

## TODO

- [ ] Generate code to std out: https://www.calhoun.io/using-code-generation-to-survive-without-generics-in-go/
- [x] Migrate sql queries to pkg/postgres (maybe rename to querybuilder)
- [x] Refresh token
- [ ] Message endpoint - WIP
- [x] Rename CreatedOn in message to CreatedAt
- [ ] Pagination + Default result object
- [ ] Etag
- [x] Api versioning: https://restfulapi.net/versioning/
- [ ] Full text search for messages https://www.postgresql.org/docs/9.5/textsearch.html
- [ ] Move create_at, updated_at and UID(?) set logic to service.
- [ ] Verify email for users
- [x] Add user to context
- [ ] check if test with db actually works by running it locally
- [ ] disable http log when testing
- [ ] Keep track of sessions?
  - max 100 - sessions and delete older ones if exceeding the max.
- [ ] maybe check how many logins are done for a single user.
- [ ] Cloud Vision API

merge packages
pkg/api/beta/router
pkg/api/beta/handler
