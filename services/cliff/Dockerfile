FROM golang:1.18.0-buster@sha256:e4950365f46f107ba92f8a9557defccb38f151a79fd5453e9424ec34de08c35d as builder

WORKDIR /go/src/cliff
COPY . .

RUN go build -o build/cliff

FROM debian:11.2-slim@sha256:d5cd7e54530a8523168473a2dcc30215f2c863bfa71e09f77f58a085c419155b

WORKDIR /cliff
COPY --from=builder /go/src/cliff/build/cliff /cliff

CMD ["./cliff"]
