#!/bin/make

GO?=go

DIRS:=\
  io \
  net \
  net/http/client \
  net/http/server \
  os \
  os/exec \
  path \
  path/filepath \
  sync

# FILES:=$(foreach dir,$(DIRS),$(wildcard $(dir)/*.go))

.PHONY: test
test:
	$(GO) test './...'

# $(GO) test $(addprefix ./,$(DIRS))

.PHONY: fmt
fmt:
	$(GO) fmt ./...

# $(GO) fmt $(addprefix ./,$(DIRS))
	
