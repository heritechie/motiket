# Build stage
FROM golang:1.18-alpine3.15 AS builder
WORKDIR /app
COPY . . 
RUN go build -o main cmd/server/main.go

# Run Stage
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8080
CMD ["/app/main"]