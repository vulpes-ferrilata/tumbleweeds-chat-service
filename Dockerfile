FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./chat-service ./cmd

FROM alpine:latest

WORKDIR /

COPY --from=builder /app/chat-service .

COPY ./locales ./locales
COPY ./config.yaml .

EXPOSE 8080

ENTRYPOINT ["/chat-service"]