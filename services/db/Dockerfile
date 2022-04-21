FROM golang:1.18.1-buster@sha256:6ed1375f000a9086969581dd1dcc676495599228fa6b27c3219c4a76b8f09cc4 as builder

WORKDIR /go/src/db
COPY . .

RUN go build -o build/db

FROM debian:11.3-slim@sha256:f75d8a3ac10acdaa9be6052ea5f28bcfa56015ff02298831994bd3e6d66f7e57

WORKDIR /db
COPY --from=builder /go/src/db/build/db /db

CMD ["./db"]
