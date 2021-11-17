FROM golang:1.17.3-buster@sha256:eac5f1bc63ff9a91539f0048afc710e690623f75995e79a5da515004ea8de826 as builder

WORKDIR /go/src/db
COPY . .

RUN go build -o build/db

FROM debian:10.11-slim@sha256:cfadaa70a452076a478dcfcf2da3a3dc3b741450afed5ab67a7f3df161387c48

WORKDIR /db
COPY --from=builder /go/src/db/build/db /db

CMD ["./db"]
