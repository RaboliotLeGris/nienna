FROM rust:1.59.0-buster@sha256:5e7c9c3f5b837e20a8b0decab5238a8593644f14fe50afa5251233dbd0aae63c as builder

WORKDIR /usr/src/backburner
COPY . .
RUN cargo build --release

FROM debian:11.2-slim@sha256:d5cd7e54530a8523168473a2dcc30215f2c863bfa71e09f77f58a085c419155b

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.3.3-0+deb11u1 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /backburner
COPY --from=builder /usr/src/backburner/target/release/backburner /backburner

CMD ["./backburner"]