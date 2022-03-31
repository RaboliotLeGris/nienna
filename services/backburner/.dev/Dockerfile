FROM rust:1.59.0-buster@sha256:97d3ed3fe650d672771f9af901bb672b44968aa0af7a62462afec86e71c02989

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.1.8-0+deb10u1 && \
    rm -rf /var/lib/apt/lists/* && \
    rustup component add clippy

WORKDIR /usr/src/backburner

CMD ["bash"]