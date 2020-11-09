.PHONY: all

all: server

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GORUN=$(GOCMD) run
BIN=bin

server: 
	@echo build server
	$(GOBUILD) -v -o $(BIN)/atami-server ./cmd/server/

serve: 
	@echo serve server
	ACCESS_SECRET=abcdefghijlk $(GORUN) ./cmd/server/

test:
	@echo running tests
	$(GOTEST) ./...
