FROM golang:1.19 as builder

WORKDIR /app
COPY . .

ENV CGO_ENABLED=0

RUN GOOS=linux GOARCH=amd64 go build -mod=vendor -o ./rest-api ./cmd/app/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/rest-api /app/rest-api

ENTRYPOINT ["/app/rest-api"]