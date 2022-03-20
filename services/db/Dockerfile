FROM golang:1.18.0-buster@sha256:35cf8e5001bd5b2c3fc0286e379bed24cc0dd1b5b5d50ddd524b30e09696b0cc as builder

WORKDIR /go/src/db
COPY . .

RUN go build -o build/db

FROM debian:11.2-slim@sha256:40f90ead34f61f38c69fd604878c73c75ec944055328b9883f3b643c1b8886f0

WORKDIR /db
COPY --from=builder /go/src/db/build/db /db

CMD ["./db"]
