FROM rust:1.56.1-buster@sha256:26e9b519d01ad1ee05a1facd7e3cd91ba726174fe4fe1fcc1b584e1dbe47179e as builder

WORKDIR /usr/src/backburner
COPY . .
RUN cargo build --release

FROM debian:11.1-slim@sha256:a23887a2e830b815955e010f30d4c2430cd5ef82e93c130471024bc9f808d5d3

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.3.3-0+deb11u1 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /backburner
COPY --from=builder /usr/src/backburner/target/release/backburner /backburner

CMD ["./backburner"]