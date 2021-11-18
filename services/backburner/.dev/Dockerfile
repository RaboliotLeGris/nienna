FROM rust:1.56.1-buster@sha256:05e1b389a5e66c5ee62c1371bce5d6037d5d78ce74fb1085ed51cc4a16c5ca94

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.1.8-0+deb10u1 && \
    rm -rf /var/lib/apt/lists/* && \
    rustup component add clippy

WORKDIR /usr/src/backburner

CMD ["bash"]