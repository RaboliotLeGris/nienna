FROM golang:1.18.0-buster@sha256:4d2553850f501c598d05e9a587653a5a81d6c7967f4368ade8e492fcf250f8f1 as builder

WORKDIR /go/src/db
COPY . .

RUN go build -o build/db

FROM debian:11.2-slim@sha256:40f90ead34f61f38c69fd604878c73c75ec944055328b9883f3b643c1b8886f0

WORKDIR /db
COPY --from=builder /go/src/db/build/db /db

CMD ["./db"]
