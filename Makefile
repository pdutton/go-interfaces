#!/bin/make

GO?=go

DIRS:=\
  net/http/client \
  net/http/server \
  os

# FILES:=$(foreach dir,$(DIRS),$(wildcard $(dir)/*.go))

.PHONY: test
test:
	$(GO) test $(addprefix ./,$(DIRS))

.PHONY: fmt
fmt:
	$(GO) fmt $(addprefix ./,$(DIRS))
	
