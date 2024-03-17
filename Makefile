run: build
	@./bin/a.out

build:
	@go build -o ./bin/a.out cmd/main.go
