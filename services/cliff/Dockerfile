FROM golang:1.18.1-buster@sha256:2b52df3e45d4f2c8ac8d6db12366d09394189bc13d67c27eb2b44b79dfa4db00 as builder

WORKDIR /go/src/cliff
COPY . .

RUN go build -o build/cliff

FROM debian:11.3-slim@sha256:f75d8a3ac10acdaa9be6052ea5f28bcfa56015ff02298831994bd3e6d66f7e57

WORKDIR /cliff
COPY --from=builder /go/src/cliff/build/cliff /cliff

CMD ["./cliff"]
