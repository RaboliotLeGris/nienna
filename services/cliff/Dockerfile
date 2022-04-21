FROM golang:1.18.1-buster@sha256:279e4ec837160ba80a6ce81e8ae5a5fa4723dc720e61e9b262b82df2c9b53b0f as builder

WORKDIR /go/src/cliff
COPY . .

RUN go build -o build/cliff

FROM debian:11.3-slim@sha256:f75d8a3ac10acdaa9be6052ea5f28bcfa56015ff02298831994bd3e6d66f7e57

WORKDIR /cliff
COPY --from=builder /go/src/cliff/build/cliff /cliff

CMD ["./cliff"]
