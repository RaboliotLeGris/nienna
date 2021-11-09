FROM golang:1.16.2-buster@sha256:b645e32e5582fe96675b6c396303c61bcdf3ca7a5b68c783c6321d2b36a5ea66 as builder

WORKDIR /go/src/cliff
COPY . .

RUN go build -o build/cliff

FROM debian:10.8-slim@sha256:ca9003e9899b458d46b6ee6b040cf9e0d715acbc58e0615871f019c38ada8bc1

WORKDIR /cliff
COPY --from=builder /go/src/cliff/build/cliff /cliff

CMD ["./cliff"]
