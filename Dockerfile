FROM golang:1.23.4-alpine AS builder

WORKDIR /work
COPY . ./
RUN CGO_ENABLED=0 go build -buildvcs=false -ldflags="-s -w" -trimpath -o main

FROM gcr.io/distroless/static:latest
WORKDIR /app
COPY font.ttf /app
COPY --from=builder /work/main /app

EXPOSE 8000

ENTRYPOINT ["/app/main"]
