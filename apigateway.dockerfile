FROM golang:1.22-alpine as builder
ENV GO111MODULE "on"
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o main ./cmd/api-gateway/

FROM debian:buster-slim
WORKDIR /app
COPY --from=builder /app/main /app
ENTRYPOINT ["/app/main"]