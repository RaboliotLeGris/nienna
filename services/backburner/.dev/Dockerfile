FROM rust:1.60.0-buster@sha256:774194aa0d2ad9a8ca090d02b8aedb49b6915fc9c72d02d57ed8d8c86783a913

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.1.8-0+deb10u1 && \
    rm -rf /var/lib/apt/lists/* && \
    rustup component add clippy

WORKDIR /usr/src/backburner

CMD ["bash"]