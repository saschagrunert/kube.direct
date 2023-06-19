include ../Makefile

TRUNK ?= trunk

all: build

.PHONY: build
build:
	$(TRUNK) build
	$(KN) func build -v

.PHONY: deploy
deploy:
	$(TRUNK) build
	$(KN) func deploy -v -n $(NAMESPACE) -b pack
