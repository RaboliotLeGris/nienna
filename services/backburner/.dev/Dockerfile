FROM rust:1.59.0-buster@sha256:5e7c9c3f5b837e20a8b0decab5238a8593644f14fe50afa5251233dbd0aae63c

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.1.8-0+deb10u1 && \
    rm -rf /var/lib/apt/lists/* && \
    rustup component add clippy

WORKDIR /usr/src/backburner

CMD ["bash"]