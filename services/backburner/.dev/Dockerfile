FROM rust:1.56.1-buster@sha256:16071065e8ad59bf5a4ee3695b4fff8f463c858e1cf5c2cb5091ab14d1309a6d

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.1.8-0+deb10u1 && \
    rm -rf /var/lib/apt/lists/* && \
    rustup component add clippy

WORKDIR /usr/src/backburner

CMD ["bash"]