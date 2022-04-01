FROM golang:1.18.0-buster@sha256:284da7e6ad7fd1250a785c24912fb6f8652176b588d637de5c5b3d4a4aa1693a as builder

WORKDIR /go/src/cliff
COPY . .

RUN go build -o build/cliff

FROM debian:11.3-slim@sha256:78fd65998de7a59a001d792fe2d3a6d2ea25b6f3f068e5c84881250373577414

WORKDIR /cliff
COPY --from=builder /go/src/cliff/build/cliff /cliff

CMD ["./cliff"]
