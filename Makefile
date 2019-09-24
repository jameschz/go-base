#include .env

PROJECTNAME=$(shell basename "$(PWD)")

# Go related variables.
GOBASE=$(shell pwd)
GOPATH=$(GOBASE)/vendor:$(GOBASE)
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

# Sed commands
ifeq ($(shell uname), Darwin)
SEDICMD=sed -i ""
else
SEDICMD=sed -i
endif


## compile: Compile binary files (default running logic).
compile: go-compile

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

## clean: Clean built & tmp files. Runs `go clean` internally.
clean: go-clean clean-swp

## create: Create new project
create: create-project

## synclib: Sync libraries
synclib: sync-lib

# Go internal commands.
go-compile: go-mkdir go-clean go-get go-build
go-mkdir:
	@echo "make> Prepare dirs ..."
	@mkdir -p $(GOBASE)/bin
	@mkdir -p $(GOBASE)/log
	@mkdir -p $(GOBASE)/vendor
go-build:
	@echo "make> Building binary files ..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)
go-generate:
	@echo "make> Generating dependency files ..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go generate $(generate)
go-get:
	@echo "make> Checking missing dependencies ..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get $(get)
go-install:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install $(GOFILES)
go-clean:
	@echo "make> Cleaning build cache ..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

# Development commands
clean-swp:
	@echo "make> Cleaning .swp files ..."
	@find . -type f -name '*.swp' | xargs rm -f
create-project: proj_name := $(notdir $(path))
create-project: proj_name_u := `echo $(proj_name) | tr a-z- A-Z_`
create-project:
ifdef path
	@echo "make> Create project '$(proj_name)' at $(path) ..."
	@mkdir -p $(path); cp ./.gitignore $(path)/.gitignore; cp ./doc/tpl/main.t $(path)/main.go; cp ./doc/tpl/Makefile $(path)/Makefile
	@mkdir -p $(path)/etc/; cp -r ./etc/* $(path)/etc/
	@mkdir -p $(path)/src/base; cp -r ./src/base/* $(path)/src/base/
	@echo "make> Change project base ENV to $(proj_name_u)_ENV ..."
	@$(SEDICMD) "s/GO_BASE_/$(proj_name_u)_/g" $(path)/src/base/util/base.go
	@$(SEDICMD) "s/GO_BASE_/$(proj_name_u)_/g" $(path)/Makefile
else
	@echo "make> Usage : make create path=/path/to/project"
endif
sync-lib: proj_name := $(notdir $(path))
sync-lib: proj_name_u := `echo $(proj_name) | tr a-z- A-Z_`
sync-lib:
ifdef path
	@echo "make> Sync all libraries to $(path) ..."
	@mkdir -p $(path)/src/base; cp -r ./src/base/* $(path)/src/base/
	@echo "make> Change project base ENV to $(proj_name_u)_ENV ..."
	@$(SEDICMD) "s/GO_BASE_/$(proj_name_u)_/g" $(path)/src/base/util/base.go
	@$(SEDICMD) "s/GO_BASE_/$(proj_name_u)_/g" $(path)/Makefile
else
	@echo "make> Usage : make synclib path=/path/to/project"
endif

#Help commands
.PHONY: help
all: help
help: Makefile
	@echo
	@echo "make> Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
