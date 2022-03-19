FROM golang:1.18.0-buster@sha256:a22f80e34a9f21a7e11a8ef8d29e04366da6993bf4f8efa41202899250130e6c as builder

WORKDIR /go/src/db
COPY . .

RUN go build -o build/db

FROM debian:11.2-slim@sha256:40f90ead34f61f38c69fd604878c73c75ec944055328b9883f3b643c1b8886f0

WORKDIR /db
COPY --from=builder /go/src/db/build/db /db

CMD ["./db"]
