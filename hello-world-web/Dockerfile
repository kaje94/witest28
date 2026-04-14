# Stage 1: build
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod .
COPY main.go .
RUN go build -o server main.go

# Stage 2: minimal runtime
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 9090
CMD ["./server"]
