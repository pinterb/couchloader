VERSION=0.1.0
MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(shell dirname $(MKFILE_PATH))
DOCKER_BIN := $(shell which docker)
EXE_FILE := $(notdir $(shell pwd))
GIT_USER := $(shell git config --get user.name)

all: build

.PHONY: check.env
check.env:
ifndef DOCKER_BIN
   $(error The docker command is not found. Verify that Docker is installed and accessible)
endif

.PHONY: test
test:
	$(DOCKER_BIN) run --rm \
	-v "$(CURRENT_DIR):/src" \
	centurylink/golang-tester

.PHONY: build
build: test
	$(DOCKER_BIN) run --rm \
	-v "$(CURRENT_DIR):/src" \
	centurylink/golang-builder

.PHONY: container
container: test
	$(DOCKER_BIN) run --rm \
	-v "$(CURRENT_DIR):/src" \
	-v /var/run/docker.sock:/var/run/docker.sock \
	centurylink/golang-builder \
	$(GIT_USER)/$(EXE_FILE):$(VERSION)

.PHONY: refresh
refresh: container

.PHONY: clean
clean: docker.gc
	rm -rf $(EXE_FILE)

.PHONY: docker.gc
docker.gc:
	for i in `docker ps --no-trunc -a -q`;do docker rm $$i;done
	docker images --no-trunc | grep none | awk '{print $$3}' | xargs -r docker rmi

.PHONY: check.env build remove refresh clean docker.gc
