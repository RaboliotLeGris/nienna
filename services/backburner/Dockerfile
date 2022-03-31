FROM rust:1.59.0-buster@sha256:97d3ed3fe650d672771f9af901bb672b44968aa0af7a62462afec86e71c02989 as builder

WORKDIR /usr/src/backburner
COPY . .
RUN cargo build --release

FROM debian:11.3-slim@sha256:78fd65998de7a59a001d792fe2d3a6d2ea25b6f3f068e5c84881250373577414

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.3.3-0+deb11u1 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /backburner
COPY --from=builder /usr/src/backburner/target/release/backburner /backburner

CMD ["./backburner"]