FROM golang:1.17.8-buster@sha256:267cd0d6449f47031b4a6b386e5652ee572ced4ee7e2ec846352b0132bf0d3e5 as builder

WORKDIR /go/src/cliff
COPY . .

RUN go build -o build/cliff

FROM debian:11.2-slim@sha256:d5cd7e54530a8523168473a2dcc30215f2c863bfa71e09f77f58a085c419155b

WORKDIR /cliff
COPY --from=builder /go/src/cliff/build/cliff /cliff

CMD ["./cliff"]
