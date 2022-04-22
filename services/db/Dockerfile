FROM golang:1.18.1-buster@sha256:d8f864bac466e488bb425a59ff9bc104828cd44d0985a8d2fcb5570a9250258a as builder

WORKDIR /go/src/db
COPY . .

RUN go build -o build/db

FROM debian:11.3-slim@sha256:f75d8a3ac10acdaa9be6052ea5f28bcfa56015ff02298831994bd3e6d66f7e57

WORKDIR /db
COPY --from=builder /go/src/db/build/db /db

CMD ["./db"]
