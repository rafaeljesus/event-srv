NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

IGNORED_PACKAGES := /vendor/

.PHONY: all clean deps build

all: clean deps test build

deps:
	@echo "$(OK_COLOR)==> Installing glide dependencies$(NO_COLOR)"
	@go get -u github.com/Masterminds/glide
	@glide install

build:
	@echo "$(OK_COLOR)==> Building... $(NO_COLOR)"
	/bin/sh -c "VERSION=${VERSION} ./build/build.sh"

test:
	@/bin/sh -c "./build/test.sh $(allpackages)"

clean:
	@echo "$(OK_COLOR)==> Cleaning project$(NO_COLOR)"
	@go clean
	@rm -rf dist

_allpackages = $(shell ( go list ./... 2>&1 1>&3 | \
    grep -v -e "^$$" $(addprefix -e ,$(IGNORED_PACKAGES)) 1>&2 ) 3>&1 | \
    grep -v -e "^$$" $(addprefix -e ,$(IGNORED_PACKAGES)))

allpackages = $(if $(__allpackages),,$(eval __allpackages := $$(_allpackages)))$(__allpackages)
