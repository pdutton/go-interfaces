#!/bin/make

GO?=go

DIRS:=\
  io \
  net/http/client \
  net/http/server \
  os \
  os/exec \
  path \
  path\filepath

# FILES:=$(foreach dir,$(DIRS),$(wildcard $(dir)/*.go))

.PHONY: test
test:
	$(GO) test $(addprefix ./,$(DIRS))

.PHONY: fmt
fmt:
	$(GO) fmt $(addprefix ./,$(DIRS))
	
