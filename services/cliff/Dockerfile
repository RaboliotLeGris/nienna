FROM golang:1.18.0-buster@sha256:1893d1126ad939d129d86a9dd8824f69548aac3ae9610ea8d45dbfc78bd5c8be as builder

WORKDIR /go/src/cliff
COPY . .

RUN go build -o build/cliff

FROM debian:11.3-slim@sha256:78fd65998de7a59a001d792fe2d3a6d2ea25b6f3f068e5c84881250373577414

WORKDIR /cliff
COPY --from=builder /go/src/cliff/build/cliff /cliff

CMD ["./cliff"]
