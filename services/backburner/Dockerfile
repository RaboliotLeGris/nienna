FROM rust:1.59.0-buster@sha256:80f2e747f4d6b572e79b845df87b47dd9102f236113efd8921a115f4515b7df1 as builder

WORKDIR /usr/src/backburner
COPY . .
RUN cargo build --release

FROM debian:11.2-slim@sha256:40f90ead34f61f38c69fd604878c73c75ec944055328b9883f3b643c1b8886f0

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.3.3-0+deb11u1 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /backburner
COPY --from=builder /usr/src/backburner/target/release/backburner /backburner

CMD ["./backburner"]