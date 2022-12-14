BASH_PATH:=$(shell which bash)
SHELL=$(BASH_PATH)
ROOT := $(shell realpath $(dir $(lastword $(MAKEFILE_LIST))))
APP := crafting-table

check-goimport:
	which goimports || GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports

format:
	- find $(ROOT) -type f -name "*.go" -not -path "$(ROOT)/vendor/*" | xargs -I R -n 1 goimports -w R
	- find $(ROOT) -type f -name "*.go" -not -path "$(ROOT)/vendor/*" | xargs -I R -n 1 gofmt -s

vendor:
	- go mod tidy
	- go mod vendor

build: format vendor
	- go build

it: build
	- sh integration_test.sh $(ROOT)
