# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=test.exe
BINARY_UNIX=$(BINARY_NAME)_unix
MAINPATH=cmd/invoices/main.go

all:build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(M).go
run:
	$(GOBUILD) -o $(BINARY_NAME) -v $(M).go
	./$(BINARY_NAME)