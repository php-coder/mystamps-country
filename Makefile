.PHONY: build build-linux fmt clean

build: my-country

build-linux: my-country-linux

my-country: main.go
	go build -o my-country .

my-country-linux: main.go
	GOOS=linux GOARH=amd64 go build -o my-country-linux .

fmt:
	gofmt -s -d -w .

clean:
	rm -f my-country my-country-linux
