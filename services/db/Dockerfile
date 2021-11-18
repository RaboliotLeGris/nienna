FROM golang:1.17.3-buster@sha256:29f41d09f4e91ac5c516c0b92610dfc372f8713db81feb24d45b67eb9232d260 as builder

WORKDIR /go/src/db
COPY . .

RUN go build -o build/db

FROM debian:11.1-slim@sha256:a23887a2e830b815955e010f30d4c2430cd5ef82e93c130471024bc9f808d5d3

WORKDIR /db
COPY --from=builder /go/src/db/build/db /db

CMD ["./db"]
