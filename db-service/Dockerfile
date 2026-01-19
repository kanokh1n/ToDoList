FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/migrate && \
    chmod +x /usr/local/bin/migrate

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o db-service .

FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate

COPY --from=builder /app/db-service .
COPY migrations /app/migrations

EXPOSE 8081

CMD ["./db-service"]