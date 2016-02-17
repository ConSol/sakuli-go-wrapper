default:build

.PHONY: build

build:
	GOOS=linux GOARCH=386 go build github.com/ConSol/sakuli-go-wrapper
	GOOS=windows GOARCH=386 go build github.com/ConSol/sakuli-go-wrapper

