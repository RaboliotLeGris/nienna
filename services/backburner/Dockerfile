FROM rust:1.56.1-buster@sha256:16071065e8ad59bf5a4ee3695b4fff8f463c858e1cf5c2cb5091ab14d1309a6d as builder

WORKDIR /usr/src/backburner
COPY . .
RUN cargo build --release

FROM debian:11.1-slim@sha256:656b29915fc8a8c6d870a1247aad7796ce296729b0ae95e168d1cfe30dd2fb3c

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.3.3-0+deb11u1 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /backburner
COPY --from=builder /usr/src/backburner/target/release/backburner /backburner

CMD ["./backburner"]