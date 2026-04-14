FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o hello-world-service .

FROM gcr.io/distroless/static:nonroot

COPY --from=builder /app/hello-world-service /hello-world-service

EXPOSE 9090

USER nonroot:nonroot

CMD ["/hello-world-service"]
