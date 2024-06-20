BINARY_NAME=apcore

MAIN_PACKAGE=./

SOURCES=$(shell find . -type f -name '*.go')

TEST_PACKAGES=$(shell go list ./... | grep -v /vendor/)

.PHONY: all dev build docker run test clean format lint vet

all: build

dev:
	@echo "Running code in dev mode"
	go run main.go

build: $(SOURCES)
	@echo "Building the binary..."
	go build -o $(BINARY_NAME) $(MAIN_PACKAGE)

docker:
	@echo "Building Docker image..."
	docker build --network host -t apcore . --progress=plain

run: build
	@echo "Running the binary..."
	./$(BINARY_NAME)

test:
	@echo "Running tests..."
	go test -v $(TEST_PACKAGES)

clean:
	@echo "Cleaning up..."
	go clean
	rm -f $(BINARY_NAME)

format:
	@echo "Formatting the code..."
	gofmt -s -w $(SOURCES)

lint:
	@echo "Linting the code..."
	@golint $(TEST_PACKAGES)

vet:
	@echo "Vetting the code..."
	@go vet $(TEST_PACKAGES)

deps:
	@echo "Installing dependencies..."
	go mod tidy

docs:
	@echo "Generating documentation..."
	swag init --parseDependency

check: format lint vet test

help:
	@echo "Makefile commands:"
	@echo "  make all       - Build the binary"
	@echo "  make dev       - Run the code in dev mode"
	@echo "  make build     - Build the binary"
	@echo "  make docker    - Build the Docker image"
	@echo "  make run       - Run the binary"
	@echo "  make test      - Run tests"
	@echo "  make clean     - Clean up generated files"
	@echo "  make format    - Format the code"
	@echo "  make lint      - Lint the code"
	@echo "  make vet       - Vet the code"
	@echo "  make deps      - Install dependencies"
	@echo "  make docs      - Generate documentation"
	@echo "  make check     - Run all checks (format, lint, vet, test)"
	@echo "  make help      - Show this help message"