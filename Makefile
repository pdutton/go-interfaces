#!/bin/make

GO?=go

DIRS:=\
  encoding/json \
  io \
  net \
  net/http/client \
  net/http/server \
  os \
  os/exec \
  os/signal \
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
	
