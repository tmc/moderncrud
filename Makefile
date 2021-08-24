# This file contains useful targets for development and testing purposes

.PHONY: test
test:
	go test -v ./...
