FROM rust:1.56.1-buster@sha256:26e9b519d01ad1ee05a1facd7e3cd91ba726174fe4fe1fcc1b584e1dbe47179e as builder

WORKDIR /usr/src/backburner
COPY . .
RUN cargo build --release

FROM debian:10.11-slim@sha256:cfadaa70a452076a478dcfcf2da3a3dc3b741450afed5ab67a7f3df161387c48

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.1.8-0+deb10u1 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /backburner
COPY --from=builder /usr/src/backburner/target/release/backburner /backburner

CMD ["./backburner"]