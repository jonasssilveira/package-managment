APP_NAME=order-package
BINARY_NAME=order-package
PKG=./...
COVERAGE_FILE=coverage.out

.PHONY: all build run test coverage clean docker-up docker-down

all: build

## Build the Go application
build:
	go build -o $(BINARY_NAME) main.go

## Run the application locally
run:
	go run main.go

## Run tests
test:
	go test $(PKG)

## Run tests with coverage and generate HTML report
coverage:
	go test -coverprofile=$(COVERAGE_FILE) $(PKG)
	go tool cover -html=$(COVERAGE_FILE) -o coverage.html

## Clean up generated files
clean:
	rm -f $(BINARY_NAME) $(COVERAGE_FILE) coverage.html

## Run the app and MongoDB with Docker
docker-up:
	docker-compose -f script/docker-compose.yaml up --build

## Stop Docker containers
docker-down:
	docker-compose -f script/docker-compose.yaml down
