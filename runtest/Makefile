# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=invoices
BINARY_UNIX=$(BINARY_NAME)_unix
MAINPATH=cmd/invoices/main.go

all:build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAINPATH)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAINPATH) 
	./$(BINARY_NAME)