BINARY_NAME = asteroids

.DEFAULT_GOAL: build

.PHONY: build
build:
	@go build -o bin/${BINARY_NAME} ./*.go

.PHONY: run
run: build
	@./bin/${BINARY_NAME}

.PHONY: clean
clean:
	@rm -rf bin