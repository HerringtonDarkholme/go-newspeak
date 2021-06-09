GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

test_unit:
	$(GOTEST) -race -v ./...

test: test_unit
