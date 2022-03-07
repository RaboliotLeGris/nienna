FROM rust:1.57.0-buster@sha256:e6bb347e01ca274c3d4511b765f7cb9ce495d8bad3a3a99b46ae1337f7af52f6 as builder

WORKDIR /usr/src/backburner
COPY . .
RUN cargo build --release

FROM debian:11.2-slim@sha256:d5cd7e54530a8523168473a2dcc30215f2c863bfa71e09f77f58a085c419155b

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.3.3-0+deb11u1 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /backburner
COPY --from=builder /usr/src/backburner/target/release/backburner /backburner

CMD ["./backburner"]