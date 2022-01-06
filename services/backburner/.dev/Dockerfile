FROM rust:1.56.1-buster@sha256:242e85e59adbc82596282dccfb09cd7d9c71b539e02030493eae2355319465f7

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.1.8-0+deb10u1 && \
    rm -rf /var/lib/apt/lists/* && \
    rustup component add clippy

WORKDIR /usr/src/backburner

CMD ["bash"]