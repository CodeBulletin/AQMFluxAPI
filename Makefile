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

generate:
	sqlc generate

# Build the application
build: clean generate $(BINARY)

$(BINARY): $(GO_FILES)
	mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BINARY) -ldflags="-X 'main.Debug=true' -X 'main.Version=dev'" $(SRC_DIR)/cmd/main.go 

# Run the application
run: clean build
	$(BINARY)

# Run migrations
migrate:
	$(GO) run $(MIGRATION_DIR)/migrate.go up

migrate-fix:
	$(GO) run $(MIGRATION_DIR)/migrate.go $(version) fix-up

# Rollback migrations
rollback:
	$(GO) run $(MIGRATION_DIR)/migrate.go down

rollback-force:
	$(GO) run $(MIGRATION_DIR)/migrate.go $(version) fix-down

# Create a new migration
migration:
	$(MIGRATION) create -ext sql -dir $(MIGRATION_DIR)/migrations $(name)