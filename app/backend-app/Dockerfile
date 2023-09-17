# Build stage
FROM golang:1.20.1-alpine3.16 AS builder
WORKDIR /app

COPY . .
RUN go build -o main ./cmd/main.go

# Run stage
FROM alpine:3.16
WORKDIR /app

COPY --from=builder /app/main .
COPY ./cmd/app.env .

EXPOSE 9090
CMD [ "/app/main" ]

