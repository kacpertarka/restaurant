build:
	@go build -o bin/api cmd/main.go

run: build
	@./bin/restaurant	

build-docker:
	@docker build -t api .

run-docker:
	@docker run -it --rm -p 8080:8080 api