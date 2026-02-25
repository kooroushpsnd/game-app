# --- build stage ---
FROM golang:1.24.1-alpine3.20 AS builder

WORKDIR /app

RUN set -eux; \
  for i in 1 2 3 4 5; do \
    apk update && apk add --no-cache git ca-certificates && break; \
    echo "apk failed... retrying ($i)"; \
    sleep 2; \
  done

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/api
RUN go install github.com/rubenv/sql-migrate/...@latest

# --- runtime stage ---
FROM alpine:3.20

WORKDIR /app
RUN apk add --no-cache ca-certificates netcat-openbsd bash

COPY --from=builder /app/server /app/server
COPY --from=builder /go/bin/sql-migrate /usr/local/bin/sql-migrate

COPY --from=builder /app/internal/repository/postgres/migrations /app/internal/repository/postgres/migrations
COPY --from=builder /app/dbconfig.yml /app/dbconfig.yml
COPY --from=builder /app/config.yml /app/config.yml

EXPOSE 8080 8081
CMD ["/app/server"]