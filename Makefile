.PHONY: build build-linux fmt clean

SRCS := main.go db/db.go rest/rest.go

build: my-country

build-linux: my-country-linux

my-country: $(SRCS)
	go build -o my-country .

my-country-linux: $(SRCS)
	GOOS=linux GOARH=amd64 go build -o my-country-linux .

fmt:
	gofmt -s -d -w .

clean:
	rm -f my-country my-country-linux
