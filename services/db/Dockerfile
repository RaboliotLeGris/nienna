FROM golang:1.17.3-buster@sha256:6d0772539cd83e5a8b24a0003e82e042148f89762d36c610c220536322541e94 as builder

WORKDIR /go/src/db
COPY . .

RUN go build -o build/db

FROM debian:11.1-slim@sha256:a23887a2e830b815955e010f30d4c2430cd5ef82e93c130471024bc9f808d5d3

WORKDIR /db
COPY --from=builder /go/src/db/build/db /db

CMD ["./db"]
