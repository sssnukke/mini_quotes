FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o mini_quotes .

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/mini_quotes .

EXPOSE 8080

CMD ["./mini_quotes"]