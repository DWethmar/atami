FROM golang:1.15-alpine AS build

RUN apk add git gcc g++ make

WORKDIR $GOPATH/src/github.com/dwethmar/atami

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod go.sum ./

RUN go mod verify 

COPY . .

# Build the Go app
RUN make server

# Start fresh from a smaller image
FROM alpine:3.9 
RUN apk add ca-certificates

COPY --from=build /go/src/github.com/dwethmar/atami/bin/server /app/server/

COPY ./migrations/ /app/migrations/

EXPOSE 8080

CMD ["/app/server/server"]
