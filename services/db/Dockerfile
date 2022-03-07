FROM golang:1.17.6-buster@sha256:ad61e922fb389384f00ba98dba8dae56f1dcf9ff731506857ab8f1312c47c212 as builder

WORKDIR /go/src/db
COPY . .

RUN go build -o build/db

FROM debian:11.2-slim@sha256:d5cd7e54530a8523168473a2dcc30215f2c863bfa71e09f77f58a085c419155b

WORKDIR /db
COPY --from=builder /go/src/db/build/db /db

CMD ["./db"]
