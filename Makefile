BINARY_NAME=apcore

MAIN_PACKAGE=./

SOURCES=$(shell find . -type f -name '*.go')

TEST_PACKAGES=$(shell go list ./... | grep -v /vendor/)

MIGRATE=migrate
MIGRATION_DIR=./migrations
DATABASE_URL=postgres://apcore_user:apcore_pass@localhost:5432/apcore_db?sslmode=disable


.PHONY: all dev build docker run test clean lint

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

lint:
	@echo "Linting the code..."
	@golangci-lint run

deps:
	@echo "Installing dependencies..."
	go mod tidy

docs:
	@echo "Generating documentation..."
	swag init --parseDependency

migrate-create:
	@read -p "Enter migration name: " name; \
	$(MIGRATE) create -ext sql -dir $(MIGRATION_DIR) -seq $$name

migrate-up:
	$(MIGRATE) -path $(MIGRATION_DIR) -database $(DATABASE_URL) up

migrate-down:
	$(MIGRATE) -path $(MIGRATION_DIR) -database $(DATABASE_URL) down 1

help:
	@echo "Makefile commands:"
	@echo "  make all            - Build the binary"
	@echo "  make dev            - Run the code in dev mode"
	@echo "  make build          - Build the binary"
	@echo "  make docker         - Build the Docker image"
	@echo "  make run            - Run the binary"
	@echo "  make test           - Run tests"
	@echo "  make clean          - Clean up generated files"
	@echo "  make lint           - Lint the code"
	@echo "  make deps           - Install dependencies"
	@echo "  make docs           - Generate documentation"
	@echo "  make migrate-create - Generate migration"
	@echo "  make migrate-up     - Apply migration"
	@echo "  make migrate-down   - Rollback migration"
	@echo "  make help           - Show this help message"
