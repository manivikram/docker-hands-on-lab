# Stage 1: Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o docker-app main.go

# Stage 2: Runtime stage
FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/docker-app /app/docker-app

CMD ["/app/docker-app"]
