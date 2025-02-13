#!/bin/make

GO?=go

DIRS:=\
	net/http

FILES:=$(foreach dir,$(DIRS),$(wildcard $(dir)/*.go))

.PHONY: fmt
fmt:
	$(GO) fmt $(FILES)
	
