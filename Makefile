BINARY := bhambins

.PHONY: build test vet ci clean

build:
	go build -o $(BINARY) ./cmd/bhambins

test:
	go test ./test -v

vet:
	go vet ./...

ci: vet build test

clean:
	rm -f $(BINARY)
