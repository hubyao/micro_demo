GOPATH:=$(shell go env GOPATH)

.PHONY: build
build: proto
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o user-srv main.go plugin.go

.PHONY: docker
docker:docker
	docker build . -t user-srv:latest

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o user-web main.go plugin.go


# MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
# CURRENT_DIR := $(notdir $(patsubst %/,%,$(MAKEFILE_DIR))
MAKEFILE_LIST_LASTWORD = $(lastword $(MAKEFILE_LIST))
MAKEFILE_PATH := $(abspath $(MAKEFILE_LIST_LASTWORD))
MAKEFILE_DIR := $(dir $(MAKEFILE_PATH))
MAKEFILE_DIR_PATSUBST := $(patsubst %/,%,$(MAKEFILE_DIR))
CURRENT_DIR := $(notdir $(MAKEFILE_DIR_PATSUBST))

.PHONY: cur_dir
cur_dir:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o user-web main.go plugin.go
	@echo MAKEFILE_LIST=$(MAKEFILE_LIST)
	@echo MAKEFILE_LIST_LASTWORD=$(MAKEFILE_LIST_LASTWORD)
	@echo MAKEFILE_PATH=$(MAKEFILE_PATH)
	@echo MAKEFILE_DIR=$(MAKEFILE_DIR)
	@echo MAKEFILE_DIR_PATSUBST=$(MAKEFILE_DIR_PATSUBST)
	@echo CURRENT_DIR=$(CURRENT_DIR)