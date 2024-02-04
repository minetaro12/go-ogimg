FROM golang:1.21.6-alpine3.19 AS builder

WORKDIR /work
COPY . ./
RUN CGO_ENABLED=0 go build -buildvcs=false -o main

FROM gcr.io/distroless/static:latest
WORKDIR /app
COPY font.ttf /app
COPY --from=builder /work/main /app

EXPOSE 8000

ENTRYPOINT ["/app/main"]