FROM rust:1.56.1-buster@sha256:26e9b519d01ad1ee05a1facd7e3cd91ba726174fe4fe1fcc1b584e1dbe47179e as builder

WORKDIR /usr/src/backburner
COPY . .
RUN cargo build --release

FROM debian:10.11-slim@sha256:544c93597c784cf68dbe492ef35c00de7f4f6a990955c7144a40b20d86a3475f

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.1.8-0+deb10u1 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /backburner
COPY --from=builder /usr/src/backburner/target/release/backburner /backburner

CMD ["./backburner"]