# Atami - Messaging platform


## setup
- `go mod download`
- `make serve`

## TODO

- [ ] Message endpoint - WIP
- [ ] Rename CreatedOn in message to CreatedAt
- [ ] Full text search for messages https://www.postgresql.org/docs/9.5/textsearch.html
- [ ] Move create_at, updated_at  and ID(?) set logic to service. 
    - Later we can create a service that can insert data with old dates for testing.
- [ ] create Channel? Or other name
- [ ] Verify email for users
- [x] Add user to context
- [ ] check if test with db actually works by running it locally 
- [ ] disable http log when testing
- [ ] Keep track of sessions?
    - max 100 - sessions and delete older ones if exceeding the max.
- [ ] maybe check how many logins are done for a single user.