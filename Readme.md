# Card Validator API

## Project Structure

```
├── cmd
│   └── main.go                   # Entry point of the application
├── config.example.yaml           # Example configuration file
├── docker-compose.yaml           # Docker Compose for multi-container orchestration
├── Dockerfile                    # Dockerfile for containerizing the app
├── docs                          # Swagger documentation
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod                        # Go module dependencies
├── go.sum                        # Go module versions
├── internal                      # Application core
│   ├── config                    # Configuration management
│   ├── di                        # Dependency injection
│   ├── entities                  # Domain entities
│   ├── errs                      # Custom error handling
│   ├── logger                    # Logging utility
│   ├── pkg                       # Packages
│   ├── repositories              # Data access layer
│   ├── services                  # Business logic services
│   └── transport                 # HTTP API layer
│       └── http                  # HTTP server setup and routes
│           └── handlers          # HTTP handlers for the API
└── Makefile                      # Makefile for common tasks

```

## Installation
Clone the Repository:
```shell
git clone https://github.com/timofeyka25/card-validator.git
cd card-validator
```

## Running

### Run Locally

To run the project locally, build and start the application:

```shell
go build -o bin/app cmd/main.go
./bin/app
```
Or using Docker Compose:
```shell
docker-compose up
```
The server will be running at 
```url 
http://localhost:8000
```

## API Documentation
The API documentation is generated with Swagger and is available at:
```url
http://localhost:8000/docs/swagger/index.html
```
To regenerate the Swagger docs:
```shell
make swag
```

## Development
### File Structure
    cmd/: The main entry point of the application.
    internal/: Contains the core business logic, services, entities, and configuration.
    transport/: Contains HTTP handlers and API logic.
    pkg/: Contains utilities such as validation.

### Linting
To lint the project:
```bash
make lint
```