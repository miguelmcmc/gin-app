.PHONY: clean client server run

# Build variables
AUTHOR := $(shell git log -1 --pretty=format:'%ae')
DATE := $(shell date +'%d.%m.%Y@%H:%M:%S')
GIT_COMMIT := $(shell git rev-list -1 HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
CLIENT_NAME = kydom-dash
VERSION := 0.0.1
CPACKAGE := "ironchip.net/kydom/core/version"


# Detect operating system
ifeq ($(OS),Windows_NT)
	GOOS += windows
	ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
		GOARCH += amd64
	endif
	ifeq ($(PROCESSOR_ARCHITECTURE),x86)
		GOARCH += 386
	endif
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		GOOS += linux
	endif
	ifeq ($(UNAME_S),Darwin)
		GOOS += darwin
	endif
		UNAME_P := $(shell uname -p)
	ifeq ($(UNAME_P),x86_64)
		GOARCH += amd64
	endif
		ifneq ($(filter %86,$(UNAME_P)),)
		GOARCH += 386
		endif
	ifneq ($(filter arm%,$(UNAME_P)),)
		GOARCH +=arm
	endif
endif

all: clean core

core:
	@echo -n "Building core..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X $(CPACKAGE).ProgramName=$(CLIENT_NAME) -X $(CPACKAGE).GitCommit=$(GIT_COMMIT) -X $(CPACKAGE).GitBranch=$(GIT_BRANCH) -X $(CPACKAGE).ProgramVersion=$(VERSION) -X $(CPACKAGE).ProgramAuthor=$(AUTHOR) -X $(CPACKAGE).ProgramBuildDate=$(DATE)" -o $(CLIENT_NAME) .
	@echo " Done!"
	@echo Executable $(CLIENT_NAME) builded successfully!

clean:
	@echo -n "Cleaning all..."
	@rm -rf $(CURDIR)/$(CLIENT_NAME)
	@echo " Done!"


