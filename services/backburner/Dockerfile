FROM rust:1.50.0-buster@sha256:25a32551f42722169ea0b50b9ea752bd43a31b3e52371e735ea675866e4eb164 as builder

WORKDIR /usr/src/backburner
COPY . .
RUN cargo build --release

FROM debian:10.11-slim@sha256:544c93597c784cf68dbe492ef35c00de7f4f6a990955c7144a40b20d86a3475f

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.1.8-0+deb10u1 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /backburner
COPY --from=builder /usr/src/backburner/target/release/backburner /backburner

CMD ["./backburner"]