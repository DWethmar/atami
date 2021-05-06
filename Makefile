.PHONY: all

all: server

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GORUN=$(GOCMD) run
GOGEN=$(GOCMD) generate
BIN=bin

server: 
	@echo build server
	$(GOBUILD) -v -o $(BIN)/server ./cmd/server/

generate:
	@echo generating
	$(GOGEN) ./...

serve: 
	@echo serve server
	ACCESS_SECRET=abcdefghijlk $(GORUN) ./cmd/server/ -in-memory=1

test:
	@echo running tests
	ACCESS_SECRET=abcdefghijlk $(GOTEST) ./... 

test-with-db:
	@echo running tests with database
	./run-tests-with-db.sh
