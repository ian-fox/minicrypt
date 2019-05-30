shell := /bin/bash

.PHONY: build install test lint clean

lint:
	revive ./...

test:
	go test ./...

build:
	go build

install:
	go install

clean:
	go clean
