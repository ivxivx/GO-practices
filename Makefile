.PHONY: all clean help confirm \
	test


SHELL = /bin/bash

help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

confirm:
  @echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

########
# test #
########
test:
	go test ./... -cover -race -failfast -timeout=30s

test-cover-package:
	go test ./... -coverpkg=./... -cover -race -failfast -timeout=30s

