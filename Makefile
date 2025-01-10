SOURCES=$(shell find . -type f -name '*.go' -not -path "./vendor/*")
TARGET=astrobotum

.PHONY: all
all: $(TARGET)

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -race ./...

.PHONY: validate
validate: lint test

$(TARGET): $(SOURCES)
	go build -o $@ ./cmd/main
