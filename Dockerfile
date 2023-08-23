FROM golang:1.21.0-alpine3.18 AS builder

WORKDIR /work
COPY . ./
RUN go build -buildvcs=false -o main

FROM gcr.io/distroless/base:latest
WORKDIR /app
ADD https://github.com/IBM/plex/raw/master/IBM-Plex-Sans-JP/fonts/complete/ttf/hinted/IBMPlexSansJP-Bold.ttf /app/font.ttf
COPY --from=builder /work/main /app

EXPOSE 8000

ENTRYPOINT ["/app/main"]