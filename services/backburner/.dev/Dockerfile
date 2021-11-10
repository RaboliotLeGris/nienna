FROM rust:1.56.1-buster@sha256:26e9b519d01ad1ee05a1facd7e3cd91ba726174fe4fe1fcc1b584e1dbe47179e

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.1.8-0+deb10u1 && \
    rm -rf /var/lib/apt/lists/* && \
    rustup component add clippy

WORKDIR /usr/src/backburner

CMD ["bash"]