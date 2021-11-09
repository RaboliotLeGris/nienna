FROM rust:1.50.0-buster@sha256:25a32551f42722169ea0b50b9ea752bd43a31b3e52371e735ea675866e4eb164 as builder

WORKDIR /usr/src/backburner
COPY . .
RUN cargo build --release

FROM debian:10.8-slim@sha256:ca9003e9899b458d46b6ee6b040cf9e0d715acbc58e0615871f019c38ada8bc1

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.1.8-0+deb10u1 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /backburner
COPY --from=builder /usr/src/backburner/target/release/backburner /backburner

CMD ["./backburner"]