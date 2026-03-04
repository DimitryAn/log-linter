test:
	go test -count=1 -v ./analyzer/	
.PHONY: test

build:
	go build cmd/loglinter/main.go
.PHONY: build