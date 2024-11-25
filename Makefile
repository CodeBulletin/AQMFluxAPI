# Variables
APP_NAME := AQMFluxAPI
BUILD_DIR := build
SRC_DIR := .
MIGRATION_DIR := ./cmd/migrate
GO_FILES := $(shell find $(SRC_DIR)/cmd -name '*.go')
BINARY := $(BUILD_DIR)/$(APP_NAME)

# Commands
GO := go
GOFMT := gofmt
MIGRATION := migrate

# Targets
.PHONY: all build clean test fmt run

all: build

# Clean build artifacts
clean:
	@if [ -d "$(BUILD_DIR)" ]; then \
		rm -rf $(BUILD_DIR); \
		echo "Cleaned $(BUILD_DIR)"; \
	else \
		echo "Nothing to clean"; \
	fi

# Build the application
build: clean $(BINARY)

$(BINARY): $(GO_FILES)
	mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BINARY) $(SRC_DIR)/cmd/main.go

# Run the application
run: clean build
	$(BINARY)

# Format the code
fmt:
	$(GOFMT) -w $(SRC_DIR)

# Run tests
test:
	$(GO) test ./...

# Run migrations
migrate:
	$(GO) run $(MIGRATION_DIR)/main.go up

migrate-fix:
	$(GO) run $(MIGRATION_DIR)/main.go $(version) fix 

# Rollback migrations
rollback:
	$(GO) run $(MIGRATION_DIR)/main.go down

rollback-force:
	$(GO) run $(MIGRATION_DIR)/main.go $(version) fixdown

# Create a new migration
migration:
	$(MIGRATION) create -ext sql -dir $(MIGRATION_DIR)/migrations $(name)