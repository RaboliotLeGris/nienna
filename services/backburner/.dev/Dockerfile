FROM rust:1.60.0-buster@sha256:141ac79d31cb2d2583428d34197486d2f0974d53464811d9ff22b8bc14da4d05

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.1.8-0+deb10u1 && \
    rm -rf /var/lib/apt/lists/* && \
    rustup component add clippy

WORKDIR /usr/src/backburner

CMD ["bash"]