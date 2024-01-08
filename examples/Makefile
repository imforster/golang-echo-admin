# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Name of the binary to be built
BINARY_NAME=adminTest

all: test build

build:
	$(GOBUILD) -ldflags "-X main.Version=`git describe --tags --abbrev=0`" -o $(BINARY_NAME) *.go

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
