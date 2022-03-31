FROM golang:1.18.0-buster@sha256:87126aede310b29a5ee8fb7c655666e6fa1d8656fd5a26b3364021137cfcc453 as builder

WORKDIR /go/src/db
COPY . .

RUN go build -o build/db

FROM debian:11.3-slim@sha256:78fd65998de7a59a001d792fe2d3a6d2ea25b6f3f068e5c84881250373577414

WORKDIR /db
COPY --from=builder /go/src/db/build/db /db

CMD ["./db"]
