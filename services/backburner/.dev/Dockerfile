FROM rust:1.57.0-buster@sha256:e6bb347e01ca274c3d4511b765f7cb9ce495d8bad3a3a99b46ae1337f7af52f6

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.1.8-0+deb10u1 && \
    rm -rf /var/lib/apt/lists/* && \
    rustup component add clippy

WORKDIR /usr/src/backburner

CMD ["bash"]