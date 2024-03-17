run: build
	@./bin/a.out

build:
	@go build -o ./bin/a.out cmd/app/main.go

install:
	@go install ./cmd/app/main.go
