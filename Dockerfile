FROM golang:1.23.2-alpine3.20 AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/app cmd/main.go

FROM alpine:3.20

WORKDIR /

COPY --from=builder /app/bin .
COPY --from=builder /app/config.example.yaml config.yaml

RUN chmod +x ./app

ENTRYPOINT ["/app"]