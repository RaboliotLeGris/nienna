FROM rust:1.50.0-buster@sha256:91a70ac97ddb82978060db53b994f5edac1253f33dd417551810a4feac0b2245

RUN apt-get update && apt-get install --no-install-recommends -y ffmpeg=7:4.1.8-0+deb10u1 && \
    rm -rf /var/lib/apt/lists/* && \
    rustup component add clippy

WORKDIR /usr/src/backburner

CMD ["bash"]