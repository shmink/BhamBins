BINARY := bhambins
POSTCODE ?=
UPRN ?=
PORT ?= 8080

.PHONY: build test vet ci clean docker run serve

build:
	go build -o $(BINARY) ./cmd/bhambins

test:
	go test ./test -v

vet:
	go vet ./...

ci: vet build test

clean:
	rm -f $(BINARY)

docker:
	docker build -t bhambins .

run:
	@if [ -z "$(POSTCODE)" ] || [ -z "$(UPRN)" ]; then \
		echo "Usage: make run POSTCODE='B17 0LY' UPRN='100070285236'"; \
		exit 1; \
	fi
	go run ./cmd/bhambins -p "$(POSTCODE)" -u "$(UPRN)"

serve:
	PORT=$(PORT) go run ./cmd/bhambins