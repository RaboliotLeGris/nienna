FROM rust:1.60.0-buster@sha256:141ac79d31cb2d2583428d34197486d2f0974d53464811d9ff22b8bc14da4d05 as builder

WORKDIR /usr/src/backburner
COPY . .
RUN cargo build --release

FROM debian:11.3-slim@sha256:f75d8a3ac10acdaa9be6052ea5f28bcfa56015ff02298831994bd3e6d66f7e57

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.3.3-0+deb11u1 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /backburner
COPY --from=builder /usr/src/backburner/target/release/backburner /backburner

CMD ["./backburner"]