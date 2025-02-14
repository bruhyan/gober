FROM golang:1.23-alpine AS builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/gober
RUN go build -o /gober

FROM alpine:latest

COPY --from=builder /gober /gober

EXPOSE 8080

ENTRYPOINT ["/gober"]

