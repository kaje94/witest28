FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /server .

FROM gcr.io/distroless/static

COPY --from=builder /server /server

EXPOSE 8080

CMD ["/server"]
