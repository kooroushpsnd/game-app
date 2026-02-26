# ---------- build stage ----------
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install git (sometimes needed for go modules)
RUN apk add --no-cache git

# Cache deps
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build a static-ish binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/api

# ---------- runtime stage ----------
FROM alpine:3.20

WORKDIR /app

# (optional) add CA certs for HTTPS calls
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server /app/server

EXPOSE 8080

CMD ["/app/server"]