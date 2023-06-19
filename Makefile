KN ?= kn
KUBECTL ?= kubectl
NAMESPACE ?= kube-direct

ZEITGEIST_VERSION = v0.4.1
BIN_DIR := bin

all: build

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

$(BIN_DIR)/zeitgeist: $(BIN_DIR)
	curl -sSfL -o $(BIN_DIR)/zeitgeist \
		https://github.com/kubernetes-sigs/zeitgeist/releases/download/$(ZEITGEIST_VERSION)/zeitgeist_$(ZEITGEIST_VERSION:v%=%)_linux_amd64
	chmod +x $(BIN_DIR)/zeitgeist

.PHONY: verify-dependencies
verify-dependencies: $(BIN_DIR)/zeitgeist
	$(BIN_DIR)/zeitgeist validate --local-only --base-path . --config dependencies.yaml

.PHONY: api
api:
	.github/generate-api

.PHONY: verify-api
verify-api: api
	.github/tree-status
