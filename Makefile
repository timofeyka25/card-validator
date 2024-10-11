swag:
	swag init --parseDependency -d ./internal/transport/http -g ./server.go

lint:
	golangci-lint run -v