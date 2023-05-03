FROM golang:1.20.4 AS builder

WORKDIR /work
COPY . ./
RUN go build -buildvcs=false

FROM gcr.io/distroless/base:latest
WORKDIR /app
ADD https://github.com/IBM/plex/raw/master/IBM-Plex-Sans-JP/fonts/complete/ttf/hinted/IBMPlexSansJP-Bold.ttf /app/font.ttf
COPY --from=builder /work/go-ogimg /app

EXPOSE 8000

ENTRYPOINT ["/app/go-ogimg"]