# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Name of the binary to be built
BINARY_NAME=myapp

all: test build

build:
	$(GOBUILD) -ldflags "-X main.Version=`git describe --tags --abbrev=0`" -o $(BINARY_NAME) main.go version.go

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
