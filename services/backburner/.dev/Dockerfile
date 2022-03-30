FROM rust:1.59.0-buster@sha256:bf759cad8c6d7b3a5f879ac55f753e3b60fa010a955a8f6ccf4e0c6a5d3a45f8

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.1.8-0+deb10u1 && \
    rm -rf /var/lib/apt/lists/* && \
    rustup component add clippy

WORKDIR /usr/src/backburner

CMD ["bash"]