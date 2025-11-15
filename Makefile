BIN := app

all: clean build
	./$(BIN)

test: clean
	go test -v

build:
	@go build -o $(BIN) *.go

format:
	@go fmt

clean:
	@rm -f $(BIN)