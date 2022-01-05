FROM golang:1.17.3-buster@sha256:ee3a388a872237ddb600de3ab9512e73df0043f8878f0f355baeb5b723ef16ec as builder

WORKDIR /go/src/db
COPY . .

RUN go build -o build/db

FROM debian:11.1-slim@sha256:656b29915fc8a8c6d870a1247aad7796ce296729b0ae95e168d1cfe30dd2fb3c

WORKDIR /db
COPY --from=builder /go/src/db/build/db /db

CMD ["./db"]
