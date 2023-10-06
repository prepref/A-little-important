FROM golang:alpine AS builder

ENV CGO_ENABLED 0

ENV GOOS linux

WORKDIR /app

ADD go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -o quotes ./cmd/quotes/quotes.go

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app /app

CMD ["./quotes"]

