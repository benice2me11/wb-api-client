SHELL := /bin/bash

.PHONY: generate test lint verify-generated

generate:
	./scripts/generate.sh

test:
	go test ./...

lint:
	go vet ./...

verify-generated:
	./scripts/generate.sh
	git diff --exit-code
