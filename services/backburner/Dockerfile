FROM rust:1.56.1-buster@sha256:16071065e8ad59bf5a4ee3695b4fff8f463c858e1cf5c2cb5091ab14d1309a6d as builder

WORKDIR /usr/src/backburner
COPY . .
RUN cargo build --release

FROM debian:11.1-slim@sha256:a23887a2e830b815955e010f30d4c2430cd5ef82e93c130471024bc9f808d5d3

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.3.3-0+deb11u1 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /backburner
COPY --from=builder /usr/src/backburner/target/release/backburner /backburner

CMD ["./backburner"]