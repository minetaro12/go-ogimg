FROM golang:1.20.0-alpine3.17 AS builder

WORKDIR /work
COPY . ./
RUN go build

FROM alpine:3.17.2
RUN apk add --no-cache chromium
WORKDIR /app
COPY --from=builder /work/go-ogimg /app

EXPOSE 8000

ENTRYPOINT ["/app/go-ogimg"]