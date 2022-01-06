FROM golang:1.17.5-buster@sha256:7f89be6ec02f7eb4b429e9a43f4ac7d49d2c0cb79a99a28e2d1efe1e7e8a3473 as builder

WORKDIR /go/src/cliff
COPY . .

RUN go build -o build/cliff

FROM debian:11.1-slim@sha256:656b29915fc8a8c6d870a1247aad7796ce296729b0ae95e168d1cfe30dd2fb3c

WORKDIR /cliff
COPY --from=builder /go/src/cliff/build/cliff /cliff

CMD ["./cliff"]
