BIN := app
EVERYWHERE := ./...

all: clean build
	./$(BIN)

test: clean build
	go test $(EVERYWHERE)

build:
	@go build -o $(BIN) *.go

format:
	@go fmt $(EVERYWHERE)

clean:
	@rm -f $(BIN)
