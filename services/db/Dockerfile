FROM golang:1.18.0-buster@sha256:1893d1126ad939d129d86a9dd8824f69548aac3ae9610ea8d45dbfc78bd5c8be as builder

WORKDIR /go/src/db
COPY . .

RUN go build -o build/db

FROM debian:11.2-slim@sha256:40f90ead34f61f38c69fd604878c73c75ec944055328b9883f3b643c1b8886f0

WORKDIR /db
COPY --from=builder /go/src/db/build/db /db

CMD ["./db"]
