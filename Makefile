# Simple Makefile for blog
# Author: Matthew Zember


COMP=go build
FLAGS=-v

main: go.mod go.sum main.go
	$(COMP)
	
