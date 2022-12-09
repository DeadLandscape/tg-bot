.PHONY: run, build
run:
	go run cmd/bot/main.go

build:
	go build -o bin/ cmd/bot/main.go