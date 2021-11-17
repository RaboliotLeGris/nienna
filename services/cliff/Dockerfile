FROM golang:1.17.3-buster@sha256:099de80fda2a712d3ee9f66f7f8b62061aad4b5e0ff43f66c86f16ba6a3a89cd as builder

WORKDIR /go/src/cliff
COPY . .

RUN go build -o build/cliff

FROM debian:11.1-slim@sha256:a23887a2e830b815955e010f30d4c2430cd5ef82e93c130471024bc9f808d5d3

WORKDIR /cliff
COPY --from=builder /go/src/cliff/build/cliff /cliff

CMD ["./cliff"]
