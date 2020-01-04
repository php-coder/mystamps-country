.PHONY: build build-linux fmt test generate-coverage open-coverage clean

SRCS := main.go db/db.go rest/rest.go

build: my-country

build-linux: my-country-linux

my-country: $(SRCS)
	go build -o my-country .

my-country-linux: $(SRCS)
	GOOS=linux GOARH=amd64 go build -o my-country-linux .

fmt:
	gofmt -s -d -w .

test:
	go test ./...

coverage.out:
	go test ./... -cover -coverprofile=$@

coverage.html: coverage.out
	go tool cover -html=$< -o $@

generate-coverage: coverage.html

open-coverage: coverage.html
	open $<

clean:
	rm -f my-country my-country-linux coverage.out coverage.html
