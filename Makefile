#!/usr/bin/make -f
.DEFAULT_GOAL := proto-all

VERSION := $(shell git describe --tags 2>/dev/null || echo "v0.0.0")
COMMIT := $(shell git log -1 --format='%H' 2>/dev/null || echo "unknown")
DOCKER := $(shell which docker)

BUILDDIR ?= $(CURDIR)/build
INVARIANT_CHECK_INTERVAL ?= $(INVARIANT_CHECK_INTERVAL:-0)
export PROJECT_HOME=$(shell git rev-parse --show-toplevel)
export GO_PKG_PATH=$(HOME)/go/pkg
export GO111MODULE = on

# process build tags

LEDGER_ENABLED ?= true
build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
	ifeq ($(OS),Windows_NT)
		GCCEXE = $(shell where gcc.exe 2> NUL)
		ifeq ($(GCCEXE),)
			$(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
		else
			build_tags += ledger
		endif
	else
		UNAME_S = $(shell uname -s)
		ifeq ($(UNAME_S),OpenBSD)
			$(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
		else
			GCC = $(shell command -v gcc 2> /dev/null)
			ifeq ($(GCC),)
				$(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
			else
				build_tags += ledger
			endif
		endif
	endif
endif

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=kiichain \
	-X github.com/cosmos/cosmos-sdk/version.AppName=kiichaind \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
	-X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"

ifeq ($(LINK_STATICALLY),true)
	ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

# BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)' -race
BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'


##################
###  Protobuf  ###
##################

protoVer=0.14.0
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)

#### Command List ####

all: lint install

install:
	# @echo "--> ensure dependencies have not been modified"
	# @go mod verify
	# @go mod tidy
	# @echo "--> installing kiichaind"
	@go install $(BUILD_FLAGS) -mod=readonly ./cmd/kiichaind

install-with-race-detector: go.sum
		go install -race $(BUILD_FLAGS) ./cmd/kiichaind

install-price-feeder: go.sum
		go install $(BUILD_FLAGS) ./oracle/price-feeder

loadtest: go.sum
		go build $(BUILD_FLAGS) -o ./build/loadtest ./loadtest/

price-feeder: go.sum
		go build $(BUILD_FLAGS) -o ./build/price-feeder ./oracle/price-feeder

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		@go mod verify

lint:
	golangci-lint run
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s
	go mod verify

build:
	go build $(BUILD_FLAGS) -o ./build/kiichaind ./cmd/kiichaind

build-price-feeder:
	go build $(BUILD_FLAGS) -o ./build/price-feeder ./oracle/price-feeder

clean:
	rm -rf ./build

build-loadtest:
	go build -o build/loadtest ./loadtest/

proto-gen:
	@echo "Generating protobuf files..."
	@$(protoImage) sh ./scripts/protogen.sh
	@go mod tidy
.PHONY: proto-gen

###############################################################################
###                       Local testing using docker container              ###
###############################################################################
# To start a 4-node cluster from scratch:
# make clean && make docker-cluster-start
# To stop the 4-node cluster:
# make docker-cluster-stop
# If you have already built the binary, you can skip the build:
# make docker-cluster-start-skipbuild
###############################################################################


# Build linux binary on other platforms
build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-linux-gnu-gcc make build
.PHONY: build-linux

build-price-feeder-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-linux-gnu-gcc make build-price-feeder
.PHONY: build-price-feeder-linux

# Build docker image
build-docker-node:
	@cd docker && docker build --tag kii-chain/localnode localnode --platform linux/x86_64
.PHONY: build-docker-node

build-rpc-node:
	@cd docker && docker build --tag kii-chain/rpcnode rpcnode --platform linux/x86_64
.PHONY: build-rpc-node

# Run a single node docker container
run-local-node: kill-sei-node build-docker-node
	@rm -rf $(PROJECT_HOME)/build/generated
	docker run --rm \
	--name sei-node \
	--network host \
	--user="$(shell id -u):$(shell id -g)" \
	-v $(PROJECT_HOME):/kiiChain/kii-chain:Z \
	-v $(GO_PKG_PATH)/mod:/root/go/pkg/mod:Z \
	-v $(shell go env GOCACHE):/root/.cache/go-build:Z \
	--platform linux/x86_64 \
	kii-chain/localnode
.PHONY: run-local-node

# Run a single rpc state sync node docker container
run-rpc-node: build-rpc-node
	docker run --rm \
	--name sei-rpc-node \
	--network docker_localnet \
	--user="$(shell id -u):$(shell id -g)" \
	-v $(PROJECT_HOME):/sei-protocol/kii-chain:Z \
	-v $(PROJECT_HOME)/../kiichain-tendermint:/sei-protocol/sei-tendermint:Z \
    -v $(PROJECT_HOME)/../kiichain-cosmos:/sei-protocol/sei-cosmos:Z \
    -v $(PROJECT_HOME)/../kiichain-db:/sei-protocol/sei-db:Z \
	-v $(GO_PKG_PATH)/mod:/root/go/pkg/mod:Z \
	-v $(shell go env GOCACHE):/root/.cache/go-build:Z \
	-p 26668-26670:26656-26658 \
	--platform linux/x86_64 \
	kii-chain/rpcnode
.PHONY: run-rpc-node

run-rpc-node-skipbuild: build-rpc-node
	docker run --rm \
	--name sei-rpc-node \
	--network docker_localnet \
	--user="$(shell id -u):$(shell id -g)" \
	-v $(PROJECT_HOME):/sei-protocol/kii-chain:Z \
	-v $(PROJECT_HOME)/../kiichain-tendermint:/sei-protocol/sei-tendermint:Z \
    -v $(PROJECT_HOME)/../kiichain-cosmos:/sei-protocol/sei-cosmos:Z \
    -v $(PROJECT_HOME)/../kiichain-db:/sei-protocol/sei-db:Z \
	-v $(GO_PKG_PATH)/mod:/root/go/pkg/mod:Z \
	-v $(shell go env GOCACHE):/root/.cache/go-build:Z \
	-p 26668-26670:26656-26658 \
	--platform linux/x86_64 \
	--env SKIP_BUILD=true \
	kii-chain/rpcnode
.PHONY: run-rpc-node

kill-sei-node:
	docker ps --filter name=sei-node --filter status=running -aq | xargs docker kill 2> /dev/null || true

kill-rpc-node:
	docker ps --filter name=sei-rpc-node --filter status=running -aq | xargs docker kill 2> /dev/null || true

# Run a 4-node docker containers
docker-cluster-start: docker-cluster-stop build-docker-node
	@rm -rf $(PROJECT_HOME)/build/generated
	@mkdir -p $(shell go env GOPATH)/pkg/mod
	@mkdir -p $(shell go env GOCACHE)
	@cd docker && USERID=$(shell id -u) GROUPID=$(shell id -g) GOCACHE=$(shell go env GOCACHE) NUM_ACCOUNTS=10 INVARIANT_CHECK_INTERVAL=${INVARIANT_CHECK_INTERVAL} UPGRADE_VERSION_LIST=${UPGRADE_VERSION_LIST} docker compose up

.PHONY: localnet-start

# Use this to skip the kiichaind build process
docker-cluster-start-skipbuild: docker-cluster-stop build-docker-node
	@rm -rf $(PROJECT_HOME)/build/generated
	@cd docker && USERID=$(shell id -u) GROUPID=$(shell id -g) GOCACHE=$(shell go env GOCACHE) NUM_ACCOUNTS=10 SKIP_BUILD=true docker compose up
.PHONY: localnet-start

# Stop 4-node docker containers
docker-cluster-stop:
	@cd docker && USERID=$(shell id -u) GROUPID=$(shell id -g) GOCACHE=$(shell go env GOCACHE) docker compose down
.PHONY: localnet-stop


# Implements test splitting and running. This is pulled directly from
# the github action workflows for better local reproducibility.

GO_TEST_FILES != find $(CURDIR) -name "*_test.go"

# default to four splits by default
NUM_SPLIT ?= 4

$(BUILDDIR):
	mkdir -p $@

# The format statement filters out all packages that don't have tests.
# Note we need to check for both in-package tests (.TestGoFiles) and
# out-of-package tests (.XTestGoFiles).
$(BUILDDIR)/packages.txt:$(GO_TEST_FILES) $(BUILDDIR)
	go list -f "{{ if (or .TestGoFiles .XTestGoFiles) }}{{ .ImportPath }}{{ end }}" ./... | sort > $@

split-test-packages:$(BUILDDIR)/packages.txt
	split -d -n l/$(NUM_SPLIT) $< $<.
test-group-%:split-test-packages
	cat $(BUILDDIR)/packages.txt.$* | xargs go test -parallel 4 -mod=readonly -timeout=10m -race -coverprofile=$*.profile.out -covermode=atomic
