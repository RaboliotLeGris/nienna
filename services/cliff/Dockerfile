FROM golang:1.18.1-buster@sha256:2b52df3e45d4f2c8ac8d6db12366d09394189bc13d67c27eb2b44b79dfa4db00 as builder

WORKDIR /go/src/cliff
COPY . .

RUN go build -o build/cliff

FROM debian:11.3-slim@sha256:78fd65998de7a59a001d792fe2d3a6d2ea25b6f3f068e5c84881250373577414

WORKDIR /cliff
COPY --from=builder /go/src/cliff/build/cliff /cliff

CMD ["./cliff"]
