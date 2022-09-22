
BIN_DIR=bin
LIN_DIR=$(BIN_DIR)/lin
WIN_DIR=$(BIN_DIR)/win
BIN=haipeaghr7

run:
	@sudo go run ./...

build: build-lin build-win

build-lin:
	@echo '--> Building linux binary...'
	@GOOS=linux go build ./... -o $(LIN_DIR)/$(BIN)

build-win:
	@echo '--> Building windows binary...'
	@GOOS=windows go build ./... -o $(WIN_DIR)/$(BIN)

.PHONY: build build-lin build-win
