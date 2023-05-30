KN ?= kn
KUBECTL ?= kubectl
NAMESPACE ?= kube-direct

all: build

.PHONY: build
build:
	$(KN) func build

.PHONY: deploy
deploy:
	$(KN) func deploy -n $(NAMESPACE) -b pack
	$(KUBECTL) apply -f deploy/domainmapping.yaml

.PHONY: delete
delete:
	$(KN) func delete -n $(NAMESPACE)
	$(KUBECTL) delete -f deploy/domainmapping.yaml
