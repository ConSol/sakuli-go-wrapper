#@IgnoreInspection BashAddShebang
#
# Makefile for quickly build needed binaries
default:build

.PHONY: build

build:
	GOOS=linux GOARCH=386 go build -o target/classes/sakuli github.com/ConSol/sakuli-go-wrapper
	GOOS=windows GOARCH=386 go build -o target/classes/sakuli.exe github.com/ConSol/sakuli-go-wrapper
