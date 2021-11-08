FROM rust:1.50.0-buster@sha256:91a70ac97ddb82978060db53b994f5edac1253f33dd417551810a4feac0b2245 as builder

WORKDIR /usr/src/backburner
COPY . .
RUN cargo build --release

FROM debian:10.8-slim@sha256:13f0764262a064b2dd9f8a828bbaab29bdb1a1a0ac6adc8610a0a5f37e514955

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.1.8-0+deb10u1 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /backburner
COPY --from=builder /usr/src/backburner/target/release/backburner /backburner

CMD ["./backburner"]