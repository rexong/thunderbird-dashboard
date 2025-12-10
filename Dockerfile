FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o /usr/local/bin/api ./cmd

FROM alpine:latest
COPY --from=builder /usr/local/bin/api /usr/local/bin/api

COPY templates/ /templates/
COPY static/ /static/

ENTRYPOINT ["/usr/local/bin/api"]

