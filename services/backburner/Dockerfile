FROM rust:1.60.0-buster@sha256:774194aa0d2ad9a8ca090d02b8aedb49b6915fc9c72d02d57ed8d8c86783a913 as builder

WORKDIR /usr/src/backburner
COPY . .
RUN cargo build --release

FROM debian:11.3-slim@sha256:f75d8a3ac10acdaa9be6052ea5f28bcfa56015ff02298831994bd3e6d66f7e57

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.3.3-0+deb11u1 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /backburner
COPY --from=builder /usr/src/backburner/target/release/backburner /backburner

CMD ["./backburner"]