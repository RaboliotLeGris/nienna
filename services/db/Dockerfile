FROM golang:1.17.3-buster@sha256:970e472a45917b6bb7a20b9d269216cebadc7fd3afd4ce6e4690636b562afa7c as builder

WORKDIR /go/src/db
COPY . .

RUN go build -o build/db

FROM debian:11.1-slim@sha256:a23887a2e830b815955e010f30d4c2430cd5ef82e93c130471024bc9f808d5d3

WORKDIR /db
COPY --from=builder /go/src/db/build/db /db

CMD ["./db"]
