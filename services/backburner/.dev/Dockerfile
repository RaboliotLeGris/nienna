FROM rust:1.50.0-buster@sha256:25a32551f42722169ea0b50b9ea752bd43a31b3e52371e735ea675866e4eb164

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.1.8-0+deb10u1 && \
    rm -rf /var/lib/apt/lists/* && \
    rustup component add clippy

WORKDIR /usr/src/backburner

CMD ["bash"]