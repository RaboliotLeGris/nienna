FROM rust:1.59.0-buster@sha256:80f2e747f4d6b572e79b845df87b47dd9102f236113efd8921a115f4515b7df1

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.1.8-0+deb10u1 && \
    rm -rf /var/lib/apt/lists/* && \
    rustup component add clippy

WORKDIR /usr/src/backburner

CMD ["bash"]