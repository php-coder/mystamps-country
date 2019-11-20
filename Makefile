.PHONY: build fmt clean

build:
	go build -o my-country .

fmt:
	gofmt -s -d .
	gofmt -s -d -w .

clean:
	rm my-country
