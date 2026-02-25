# --- build stage ---
FROM golang:1.22-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# build from cmd/api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/api

# install sql-migrate
RUN go install github.com/rubenv/sql-migrate/...@latest


# --- runtime stage ---
FROM alpine:3.20

WORKDIR /app
RUN apk add --no-cache ca-certificates netcat-openbsd bash

# binaries
COPY --from=builder /app/server /app/server
COPY --from=builder /go/bin/sql-migrate /usr/local/bin/sql-migrate

# runtime files needed:
# migrations + sql-migrate config + app config
COPY --from=builder /app/internal/repository/postgres/migrations /app/internal/repository/postgres/migrations
COPY --from=builder /app/dbconfig.yml /app/dbconfig.yml
COPY --from=builder /app/config.yml /app/config.yml

# expose both ports your app uses
EXPOSE 8080 8081

CMD ["/app/server"]