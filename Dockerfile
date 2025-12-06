FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bhambins ./cmd/bhambins

FROM alpine:3.21

RUN apk add --no-cache ca-certificates

COPY --from=builder /bhambins /usr/local/bin/bhambins

ENTRYPOINT ["bhambins"]


