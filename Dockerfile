# Stage 1: Build binary
FROM golang:1.25.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/engine ./cmd/server/main.go

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/engine .

EXPOSE 8080

CMD ["./engine"]