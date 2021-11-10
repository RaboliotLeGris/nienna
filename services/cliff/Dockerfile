FROM golang:1.16.2-buster@sha256:b645e32e5582fe96675b6c396303c61bcdf3ca7a5b68c783c6321d2b36a5ea66 as builder

WORKDIR /go/src/cliff
COPY . .

RUN go build -o build/cliff

FROM debian:10.11-slim@sha256:544c93597c784cf68dbe492ef35c00de7f4f6a990955c7144a40b20d86a3475f

WORKDIR /cliff
COPY --from=builder /go/src/cliff/build/cliff /cliff

CMD ["./cliff"]
