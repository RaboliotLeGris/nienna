FROM golang:1.17.3-buster@sha256:622a414b498370198e16b23f2d174b8aefcd1fd96958b7b6e875dc37bea114ba as builder

WORKDIR /go/src/db
COPY . .

RUN go build -o build/db

FROM debian:11.1-slim@sha256:656b29915fc8a8c6d870a1247aad7796ce296729b0ae95e168d1cfe30dd2fb3c

WORKDIR /db
COPY --from=builder /go/src/db/build/db /db

CMD ["./db"]
