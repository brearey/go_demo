BIN := app

all:
	./$(BIN)

build:
	@go build -o $(BIN) *.go

clean:
	@rm $(BIN)