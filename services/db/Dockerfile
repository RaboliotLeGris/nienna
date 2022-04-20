FROM golang:1.18.1-buster@sha256:2b52df3e45d4f2c8ac8d6db12366d09394189bc13d67c27eb2b44b79dfa4db00 as builder

WORKDIR /go/src/db
COPY . .

RUN go build -o build/db

FROM debian:11.3-slim@sha256:f75d8a3ac10acdaa9be6052ea5f28bcfa56015ff02298831994bd3e6d66f7e57

WORKDIR /db
COPY --from=builder /go/src/db/build/db /db

CMD ["./db"]
