include ../Makefile

all: build

.PHONY: build
build:
	$(KN) func build -v

.PHONY: deploy
deploy:
	$(KN) func deploy -v -n $(NAMESPACE) -b pack
