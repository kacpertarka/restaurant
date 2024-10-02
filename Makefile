build:
	@go build -o bin/restaurant cmd/main.go

run: build
	@./bin/restaurant	