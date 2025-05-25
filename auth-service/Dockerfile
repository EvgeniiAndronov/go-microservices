#auth-service/Dockerfile

FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /auth-service ./cmd/main.go

FROM alpine:latest
COPY --from=builder /auth-service /auth-service
COPY --from=builder /app/.env /.env
EXPOSE 8080
CMD ["/auth-service"]