swagger: 
	@swag init -g cmd/router/main.go -o api

build: swagger
	@go build -o bin/server -v

run: swagger
	@./bin/server
