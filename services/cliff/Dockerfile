FROM golang:1.18.0-buster@sha256:2845d6d719c715353312c059c816d9e4a21cdf65c98a686d7c35429db9b41a36 as builder

WORKDIR /go/src/cliff
COPY . .

RUN go build -o build/cliff

FROM debian:11.2-slim@sha256:40f90ead34f61f38c69fd604878c73c75ec944055328b9883f3b643c1b8886f0

WORKDIR /cliff
COPY --from=builder /go/src/cliff/build/cliff /cliff

CMD ["./cliff"]
