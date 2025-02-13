#!/bin/make

GO?=go

DIRS:=\
  net/http \
  os

# FILES:=$(foreach dir,$(DIRS),$(wildcard $(dir)/*.go))

.PHONY: fmt
fmt:
	$(GO) fmt $(addprefix ./,$(DIRS))
	
