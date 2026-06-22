BINARY=wtf
BUILD_DIR=./bin

.PHONY: build install run clean

build:
	go build -o $(BUILD_DIR)/$(BINARY) ./...

install:
	go install ./...

run:
	go run main.go

clean:
	rm -rf $(BUILD_DIR)