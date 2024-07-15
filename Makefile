run: build
	@./bin/app

build:
	@go build -o bin/app ./cmd/main.go

race:
	@go run -race ./cmd/main.go  
