.PHONY: build build-linux fmt clean

build:
	go build -o my-country .

build-linux:
	GOOS=linux GOARH=amd64 go build -o my-country-linux .

fmt:
	gofmt -s -d -w .

clean:
	rm -f my-country my-country-linux
