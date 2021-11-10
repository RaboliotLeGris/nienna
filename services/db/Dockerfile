FROM golang:1.17.3-buster@sha256:eac5f1bc63ff9a91539f0048afc710e690623f75995e79a5da515004ea8de826 as builder

WORKDIR /go/src/db
COPY . .

RUN go build -o build/db

FROM debian:10.11-slim@sha256:544c93597c784cf68dbe492ef35c00de7f4f6a990955c7144a40b20d86a3475f

WORKDIR /db
COPY --from=builder /go/src/db/build/db /db

CMD ["./db"]
